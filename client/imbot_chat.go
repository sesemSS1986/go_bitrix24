package client

func (c *Client) ImbotChatAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatSetowner(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.setowner", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatSetmanager(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.setmanager", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUpdatecolor(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.updatecolor", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUpdatetitle(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.updatetitle", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUpdateavatar(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.updateavatar", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatLeave(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.leave", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUserAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUserDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUserList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.user.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatSendtyping(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.chat.sendtyping", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
