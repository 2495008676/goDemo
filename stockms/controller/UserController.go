package controller

import (
	"net/http"
	"stockms/framework"
	"stockms/service"
	"stockms/utils"
)

/**
 * r.PostFormValue  : 可以解析 Post/PUT Content-Type=application/x-www-form-urlencoded 或 Content-Type=multipart/form-data
 */

type UserController struct {
}

var userService = new(service.UserService)

func (p *UserController) Router(router *framework.RouterHandler) {
	router.Router("/register", p.register)
	router.Router("/login", p.login)
	router.Router("/findAll", p.findAll)
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserController) register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		framework.ResultFail(w, "username or password can not be empty")
		return
	}
	id := userService.Insert(username, password)
	if id <= 0 {
		framework.ResultFail(w, "register fail")
		return
	}
	framework.ResultOk(w, "register success")
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserController) login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		framework.ResultFail(w, "username or password can not be empty")
		return
	}
	users := userService.SelectUserByName(username)
	if len(users) == 0 {
		framework.ResultFail(w, "user does not exist")
		return
	}
	if users[0].Password != password {
		framework.ResultFail(w, "password error")
		return
	}

	framework.ResultOk(w, "login success")
}

// GET/POST
func (p *UserController) findAll(w http.ResponseWriter, r *http.Request) {
	users := userService.SelectAllUser()
	framework.ResultJsonOk(w, users)
}
