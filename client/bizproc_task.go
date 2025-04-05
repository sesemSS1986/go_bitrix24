package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) BizprocTaskList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.task.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocTaskComplete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.task.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
