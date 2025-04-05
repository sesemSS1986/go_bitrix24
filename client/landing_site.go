package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LandingSiteGetadditionalfields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.getadditionalfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteGetpublicurl(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.getpublicurl", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteMarkdelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.markdelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteMarkundelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.markundelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSitePublication(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.publication", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteUnpublic(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.unpublic", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteFullexport(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.fullexport", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteSetrights(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.setrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteGetrights(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.getrights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSiteSetscope(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.site.setscope", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
