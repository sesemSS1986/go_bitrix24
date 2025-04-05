package client

func (c *Client) ListsGetIblockTypeId(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "lists.get.iblock.type.id", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
