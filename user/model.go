package user

import (
	"log"
	"requestencrypt/database"
	"strconv"
)

func (p *Pengunjung) PengunjungPost()error {
	db, err := database.ConnectDB()
	if err != nil{
		log.Println("connectDB")
		return err
	}
	db.Create(&p)
	return nil
}

func (p *Pengunjungs) PengunjungGet() error {
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	db.Find(&p)
	return nil
}

func (p *Pengunjung) PengunjungUpdate(id string) error{
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	i,_ := strconv.Atoi(id)
	db.Model(&Pengunjung{}).Where("id = ?", i).Updates(&p)
	return nil
}

func (p *Pengunjung) PengunjungDelete(id string) error{
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	i,_ := strconv.Atoi(id)
	db.Delete(&p, i)
	return nil
}