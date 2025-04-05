package client

func (c *Client) TaskCtaskitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetexecutiveuserid(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getexecutiveuserid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetdata(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getdata", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetdescription(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getdescription", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetfiles(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getfiles", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemAddfile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.addfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDeletefile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.deletefile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGettags(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.gettags", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetdependson(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getdependson", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetallowedtaskactions(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getallowedtaskactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetallowedactions(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getallowedactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetallowedtaskactionsasstrings(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.getallowedtaskactionsasstrings", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDelegate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.delegate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemStartexecution(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.startexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemPauseexecution(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.pauseexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDefer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.defer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemComplete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemRenew(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemApprove(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.approve", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDisapprove(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.disapprove", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemAddtofavorite(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.addtofavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDeletefromfavorite(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskitem.deletefromfavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
