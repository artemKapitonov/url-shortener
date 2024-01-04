package grpc_api

import (
	"context"

	url_shortner_v1 "github.com/artemKapitonov/url-shortner/pkg/url-shortner_v1"
)

func (api *GrpcServerApi) Get(context.Context, *url_shortner_v1.ShortURL) (*url_shortner_v1.FullURL, error) {
	return &url_shortner_v1.FullURL{Url: "eeeee"}, nil
}

func (api *GrpcServerApi) Create(context.Context, *url_shortner_v1.FullURL) (*url_shortner_v1.ShortURL, error) {
	panic("implement me")
}
