package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmTypeFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.type.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTypeAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.type.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTypeGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.type.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTypeList(data interface{}) (*types.TypesResponse, error) {
	resp, err := c.DoRaw("crm.type.list", data, &types.TypesResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.TypesResponse), err
}

func (c *Client) CrmTypeUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.type.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTypeDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.type.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
