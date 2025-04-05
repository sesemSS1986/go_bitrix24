package client

func (c *Client) PullWatchExtend(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"pull.watch.extend", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
