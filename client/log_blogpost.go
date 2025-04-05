package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LogBlogpostGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostShare(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.share", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostGetusersImportant(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogpost.getusers.important", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
