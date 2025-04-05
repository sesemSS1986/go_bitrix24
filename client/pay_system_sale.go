package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) SalePaysystemHandlerAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.handler.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemHandlerUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.handler.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemHandlerDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.handler.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemHandlerList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.handler.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.settings.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsInvoiceGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.settings.invoice.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemSettingsPaymentGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.settings.payment.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemPayInvoice(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.pay.invoice", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) SalePaysystemPayPayment(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sale.paysystem.pay.payment", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
