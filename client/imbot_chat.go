package client

func (c *Client) ImbotChatAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatSetowner(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.setowner", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatSetmanager(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.setmanager", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUpdatecolor(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.updatecolor", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUpdatetitle(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.updatetitle", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUpdateavatar(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.updateavatar", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatLeave(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.leave", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUserAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUserDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatUserList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.user.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotChatSendtyping(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.chat.sendtyping", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
