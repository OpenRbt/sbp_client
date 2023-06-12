package pay

//go:generate rm -rf ./model ./client
//go:generate swagger generate client -t ./ -f ./swagger.yaml --strict-responders --strict-additional-properties
