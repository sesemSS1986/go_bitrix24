package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) MobileIntranetDepartmentsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.intranet.departments.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileIntranetStresslevelSharedataGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.intranet.stresslevel.sharedata.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
