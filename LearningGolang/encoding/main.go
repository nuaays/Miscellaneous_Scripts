package main

import	(
	"fmt"
	"encoding/base64"
	_ "encoding/hex"
	_ "encoding/json"
	_ "encoding/xml"
	_ "encoding/csv"
	_ "os"
	"encoding/hex"
)


func demoBase64(message	string)	{
	fmt.Println("--------Demo	encoding	base64--------")
	fmt.Printf("plaintext          : %s\n", message)

	encoding	:=	base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Printf("base64	msg        : %s\n", encoding)

	decoding,	_	:=	base64.StdEncoding.DecodeString(encoding)
	fmt.Printf("decoding base64	msg: %s\n", string(decoding))

}

func	demoHex(message	string)	{
	fmt.Println("--------Demo	encoding	Hex--------")
	fmt.Printf("plaintext          : %s\n", message)

	encoding	:=	hex.EncodeToString([]byte(message))
	fmt.Printf("hex	    msg        : %s\n", encoding)

	decoding,	_	:=	hex.DecodeString(encoding)
	fmt.Printf("decoding hex	msg: %s\n", string(decoding))

}



func main() {
	message	:=	"hello,go	(*w3hu%#"
	demoBase64(message)
	fmt.Println("================================================================================")
	demoHex(message)


}
