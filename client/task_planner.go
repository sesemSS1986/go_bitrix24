package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskPlannerGetcurrenttaskslist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.planner.getcurrenttaskslist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskPlannerGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.planner.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
