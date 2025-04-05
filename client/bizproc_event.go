package client

func (c *Client) BizprocEventSend(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.event.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
