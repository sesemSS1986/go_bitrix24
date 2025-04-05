package client

func (c *Client) MobileIntranetDepartmentsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.intranet.departments.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileIntranetStresslevelSharedataGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.intranet.stresslevel.sharedata.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
