package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TelephonyExternallineAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalline.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternallineUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalline.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternallineDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalline.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternallineGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalline.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
