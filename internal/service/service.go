package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/jsiebens/ionscale/internal/broker"
	"github.com/jsiebens/ionscale/internal/config"
	"github.com/jsiebens/ionscale/internal/domain"
	"github.com/jsiebens/ionscale/internal/provider"
	"github.com/jsiebens/ionscale/internal/version"
	api "github.com/jsiebens/ionscale/pkg/gen/ionscale/v1"
)

func NewService(config *config.Config, authProvider provider.AuthProvider, repository domain.Repository, pubsub broker.Pubsub) *Service {
	return &Service{
		config:       config,
		authProvider: authProvider,
		repository:   repository,
		pubsub:       pubsub,
	}
}

type Service struct {
	config       *config.Config
	authProvider provider.AuthProvider
	repository   domain.Repository
	pubsub       broker.Pubsub
}

func (s *Service) GetVersion(_ context.Context, _ *connect.Request[api.GetVersionRequest]) (*connect.Response[api.GetVersionResponse], error) {
	v, revision := version.GetReleaseInfo()
	return connect.NewResponse(&api.GetVersionResponse{
		Version:  v,
		Revision: revision,
	}), nil
}
