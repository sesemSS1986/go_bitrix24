package client

func (c *Client) BizprocEventSend(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.event.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
