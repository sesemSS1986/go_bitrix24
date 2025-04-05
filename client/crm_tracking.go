package client

func (c *Client) CrmTrackingTraceAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.tracking.trace.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTrackingTraceDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.tracking.trace.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
