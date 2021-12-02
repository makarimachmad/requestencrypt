package kriptografi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (p *Pengunjung) Create() error {
	e := errors.New("cannt request pengunjung at this time")
	resp := new(Response)
	resp.Error = true
	resp.Message = "error post"

	p.PlainText = p.Req
	bantu := new(Bantu)
	json.Unmarshal(p.PlainText, &bantu)

	err := p.Encrypt(true)
	if err != nil{
		log.Println("encrypt")
		return e
	}
	
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:1323/v1/pengunjung/registrasi", bytes.NewBuffer(p.ReqBuffer))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil{
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("error reading another server response")
	}
	
	resdek := new(Resdek)
	json.Unmarshal(body, &resdek)
	fmt.Println("body: ", body)
	fmt.Println("resdek: ", resdek)

	return nil	
}

func GetPgjg(p *Pengunjung) error{
	e := errors.New("cant get pengunjung at this time")

	err := p.Encrypt(true)
	if err != nil{
		log.Println("encrypt")
		return e
	}

	req, _ := http.NewRequest(http.MethodGet, "localhost:1323/v1/pengunjung/", bytes.NewBuffer(p.ReqBuffer))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		return e
	}
	fmt.Println("lewat aman")
	defer res.Body.Close()

	return nil
}