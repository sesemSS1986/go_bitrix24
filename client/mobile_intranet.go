package client

func (c *Client) MobileIntranetDepartmentsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.intranet.departments.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileIntranetStresslevelSharedataGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.intranet.stresslevel.sharedata.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
