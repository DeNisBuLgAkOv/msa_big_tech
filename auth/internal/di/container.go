package container

import (
	"database/sql"
	auth_grpc "msa_big_tech/auth/internal/delivery/grpc"
	middleware_validation "msa_big_tech/auth/internal/middleware/validation"
	user_repo "msa_big_tech/auth/internal/repository/postgress/auth"
	token_repo "msa_big_tech/auth/internal/repository/postgress/token"
	"msa_big_tech/auth/internal/usecases"

	"buf.build/go/protovalidate"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Dependencies struct {
	dig.In
	GRPCServer *grpc.Server
	AuthImpl   *auth_grpc.Implementation
}

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	// База данных
	if err := container.Provide(provideDatabase); err != nil {
		return nil, err
	}

	// Репозитории
	if err := container.Provide(user_repo.NewRepository); err != nil {
		return nil, err
	}

	if err := container.Provide(token_repo.NewRepository); err != nil {
		return nil, err
	}

	// UseCase
	if err := container.Provide(usecases.NewUseCase); err != nil {
		return nil, err
	}

	// gRPC Implementation
	if err := container.Provide(auth_grpc.NewImplementation); err != nil {
		return nil, err
	}

	// Валидатор
	if err := container.Provide(provideValidator); err != nil {
		return nil, err
	}

	// gRPC сервер
	if err := container.Provide(provideGRPCServer); err != nil {
		return nil, err
	}

	return container, nil
}

func provideDatabase() (*sql.DB, error) {
	return &sql.DB{}, nil
}

func provideValidator() (protovalidate.Validator, error) {
	return protovalidate.New()
}

func provideGRPCServer(validator protovalidate.Validator) *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(middleware_validation.ValidationInterceptor(validator)))
}
