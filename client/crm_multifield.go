package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmMultifieldFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.multifield.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
