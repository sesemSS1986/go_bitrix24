package client

func (c *Client) ImopenlinesBotSessionOperator(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.bot.session.operator", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesBotSessionSendMessage(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.bot.session.send.message", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesBotSessionMessageSend(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.bot.session.message.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesBotSessionTransfer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.bot.session.transfer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesBotSessionFinish(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.bot.session.finish", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
