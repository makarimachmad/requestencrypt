package router

import(
	"net/http"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(r *echo.Echo){
	r.GET("/", func (c echo.Context) error{
		return c.String(http.StatusOK, "berhasil terhubung")
	})

	// app := r.Group("/v1")

	// p := app.Group("/pengunjung")
	// p.GET("/login").Name = "get-login"
	// p.POST("/registrasi").Name = "post-regist"
}