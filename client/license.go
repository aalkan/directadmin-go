package client

var responseMap = make(map[string]string)

func (c *client) GetLicense(params *GetLicenseParam) (map[string]string, error) {
	response := responseMap
	if err := c.request("GET", "/list.php", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) CreateLicense(params *CreateLicenseParam) (map[string]string, error) {
	response := responseMap
	if err := c.request("GET", "/createlicense.php", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) DeleteLicense(params *DeleteLicenseParam) (map[string]string, error) {
	response := responseMap
	if err := c.request("POST", "/deletelicense.php", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) SuspendChange(params *SuspendChangeParam) (map[string]string, error) {
	response := responseMap
	if err := c.request("POST", "/special.php", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RenewLicense renews/active payment for a license.
func (c *client) RenewLicensePayment(params *RenewLicensePaymentParam) (map[string]string, error) {
	response := responseMap
	if err := c.request("POST", "/clients/api/makepayment.php", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteLicense delete payment for a license.
func (c *client) DeleteLicensePayment(params *DeleteLicensePaymentParam) (map[string]string, error) {
	response := responseMap
	if err := c.request("POST", "/clients/api/deletelicense.php", params, response); err != nil {
		return nil, err
	}
	return response, nil
}
