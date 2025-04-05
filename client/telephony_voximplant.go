package client

func (c *Client) VoximplantUrlGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.url.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.sip.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.sip.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.sip.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.sip.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipStatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.sip.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipConnectorStatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.sip.connector.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantStatisticGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.statistic.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineOutgoingSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.line.outgoing.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineOutgoingGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.line.outgoing.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineOutgoingSipSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.line.outgoing.sip.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.line.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantTtsVoicesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.tts.voices.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserGetdefaultlineid(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.user.getdefaultlineid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserActivatephone(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.user.activatephone", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserDeactivatephone(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.user.deactivatephone", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantAuthorizationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.authorization.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantAuthorizationSignonetimekey(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.authorization.signonetimekey", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantAuthorizationOnerror(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.authorization.onerror", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallInit(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.init", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallStartwithdevice(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.startwithdevice", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallHangupdevice(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.hangupdevice", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSendwait(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.sendwait", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSendready(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.sendready", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallAnswer(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.answer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSkip(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.skip", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallHold(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.hold", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallUnhold(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.unhold", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallStartviarest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.startviarest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallIntercept(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.intercept", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSavecomment(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"voximplant.call.savecomment", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
