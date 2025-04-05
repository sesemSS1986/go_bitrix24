package client

func (c *Client) ImbotDialogGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
