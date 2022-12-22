package main

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL   = "https://www.newyorkcares.org"
	loginURL  = "/user"
	rlc       = "rescue-good-food-becoming-waste"
	coalition = "staff-mobile-soup-kitchen"
)

var (
	username = ""
	password = ""
	formID   = "user_login"
	op       = "Log in"
)

type App struct {
	Client *http.Client
}

type Project struct {
	Name string
}

func (app *App) login() {
	client := app.Client

	loginURL := baseURL + loginURL

	data := url.Values{
		"name":    {username},
		"pass":    {password},
		"form_id": {formID},
		"op":      {op},
	}

	response, err := client.PostForm(loginURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
}

func (app *App) getProjects() {
	projectSearchURL := baseURL + "/search/projects/results?location=496"
	client := app.Client
	response, err := client.Get(projectSearchURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

}

func main() {
	jar, _ := cookiejar.New(nil)

	app := App{
		Client: &http.Client{Jar: jar},
	}

	app.login()
}
