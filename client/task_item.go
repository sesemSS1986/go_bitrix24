package client

func (c *Client) TaskItemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetexecutiveuserid(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getexecutiveuserid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetdata(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getdata", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetdescription(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getdescription", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetfiles(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getfiles", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemAddfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.addfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDeletefile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.deletefile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGettags(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.gettags", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetdependson(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getdependson", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetallowedtaskactions(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getallowedtaskactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetallowedactions(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getallowedactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetallowedtaskactionsasstrings(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.getallowedtaskactionsasstrings", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDelegate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.delegate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemStartexecution(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.startexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemPauseexecution(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.pauseexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDefer(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.defer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemComplete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemRenew(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemApprove(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.approve", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDisapprove(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.disapprove", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemAddtofavorite(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.addtofavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDeletefromfavorite(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.deletefromfavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGetfields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGettypes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.item.userfield.gettypes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
