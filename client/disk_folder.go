package client

func (c *Client) DiskFolderGetfields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderGetchildren(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.getchildren", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderAddsubfolder(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.addsubfolder", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderCopyto(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.copyto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderMoveto(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.moveto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderRename(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.rename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderDeletetree(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.deletetree", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderMarkdeleted(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.markdeleted", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderRestore(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.restore", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderUploadfile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.uploadfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderGetexternallink(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.getexternallink", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderSharetouser(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.sharetouser", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFolderListallowedoperations(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.folder.listallowedoperations", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
