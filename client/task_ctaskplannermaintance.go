package client

func (c *Client) TaskCtaskplannermaintanceGetcurrenttaskslist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskplannermaintance.getcurrenttaskslist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskplannermaintanceGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskplannermaintance.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
