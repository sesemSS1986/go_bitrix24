package client

func (c *Client) TaskChecklistitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemComplete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemRenew(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemMoveafteritem(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.moveafteritem", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.checklistitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
