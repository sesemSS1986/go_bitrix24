package client

func (c *Client) ImbotBotList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.bot.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
