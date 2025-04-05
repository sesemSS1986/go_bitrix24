package client

func (c *Client) ImDialogGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogMessagesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.messages.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogUsersGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.users.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogRead(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.read", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogReadall(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.readall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogUnread(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.unread", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDialogWriting(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.dialog.writing", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
