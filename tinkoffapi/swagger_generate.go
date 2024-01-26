package tinkoff

//go:generate rm -rf ./models ./client
//go:generate swagger generate client -t ./ -f ./swagger.yaml --strict-responders
