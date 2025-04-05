package client

func (c *Client) ImDialogGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogMessagesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.messages.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogUsersGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.users.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogRead(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.read", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogReadall(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.readall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogUnread(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.unread", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogWriting(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.dialog.writing", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
