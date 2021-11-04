package user

import (
	"crypto/md5"
	"encoding/hex"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type(
	Pengunjung struct{
		ID		int	`json:"id"`
		Nama		string	`json:"nama"`
		Umur		int	`json:"umur"`
		Alamat		string	`json:"alamat"`
		Pekerjaan	string	`json:"pekerjaan"`
		TanggalLahir	string	`json:"tanggal_lahir"`
		CreatedDate	string	`json:"created_date"`
	}
	Pengunjungs	[]Pengunjung

	DatabaseKeys struct{
		ID	int	`json:"id"`
		UserID	int	`json:"user_id"`
		Username	string	`json:"username"`
		Password	string	`json:"password"`
		EncryptionKey	string	`json:"encryption_key"`
	}
)

func (p Pengunjung) Validate() error{
	return v.ValidateStruct(&p,
		v.Field(&p.Nama, v.Required),
		v.Field(&p.Umur, v.Required),
		v.Field(&p.Alamat, v.Required),
		v.Field(&p.Pekerjaan, v.Required),
		v.Field(&p.TanggalLahir, v.Required),
	)
}

func (u *DatabaseKeys) GenPwd() {
	// decPass, _ := base64.StdEncoding.DecodeString(u.Password)
	// u.Password = Md5(decPass)
	hash := md5.New()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

func Md5(p []byte) string {
	hash := md5.New()
	hash.Write(p)
	return hex.EncodeToString(hash.Sum(nil))
}