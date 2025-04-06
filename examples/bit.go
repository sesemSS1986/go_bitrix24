package main

import (
	"fmt"
	"github.com/sesemSS1986/go_bitrix24/client"
	"log"
)

const (
	UrlWebHookIn  = "https://b24-t5rbna.bitrix24.ru/rest/1/f3s1ah2p3r1afwli/"
	UrlWebHookOut = "UrlWebHookOut"
	UrlAuth       = "https://oauth.bitrix.info/rest/"
	AuthToken     = "a607f3670077820c0077761a000000010000073a1072f43128a489d98fa45fe91ab9f2"
	RefreshToken  = "b6771a680077820c0077761a00000001000007237d81de8767a31404204aeae1f98759"
)

func main() {
	FullAccesReq()
	//WebHookReq()
}

func FullAccesReq() {

	a := client.OAuthData{AuthToken: AuthToken, RefreshToken: RefreshToken}
	c := client.Client{&a}
	p := client.Parameters{}

	resp, err := c.AppInfo(true, UrlAuth, p)

	if err != nil {
		log.Println(err)
	} else {
		//example all data response
		log.Println(resp["result"])

		//example get specific data response
		spmap, ok := resp["result"].(map[string]interface{})
		if ok {
			log.Println(spmap["client_id"])
		}

		//example data convert new map
		var convDataMap = make(map[string]string)
		val, ok := resp["result"].(map[string]interface{})
		if ok {
			for k, v := range val {
				a1, ok := val[k].(map[string]interface{})
				if ok {
					for k1, v1 := range a1 {
						b1, ok := a1[k].(map[string]interface{})
						if ok {
							for k2, v2 := range b1 {
								convDataMap[k2] = fmt.Sprint(v2)
							}
						} else {
							convDataMap[k1] = fmt.Sprint(v1)
						}
					}
				} else {
					convDataMap[k] = fmt.Sprint(v)
				}
			}
		}
		for k, v := range convDataMap {
			log.Println(k, v)
		}
	}

}

func WebHookReq() {
	c := client.Client{}
	p := client.Parameters{}

	resp, err := c.Profile(false, UrlWebHookIn, p)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp["result"])
	}
}
