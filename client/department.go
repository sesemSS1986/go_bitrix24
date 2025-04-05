package client

func (c *Client) DepartmentFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "department.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "department.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "department.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "department.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "department.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
