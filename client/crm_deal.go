package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmDealFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealProductrowsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.productrows.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealProductrowsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.productrows.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactItemsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.contact.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactItemsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.contact.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactItemsDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.contact.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringExpose(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.recurring.expose", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationReset(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationForcecommonscopeforall(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.deal.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
