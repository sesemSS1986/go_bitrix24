package client

func (c *Client) ImDepartmentGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.department.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentColleaguesList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.department.colleagues.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentColleaguesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.department.colleagues.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentManagersGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.department.managers.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentEmployeesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.department.employees.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
