package main

import (
	"github.com/sesemSS1986/go_bitrix24/client"
	"log"
)

const (
	UrlWebHookIn  = "https://b24-t5rbna.bitrix24.ru/rest/1/f3s1ah2p3r1afwli/"
	UrlWebHookOut = "UrlWebHookOut"
	UrlAuth       = "https://oauth.bitrix.info/rest/"
	AuthToken     = "80a0f1670077820c0077761a00000001000007ed97fc5d34b82d09a7a12e77690b9309"
	RefreshToken  = "701f19680077820c0077761a0000000100000729cec85eef5dd68f60393c635882e9be"
)

func main() {
	FullAccesReq()
}

func FullAccesReq() {
	a := client.OAuthData{AuthToken: AuthToken, RefreshToken: RefreshToken}
	c := client.Client{&a}
	p := client.Parameters{}

	resp, err := c.Profile(true, UrlAuth, p)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
}

func WebHookReq() {
	c := client.Client{}
	p := client.Parameters{}

	resp, err := c.TaskItemList(false, UrlWebHookIn, p)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
}
