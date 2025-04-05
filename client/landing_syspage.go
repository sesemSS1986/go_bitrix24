package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LandingSyspageSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.syspage.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.syspage.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageDeleteforsite(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.syspage.deleteforsite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageDeleteforlanding(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.syspage.deleteforlanding", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LandingSyspageGetspecialpage(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing.syspage.getspecialpage", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
