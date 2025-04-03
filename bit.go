package main

import (
	"github.com/sviridoves/go-bitrix/client"
	"github.com/sviridoves/go-bitrix/types"
	"log"
)

func main() {
	GetToken()
}

func GetToken() {
	c, err := client.NewClientWithFirstOAuth("https://b24-t5rbna.bitrix24.ru/", "local.67eeea5d72f9b2.07847921", "8pd6hV8bv545EBQ5wsvadT6Zgf2Rr0VwjXCMCj5t0x8ygl13zB")
	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(true)
	c.SetDebug(true)

	resp, err := c.Methods(&types.MethodsRequest{
		Full: true,
	})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)
}

func Bitrix() {
	c, err := client.NewClientWithOAuth("https://b24-t5rbna.bitrix24.ru/", "AutTOKEN", "RefrAutTOKEN")

	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(true)
	c.SetDebug(true)

	resp, err := c.Methods(&types.MethodsRequest{
		Full:  true,
		Scope: "landing",
	})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)
}
