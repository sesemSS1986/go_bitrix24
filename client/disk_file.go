package client

func (c *Client) DiskFileGetfields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileCopyto(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.copyto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileMoveto(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.moveto", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileRename(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.rename", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileMarkdeleted(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.markdeleted", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileRestore(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.restore", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileUploadversion(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.uploadversion", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileGetexternallink(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.getexternallink", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileGetversions(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.getversions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileRestorefromversion(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.restorefromversion", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DiskFileListallowedoperations(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.file.listallowedoperations", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
