package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmActivityFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityCommunicationFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.communication.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityTypeAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.type.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityTypeList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.type.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmActivityTypeDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.activity.type.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
