# directadmin-go

### Usage

```go
c := NewClient(&directadmin.Config
{
Username: "your_username",
Password: "your_passoword",
})
params := &GetLicenseParam{
LicenceId: "1234",
}
_, err := c.GetLicense(params)
if err != nil {
fmt.Println(err)
}
```