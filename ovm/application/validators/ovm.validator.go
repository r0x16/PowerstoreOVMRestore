package validators

import (
	"strconv"

	"github.com/asaskevich/govalidator"
)

// Validates Ovm data
type OvmValidator struct{}

// Checks if a string represents a hostname or IP address
func (v *OvmValidator) ValidateHostname(hostname string) bool {
	return govalidator.IsDNSName(hostname)
}

// Checks if a string represents a valid port number
func (v *OvmValidator) ValidatePort(port string) bool {
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	return portNumber >= 0 && portNumber <= 65535
}
