package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmSiteFormFill(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.site.form.fill", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmSiteFormUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.site.form.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
