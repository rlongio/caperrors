test:
	go test -v -failfast -coverprofile=prof.out ./...
	go tool cover -html=prof.out -o ./testdata/coverage.html
	rm prof.out

manual_test:
	go run ./cmd/caperrors/main.go -path=./testdata/error_files -log_path=./testdata/error_files/CapHandler.log

manual_test_diagnostic:
	go run ./cmd/caperrors/main.go -path=./testdata/error_files -log_path=./testdata/error_files/CapHandler.log -diagnostic