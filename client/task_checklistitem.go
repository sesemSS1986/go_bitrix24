package client

func (c *Client) TaskChecklistitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemComplete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemRenew(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemMoveafteritem(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.moveafteritem", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskChecklistitemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.checklistitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
