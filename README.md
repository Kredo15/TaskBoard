бэкенд для веб-приложения управления задачами (аналог Trello)

# Структура

├── cmd/
│   └── main.go # Точка входа в приложение
├── config/
│   ├── app.yaml  # Конфигурационный файл приложения
│   └── config.go # Загрузка конфигурации
├── internal/
│   ├── domain/       # Доменная логика (DDD)
│   │   ├── task/
│   │   │   ├── model.go
│   │   │   ├── repository.go
│   │   │   └── storage.go
│   │   ├── board/
│   │   │   ├── model.go
│   │   │   ├── repository.go
│   │   │   └── storage.go
│   │   └── column/
│   │       ├── model.go
│   │       ├── repository.go
│   │       └── storage.go
│   ├── infrastructure/  # Инфраструктурный слой (Postgres, Redis, gRPC)
│   │   ├── db/
│   │   │   ├── postgres/
│   │   │   │   ├── repository.go  # Реализация PostgreSQL-хранилища
│   │   │   │   └── migrations.sql # SQL миграции
│   │   │   └── redis/
│   │   │       ├── cache.go       # Реализация кеширования Redis
│   │   │       └── pubsub.go      # PubSub реализация Redis
│   │   └──api/
│   │       ├── http/
│   │       │   ├── handler/
│   │       │   │   ├── board.go
│   │       │   │   ├── column.go
│   │       │   │   └── task.go
│   │       │   ├── middleware/
│   │       │   ├── server.go
│   │       │   └── router.go
│   │       └── grpc/
│   │           ├── client.go          # Клиентские реализации gRPC
│   │           └── server.go          # Серверные реализации gRPC
│   └── application/              # Use Cases / Application Layer
│       ├── command/              # CQRS Commands
│       │   ├── create_board.go
│       │   ├── move_task.go
│       │   └── handler.go
│       ├── query/                # CQRS Queries
│       │   ├── get_board.go
│       │   └── handler.go
│       └── dto/                  # Правильное место для DTO
│           ├── request/
│           └── response/
├── pkg/              # Пакеты общего назначения (общие библиотеки и утилиты)
│   ├── logger/
│   │   └── logger.go
│   └── validator/
│       └── validate.go
│   
├── proto/            # Протоколы gRPC
│   ├── board.proto
│   ├── column.proto
│   └── task.proto