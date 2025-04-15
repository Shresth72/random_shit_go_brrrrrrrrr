test_decode:
	cd pkg/decode && \
		go test . -v

run:
	go run cmd/client/main.go decode 4:spam
