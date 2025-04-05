package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) DepartmentFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("department.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("department.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("department.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("department.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) DepartmentDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("department.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
