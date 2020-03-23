package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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

func AddThread(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprint(resp, "添加贴子信息失败，请联系系统管理员")
		return
	}

	userId, _ := strconv.ParseInt(req.FormValue("userId"), 10, 64)
	thread := dao.Thread{
		Uuid:    uuid.New().String(),
		Title:   req.FormValue("title"),
		Content: req.FormValue("content"),
		UserId:  userId,
	}

	id, err := dao.Insert(&thread)
	if err != nil {
		fmt.Fprint(resp, "添加贴子信息失败，请联系系统管理员")
		return
	}
	fmt.Fprint(resp, fmt.Sprintf("添加帖子信息成功，帖子ID : %d", id))

}
