package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Landing_cloudCloudGetrepository(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing_cloud.cloud.getrepository", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetdemositelist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing_cloud.cloud.getdemositelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetdemopagelist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing_cloud.cloud.getdemopagelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGeturlpreview(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing_cloud.cloud.geturlpreview", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetappitems(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing_cloud.cloud.getappitems", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetappitemmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("landing_cloud.cloud.getappitemmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
