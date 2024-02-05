package address

import "testing"

type testScenario struct {
	receivedAddressType string
	expectedAddressType string
}

func TestAddressType(t *testing.T) {
	testsScenarios := []testScenario {
		{ "Street of the Silly","Street" },
		{ "Avenue 123","Avenue" },
		{ "Boulevard of Glass","Boulevard" },
		{ "Highway 53","Highway" },
		{ "Saint Paul Square","Type unknown" },
	}

	for _, scenario := range testsScenarios {
		receivedAddressType := AddressType(scenario.receivedAddressType)

		if receivedAddressType != scenario.expectedAddressType {
			t.Errorf(
				"The address type received is different from the expected. Expected: %s but Received: %s", 
				scenario.receivedAddressType, 
				scenario.expectedAddressType,
			)
		}
	}
}