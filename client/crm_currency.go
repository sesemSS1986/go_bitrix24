package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmCurrencyFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.localizations.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.localizations.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.localizations.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.localizations.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyBaseSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.base.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyBaseGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.currency.base.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
