package main

import (
	"github.com/sesemSS1986/go_bitrix24/client"
	"log"
)

const (
	UrlWebHookIn  = "https://b24-t5rbna.bitrix24.ru/rest/1/f3s1ah2p3r1afwli/"
	UrlWebHookOut = "UrlWebHookOut"
	UrlAuth       = "https://oauth.bitrix.info/rest/"
	AuthToken     = "5ae7f2670077820c0077761a00000001000007dabf86e5adfdbfaaf44e9c6fd237d033"
	RefreshToken  = "4a661a680077820c0077761a000000010000074529eb10132f9d5f8a6e9571d3e62688"
)

func main() {
	FullAccesReq()
	//WebHookReq()
}

func FullAccesReq() {
	a := client.OAuthData{AuthToken: AuthToken, RefreshToken: RefreshToken}
	c := client.Client{&a}
	p := client.Parameters{}

	resp, err := c.UserAdmin(true, UrlAuth, p)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
}

func WebHookReq() {
	c := client.Client{}
	p := client.Parameters{}

	resp, err := c.Profile(false, UrlWebHookIn, p)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
}
