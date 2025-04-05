package client

func (c *Client) TelephonyCallAttachtranscription(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "telephony.call.attachtranscription", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
