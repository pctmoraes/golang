package address

import "testing"

func TestAddressType(t *testing.T) {
	testAddress := "Street of the Silly"
	expectedType := "Street"
	receivedType := AddressType(testAddress)

	if receivedType != expectedType {
		t.Errorf(
			"The address type received is different from the expected. Expected: %s but Received: %s", 
			expectedType, 
			receivedType,
		)
	}
}