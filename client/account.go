package client

func (c *client) GetAccountInfo() (map[string]string, error) {
	response := responseMap
	if err := c.request("GET", "/user_info.php",nil, response); err != nil {
		return nil, err
	}
	return response, nil
}