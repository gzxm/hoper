package fs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"time"
)

func FetchFile(url string) (*FileInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	vbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	var file FileInfo
	file.Binary = vbytes
	file.name = path.Base(resp.Request.URL.Path)
	file.modTime, _ = time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	file.size, _ = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	return &file, nil
}
