package client

func (c *Client) LandingSiteGetadditionalfields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.getadditionalfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteGetpublicurl(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.getpublicurl", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteMarkdelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.markdelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteMarkundelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.markundelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSitePublication(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.publication", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteUnpublic(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.unpublic", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteFullexport(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.fullexport", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteSetrights(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.setrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteGetrights(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.getrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteSetscope(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.site.setscope", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
