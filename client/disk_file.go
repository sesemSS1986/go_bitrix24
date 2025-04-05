package client

func (c *Client) DiskFileGetfields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileCopyto(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.copyto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileMoveto(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.moveto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileRename(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.rename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileMarkdeleted(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.markdeleted", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileRestore(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.restore", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileUploadversion(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.uploadversion", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileGetexternallink(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.getexternallink", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileGetversions(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.getversions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileRestorefromversion(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.restorefromversion", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileListallowedoperations(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "disk.file.listallowedoperations", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
