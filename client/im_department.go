package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImDepartmentGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.department.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentColleaguesList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.department.colleagues.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentColleaguesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.department.colleagues.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentManagersGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.department.managers.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImDepartmentEmployeesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.department.employees.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
