package client

func (c *Client) LandingRoleGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.role.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleGetrights(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.role.getrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleSetrights(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.role.setrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleSetaccesscodes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.role.setaccesscodes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleIsenabled(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.role.isenabled", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleEnable(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.role.enable", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
