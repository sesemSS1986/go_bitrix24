package client

func (c *Client) CrmSiteFormFill(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.site.form.fill", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmSiteFormUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.site.form.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
