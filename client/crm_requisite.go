package client

func (c *Client) CrmRequisiteFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetCountries(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.countries", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldAvailabletoadd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.availabletoadd", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisitePresetFieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.preset.field.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.bankdetail.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.bankdetail.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.bankdetail.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.bankdetail.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.bankdetail.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteBankdetailDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.bankdetail.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.link.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.link.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.link.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.link.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmRequisiteLinkUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.requisite.link.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
