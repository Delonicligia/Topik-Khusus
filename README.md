# Sampel Kode Redis dengan Golang - Clean Architecture

Sampel kode Redis dengan Golang menggunakan prinsip Clean Architecture untuk pemisahan lapisan yang jelas.

## Struktur Proyek

```
redis-delonic/
├── domain/
│   ├── entity/          # Domain entities (User)
│   ├── repository/      # Repository interfaces
│   └── usecase/         # Business logic
├── infrastructure/      # External concerns (Redis implementation)
├── interfaces/          # Interface adapters (CLI handler)
├── main.go              # Application entry point
├── go.mod
└── README.md
```

## Lapisan Clean Architecture

- **Domain**: Entities, Use Cases, Repository Interfaces
- **Infrastructure**: Implementasi konkrit (Redis client)
- **Interfaces**: Adapters untuk interaksi eksternal (CLI)

## Persyaratan

- Go 1.21 atau versi lebih baru
- Redis server berjalan di localhost:6379

## Instalasi

1. Pastikan Go sudah terinstall.
2. Clone atau download kode ini.
3. Jalankan `go mod tidy`.

## Menjalankan Kode

1. Jalankan Redis server.
2. Jalankan: `go run main.go`

## Output yang Diharapkan

```
User saved as string
Retrieved user: ID=1, Name=John Doe, Age=30
User saved as hash
Retrieved user from hash: ID=1, Name=John Doe, Age=30
```

## Penjelasan

- **Entity**: Representasi data domain (User)
- **Repository**: Interface untuk data access
- **Usecase**: Business logic
- **Infrastructure**: Implementasi Redis
- **Interfaces**: Handler untuk CLI

## Dependencies

- `github.com/redis/go-redis/v9`: Redis client untuk Go