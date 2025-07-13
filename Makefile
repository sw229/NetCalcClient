build:
	@go build -o bin/net_calc_client
run: build
	@./bin/net_calc_client
