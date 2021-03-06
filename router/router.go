package router

import (
	"net/http"
	"requestencrypt/user"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(r *echo.Echo){
	r.GET("/", func (c echo.Context) error{
		return c.String(http.StatusOK, "berhasil terhubung")
	})

	app := r.Group("/v1")

	p := app.Group("/pengunjung")
	p.GET("/login", user.GetLogin).Name = "get-login"
	p.POST("/registrasi", user.PostPengunjung).Name = "post-pengunjung"
	p.POST("/coba", user.Coba).Name = "post-pengunjung"
	p.GET("/", user.GetPengunjung).Name="get-pengunjung"
	p.PATCH("/:idx", user.UpdatePengunjung).Name="update-pengunjung"
	p.DELETE("/:idx", user.DeletePengunjung).Name="delete-pengunjung"

	//encrypt
	e := app.Group("/enkrip")
	e.GET("/login", user.GetLogin).Name = "get-login"
	e.POST("/registrasi", user.PostEnkrip).Name = "post-pengunjung"
	e.GET("/", user.GetEnkrip).Name="get-pengunjung"
	e.PATCH("/:idx", user.UpdatePengunjung).Name="update-pengunjung"
	e.DELETE("/:idx", user.DeletePengunjung).Name="delete-pengunjung"
}