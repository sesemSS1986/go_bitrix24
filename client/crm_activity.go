package client

func (c *Client) CrmActivityFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityCommunicationFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.communication.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityTypeAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.type.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityTypeList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.type.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityTypeDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.activity.type.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
