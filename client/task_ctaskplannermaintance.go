package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCtaskplannermaintanceGetcurrenttaskslist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskplannermaintance.getcurrenttaskslist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskplannermaintanceGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskplannermaintance.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
