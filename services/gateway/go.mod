module github.com/CutyDog/mint-flea/services/gateway

go 1.24.5

replace github.com/CutyDog/mint-flea/proto => ../../proto

require (
	github.com/99designs/gqlgen v0.17.81
	github.com/CutyDog/mint-flea/proto v0.0.0-20251021151625-a2bf86454958
	github.com/vektah/gqlparser/v2 v2.5.30
	google.golang.org/grpc v1.76.0
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251020155222-88f65dc88635 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
