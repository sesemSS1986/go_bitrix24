package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskElapseditemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.elapseditem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
