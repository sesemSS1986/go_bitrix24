package client

func (c *Client) DiskFolderGetfields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderGetchildren(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.getchildren", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderAddsubfolder(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.addsubfolder", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderCopyto(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.copyto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderMoveto(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.moveto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderRename(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.rename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderDeletetree(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.deletetree", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderMarkdeleted(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.markdeleted", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderRestore(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.restore", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderUploadfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.uploadfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderGetexternallink(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.getexternallink", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderSharetouser(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.sharetouser", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderListallowedoperations(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.folder.listallowedoperations", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
