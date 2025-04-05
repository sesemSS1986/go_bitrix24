package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) FaceClientIdentify(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("face.client.identify", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceClientAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("face.client.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceUserIdentify(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("face.user.identify", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceUserAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("face.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) FaceUserDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("face.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
