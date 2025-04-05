package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) VoximplantUrlGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.url.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.sip.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.sip.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.sip.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.sip.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipStatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.sip.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantSipConnectorStatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.sip.connector.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantStatisticGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.statistic.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineOutgoingSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.line.outgoing.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineOutgoingGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.line.outgoing.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineOutgoingSipSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.line.outgoing.sip.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantLineGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.line.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantTtsVoicesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.tts.voices.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserGetdefaultlineid(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.user.getdefaultlineid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserActivatephone(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.user.activatephone", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantUserDeactivatephone(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.user.deactivatephone", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantAuthorizationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.authorization.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantAuthorizationSignonetimekey(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.authorization.signonetimekey", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantAuthorizationOnerror(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.authorization.onerror", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallInit(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.init", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallStartwithdevice(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.startwithdevice", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallHangupdevice(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.hangupdevice", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSendwait(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.sendwait", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSendready(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.sendready", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallAnswer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.answer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSkip(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.skip", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallHold(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.hold", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallUnhold(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.unhold", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallStartviarest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.startviarest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallIntercept(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.intercept", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantCallSavecomment(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("voximplant.call.savecomment", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
