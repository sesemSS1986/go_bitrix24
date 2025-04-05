package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImSearchUserList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.user.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchUser(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.user", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchChatList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.chat.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchChat(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.chat", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchDepartmentList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.department.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchDepartment(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.department", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchLastGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.last.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchLastAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.last.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchLastDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.search.last.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
