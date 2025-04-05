package client

func (c *Client) BizprocWorkflowTerminate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.terminate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowStart(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.start", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowInstanceList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.instance.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.template.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.template.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.template.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.template.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowInstances(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "bizproc.workflow.instances", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
