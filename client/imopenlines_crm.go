package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImopenlinesCrmChatUserAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.crm.chat.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesCrmChatGetlastid(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.crm.chat.getlastid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
