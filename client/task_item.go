package client

func (c *Client) TaskItemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetexecutiveuserid(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getexecutiveuserid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetdata(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getdata", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetdescription(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getdescription", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetfiles(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getfiles", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemAddfile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.addfile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDeletefile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.deletefile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGettags(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.gettags", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetdependson(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getdependson", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetallowedtaskactions(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getallowedtaskactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetallowedactions(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getallowedactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemGetallowedtaskactionsasstrings(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.getallowedtaskactionsasstrings", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemIsactionallowed(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDelegate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.delegate", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemStartexecution(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.startexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemPauseexecution(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.pauseexecution", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDefer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.defer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemComplete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemRenew(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.renew", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemApprove(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.approve", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDisapprove(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.disapprove", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemAddtofavorite(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.addtofavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemDeletefromfavorite(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.deletefromfavorite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGetfields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.getfields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskItemUserfieldGettypes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.item.userfield.gettypes", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
