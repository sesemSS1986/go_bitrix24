package client

func (c *Client) DocumentgeneratorStub(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "documentgenerator.stub", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
