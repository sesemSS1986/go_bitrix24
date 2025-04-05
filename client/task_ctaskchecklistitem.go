package client

func (c *Client) TaskCtaskchecklistitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemComplete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemRenew(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemMoveafteritem(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.moveafteritem", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskchecklistitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
