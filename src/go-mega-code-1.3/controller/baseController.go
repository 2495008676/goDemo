package controller

import (
	"go-mega-code-1.3/utils"
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
	pageLimit      int
)

func init() {
	templates = utils1.PopulateTemplates()
	pageLimit = 5
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
