package client

func (c *Client) ImconnectorActivate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.activate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorStatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorConnectorDataSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.connector.data.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSendMessages(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.send.messages", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorUpdateMessages(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.update.messages", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorDeleteMessages(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.delete.messages", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSendStatusDelivery(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.send.status.delivery", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSendStatusReading(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.send.status.reading", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImconnectorSetError(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imconnector.set.error", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
