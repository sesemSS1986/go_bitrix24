package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImopenlinesWidgetConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetDialogGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetUserRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.user.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetUserConsentApply(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.user.consent.apply", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetOperatorGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.operator.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetVoteSend(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.vote.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetFormSend(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.widget.form.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
