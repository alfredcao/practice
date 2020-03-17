package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gowebprograming/chitchat/dao"
	"net/http"
	"strconv"
)

func GetThreadList(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	threads, err := dao.GetThreadList()
	if err != nil {
		fmt.Fprint(resp, "获取贴子列表失败，请联系系统管理员")
	} else {
		jsonByteArr, _ := json.Marshal(threads)
		fmt.Fprint(resp, string(jsonByteArr))
	}
}

func GetThreadById(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		fmt.Fprint(resp, "传参有误")
		return
	}

	thread, err := dao.GetThreadById(id)
	if err != nil {
		fmt.Fprint(resp, "获取贴子信息失败，请联系系统管理员")
	} else {
		jsonByteArr, _ := json.Marshal(thread)
		fmt.Fprint(resp, string(jsonByteArr))
	}
}
