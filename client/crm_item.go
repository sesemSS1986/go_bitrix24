package client

import "github.com/sviridoves/go-bitrix/types"

func (c *Client) CrmItemFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmItemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmItemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmItemList(data interface{}) (*types.ItemsResponse, error) {
	resp, err := c.DoRaw("crm.item.list", data, &types.ItemsResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.ItemsResponse), err
}

func (c *Client) CrmItemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmItemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmItemImport(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.import", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmItemBatchImport(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.item.batchImport", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
