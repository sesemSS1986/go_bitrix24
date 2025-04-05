package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCtaskelapseditemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskelapseditemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskelapseditemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskelapseditemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskelapseditemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskelapseditemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskelapseditemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskelapseditem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
