package client

func (c *Client) ImconnectorActivate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.activate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorStatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorConnectorDataSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.connector.data.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSendMessages(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.send.messages", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorUpdateMessages(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.update.messages", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorDeleteMessages(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.delete.messages", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSendStatusDelivery(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.send.status.delivery", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSendStatusReading(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.send.status.reading", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSetError(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imconnector.set.error", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
