package client

func (c *Client) MobileDiskFolderGetchildren(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.disk.folder.getchildren", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileDiskGetattachmentsdata(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.disk.getattachmentsdata", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
