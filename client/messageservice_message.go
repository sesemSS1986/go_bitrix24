package client

func (c *Client) MessageserviceMessageStatusUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("messageservice.message.status.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
