package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LandingLandingGetpreview(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.getpreview", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingGetpublicurl(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.getpublicurl", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingGetadditionalfields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.getadditionalfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingPublication(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.publication", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingUnpublic(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.unpublic", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingAddblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.addblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingDeleteblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.deleteblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkdeletedblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.markdeletedblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkundeletedblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.markundeletedblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingUpblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.upblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingDownblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.downblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingShowblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.showblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingHideblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.hideblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingCopyblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.copyblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMoveblock(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.moveblock", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingRemoveentities(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.removeentities", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingAddbytemplate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.addbytemplate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingCopy(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.copy", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkdelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.markdelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingLandingMarkundelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.landing.markundelete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
