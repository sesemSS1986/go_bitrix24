package client

func (c *Client) LandingBlockClonecard(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.clonecard", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockAddcard(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.addcard", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockRemovecard(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.removecard", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatecards(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.updatecards", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockChangenodename(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.changenodename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockChangeanchor(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.changeanchor", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatenodes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.updatenodes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatestyles(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.updatestyles", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdateattrs(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.updateattrs", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetcontent(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getcontent", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatecontent(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.updatecontent", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetbyid(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getbyid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetmanifestfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getmanifestfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetcontentfromrepository(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getcontentfromrepository", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetrepository(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.getrepository", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUploadfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing.block.uploadfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
