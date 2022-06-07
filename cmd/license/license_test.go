package license

import (
	"github.com/aalkan/directadmin-go"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	c := NewClient(&directadmin.Config{
		Username: "#",
		Password: "#",
	})

	response := responseMap
	if err := c.request("GET", "/user_info.php",nil, response); err != nil {
		t.Error(err)
	}
}
