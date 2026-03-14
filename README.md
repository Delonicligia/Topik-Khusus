# Redis Implementation with Golang

### redis-delonic

## Deskripsi

Project ini merupakan contoh penggunaan **Redis dengan Golang** menggunakan pendekatan **Clean Architecture**.
Aplikasi ini menunjukkan cara **menyimpan dan mengambil data user di Redis** menggunakan library Go Redis.

Redis digunakan sebagai **in-memory database** sehingga proses penyimpanan dan pengambilan data menjadi sangat cepat.

---

## Library yang Digunakan

```
github.com/redis/go-redis/v9
```

Contoh koneksi Redis:

```
rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})
```

---

## Struktur Project (Clean Architecture)

```
redis-delonic
│
├── main.go
├── domain
├── infrastructure
├── interfaces
└── tests
```

Penjelasan singkat:

* **Domain** → entity dan business logic
* **Infrastructure** → implementasi Redis repository
* **Interfaces** → CLI handler
* **Main** → dependency injection dan menjalankan aplikasi

---

## Cara Menjalankan

1. Install Golang

2. Jalankan Redis

```
docker run -d -p 6379:6379 redis
```

3. Install dependency

```
go mod tidy
```

4. Jalankan program

```
go run main.go
```

---

## Unit Testing

Menjalankan unit test:

```
go test ./... -v
```

Melihat test coverage:

```
go test ./... -cover
```

Target coverage minimal:

```
80%
```

---

## Prompt AI yang Digunakan

Prompt yang digunakan pada AI di VS Code:

```
Buatkan contoh implementasi Redis menggunakan Golang.
```

```
Restrukturisasi kode menggunakan Clean Architecture.
```

```
Buatkan unit testing dengan minimal 80% test coverage.
```
