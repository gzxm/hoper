package timepill

import (
	"encoding/json"
	"fmt"
	"github.com/actliboy/hoper/server/go/lib/utils/log"
	"github.com/actliboy/hoper/server/go/lib/utils/net/http/client"
	"net/http"
	"strconv"
)

const baseUrl = "https://open.timepill.net/api"
const v2Url = "https://v2.timepill.net/api"

var ApiService = &apiService{}

type apiService struct{}

func (api *apiService) GetSelfInfo() *User {
	var selfInfo User
	err := getV2("/users/my", nil, &selfInfo)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &selfInfo
}

func (api *apiService) GetUserInfo(id int) *User {
	var selfInfo User
	err := getV2("/users/"+strconv.Itoa(id), nil, &selfInfo)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &selfInfo
}

type Page struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type TodayDiariesReq struct {
	Page
	FirstId string `json:"first_id"`
}

type TodayDiaries struct {
	Count    int      `json:"count"`
	Page     string   `json:"page"`
	PageSize string   `json:"page_size"`
	Diaries  []*Diary `json:"diaries"`
}

func (api *apiService) GetTodayDiaries(page, pageSize int, firstId string) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1("/diaries/today", &TodayDiariesReq{Page{page, pageSize}, firstId}, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

func (api *apiService) GetTodayTopicDiaries(page, pageSize int, firstId string) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1("/topic/diaries", &TodayDiariesReq{Page{page, pageSize}, firstId}, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

func (api *apiService) GetFollowDiaries(page, pageSize int, firstId string) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1("/diaries/follow", &TodayDiariesReq{Page{page, pageSize}, firstId}, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

type NotebookDiaries struct {
	Count    int      `json:"count"`
	Page     string   `json:"page"`
	PageSize string   `json:"page_size"`
	Items    []*Diary `json:"items"`
}

func (api *apiService) GetNotebookDiaries(id, page, pageSize int) *NotebookDiaries {
	var notebookDiaries NotebookDiaries
	err := getV1(fmt.Sprintf("/notebooks/%d/diaries", id), &Page{page, pageSize}, &notebookDiaries)
	if err != nil {
		log.Error(err)
	}
	return &notebookDiaries
}

func (api *apiService) GetNotebook(id int) *NoteBook {
	var notebook NoteBook
	err := getV1(fmt.Sprintf("/notebooks/%d", id), nil, &notebook)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &notebook
}

func (api *apiService) GetUserTodayDiaries(userId int) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1(fmt.Sprintf("/users/%d/diaries", userId), nil, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

func (api *apiService) GetDiaryComments(diaryId int) []*Comment {
	var comments []*Comment
	err := getV1(fmt.Sprintf("/diaries/%d/comments", diaryId), nil, &comments)
	if err != nil {
		log.Error(err)
		return nil
	}
	return comments
}

func (api *apiService) GetUserNotebooks(userId int) []*NoteBook {
	var notebooks []*NoteBook
	err := getV1(fmt.Sprintf("/users/%d/notebooks", userId), nil, &notebooks)
	if err != nil {
		log.Error(err)
		return nil
	}
	return notebooks
}

func (api *apiService) GetRelationUsers(page, pageSize int) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1("/relation", &Page{page, pageSize}, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

func (api *apiService) GetRelationReverseUsers(page, pageSize int) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1("/relation/reverse", &Page{page, pageSize}, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

func (api *apiService) DeleteDiary(diaryId int) *Response {
	var res Response
	err := call(http.MethodDelete, baseUrl+fmt.Sprintf("/diaries/%d", diaryId), nil, &res)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &res
}

func (api *apiService) DeleteNotebook(noteBookId int) *Response {
	var res Response
	err := call(http.MethodDelete, baseUrl+fmt.Sprintf("/notebooks/%d", noteBookId), nil, &res)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &res
}

func (api *apiService) GetRelation(userId int) *TodayDiaries {
	var todayDiaries TodayDiaries
	err := getV1(fmt.Sprintf("/relation/%d", userId), nil, &todayDiaries)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &todayDiaries
}

func (api *apiService) GetDiary(diaryId int) *Diary {
	var diary Diary
	err := getV1(fmt.Sprintf("/diaries/%d", diaryId), nil, &diary)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &diary
}

func getV1(api string, param, result any) error {
	return call(http.MethodGet, baseUrl+api, param, result)
}
func postV1(api string, param, result any) error {
	return call(http.MethodPost, baseUrl+api, param, result)
}

func getV2(api string, param, result any) error {
	return call(http.MethodGet, v2Url+api, param, result)
}
func postV2(api string, param, result any) error {
	return call(http.MethodPost, v2Url+api, param, result)
}

func callV1(method, api string, param, result any) error {
	return call(method, baseUrl+api, param, result)
}

func callV2(method, api string, param, result any) error {
	return call(method, v2Url+api, param, result)
}

func call(method, api string, param, result any) error {
	return client.NewRequest(api, method, param).SetHeader("Authorization", Token).SetLogger(nil).Do(result)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func LikeDiary(id int) *Response {
	var res Response
	err := callV2("PUT", "/like/diaries/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		log.Error(err)
	}
	return &res
}

func UpdateUserIcon(photoUri string) *Response {
	var res Response
	err := upload("POST", "/users/icon", json.RawMessage(`{
icon: {uri: photoUri, name: 'image.jpg', type: 'image/jpg'}
})`), &res)
	if err != nil {
		log.Error(err)
	}
	return &res
}

func upload(method, api string, param, result any) error {
	return client.NewRequest(api, method, param).SetContentType(client.ContentTypeForm).SetHeader("Authorization", Token).SetLogger(nil).Do(result)
}

func UpdateUserInfo(name, intro string) *Response {
	var res Response
	err := call("PUT", "/users", json.RawMessage(`{
name: name,
intro: intro
}`), &res)
	if err != nil {
		log.Error(err)
	}
	return &res
}
