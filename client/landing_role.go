package client

func (c *Client) LandingRoleGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.role.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleGetrights(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.role.getrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleSetrights(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.role.setrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleSetaccesscodes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.role.setaccesscodes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleIsenabled(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.role.isenabled", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRoleEnable(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.role.enable", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
