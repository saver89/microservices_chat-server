package app

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/saver89/microservices_chat-server/internal/api/chat"
	"github.com/saver89/microservices_chat-server/internal/client/db"
	"github.com/saver89/microservices_chat-server/internal/client/db/pg"
	"github.com/saver89/microservices_chat-server/internal/client/db/transaction"
	"github.com/saver89/microservices_chat-server/internal/closer"
	"github.com/saver89/microservices_chat-server/internal/config"
	"github.com/saver89/microservices_chat-server/internal/repository"
	chatRepository "github.com/saver89/microservices_chat-server/internal/repository/chat"
	chatLogRepository "github.com/saver89/microservices_chat-server/internal/repository/chat_log"
	chatUserRepository "github.com/saver89/microservices_chat-server/internal/repository/chat_user"
	"github.com/saver89/microservices_chat-server/internal/service"
	chatService "github.com/saver89/microservices_chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	chatService        service.ChatService
	chatRepository     repository.ChatRepository
	chatUserRepository repository.ChatUserRepository
	chatLogRepository  repository.ChatLogRepository
	chatImplementation *chat.Implementation

	log *slog.Logger
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN(), s.Log())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatService(
			s.ChatRepository(ctx),
			s.ChatUserRepository(ctx),
			s.ChatLogRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewChatRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatUserRepository(ctx context.Context) repository.ChatUserRepository {
	if s.chatUserRepository == nil {
		s.chatUserRepository = chatUserRepository.NewChatUserRepository(s.DBClient(ctx))
	}

	return s.chatUserRepository
}

func (s *serviceProvider) ChatLogRepository(ctx context.Context) repository.ChatLogRepository {
	if s.chatLogRepository == nil {
		s.chatLogRepository = chatLogRepository.NewChatLogRepository(s.DBClient(ctx))
	}

	return s.chatLogRepository
}

func (s *serviceProvider) ChatImplementation(ctx context.Context) *chat.Implementation {
	if s.chatImplementation == nil {
		s.chatImplementation = chat.NewImplementation(s.Log(), s.ChatService(ctx))
	}

	return s.chatImplementation
}

func (s *serviceProvider) Log() *slog.Logger {
	if s.log == nil {
		s.log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return s.log
}
