package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"requestencrypt/kriptografi"
	"time"

	"github.com/labstack/echo/v4"
)

func PostPengunjung(ctx echo.Context) error {
	fmt.Println("ctx: ", ctx)
	resp := new(Response)
	resp.Error = true
	resp.Message = "error post"

	pengunjung := new(Pengunjung)
	pengunjungdekrip := new(kriptografi.Pengunjung)

	err := ctx.Bind(pengunjungdekrip)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	pengunjungdekrip.CipherText = pengunjungdekrip.Data

	//start dekrip
	errors := pengunjungdekrip.Decrypt(true)
	if errors != nil {
		fmt.Println("gagal dekrip")
	}
	//end dekrip

	json.Unmarshal(pengunjungdekrip.PlainText, &pengunjung)

	lahirParse, _ := time.Parse("2006-01-02", pengunjung.TanggalLahir)

	ty, _, _ := time.Now().Date()
	by, _, _ := lahirParse.Date()

	usia := ty - by

	pengunjung.Umur = int(usia)
	pengunjung.CreatedDate = time.Now().String()

	err = pengunjung.PengunjungPost()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp.Error = false
	resp.Message = "success post dekrip"
	resp.Data = pengunjung

	fmt.Println("--------------- enkrip response ----------")
	responResp := new(kriptografi.Crypt)
	responResp.PlainText, _ = json.Marshal(resp)
	err = responResp.Encrypt(true)
	if err != nil{
		return err
	}
	fmt.Println("responResp enkrip: ", responResp)

	return ctx.JSON(http.StatusOK, responResp)
}

func Coba(ctx echo.Context) error {
	log.Println("hai")
	return nil
}

func GetPengunjung(ctx echo.Context) error {
	resp := new(Response)
	resp.Error = true
	resp.Message = "error get"

	pengunjungs := new(Pengunjungs)
	err := ctx.Bind(pengunjungs)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	err = pengunjungs.PengunjungGet()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp.Error = false
	resp.Message = "success get"
	resp.Data = pengunjungs
	return ctx.JSON(http.StatusOK, resp)
}

func UpdatePengunjung(ctx echo.Context) error {
	resp := new(Response)
	resp.Error = true
	resp.Message = "error update"

	pengunjung := new(Pengunjung)
	err := ctx.Bind(pengunjung)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	id := ctx.Param("idx")

	lahirParse, _ := time.Parse("2006-01-02", pengunjung.TanggalLahir)

	ty, _, _ := time.Now().Date()
	by, _, _ := lahirParse.Date()

	usia := ty - by

	pengunjung.Umur = int(usia)

	err = pengunjung.PengunjungUpdate(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}
	resp.Error = false
	resp.Message = "success update data"
	resp.Data = pengunjung
	return ctx.JSON(http.StatusOK, resp)
}

func DeletePengunjung(ctx echo.Context) error {
	resp := new(Response)
	resp.Error = true
	resp.Message = "error delete"

	id := ctx.Param("idx")
	pengunjung := new(Pengunjung)
	err := pengunjung.PengunjungDelete(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	resp.Error = false
	resp.Message = "success delete data"
	return ctx.JSON(http.StatusOK, resp)
}
