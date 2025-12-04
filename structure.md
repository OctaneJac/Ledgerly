ledger/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── ledger/
│   │   ├── account.go
│   │   ├── transaction.go
│   │   └── ledger.go
│   │
│   ├── store/
│   │   ├── postgres.go
│   │   └── models.go
│   │
│   ├── api/
│   │   ├── router.go
│   │   ├── handlers.go
│   │   └── middleware.go
│   │
│   ├── service/
│   │   └── ledger_service.go
│   │
│   ├── jobs/
│   │   └── reconciliation.go
│   │
│   └── config/
│       └── config.go
│
├── scripts/
│   └── run_local.sh
│
├── go.mod
└── README.md

cmd/api
You will run this as your API/HTTP server.

File:

main.go — wires config → DB → router → runs the HTTP server.

cmd/worker
Will run background jobs (batch jobs, reconciliation, asynchronous tasks).

File:

main.go — initializes a job processor loop.

internal/ledger/
This is your domain layer — pure business logic.

Files:

account.go — account entity, validation.

transaction.go — transaction entity (id, debit/credit, metadata).

ledger.go — core methods (CreateTransaction, GetBalance, ValidateEntry, etc.).

This is the heart of your ledger.

internal/store/
Database layer.

Files:

postgres.go — DB connection + methods to store/retrieve transactions.

models.go — DB models (TransactionRow, AccountRow, etc.)

Your service layer will call these.

internal/api/
HTTP interface layer (REST or gRPC).

Files:

router.go — setup mux/chi/fiber routing.

handlers.go — endpoints: create transaction, list ledger, get balance.

middleware.go — logging, auth, recovery.

internal/service/
Application logic (orchestration between domain and store).

File:

ledger_service.go

Takes requests from API handlers

Uses domain logic (internal/ledger)

Uses DB repo (internal/store)

Applies business rules (double-entry enforcement, validation, idempotency)

internal/jobs/
For background tasks.

File:

reconciliation.go

Recompute balances, detect inconsistencies, batch processes.

internal/config/
Configuration loader.

File:

config.go

Loads env vars: DB_URL, PORT, REDIS_URL, ENV.

migrations/
SQL migrations for Postgres.

File:

0001_init.sql

accounts, transactions, entries tables.

scripts/
Helper scripts for dev.

File:

run_local.sh

start DB, run migrations, run API.

