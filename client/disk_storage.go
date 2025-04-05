package client

func (c *Client) DiskStorageGetfields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageRename(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.rename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGettypes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.gettypes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageAddfolder(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.addfolder", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGetchildren(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.getchildren", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageUploadfile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.uploadfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskStorageGetforapp(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.storage.getforapp", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
