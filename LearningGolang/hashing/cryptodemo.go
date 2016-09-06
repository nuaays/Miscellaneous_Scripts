package main


import	(
	"fmt"
	"crypto/md5"
	_ "crypto/sha256"
	_ "crypto/hmac"
	_ "crypto/aes"
	_ "crypto/rsa"
	_ "crypto/cipher"
	_ "crypto/rand"
	_ "crypto/x509"
	_ "encoding/hex"
	_ "encoding/base64"
	_ "encoding/pem"
	_ "io"
	_ "io/ioutil"
	_ "encoding/hex"
	"encoding/hex"
	"crypto/sha256"
	"crypto/hmac"
)


func calHash_md5(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	message_hashed := hex.EncodeToString(h.Sum(nil))
	return message_hashed
}

func calHash_sha256(message string) string {
	h := sha256.New()
	h.Write([]byte(message))
	message_hashed := hex.EncodeToString(h.Sum(nil))
	return message_hashed
}


//Hashing	with	Key	Using	HMAC
func calHash_hmac(key, message string) string {
	hmac_key := []byte(key)
	h := hmac.New(sha256.New, hmac_key)
	h.Write([]byte(message))
	message_hashed := hex.EncodeToString(h.Sum(nil))
	return message_hashed
}

func main() {

	mystr := "nihao"
	fmt.Println("MD5   : ", calHash_md5(mystr))
	fmt.Println("SHA256: ", calHash_sha256(mystr))
	fmt.Println("HMAC  : ", calHash_hmac("za", mystr))


}
