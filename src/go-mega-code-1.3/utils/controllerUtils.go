package utils1

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"go-mega-code-1.3/config"
	"go-mega-code-1.3/dto"
	"go-mega-code-1.3/model"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

var (
	sessionName string
	flashName   string
	store       *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	sessionName = viper.GetString("sessionName")
	flashName = viper.GetString("flashName")
}

func MiddleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := GetSessionUser(r)
		log.Println("middle:", username)
		if username != "" {
			log.Println("Last seen:", username)
			model.UpdateLastSeen(username)
		}
		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

// PopulateTemplates func
// Create map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

// session

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

// Login Check
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

func checkUsername(username string) string {
	return CheckLen("Username", username, 3, 20)
}

func checkPassword(password string) string {
	return CheckLen("Password", password, 6, 50)
}

func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}

func checkUserPassword(username, password string) string {
	if !dto.CheckLogin(username, password) {
		return fmt.Sprintf("Username and password is not correct.")
	}
	return ""
}

func checkUserExist(username string) string {
	if dto.CheckUserExist(username) {
		return fmt.Sprintf("Username already exist, please choose another username")
	}
	return ""
}

func checkEmailExistRegister(email string) string {
	if dto.CheckUserExist(email) {
		return fmt.Sprintf("Email has registered by others, please choosse another email.")
	}
	return ""
}

func checkEmailExist(email string) string {
	if !dto.CheckEmailExist(email) {
		return fmt.Sprintf("Email does not register yet.Please Check email.")
	}
	return ""
}

// CheckLogin()
func CheckLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserPassword(username, password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// CheckRegister()
func CheckRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmailExistRegister(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

func CheckResetPasswordRequest(email string) []string {
	var errs []string
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmailExist(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

func CheckResetPassword(pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// AddUser()
func AddUser(username, password, email string) error {
	return dto.AddUser(username, password, email)
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

// Email

// SendEmail func
func SendEmail(target, subject, content string) {
	server, port, usr, pwd := config.GetSMTPConfig()
	d := gomail.NewDialer(server, port, usr, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", usr)
	m.SetHeader("To", target)
	m.SetAddressHeader("Cc", usr, "admin")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Email Error:", err)
		return
	}
}