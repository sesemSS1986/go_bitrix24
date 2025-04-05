package client

func (c *Client) LandingBlockClonecard(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.clonecard", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockAddcard(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.addcard", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockRemovecard(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.removecard", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatecards(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.updatecards", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockChangenodename(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.changenodename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockChangeanchor(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.changeanchor", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatenodes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.updatenodes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatestyles(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.updatestyles", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdateattrs(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.updateattrs", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetcontent(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getcontent", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUpdatecontent(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.updatecontent", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetbyid(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getbyid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetmanifestfile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getmanifestfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetcontentfromrepository(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getcontentfromrepository", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockGetrepository(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.getrepository", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingBlockUploadfile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.block.uploadfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
