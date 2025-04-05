package client

func (c *Client) BizprocRobotAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.robot.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocRobotUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.robot.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocRobotDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.robot.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocRobotList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.robot.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
