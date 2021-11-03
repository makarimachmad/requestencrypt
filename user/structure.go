package user

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
)