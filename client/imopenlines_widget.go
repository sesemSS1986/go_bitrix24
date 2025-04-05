package client

func (c *Client) ImopenlinesWidgetConfigGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetDialogGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.dialog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetUserRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.user.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetUserConsentApply(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.user.consent.apply", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetOperatorGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.operator.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetVoteSend(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.vote.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesWidgetFormSend(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.widget.form.send", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
