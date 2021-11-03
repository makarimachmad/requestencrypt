package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PostPengunjung(ctx echo.Context) error{
	resp := new(Response)
	resp.Error = true
	resp.Message = "error post"

	pengunjung := new(Pengunjung)
	err := ctx.Bind(pengunjung)
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}

	lahirParse,_ := time.Parse("2006-01-02", pengunjung.TanggalLahir)
	
	ty,_,_ := time.Now().Date()
	by,_,_ := lahirParse.Date()

	usia := ty-by
	
	pengunjung.Umur = int(usia)
	pengunjung.CreatedDate = time.Now().String()

	err = pengunjung.PengunjungPost()
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}
	resp.Error = false
	resp.Message = "success post"
	resp.Data = pengunjung
	return ctx.JSON(http.StatusOK, resp)
}

func GetPengunjung(ctx echo.Context) error{
	resp := new(Response)
	resp.Error = true
	resp.Message = "error get"

	pengunjungs := new(Pengunjungs)
	err := ctx.Bind(pengunjungs)
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}

	err = pengunjungs.PengunjungGet()
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}
	resp.Error = false
	resp.Message = "success get"
	resp.Data = pengunjungs
	return ctx.JSON(http.StatusOK, resp)
}

func UpdatePengunjung(ctx echo.Context) error{
	resp := new(Response)
	resp.Error = true
	resp.Message = "error update"

	pengunjung := new(Pengunjung)
	err := ctx.Bind(pengunjung)
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}

	id := ctx.Param("idx")

	lahirParse,_ := time.Parse("2006-01-02", pengunjung.TanggalLahir)
	
	ty,_,_ := time.Now().Date()
	by,_,_ := lahirParse.Date()

	usia := ty-by
	
	pengunjung.Umur = int(usia)

	err = pengunjung.PengunjungUpdate(id)
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}
	resp.Error = false
	resp.Message = "success update data"
	resp.Data = pengunjung
	return ctx.JSON(http.StatusOK, resp)
}

func DeletePengunjung(ctx echo.Context) error{
	resp := new(Response)
	resp.Error = true
	resp.Message = "error delete"

	id := ctx.Param("idx")
	pengunjung := new(Pengunjung)
	err := pengunjung.PengunjungDelete(id)
	if err != nil{
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp.Error = false
	resp.Message = "success delete data"
	return ctx.JSON(http.StatusOK, resp)
}