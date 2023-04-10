package server

import (
	"log"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/http/dependency"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/http/tracing"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db/drive/sqlite"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func ServerRest() {

	opt := tracing.NewZipkinOpenTel()
	opt.ExporterEndpoint = "http://localhost:9411/api/v2/spans"
	opt.ServiceName = "ServerHTTP"

	tp := opt.Zipkin()

	otel.SetTracerProvider(tp)

	mux := http.NewServeMux()

	db, err := sqlite.ConnectionSqlite()

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	dependency.Dependency(db, mux)

	handler := otelhttp.NewHandler(mux, "server",
		otelhttp.WithTracerProvider(tp),
	)

	log.Printf("Starting server on port %s...\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
