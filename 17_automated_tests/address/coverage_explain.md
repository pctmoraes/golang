to generate a cover profile run the code below in the cards folder:
`go test --coverprofile coverage.txt`

this command created this text file, which contains the tests covarage:
mode: set
tests/address/address.go:6.41,14.43 5 1
tests/address/address.go:14.43,15.38 1 1
tests/address/address.go:15.38,17.4 1 1
tests/address/address.go:20.2,20.14 1 1
tests/address/address.go:20.14,22.3 1 1
tests/address/address.go:24.2,24.23 1 1

but since it is not that simple to understand what is being covered and what not
run the code below:
`go tool cover --func=coverage.txt`

this code will print on the terminal more infos on the covarage:
tests/address/address.go:6:     AddressType     100.0%
total:                          (statements)    100.0%

but to get the maximum out of the go tools
run the code below:
`go tool cover --html=coverage.txt`

and this will open up a temporary html file containing all the functions that are being covered, the ones not being covared and what are the lines

check out the result saved in the file **covarage.html**