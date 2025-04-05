package client

func (c *Client) FaceClientIdentify(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "face.client.identify", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceClientAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "face.client.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceUserIdentify(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "face.user.identify", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceUserAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "face.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceUserDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "face.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
