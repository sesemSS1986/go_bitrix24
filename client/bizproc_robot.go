package client

func (c *Client) BizprocRobotAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.robot.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocRobotUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.robot.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocRobotDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.robot.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocRobotList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.robot.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
