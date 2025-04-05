package client

func (c *Client) LandingRepoCheckcontent(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.checkcontent", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoGetappinfo(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.getappinfo", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoBind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoUnbind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.repo.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
