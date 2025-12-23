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
│   │   │   └── repository.go
│   │   ├── board/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   └── column/
│   │       ├── model.go
│   │       └── repository.go
│   ├── infrastructure/  # Инфраструктурный слой (Postgres, Redis, gRPC)
│   │   ├── persistence/
│   │   │   ├── postgres/
│   │   │   │   ├── task_repository.go 
│   │   │   │   ├── column_repository.go 
│   │   │   │   └── board_repository.go 
│   │   │   └── redis/
│   │   │       ├── cache.go       # Реализация кеширования Redis
│   │   │       └── pubsub.go      # PubSub реализация Redis
│   │   ├── db/
│   │   │   ├── postgres/
│   │   │   │   ├── connect.go
│   │   │   │   └── migrations.sql # SQL миграции
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
│       ├── usecase/              # CQRS Commands
│       │   ├── create_board.go
│       │   ├── move_task.go
│       │   ├── get_board.go
│       │   └── handler.go
│       └── dto/
│           ├── board.go
│           ├── column.go
│           └── task.go
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