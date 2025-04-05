package client

func (c *Client) ListsGetIblockTypeId(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.get.iblock.type.id", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
