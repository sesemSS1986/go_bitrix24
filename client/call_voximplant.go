package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) VoximplantCallbackStart(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.callback.start", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantInfocallStartwithtext(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.infocall.startwithtext", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantInfocallStartwithsound(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.infocall.startwithsound", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
