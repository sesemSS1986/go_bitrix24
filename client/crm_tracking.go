package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmTrackingTraceAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.tracking.trace.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTrackingTraceDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.tracking.trace.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
