package client

func (c *Client) SalePaysystemHandlerAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.handler.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemHandlerUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.handler.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemHandlerDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.handler.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemHandlerList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.handler.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.settings.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsInvoiceGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.settings.invoice.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsPaymentGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.settings.payment.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemPayInvoice(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.pay.invoice", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemPayPayment(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sale.paysystem.pay.payment", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
