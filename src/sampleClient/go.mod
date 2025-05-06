module sampleClient

go 1.23.5

replace github.com/tashinoki/logistic_app_tentative/proto => ../proto

require (
	github.com/bufbuild/connect-go v1.10.0
	github.com/go-chi/chi v1.5.5
	github.com/tashinoki/logistic_app_tentative/proto v0.0.0-20231009120523-2f1a3b4c5d7e
	google.golang.org/protobuf v1.36.6
)
