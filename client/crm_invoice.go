package client

func (c *Client) CrmInvoiceFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceGetexternallink(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.getexternallink", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.status.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.status.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.status.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.status.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceStatusDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.status.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceRecurringExpose(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.recurring.expose", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmInvoiceUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.invoice.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
