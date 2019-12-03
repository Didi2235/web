package controllers

import (
	"fmt"
	"net/http"
	"time"

	"../models"
	"github.com/labstack/echo"
)

type Data struct {
	Number int    `json:"Number" form:"Number" query:"Number"`
	Otp    string `json:"Otp" form:"Otp" query:"Otp"`
}
type Kuki struct {
	Name    string
	Value   string
	Expires time.Time
	Path    string
}

func Otp(c echo.Context) error {
	e := new(Data)
	c.Bind(e)
	fmt.Println(e.Number)
	result := models.Otp(e.Number)
	return c.JSON(result.Status, result.Resu)
}

func Loginxl(c echo.Context) error {
	e := new(Data)
	c.Bind(e)
	fmt.Println(e.Number)
	result := models.Loginxl(e.Number, e.Otp)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	cookie := addCookie("token", result.Token, "/", time.Now().Add(1*time.Hour))
	cookie = addCookie("rtoken", result.RToken, "/", time.Now().Add(23*time.Hour))
	cookie = addCookie("msisdn", result.Msisdn, "/", time.Now().Add(23*time.Hour))
	c.SetCookie(cookie)
	c.Response().WriteHeader(result.Status)
	return c.JSON(result.Status, result.Resu)
}
func addCookie(Name, Value, Path string, Expires time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = Name
	cookie.Value = Value
	cookie.Path = Path
	cookie.Expires = Expires
	return cookie
}
