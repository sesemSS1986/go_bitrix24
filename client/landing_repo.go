package client

func (c *Client) LandingRepoCheckcontent(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.checkcontent", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoGetappinfo(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.getappinfo", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoBind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoUnbind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingRepoGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.repo.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
