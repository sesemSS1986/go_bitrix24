package client

func (c *Client) CrmDealcategoryFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryStatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryStageList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.stage.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryDefaultGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.default.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealcategoryDefaultSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.dealcategory.default.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
