package main

import (
	"github.com/sesemSS1986/go_bitrix24/client"
	"log"
)

func main() {
	GetToken()
}

func GetToken() {
	c, err := client.NewClientWithWebhookInCome("https://b24-t5rbna.bitrix24.ru/rest/1/f3s1ah2p3r1afwli/")
	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(false)
	c.SetDebug(true)

	resp, err := c.Profile("")

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)
}
