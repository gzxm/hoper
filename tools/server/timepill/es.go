package timepill

import (
	"bytes"
	"context"
	"encoding/json"
	utilv8 "github.com/actliboy/hoper/server/go/lib/utils/dao/es/v8"
	"github.com/actliboy/hoper/server/go/lib/utils/def/request"
	"github.com/actliboy/hoper/server/go/lib/utils/io/reader"
	"github.com/actliboy/hoper/server/go/lib/utils/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net/http"
	"strconv"
	"strings"
)

const DiaryIndex = "diary"

const Mapping = `
{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0
    },
    "mappings": {
        "properties": {
			"id": {
                "type": "long"
            },
            "user_id": {
                "type": "keyword",
                "fields": {
                    "raw": {
                        "type": "long"
                    }
                }
            },
            "notebook_id": {
                "type": "keyword",
                "fields": {
                    "raw": {
                        "type": "long"
                    }
                }
            },
            "notebook_subject": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart",
                "store": true,
                "fielddata": true
            },
            "content": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "created": {
                "type": "date"
            }
        }
    }
}`

type IndexDiary struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id" gorm:"index"`
	NoteBookId      int    `json:"notebook_id" gorm:"index"`
	NoteBookSubject string `json:"notebook_subject" gorm:"index"`
	Content         string `json:"content" gorm:"type:text"`
	Created         string `json:"created" gorm:"timestamptz(6);default:'0001-01-01 00:00:00';index"`
}

func (diary *Diary) DiaryIndex() *IndexDiary {

	return &IndexDiary{
		Id:              diary.Id,
		UserId:          diary.UserId,
		NoteBookId:      diary.NoteBookId,
		NoteBookSubject: diary.NoteBookSubject,
		Content:         diary.Content,
		Created:         diary.Created,
	}
}

type EsDao struct {
	ctx context.Context
	Es8 *elasticsearch.Client
}

func (dao *EsDao) MaxIdEs8() int {
	size := 1
	req := esapi.SearchRequest{
		Index: []string{DiaryIndex},
		Sort:  []string{"id:desc"},
		Size:  &size,
	}

	resp, err := utilv8.GetSearchResponse[IndexDiary](req.Do(dao.ctx, dao.Es8))
	if err != nil {
		log.Error(err)
		return 0
	}
	if resp.Hits.Total.Value > 0 {
		id, _ := strconv.Atoi(resp.Hits.Hits[0].Id)
		return id
	}
	return 0
}

func (dao *EsDao) LoadEs8() {
	req := &ListReq{
		ListReq: request.ListReq{
			PageReq: request.PageReq{PageNo: 1, PageSize: Conf.TimePill.PageSize},
			SortReq: request.SortReq{SortField: "id", SortType: request.SortTypeASC},
		},
		RangeReq: request.RangeReq{
			RangeField: "id",
			RangeStart: dao.MaxIdEs8(),
			RangeEnd:   nil,
			Include:    false,
		},
	}

	for {
		req.PageSize = Conf.TimePill.PageSize
		if req.PageSize < 1 {
			req.PageSize = 10
		}
		diaries, err := Dao.DBDao(dao.ctx).ListDB(req)
		if err != nil {
			log.Error(err)
		}
		for i, diary := range diaries {
			body, _ := json.Marshal(diary.DiaryIndex())
			esreq := esapi.CreateRequest{
				Index:      DiaryIndex,
				DocumentID: strconv.Itoa(diary.Id),
				Body:       bytes.NewReader(body),
			}
			resp, err := esreq.Do(dao.ctx, Dao.Es8)
			if err != nil {
				// Handle error
				log.Error(err)
			}
			bytes, err := reader.ReadCloser(resp.Body)
			log.Info(string(bytes))
			if i == len(diaries)-1 {
				req.RangeStart = diary.Id
			}
		}
		if len(diaries) < req.PageSize {
			break
		}
	}
}

func (dao *dao) CreateIndexEs8() {
	resp, err := dao.Es8.Indices.Exists([]string{DiaryIndex})
	if err != nil {
		// Handle error
		panic(err)
	}
	resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		log.Info("index not found")
		_, err := dao.Es8.Indices.Create(DiaryIndex, dao.Es8.Indices.Create.WithBody(strings.NewReader(Mapping)))
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
		log.Info("index created")
	} else {
		log.Info("index found")
	}
}
