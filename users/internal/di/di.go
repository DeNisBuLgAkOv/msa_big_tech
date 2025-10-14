package di

import (
	"database/sql"
	"fmt"
	"log/slog"
	users_grpc "msa_big_tech/users/internal/delivery/grpc"
	users_repo "msa_big_tech/users/internal/repository/postgress"
	"msa_big_tech/users/internal/usecases"
	users "msa_big_tech/users/pkg/proto/v1"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Container struct {
	DB         *sql.DB
	Repository *users_repo.Repository
	UseCase    *usecases.UseCase
	GRPCImpl   *users_grpc.Implementation
	GRPCServer *grpc.Server
	Listener   net.Listener
}

func NewContainer() *Container {
	container := &Container{}

	// Инициализация в правильном порядке
	container.initDB()
	container.initRepository()
	container.initUseCase()
	container.initGRPCImplementation()
	container.initGRPCServer()
	container.initListener()

	return container
}

func (c *Container) initDB() {
	// Пока используем заглушку
	c.DB = &sql.DB{}
	slog.Info("Database initialized")
}

func (c *Container) initRepository() {
	c.Repository = users_repo.NewRepository(c.DB)
	slog.Info("Repository initialized")
}

func (c *Container) initUseCase() {
	c.UseCase = usecases.NewUseCase(c.Repository)
	slog.Info("UseCase initialized")
}

func (c *Container) initGRPCImplementation() {
	c.GRPCImpl = users_grpc.NewImplementation(c.UseCase)
	slog.Info("GRPC Implementation initialized")
}

func (c *Container) initGRPCServer() {
	c.GRPCServer = grpc.NewServer()
	users.RegisterUsersServiceServer(c.GRPCServer, c.GRPCImpl)
	reflection.Register(c.GRPCServer)
	slog.Info("GRPC Server initialized")
}

func (c *Container) initListener() {
	var err error
	c.Listener, err = net.Listen("tcp", fmt.Sprintf(":%s", "50055"))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		panic(err)
	}
	slog.Info("Listener initialized on port 50055")
}

func (c *Container) Run() error {
	slog.Info("Starting gRPC server...")
	return c.GRPCServer.Serve(c.Listener)
}

func (c *Container) Close() {
	if c.Listener != nil {
		c.Listener.Close()
	}
	if c.DB != nil {
		c.DB.Close()
	}
	slog.Info("Container closed")
}
