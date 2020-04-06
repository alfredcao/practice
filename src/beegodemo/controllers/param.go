package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ParamController struct {
	beego.Controller
}

func (p *ParamController) Post() {
	// form参数解析
	//user := User{}
	//err := p.ParseForm(&user)
	//if err != nil {
	//	fmt.Println("parse form failed, err :", err)
	//	return
	//}
	//
	//fmt.Println(user)

	// 请求体解析
	//err := json.Unmarshal(p.Ctx.Input.RequestBody, &user)
	//if err != nil {
	//	fmt.Println("parse json failed, err :", err)
	//	return
	//}
	//
	//fmt.Println(user)

	// 文件上传
	file, fileHeader, err := p.GetFile("uploadfile")
	if err != nil {
		fmt.Println("get file failed, err :", err)
		return
	}
	defer file.Close()
	err = p.SaveToFile("uploadfile", "static/upload/"+fileHeader.Filename)
	if err != nil {
		fmt.Println("save file failed, err :", err)
		return
	}

}

type User struct {
	Id    string `form:"-" json:"id"`
	Name  string `form:"username" json:"username"`
	Email string `json:"email"`
}
