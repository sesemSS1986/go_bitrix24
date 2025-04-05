package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCommentitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.commentitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
