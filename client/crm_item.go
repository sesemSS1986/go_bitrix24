package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmItemFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemList(data interface{}) (*types.ItemsResponse, error) {
	resp, err := c.Request("crm.item.list", data, &types.ItemsResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.ItemsResponse), err
}

func (c *Client) CrmItemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemImport(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.import", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemBatchImport(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.item.batchImport", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
