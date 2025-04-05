package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmTrackingTraceAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.tracking.trace.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTrackingTraceDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.tracking.trace.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
