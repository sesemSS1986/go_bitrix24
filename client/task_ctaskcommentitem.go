package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCtaskcommentitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcommentitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
