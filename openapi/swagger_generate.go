package openapi

//go:generate rm -rf ./restapi ./models ./client
//go:generate swagger generate server -t ./ -f ./swagger.yaml --strict-responders --strict-additional-properties --principal sbp/internal/entities.Auth --exclude-main
//go:generate swagger generate client -t ./ -f ./swagger.yaml --strict-responders --strict-additional-properties --principal sbp/internal/entities.Auth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
//go:generate swagger generate server --target ./ --name WashSbp --spec ./swagger.yaml --principal sbp/internal/entities.Auth --exclude-main --strict-responders
