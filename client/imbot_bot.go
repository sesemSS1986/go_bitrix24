package client

func (c *Client) ImbotBotList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.bot.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
