package pro

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	py2 "github.com/actliboy/hoper/server/go/lib/utils/strings/pinyin"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/actliboy/hoper/server/go/lib/utils/fs"
	"github.com/actliboy/hoper/server/go/lib/utils/strings"
	"golang.org/x/net/html"
)

var userAgent = []string{
	`Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36`,
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36",
	"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; en-US; rv:1.9.2.15) Gecko/20110303 Firefox/3.6.15",
	`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36`,
}

func SetClient(client *http.Client, timeout time.Duration, proxyUrl string) {
	if timeout < time.Second {
		timeout = timeout * time.Second
	}
	proxyURL, _ := url.Parse(proxyUrl)
	client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).DialContext,
	}
}

type Fail chan string

func NewFail(cap int) Fail {
	return make(chan string, cap)
}

func (f Fail) Do(name string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		file, _ := os.Create(Conf.Pro.CommonDir + name + time.Now().Format("2006_01_02_15_04_05") + Conf.Pro.Ext)
		for txt := range f {
			file.WriteString(txt + "\n")
		}
		file.Close()
		wg.Done()
	}()
}

type Speed struct {
	wg                    *sync.WaitGroup
	web, pic              chan struct{}
	Fail, FailPic, FailDB Fail
}

func (s *Speed) Add(i int) {
	s.wg.Add(i)
	s.pic <- struct{}{}
}

func (s *Speed) WebAdd(i int) {
	s.wg.Add(i)
	s.web <- struct{}{}
}

func (s *Speed) Done() {
	s.wg.Done()
	<-s.pic
}

func (s *Speed) WebDone() {
	s.wg.Done()
	<-s.web
}

func (s *Speed) Wait() {
	s.wg.Wait()
}

func NewSpeed(cap int) *Speed {
	return &Speed{
		wg:      new(sync.WaitGroup),
		pic:     make(chan struct{}, cap),
		web:     make(chan struct{}, cap),
		Fail:    NewFail(cap),
		FailPic: NewFail(cap),
		FailDB:  NewFail(cap),
	}
}

func Fetch(id int, sd *Speed) {
	defer sd.WebDone()
	tid := strconv.Itoa(id)
	reader, err := Request(http.DefaultClient, Conf.Pro.CommonUrl+tid)
	if err != nil {
		log.Println(err, "id:", tid)
		if !strings.HasPrefix(err.Error(), "????????????") {
			sd.Fail <- tid
		}
		invalidPost := &Post{TId: id, Status: 2}
		err := Dao.DB.Save(invalidPost).Error
		if err != nil && !strings.HasPrefix(err.Error(), "ERROR: duplicate key") {
			sd.FailDB <- tid + " 2"
		}
		return
	}
	defer reader.Close()
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
		sd.Fail <- tid
		return
	}
	s := doc.Find(`img[src="images/common/none.gif"]`)

	auth, title, text, postTime, htl, post := ParseHtml(doc)
	post.TId = id
	post.PicNum = uint32(s.Length())
	status := "0"
	if post.PicNum == 0 {
		status = "1"
	}

	dir := Conf.Pro.CommonDir + "pic_" + strconv.Itoa(id/100000) + "/"

	if auth != "" {
		dir += py2.FistLetter(auth) + Sep + auth + Sep
	}
	if title != "" {
		dir += title + `_` + tid + Sep
	}
	dir = fs.PathClean(dir)

	post.Path = dir[CommonDirLen:]
	err = Dao.DB.Save(post).Error
	if err != nil && !strings.HasPrefix(err.Error(), "ERROR: duplicate key") {
		sd.FailDB <- tid + " " + status
	}

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0666)
		if err != nil {
			log.Println(err, dir)
			sd.Fail <- tid
			return
		}
	}
	if text != "" {
		f, err := os.Create(dir + postTime + Conf.Pro.Ext)
		f.WriteString(text)
		f.Close()
		if err != nil {
			log.Println(err)
		}
		if htl.Length() > 0 {
			f, err = os.Create(dir + `index.html`)
			for c := htl.Nodes[0].FirstChild; c != nil; c = c.NextSibling {
				html.Render(f, c)
			}
			f.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}

	s.Each(func(i int, s *goquery.Selection) {
		if url, ok := s.Attr("file"); ok {
			sd.Add(1)
			go Download(url, dir, sd)
			time.Sleep(Conf.Pro.Interval)
		}
	})
}

