package license

import (
	"github.com/aalkan/directadmin-go"
	"testing"
)

func TestGetLicense(t *testing.T) {
	c := NewClient(&directadmin.Config{
		Username: "#",
		Password: "#",
	})
	params := &GetLicenseParam{
		//LicenceId: "1234",
		Ip: "192.1682.1.1",
		//Active: "Y",
		//SortBy: "ip",
	}

	_, err := c.GetLicense(params)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateLicense(t *testing.T) {
	c := NewClient(&directadmin.Config{
		Username: "#",
		Password: "#",
	})
	params := &CreateLicenseParam{}

	_, err := c.CreateLicense(params)
	if err != nil {
		t.Error(err)
	}
}
