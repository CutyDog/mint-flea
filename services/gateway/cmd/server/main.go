package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/CutyDog/mint-flea/services/gateway/auth"
	"github.com/CutyDog/mint-flea/services/gateway/client"
	"github.com/CutyDog/mint-flea/services/gateway/graph"
)

func main() {
	// 環境変数から設定を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	accountServiceAddr := os.Getenv("ACCOUNT_SERVICE_ADDR")
	if accountServiceAddr == "" {
		accountServiceAddr = "account:9090"
	}

	// gRPCクライアントを初期化
	accountClient, err := client.NewAccountClient(accountServiceAddr)
	if err != nil {
		log.Fatalf("failed to create account client: %v", err)
	}
	defer accountClient.Close()

	// AuthMiddlewareを初期化
	authMiddleware, err := auth.NewAuthMiddleware()
	if err != nil {
		log.Fatalf("failed to create auth middleware: %v", err)
	}

	// GraphQLサーバーを初期化
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			AccountClient: accountClient,
		},
	}))

	// HTTPハンドラーをセットアップ
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authMiddleware.AuthMiddleware(srv))

	log.Printf("🚀 GraphQL server ready at http://localhost:%s/", port)
	log.Printf("🎮 Playground available at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
