package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) UserconsentAgreementList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userconsent.agreement.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserconsentAgreementText(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userconsent.agreement.text", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
