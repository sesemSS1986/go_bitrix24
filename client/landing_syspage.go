package client

func (c *Client) LandingSyspageSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.syspage.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.syspage.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageDeleteforsite(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.syspage.deleteforsite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageDeleteforlanding(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.syspage.deleteforlanding", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageGetspecialpage(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.syspage.getspecialpage", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
