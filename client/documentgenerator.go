package client

func (c *Client) DocumentgeneratorStub(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("documentgenerator.stub", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
