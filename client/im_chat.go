package client

func (c *Client) ImChatAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatSetowner(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.setowner", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatSetmanager(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.setmanager", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUpdatecolor(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.updatecolor", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUpdatetitle(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.updatetitle", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUpdateavatar(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.updateavatar", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatLeave(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.leave", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUserAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUserDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUserList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.user.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatSendtyping(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.sendtyping", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatMute(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.mute", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatParentJoin(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.chat.parent.join", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
