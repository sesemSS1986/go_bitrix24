package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImbotDialogGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
