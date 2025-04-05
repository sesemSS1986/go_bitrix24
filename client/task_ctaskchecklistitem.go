package client

func (c *Client) TaskCtaskchecklistitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemComplete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemRenew(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemMoveafteritem(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.moveafteritem", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskchecklistitemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskchecklistitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
