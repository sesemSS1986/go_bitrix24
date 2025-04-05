package client

func (c *Client) LandingDemosGetsitelist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.demos.getsitelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosGetpagelist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.demos.getpagelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosGeturlpreview(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.demos.geturlpreview", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.demos.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.demos.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.demos.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
