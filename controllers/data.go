package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/labstack/echo"
)

type Site struct {
	Response    int
	Information string
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func GetEmployees(c echo.Context) error {
	secure := c.Request().Header.Get("PrivAuth")
	if secure != "Akbarfa" {
		report := new(Site)
		report.Response = http.StatusBadRequest
		report.Information = "Only Admin Allowed"
		return c.JSON(http.StatusBadRequest, report)
	}
	cookie, err := c.Request().Cookie("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cookie)
	result := models.GetEmployee()
	return c.JSON(http.StatusOK, result)
}
func Register(c echo.Context) error {
	e := new(models.Employee)
	if err := c.Bind(e); err != nil {
		return err
	}
	result := models.Regist(e.Firstname, e.Lastname, e.User, e.Email, e.Pass)
	return c.JSON(result.Code, result.Status)
}
