package kriptografi

import "crypto/cipher"

var (
	SecretKey                  string
	SaltLen                    = 4
	Key                        []byte
	Salt                       string
	EncSalt, PreSalt, PostSalt []byte
	Block                      cipher.Block
)

type (
	Crypt struct {
		PlainText  []byte
		CipherText string
		Req        Request
		ReqBuffer  []byte
		Res        Response
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
	}
)
