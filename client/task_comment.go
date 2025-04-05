package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCommentGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.comment.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.comment.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
