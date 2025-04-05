package client

func (c *Client) TaskCtaskplannermaintanceGetcurrenttaskslist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskplannermaintance.getcurrenttaskslist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskplannermaintanceGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskplannermaintance.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
