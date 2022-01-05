package controllers

import (
	"beego-demo/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"strings"
)

type MainController struct {
	beego.Controller
}

// RegisterGet 注册展示页面
func (c *MainController) RegisterGet() {
	c.TplName = "register.html"
}

// RegisterPost 注册获取数据页面
func (c *MainController) RegisterPost() {
	// 1.拿到数据，并去除两边的空格
	username := strings.TrimSpace(c.GetString("username"))
	password := strings.TrimSpace(c.GetString("password"))
	// 2.对数据进行校验
	if username == "" || password == "" {
		logs.Error("用户名或密码不能为空")
		c.Redirect("/register", http.StatusFound)
		return
	}
	// 3.插入数据库
	o := orm.NewOrm()
	user := models.User{
		Name: username,
		Pwd:  password,
	}
	_, err := o.Insert(&user)
	if err != nil {
		logs.Error("insert failed, err:", err)
		c.Redirect("/register", http.StatusNotFound)
		return
	}
	// 4.返回登陆界面
	// c.TplName = "login.html" // 指定视图文件，同时可以给这个视图传递一些数据
	c.Redirect("/login", http.StatusFound) // 跳转，不能传递数据，速度快
}

// LoginGet 登陆页面get方法
func (c *MainController) LoginGet() {
	c.TplName = "login.html"
}

// LoginPost 登陆页面post方法
func (c *MainController) LoginPost() {
	// 1.拿到数据，并去除两边的空格
	username := strings.TrimSpace(c.GetString("username"))
	password := strings.TrimSpace(c.GetString("password"))
	logs.Info("username:", username, "password:", password)
	// 2.判断数据是否合法
	if username == "" || password == "" {
		logs.Error("用户名或密码不能为空")
		c.TplName = "login.html"
		c.Data["errMsg"] = "登陆失败"
	}
	// 3.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	if err := o.Read(&user, "Name"); err != nil {
		logs.Info("用户名错误")
		c.TplName = "login.html"
		c.Data["errMsg"] = "用户名或密码错误"
		return
	}
	// 判断密码是否一致
	if user.Pwd != password {
		logs.Info("密码错误")
		c.TplName = "login.html"
		c.Data["errMsg"] = "用户名或密码错误"
		return
	}
	// 4.跳转
	c.Ctx.WriteString("登陆成功")
}
