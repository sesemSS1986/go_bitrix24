package client

func (c *Client) UserconsentAgreementList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userconsent.agreement.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserconsentAgreementText(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userconsent.agreement.text", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
