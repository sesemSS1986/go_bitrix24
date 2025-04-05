package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) UserconsentConsentAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userconsent.consent.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
