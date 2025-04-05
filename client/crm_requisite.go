package client

func (c *Client) CrmRequisiteFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetCountries(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.countries", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldAvailabletoadd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.availabletoadd", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.preset.field.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.bankdetail.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.bankdetail.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.bankdetail.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.bankdetail.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.bankdetail.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.bankdetail.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.link.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.link.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.link.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.link.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.requisite.link.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
