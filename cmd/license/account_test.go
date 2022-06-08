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

	_, err := c.GetAccountInfo()
	if err != nil {
		t.Error(err)
	}
}
