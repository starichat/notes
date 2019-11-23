package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func (c echo.Context) dologin() (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	// todo
	// 校验数据
	return c.JSON(http.StatusOK, u)

}

func main() {
	e := echo.New()
	s := &http.Server{
		Addr:         ":1234",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Get("/", dologin)

	e.Logger.Fatal(e.StartServer(s))

}
