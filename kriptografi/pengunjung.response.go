package kriptografi

import (
	"encoding/json"
	"errors"
	"log"
)

func (p *Pengunjung) CreateResponse() error {
	e := errors.New("cannt request pengunjung at this time")
	resp := new(Response)
	resp.Error = true
	resp.Message = "error post"

	p.PlainText = p.Req
	bantu := new(Bantu)
	json.Unmarshal(p.PlainText, &bantu)

	err := p.Encrypt(true)
	if err != nil {
		log.Println("encrypt")
		return e
	}
	return nil
}