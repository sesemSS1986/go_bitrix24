package client

func (c *Client) CrmTimelineCommentFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.comment.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.comment.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.comment.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.comment.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.comment.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.comment.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.bindings.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.bindings.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsBind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.bindings.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsUnbind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.timeline.bindings.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
