package license

import "fmt"

var responseMap =make(map[string]string)

func (c *client) GetAccountInfo() (map[string]string, error) {
	response := responseMap
	if err := c.request("GET", "/user_info.php",nil, response); err != nil {
		return nil, err
	}
	fmt.Println("response:", response)
	return response, nil
}
