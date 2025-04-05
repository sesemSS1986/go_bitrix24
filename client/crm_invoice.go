package client

func (c *Client) CrmInvoiceFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceGetexternallink(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.getexternallink", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.status.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.status.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.status.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.status.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.status.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringExpose(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.recurring.expose", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.invoice.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
