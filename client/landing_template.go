package client

func (c *Client) LandingTemplateGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.template.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateSetsiteref(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.template.setsiteref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateSetlandingref(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.template.setlandingref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateGetsiteref(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.template.getsiteref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateGetlandingref(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.template.getlandingref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
