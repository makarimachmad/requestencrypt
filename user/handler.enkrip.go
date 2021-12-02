package user

import (
	// "bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"requestencrypt/kriptografi"
	"time"

	"github.com/labstack/echo/v4"
)

func PostEnkrip(ctx echo.Context) error {
	resp := new(Response)
	resp.Error = true
	resp.Message = "error post"

	pengunjung := new(Pengunjung)
	err := ctx.Bind(pengunjung)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, resp)
	}

	lahirParse, _ := time.Parse("2006-01-02", pengunjung.TanggalLahir)

	ty, _, _ := time.Now().Date()
	by, _, _ := lahirParse.Date()

	usia := ty - by

	pengunjung.Umur = int(usia)
	pengunjung.CreatedDate = time.Now().String()

	//encrypt
	reqEnc := new(kriptografi.Pengunjung)
	reqEnc.Req, _ = json.Marshal(pengunjung)
	err = reqEnc.Create()
	if err != nil {
		return err
	}
	//end encrypt

	resp.Error = false
	resp.Message = "success post enkrip"
	resp.Data = reqEnc.Res
	
	return ctx.JSON(http.StatusOK, resp)
}

func GetEnkrip(ctx echo.Context) error{
	e := errors.New("gagal cek kontak pada saat ini")
	resp := new(Response)
	resp.Error = true
	resp.Message = "error get"

	p := new(kriptografi.Pengunjung)
	err := p.Encrypt(true)
	if err != nil{
		log.Println("enkrip krip")
		return e
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:1323/v1/pengunjung/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Source", "pengunjung")
	if err != nil{
		log.Println("localhost")
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		log.Println("client do")
		return e
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Println("readall")
		return errors.New("error membaca cek pengunjung response")
	}
	json.Unmarshal(body, &p)

	if res.StatusCode != http.StatusOK || len(p.Res.Errors) > 0{
		var errString string
		for _, e := range p.Res.Errors{
			errString = string(e)
		}
		return errors.New(errString)
	}
	if len(p.Data) == 0{
		return errors.New("tidak ada respon dari sebelah")
	}
	p.CipherText = p.Data
	if p.Decrypt(false) != nil{
		return errors.New("error ketika dekrip isi pesan dari sebelah")
	}

	json.Unmarshal(p.PlainText, &p.Res)

	resp.Error = false
	resp.Message = "success get"
	resp.Data = p
	return ctx.JSON(http.StatusOK, resp)
}
