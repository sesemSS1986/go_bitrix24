package client

func (c *Client) MessageserviceMessageStatusUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"messageservice.message.status.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
