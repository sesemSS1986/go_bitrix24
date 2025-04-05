package client

func (c *Client) ImDiskFolderGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.disk.folder.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDiskFileCommit(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.disk.file.commit", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDiskFileDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.disk.file.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDiskFileSave(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.disk.file.save", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
