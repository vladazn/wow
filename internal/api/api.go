package api

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"net/http"
	"vladazn/wow/internal/domain"
	grpcserver "vladazn/wow/internal/pkg/grpc/server"
	"vladazn/wow/proto/gen/go/proto/wow"
)

// using grpc-gateway to run server based on the described proto
func RunApiServer(ctx context.Context, wowServer *grpcserver.WowServer,
	configs domain.ListenConfigs) error {

	mux := runtime.NewServeMux()

	err := wow.RegisterWowHandlerServer(
		ctx,
		mux,
		wowServer,
	)

	handler := cors.AllowAll().Handler(mux)

	fmt.Printf("serving api at :%v\n", configs.Api)

	err = http.ListenAndServe(
		fmt.Sprintf(":%v", configs.Api),
		handler,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
