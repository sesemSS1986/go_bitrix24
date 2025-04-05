package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImChatAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatSetowner(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.setowner", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatSetmanager(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.setmanager", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUpdatecolor(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.updatecolor", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUpdatetitle(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.updatetitle", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUpdateavatar(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.updateavatar", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatLeave(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.leave", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUserAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUserDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatUserList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.user.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatSendtyping(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.sendtyping", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatMute(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.mute", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImChatParentJoin(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.chat.parent.join", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
