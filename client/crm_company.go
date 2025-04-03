package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmCompanyFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyList(data interface{}) (*types.CompaniesResponse, error) {
	resp, err := c.DoRaw("crm.company.list", data, &types.CompaniesResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.CompaniesResponse), err
}

func (c *Client) CrmCompanyUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyContactFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.contact.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyContactAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.contact.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyContactDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.contact.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyContactItemsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.contact.items.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyContactItemsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.contact.items.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyContactItemsDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.contact.items.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyDetailsConfigurationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.details.configuration.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyDetailsConfigurationSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.details.configuration.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyDetailsConfigurationReset(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.details.configuration.reset", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCompanyDetailsConfigurationForcecommonscopeforall(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.company.details.configuration.forcecommonscopeforall", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
