package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImDiskFolderGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.disk.folder.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDiskFileCommit(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.disk.file.commit", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDiskFileDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.disk.file.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDiskFileSave(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.disk.file.save", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
