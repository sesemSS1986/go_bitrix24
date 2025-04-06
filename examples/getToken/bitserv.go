package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
)

const (
	ipapp        = "81.200.148.182"
	domain       = "https://b24-t5rbna.bitrix24.ru"
	clientID     = "local.67f14de3c5f067.89813654"
	clientSecret = "wjb0y8AlIoZP7XffQvrxKn2CXEmYrxg0v7hdltIbmgljlDQmT0"
)

type link struct {
	Code  string `query:"code"`
	Scope string `query:"scope"`
}

// структурв нескольких шаблонов
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Метод рендера темплейтов общий
// Метод рендера темплейтов общий
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found - " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base", data)
}

func main() {
	e := echo.New()

	// Root level middleware
	//e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Использование нескольких темплейтов
	templates := make(map[string]*template.Template)
	templates["index"] = template.Must(template.ParseFiles("pages/index.html", "pages/base.html"))

	//Рендер теймплейтов
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Роутеры
	e.GET("/", Main)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:3535", ipapp)))
}

func Main(ctx echo.Context) error {
	var Url string
	Url = GenFirstLink()

	u := new(link)
	if err := ctx.Bind(u); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if u.Code != "" {
		Url = fmt.Sprintf("%v/oauth/token/?grant_type=authorization_code&client_id=%v&client_secret=%v&scope=%v&code=%v", domain, clientID, clientSecret, u.Code, u.Code)
	}

	return ctx.Render(http.StatusOK, "index", Url)
}

func GenFirstLink() string {

	firsUrl := fmt.Sprintf("%v/oauth/authorize/?response_type=code&client_id=%v", domain, clientID)

	return firsUrl
}
