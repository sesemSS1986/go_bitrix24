package client

func (c *Client) DiskStorageGetfields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageRename(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.rename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGettypes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.gettypes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageAddfolder(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.addfolder", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGetchildren(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.getchildren", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageUploadfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.uploadfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGetforapp(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.storage.getforapp", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
