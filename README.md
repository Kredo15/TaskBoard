бэкенд для веб-приложения управления задачами (аналог Trello)

# Структура

TaskBoard/
├── cmd/
│   └── main.go # Точка входа в приложение
├── config/
│   ├── app.yaml  # Конфигурационный файл приложения
│   └── config.go # Загрузка конфигурации
├── internal/
│   ├── app/
│   │   └── app.go 
│   ├── application/              # Use Cases / Application Layer
│   │   ├── dto/
│   │   │   ├── board.go
│   │   │   ├── column.go
│   │   │   └── task.go
│   │   └── usecase/              # CQRS Commands
│   │       ├── create_board.go
│   │       ├── move_task.go
│   │       ├── get_board.go
│   │       └── ect...
│   ├── domain/       # Доменная логика (DDD)
│   │   ├── board/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── column/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   └── task/
│   │       ├── model.go
│   │       └── repository.go
│   ├── infrastructure/  # Инфраструктурный слой
│   │   ├── geteways/
│   │   │   ├── grpc/
│   │   │   │   ├── client.go          # Клиентские реализации gRPC
│   │   │   │   └── server.go          # Серверные реализации gRPC
│   │   │   └── http/
│   │   │       ├── v1
│   │   │       │   ├── handler/
│   │   │       │   │   ├── board/
│   │	│       │   │   │   ├── crud.go
│   │	│       │   │   │   ├── actions.go
│   │	│       │   │   │   └── board_handler.go
│   │   │       │   │   ├── column/
│   │	│       │   │   │   ├── crud.go
│   │	│       │   │   │   ├── actions.go
│   │	│       │   │   │   └── column_handler.go
│   │   │       │   │   └── task/
│   │	│       │   │       ├── crud.go
│   │	│       │   │       ├── actions.go
│   │	│       │   │       └── task_handler.go
│   │   │       │   ├── middleware.go
│   │   │       │   └── router.go
│   │   │       └── server.go
│   │   └── persistence/
│   │       └── postgres/
│   │           ├── task_repository.go 
│   │           ├── column_repository.go 
│   │           └── board_repository.go 
├── pkg/              # Пакеты общего назначения (общие библиотеки и утилиты)
│   ├── db/
│   │   ├── postgres/
│   │   │   └── connect.go
│   │   └── redis/
│   │       ├── cache.go       # Реализация кеширования Redis
│   │       └── pubsub.go      # PubSub реализация Redis
│   ├── logger/
│   │   └── logger.go
│   └── validator/
│       └── validate.go
├── migrations/ # SQL миграции
│   
├── proto/            # Протоколы gRPC
│   ├── board.proto
│   ├── column.proto
│   └── task.proto
├── scripts/          # Скрипты для CI/CD, миграций БД и прочего
│   ├── migrate.sh
│   └── seed_data.sql
├── tests/            # Тестовые сценарии
│   ├── unit_tests/
│   │   ├── domain/
│   │   │   ├── board_test.go
│   │   │   └── task_test.go
│   │   └── infra/
│   │       ├── postgres_test.go
│   │       └── redis_test.go
│   └── integration_tests/
│       ├── controller_test.go
│       └── grpc_client_test.go
├── Makefile          # Файл сборки make
├── Dockerfile        # Docker-образ
├── go.mod            # Go-модуль
└── README.md         # Описание проекта