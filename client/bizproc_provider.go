package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) BizprocProviderAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.provider.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocProviderDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.provider.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocProviderList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.provider.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
