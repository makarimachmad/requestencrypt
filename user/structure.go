package user

import(
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