package utils1

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"net/http"
)

var (
	sessionName string
	store       *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	sessionName = viper.GetString("sessionName")
}

/**
获取session用户
*/
func GetSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	fmt.Println("val:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
	fmt.Println("username:", username)
	return username, nil
}

/**
设置session用户
*/
func SetSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

/**
清除session
*/
func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
