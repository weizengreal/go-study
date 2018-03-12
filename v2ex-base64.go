package main

import (
	"encoding/base64"
	"fmt"
	"flag"
)

func main() {

	codeType := flag.Int("type", 1, "please input your code type,0 means encode,else means decode")

	var codeString string

	flag.StringVar(&codeString, "code", "", "input string your want to encode or decode")

	flag.Parse()

	if codeString == "" {
		fmt.Println("please input code")
		return
	}

	if *codeType == 0 {
		// encode
		fmt.Printf("encode string is: %s \n", base64.StdEncoding.EncodeToString([] byte(codeString)))
	} else {
		// decode
		decodeString, err := base64.StdEncoding.DecodeString(codeString)
		if err != nil {
			println(err)
		} else {
			fmt.Printf("decode string is: %s \n", string(decodeString))
		}
	}

}
