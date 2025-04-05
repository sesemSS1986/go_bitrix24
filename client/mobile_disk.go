package client

func (c *Client) MobileDiskFolderGetchildren(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.disk.folder.getchildren", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileDiskGetattachmentsdata(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.disk.getattachmentsdata", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
