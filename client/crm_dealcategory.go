package client

func (c *Client) CrmDealcategoryFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryStatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryStageList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.stage.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryDefaultGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.default.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryDefaultSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.dealcategory.default.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
