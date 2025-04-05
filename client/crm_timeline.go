package client

func (c *Client) CrmTimelineCommentFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.comment.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.comment.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.comment.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.comment.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.comment.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineCommentUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.comment.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.bindings.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.bindings.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsBind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.bindings.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTimelineBindingsUnbind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.timeline.bindings.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
