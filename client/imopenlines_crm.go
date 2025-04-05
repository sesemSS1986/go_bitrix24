package client

func (c *Client) ImopenlinesCrmChatUserAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.crm.chat.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesCrmChatGetlastid(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.crm.chat.getlastid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
