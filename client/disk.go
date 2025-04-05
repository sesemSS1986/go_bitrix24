package client

func (c *Client) Download(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "download", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Upload(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "upload", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
