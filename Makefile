mockgen:
	mockgen -source=$(file) -destination=./testmock/$(file)

test:
	ENV=test go test ./... -v