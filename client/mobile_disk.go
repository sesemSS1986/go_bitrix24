package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) MobileDiskFolderGetchildren(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.disk.folder.getchildren", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileDiskGetattachmentsdata(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.disk.getattachmentsdata", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
