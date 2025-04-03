package client

import "github.com/sviridoves/go-bitrix/types"

func (c *Client) ImopenlinesCrmChatUserAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.crm.chat.user.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesCrmChatGetlastid(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.crm.chat.getlastid", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
