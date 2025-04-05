package client

func (c *Client) ImSearchUserList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.user.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchUser(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.user", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchChatList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.chat.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchChat(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.chat", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchDepartmentList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.department.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchDepartment(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.department", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchLastGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.last.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchLastAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.last.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImSearchLastDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.search.last.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