func ParseHtml(doc *goquery.Document) (string, string, string, string, *goquery.Selection, *Post) {
	auth := doc.Find("#postlist .popuserinfo a").First().Text()
	title := doc.Find("#threadtitle h1").Text()
	timenode := doc.Find(".posterinfo .authorinfo em").First()
	postTime := timenode.Text()
	if strings.HasPrefix(postTime, "?????????") {
		postTime = postTime[len(`????????? `):]
	}
	if strings.HasSuffix(postTime, "???") {
		postTime2, ok := timenode.Find("span").First().Attr("title")
		if !ok {
			postTime = time.Now().Format("2006-01-02 15:04:05")
		} else {
			postTime = postTime2
		}
	}

	if strings.Contains(postTime, "???") {
		now := time.Now()
		var day int
		if strings.Contains(postTime, "??????") {
			day, _ = strconv.Atoi(postTime[0:0])
		} else {
			describe := postTime[:6]
			switch describe {
			case "??????":
				day = 2
			case "??????":
				day = 1
			case "??????":
				day = 0
			}
		}
		now.AddDate(0, 0, -day)
		date := now.Format("2006-01-02")
		postTime = date + " " + postTime[len(postTime)-5:]
	}

	post := &Post{
		TId:   0,
		Auth:  auth,
		Title: title,
	}

	post.CreatedAt = postTime
	content := doc.Find(".t_msgfont").First()
	text := content.Contents().Not(".t_attach").Text()
	html := content.Not(".t_attach").Not("span")
	post.Content = text
	return FixPath(auth), FixPath(title), text, postTime, html, post
}

func FixPath(path string) string {
	path = stringsi.ReplaceRuneEmpty(path, []rune{'\\', '/', ':', ' '})
	if strings.HasSuffix(path, ".") {
		path += "$"
	}
	return path
}

func Download(url, dir string, sd *Speed) {
	defer sd.Done()
	reader, err := Request(picClient, url)
	if err != nil {
		log.Println(err, "url:", url)
		if !strings.HasPrefix(err.Error(), "????????????") {
			sd.FailPic <- url + "<->" + dir
		}
		return
	}
	defer reader.Close()
	s := strings.Split(url, "//")
	name := s[len(s)-1]
	if strings.Contains(name, "/") {
		s = strings.Split(url, "/")
		name = s[len(s)-1]
	}
	if strings.Contains(name, "\\") {
		s = strings.Split(url, "\\")
		name = s[len(s)-1]
	}
	f, err := os.Create(dir + name)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, reader)
	if err != nil {
		log.Printf("?????????????????????%v, ???????????????%s,?????????%s\n", err, url, dir)
		sd.FailPic <- url + "<->" + dir
		return
	}
	log.Printf("???????????????%s,?????????%s\n", url, dir)
}

func Request(client *http.Client, url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9;charset=utf-8")
	req.Header.Set("Connection", "keep-alive")
	var reader io.ReadCloser
	var resp *http.Response
	for i := 0; i < 20; i++ {
		if i > 0 {
			time.Sleep(time.Second)
		}
		n := rand.Intn(5)
		req.Header.Set("User-Agent", userAgent[n])
		resp, err = client.Do(req)
		if err != nil {
			log.Println(err, "url:", url)
			continue
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			return nil, fmt.Errorf("???????????????????????????%d,url:%s", resp.StatusCode, url)
		}

		if resp.Header.Get("Content-Encoding") == "gzip" {
			reader, err = gzip.NewReader(resp.Body)
			if err != nil {
				if resp != nil {
					resp.Body.Close()
				}
				log.Println(err, "url:", url)
				continue
			}
		} else {
			reader = resp.Body
		}
		if reader != nil {
			break
		}
	}
	if reader == nil {
		if resp != nil {
			resp.Body.Close()
		}
		msg := "???????????????" + url
		if err != nil {
			msg = err.Error() + msg
		}
		return nil, errors.New(msg)
	}
	return reader, nil
}

func newRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9;charset=utf-8")
	req.Header.Set("Connection", "keep-alive")
	http.DefaultClient.Timeout = 300 * time.Second
	return req, nil
}

func Start(job func(sd *Speed)) {
	sd := NewSpeed(Conf.Pro.Loop)
	wg := new(sync.WaitGroup)
	sd.Fail.Do("fail_post_", wg)
	sd.FailPic.Do("fail_pic_", wg)
	sd.FailDB.Do("fail_db_", wg)
	job(sd)
	sd.Wait()
	close(sd.Fail)
	close(sd.FailPic)
	close(sd.FailDB)
	wg.Wait()
}

type Post struct {
	ID        uint32
	TId       int    `gorm:"uniqueIndex"`
	Auth      string `gorm:"size:255;default:''"`
	Title     string `gorm:"size:255;default:''"`
	Content   string `gorm:"type:text"`
	CreatedAt string `gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00'"`
	PicNum    uint32 `gorm:"default:0"`
	Score     uint8  `gorm:"default:0"`
	Status    uint8  `gorm:"default:0"`
	Path      string `gorm:"size:255;default:''"`
}

func FixWeb(path string, sd *Speed, handle func(int, *Speed)) {
	f, err := os.Open(Conf.Pro.CommonDir + path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sd.WebAdd(1)
		id, _ := strconv.Atoi(scanner.Text())
		go handle(id, sd)
		time.Sleep(Conf.Pro.Interval)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
