package controllers

import (
	"../models"
	"github.com/labstack/echo"
)

type Acc struct {
	Email string `json:"username" form:"username"`
	Pass  string `json:"pass"`
}

func Login(c echo.Context) error {
	e := new(Acc)
	if err := c.Bind(e); err != nil {
		return err
	}
	result := models.Login(e.Email, e.Pass)

	return c.JSON(result.Response, result.Log)
}
