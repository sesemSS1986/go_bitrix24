package client

func (c *Client) CrmAutomationTrigger(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.automation.trigger", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.automation.trigger.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.automation.trigger.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.automation.trigger.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAutomationTriggerExecute(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.automation.trigger.execute", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
