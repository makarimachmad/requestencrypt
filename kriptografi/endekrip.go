package kriptografi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
)

func EncryptionInit() (err error) {
	a := []rune(SecretKey)
	if len(a) < 32 {
		err = errors.New("init panjang enkripsi error")
		return
	}

	salt := string(a[0:12])
	key := []byte(string(a[0:32]))

	Block, err = aes.NewCipher(key)
	if err != nil{
		err = errors.New("init blok enkripsi gagal")
		return
	}

	EncSalt = hasher256(salt)
	PreSalt = EncSalt[0:SaltLen]
	PostSalt = EncSalt[len(EncSalt)-SaltLen:]
	return
}

func hasher256(text string) []byte{
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hasher.Sum(nil)
}

func genHMAC256(ciphertext []byte) []byte {
	mac := hmac.New(sha256.New, Key)
	mac.Write(ciphertext)
	return mac.Sum(nil)
}

func (c *Crypt) Encrypt(isRequest bool) error {
	var encData []byte
	plainTextLen := len(c.PlainText)

	enc := make([]byte, plainTextLen)
	ivBlock := make([]byte, aes.BlockSize+plainTextLen)

	ivEnc := ivBlock[:aes.BlockSize]
	io.ReadFull(rand.Reader, ivEnc)

	cipher.NewCTR(Block, ivEnc).XORKeyStream(enc, c.PlainText)

	encData = append(PreSalt, ivEnc...)
	encData = append(encData, genHMAC256(enc)...)
	encData = append(encData, enc...)
	encData = append(encData, PostSalt...)

	c.CipherText = base64.StdEncoding.EncodeToString(encData)
	if isRequest {
		c.Req.Data = c.CipherText
		c.ReqBuffer, _ = json.Marshal(c.Req)
	}
	return nil
}

func (c *Crypt) Decrypt(isResponse bool) error {
	var ciphertext, iv, hmacOri []byte
	var sha2len, hmacLen, ivLen int
	e := errors.New("can't decrypt data")

	if isResponse {
		ciphertext, _ = base64.StdEncoding.DecodeString(c.Res.Data)
	} else {
		if c.CipherText == "" {
			return e
		}
		ciphertext, _ = base64.StdEncoding.DecodeString(c.CipherText)
	}

	iv = ciphertext[SaltLen : aes.BlockSize+SaltLen]

	ivLen = len(iv)
	sha2len = 32
	hmacLen = SaltLen + ivLen + sha2len
	hmacOri = ciphertext[SaltLen+ivLen : hmacLen]

	ciphertextRaw := ciphertext[hmacLen : len(ciphertext)-SaltLen]
	hmacCompare := genHMAC256(ciphertextRaw)

	plaintext := make([]byte, len(ciphertextRaw))

	stream := cipher.NewCTR(Block, iv)
	stream.XORKeyStream(plaintext, ciphertextRaw)
	if hmac.Equal(hmacOri, hmacCompare) {
		c.PlainText = plaintext
		return nil
	}
	return e
}

func Md5(p []byte) string {
	hash := md5.New()
	hash.Write(p)
	return hex.EncodeToString(hash.Sum(nil))
}