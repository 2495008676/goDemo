package utils1

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var (
	flashName string
)

func init() {
	flashName = viper.GetString("flashName")
}

/**
输入长度校验
*/
func CheckLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

func SetFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := store.Get(r, sessionName)
	session.AddFlash(message, flashName)
	session.Save(r, w)
}

func GetFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}
	session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}

func GetPage(r *http.Request) int {
	url := r.URL         // net/url.URL
	query := url.Query() // Values (map[string][]string)

	q := query.Get("page")
	if q == "" {
		return 1
	}

	page, err := strconv.Atoi(q)
	if err != nil {
		return 1
	}
	return page
}
