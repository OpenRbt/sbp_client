package rabbitmqapi

//go:generate rm -rf ./models ./client
//go:generate swagger generate client -t ./ -f swagger.yaml --strict-responders --strict-additional-properties
