package kriptografi

import "crypto/cipher"

var (
	SecretKey                  string
	SaltLen                    = 4
	Key                        []byte
	Salt                       string
	EncSalt, PreSalt, PostSalt []byte
	Block                      cipher.Block

	LisenceStatus bool
	ClientID      string
	ProductKEY    string
	SecretKEY     string
)

type (
	Crypt struct {
		PlainText  []byte
		CipherText string
		Req        Request
		ReqBuffer  []byte
		Res        Response
		Pon	   PonseRes
	}

	Request struct {
		Data string `json:"data"`
	}

	Response struct {
		Status        string      `json:"status"`
		Message       string      `json:"message"`
		Data          string      `json:"data,omitempty"`
		Errors        string      `json:"errors,omitempty"`
		DecryptedData interface{} `json:"decrypted_data,omitempty"`
		Error	      bool	  `json:"error"`
	}

	PonseRes struct {
		Status        string      `json:"status"`
		Message       string      `json:"message"`
		Data          string      `json:"data,omitempty"`
		Errors        string      `json:"errors,omitempty"`
		DecryptedData interface{} `json:"decrypted_data,omitempty"`
	}
	ResPengunjung struct {
		Status        string      `json:"status"`
		Message       string      `json:"message"`
		Data          interface{} `json:"data"`
		Errors        string      `json:"errors,omitempty"`
		DecryptedData interface{} `json:"decrypted_data,omitempty"`
		Error	      bool	  `json:"error"`
	}

	Pengunjung struct{
		Req	[]byte
		Response
		Crypt
		Res 	ResPengunjung
		Respon Bantu
	}

	Pengunjungs []Pengunjung

	Bantu struct{
		ID		int	`json:"id"`
		Nama		string	`json:"nama"`
		Umur		int	`json:"umur"`
		Alamat		string	`json:"alamat"`
		Pekerjaan	string	`json:"pekerjaan"`
		TanggalLahir	string	`json:"tanggal_lahir"`
		CreatedDate	string	`json:"created_date"`
	}

	Resdek struct{
		Error	bool		`json:"error"`
		Message	string		`json:"message"`
		Data	interface{}	`json:"data"`
		Erorrs	interface{}	`json:"errors"`
	}
)
