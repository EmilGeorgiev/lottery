mock-expected-keepers:
	mockgen -source=x/lottery/types/expected_keepers.go \
		-package testutil \
		-destination=x/lottery/testutil/expected_keepers_mocks.go