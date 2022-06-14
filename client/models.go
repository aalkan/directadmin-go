package client

type AccountInfoResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Balance string `json:"balance"`
}

type GetLicenseParam struct {
	LicenceId string `json:"lid"`
	Ip        string `json:"ip"`
	Active    string `json:"active"`
	SortBy    string `json:"sort-by"`
}

type CreateLicenseParam struct {
	Id       string `json:"id"` // Client ID
	Password string `json:"password"`
	Name     string `json:"name"`    // License Name
	Pid      string `json:"pid"`     // Product ID
	Os       string `json:"os"`      // Operating System Name
	Ip       string `json:"ip"`      // IP Address
	Payment  string `json:"payment"` // Payment Type (balance or mail)
	Email    string `json:"email"`   // Email Address
}

type DeleteLicenseParam struct {
	LicenseId string `json:"lid"` // License ID
}

type SuspendChangeParam struct {
	LicenseId string `json:"lid"`       // License ID
	Suspended string `json:"suspended"` // Suspended (Y or N)
}

type RenewLicensePaymentParam struct {
	LicenseId string `json:"lid"` // License ID
}

type DeleteLicensePaymentParam struct {
	LicenseId string `json:"lid"` // License ID
}
