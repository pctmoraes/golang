package address

import "strings"

// AddressType returns the address type of an address
func AddressType(address string) string {
	addressTypes := []string{"street", "avenue", "boulevard", "highway"}

	loweredAddress := strings.ToLower(address)

	addressFirstWord := strings.Split(loweredAddress, " ")[0]

	hasAType := false
	for _, addressType := range addressTypes {
		if addressFirstWord == addressType {
			hasAType = true
		}
	}

	if hasAType {
		return strings.Title(addressFirstWord)
	}

	return "Type unknown"
}
