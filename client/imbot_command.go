package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImbotCommandRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.command.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotCommandUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.command.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotCommandUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.command.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotCommandAnswer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.command.answer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
