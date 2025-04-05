package client

func (c *Client) TelephonyCallAttachtranscription(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.call.attachtranscription", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
