package user

import (
	"errors"
	"requestencrypt/database"
	"requestencrypt/kriptografi"
	"strings"
)

func (d *DatabaseKeys) DatabaseKeyPost() error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	db.Create(&d)
	return nil
}

func Auth(username, password string) (*DatabaseKeys, error){
	var d DatabaseKeys
	db, err := database.ConnectDB()
	if err != nil{
		return nil, err
	}
	
	err = db.Where(map[string]interface{}{"username":username}).Find(&d).Error
	
	if err != nil{
		return nil, errors.New("internal error")
	}
	if username == ""{
		return nil, errors.New("pengguna tidak terdaftar")
	}
	md5pass := kriptografi.Md5([]byte(password))
	if !strings.EqualFold(d.Password, md5pass){
		return nil, errors.New("password salah")
	}
	return &d, nil
}