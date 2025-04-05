package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmLeadFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadProductrowsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.productrows.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadProductrowsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.productrows.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationReset(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationForcecommonscopeforall(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.lead.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
