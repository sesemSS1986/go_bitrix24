package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LandingDemosGetsitelist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.demos.getsitelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosGetpagelist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.demos.getpagelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosGeturlpreview(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.demos.geturlpreview", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.demos.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.demos.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingDemosGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.demos.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
