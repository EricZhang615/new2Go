package controller

import (
	"strconv"
	"test-sample/service"
)

type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{Code: 1, Msg: "strconv error", Data: err}
	}
	pageInfo, err := service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{Code: 2, Msg: "service.QueryPageInfo error", Data: err}
	}

	return &PageData{Code: 0, Msg: "", Data: pageInfo}
}
