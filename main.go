package web

import (
	"net/http"

	"github.com/sysriot/web/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	//rendering
	e.Static("", "content")
	e.Renderer = controllers.NewRenderer("content/*.html", true)

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404nf.html", nil)
	}
	e.GET("/", controllers.Index)
	e.GET("/admin", controllers.GetEmployees)
	e.POST("/", controllers.Register)
	//myxl
	e.GET("/myxl/:Number/:Otp", controllers.Loginxl)
	e.POST("/myxl", controllers.Otp)

	e.Logger.Fatal(e.Start(":80"))

}
