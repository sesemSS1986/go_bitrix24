package client

func (c *Client) BizprocWorkflowTerminate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.terminate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowStart(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.start", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowInstanceList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.instance.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.template.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.template.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.template.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowTemplateDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.template.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocWorkflowInstances(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.workflow.instances", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
