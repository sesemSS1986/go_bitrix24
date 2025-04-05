package client

func (c *Client) TaskCtaskitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetexecutiveuserid(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getexecutiveuserid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetdata(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getdata", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetdescription(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getdescription", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetfiles(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getfiles", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemAddfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.addfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDeletefile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.deletefile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGettags(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.gettags", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetdependson(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getdependson", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetallowedtaskactions(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getallowedtaskactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetallowedactions(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getallowedactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemGetallowedtaskactionsasstrings(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.getallowedtaskactionsasstrings", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDelegate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.delegate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemStartexecution(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.startexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemPauseexecution(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.pauseexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDefer(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.defer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemComplete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemRenew(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemApprove(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.approve", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDisapprove(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.disapprove", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemAddtofavorite(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.addtofavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskitemDeletefromfavorite(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskitem.deletefromfavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
