package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strings"
)

/*
Required 不为空，即各个类型要求不为其零值
Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
Numeric 数字，有效类型：string，其他类型都将不能通过验证
AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
Base64 base64 编码，有效类型：string，其他类型都将不能通过验证
Mobile 手机号，有效类型：string，其他类型都将不能通过验证
Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证
*/

type Form struct {
	Name    string `valid:"Required;MaxSize(15)"`
	Age     int    `valid:"Range(18,100)"`
	Email   string `valid:"Email"`
	Mobile  string `valid:"Mobile"`
	Phone   string `valid:"Phone"`
	ZipCode string `valid:"ZipCode"`
	Ip      string `valid:"IP"`
}

type FormValidController struct {
	beego.Controller
}

func (p *Form) Valid(val *validation.Validation) {
	if strings.Index(p.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		val.SetError("Name", "名称里不能含有 admin")
	}
}

func (p *FormValidController) Valid() {
	form := &Form{}
	err := p.ParseForm(form)
	if err != nil {
		p.Ctx.WriteString("parse form failed!")
		return
	}

	val := validation.Validation{}
	val.Required(form.Name, "name")
	val.MaxSize(form.Name, 15, "nameMax")
	val.Range(form.Age, 18, 100, "age").Message("少儿不宜")
	val.Email(form.Email, "email")
	val.Mobile(form.Mobile, "mobile")
	val.Phone(form.Phone, "email")
	val.ZipCode(form.ZipCode, "zipcode")
	val.IP(form.Ip, "ip")

	if val.HasErrors() {
		var buf bytes.Buffer
		buf.WriteString("check failed : ")
		for _, error := range val.Errors {
			buf.WriteString(error.Message + "；")
		}
		p.Ctx.WriteString(buf.String())
		return
	}

	p.Ctx.WriteString("check success!")
}

func (p *FormValidController) ValidByStructTag() {
	form := &Form{}
	err := p.ParseForm(form)
	if err != nil {
		p.Ctx.WriteString("parse form failed!")
		return
	}

	val := validation.Validation{}
	b, err := val.Valid(form)
	if err != nil {
		p.Ctx.WriteString("valid failed!")
		return
	}

	if !b {
		var buf bytes.Buffer
		buf.WriteString("check failed : ")
		for _, error := range val.Errors {
			buf.WriteString(error.Message + "；")
		}
		p.Ctx.WriteString(buf.String())
		return
	}

	p.Ctx.WriteString("check success!")

}
