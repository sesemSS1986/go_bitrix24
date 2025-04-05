package client

func (c *Client) LandingTemplateGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.template.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateSetsiteref(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.template.setsiteref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateSetlandingref(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.template.setlandingref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateGetsiteref(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.template.getsiteref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingTemplateGetlandingref(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.template.getlandingref", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
