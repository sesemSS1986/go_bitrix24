package client

func (c *Client) LogBlogpostGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostShare(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.share", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogpostGetusersImportant(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.blogpost.getusers.important", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
