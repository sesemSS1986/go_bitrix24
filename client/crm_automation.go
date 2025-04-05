package client

func (c *Client) CrmAutomationTrigger(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.automation.trigger", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.automation.trigger.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.automation.trigger.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.automation.trigger.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerExecute(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.automation.trigger.execute", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
