package client

func (c *Client) ImbotDialogGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
