package ovm

import "github.com/r0x16/go-ovm-client/ovmHelper"

type OvmClient struct {
	Helper  *ovmHelper.Client
	manager *ovmHelper.Manager
}

// Returns a new client from ovm helper package
// This client allows access all ovm services
func NewOvmClient(user, password string, baseUrl string) *OvmClient {
	return &OvmClient{
		Helper: ovmHelper.NewClient(user, password, baseUrl),
	}
}

// returns a new or cached ovm manager details object
// This object is cached because rarelly changes, changes doesn't
// affects other services access or data
func (c *OvmClient) GetManager() (*ovmHelper.Manager, error) {
	if c.manager != nil {
		return c.manager, nil
	}

	var err error
	c.manager, err = c.Helper.Managers.Read()

	if err != nil {
		return nil, err
	}

	return c.manager, nil
}
