package main

import (
	"log"
	"sbp/internal/app"
	"sbp/internal/dal"
	"sbp/internal/firebase_authorization"
	"sbp/internal/transport/pay"
	"sbp/internal/transport/rest"
	"sbp/pkg/bootstrap"
)

//go:generate rm -rf ./openapi/restapi ./openapi/model ./openapi/client
//go:generate swagger generate server -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal sbp/internal/app.Auth --exclude-main
//go:generate swagger generate client -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal sbp/internal/app.Auth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("new config: ", err)
	}

	l, err := bootstrap.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("new logger: ", err)
	}

	dbConn, err := bootstrap.NewDbConn(cfg.DB)
	if err != nil {
		l.Fatalln("new db conn: ", err)
	}
	defer dbConn.Close()

	l.Debug("connected to db")

	err = bootstrap.UpMigrations(dbConn.DB, cfg.DB.Database, "migrations")
	if err != nil {
		l.Fatalln("up migrations: ", err)
	}

	l.Debug("applied migrations")

	// rabbitSvc, err := rabbit.New(
	// 	l,
	// 	cfg.RabbitMQConfig.Url,
	// 	cfg.RabbitMQConfig.Port,
	// 	cfg.RabbitMQConfig.User,
	// 	cfg.RabbitMQConfig.Password,
	// )
	// if err != nil {
	// 	l.Fatalln("new rabbit conn: ", err)
	// }
	// l.Debug("connected to rabbit")

	authSvc := firebase_authorization.New(cfg.FirebaseConfig.FirebaseKeyFilePath)

	repo := dal.New(dbConn, l)

	svcPay, err := pay.New(l)

	// washSvc := app.NewWashServerService(l, repo, rabbitSvc, svcPay)

	washSvc := app.NewWashServerService(l, repo, nil, svcPay)

	server, err := rest.NewServer(cfg, authSvc, l, washSvc)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	err = server.Serve()
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}
