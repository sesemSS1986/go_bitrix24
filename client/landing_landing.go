package client

func (c *Client) LandingLandingGetpreview(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.getpreview", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingGetpublicurl(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.getpublicurl", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingGetadditionalfields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.getadditionalfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingPublication(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.publication", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingUnpublic(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.unpublic", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingAddblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.addblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingDeleteblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.deleteblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkdeletedblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.markdeletedblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkundeletedblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.markundeletedblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingUpblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.upblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingDownblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.downblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingShowblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.showblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingHideblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.hideblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingCopyblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.copyblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMoveblock(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.moveblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingRemoveentities(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.removeentities", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingAddbytemplate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.addbytemplate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingCopy(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.copy", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkdelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.landing.markdelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkundelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "landing.landing.markundelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
