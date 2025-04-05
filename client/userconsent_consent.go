package client

func (c *Client) UserconsentConsentAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userconsent.consent.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
