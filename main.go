// Package main is the sqlite lesson `l2_wal_mode` homework scaffold for Vibe Learn.
//
// Задача: бенчмарк rollback journal vs WAL: 1 писатель + N читателей, p50/p99 read-латентности, ручной wal_checkpoint.
// Реализуй функции ниже — сигнатуры и тестовая поверхность фиксированы;
// CI (.github/workflows/ci.yml) гоняет `go vet` и `go test ./...`.
// Подробности и критерии приёмки — в README.md.
//
// Драйвер: modernc.org/sqlite — ЧИСТЫЙ Go, без CGO. Имя драйвера "sqlite",
// импорт blank-формой `_ "modernc.org/sqlite"`. БД встроена — сервера нет.
package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

// Note — пример доменной строки (notes/CRUD-уроки).
type Note struct {
	ID        int64
	Title     string
	Body      string
	CreatedAt string
}

// RegionRevenue — пример агрегата (join/CTE-уроки).
type RegionRevenue struct {
	Region  string
	Revenue int64
}

// Latencies — собранные перцентили для бенчмарка (WAL/concurrency-уроки).
type Latencies struct{ P50, P99 int64 }

// ----- config -----

// envOr returns the env var for `key` if set, else `fallback`.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// DatabasePath — путь к файлу БД. Дефолт ":memory:" — БД живёт в процессе.
// Для WAL-уроков переопредели на файловый путь (WAL завязан на файл).
func DatabasePath() string {
	return envOr("DATABASE_PATH", ":memory:")
}

// OpenDB открывает sqlite-БД через чистый Go-драйвер modernc и создаёт схему.
func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(Schema); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

// Schema — DDL, исполняемый идемпотентно при открытии БД (CREATE ... IF NOT EXISTS).
const Schema = `CREATE TABLE IF NOT EXISTS kv (k INTEGER PRIMARY KEY, v INTEGER NOT NULL)`

// ----- TODO #1: SetJournalMode -----
//
// PRAGMA journal_mode=<mode>; вернуть фактически установленный режим (проверь, что вернулся 'wal')
func SetJournalMode(db *sql.DB, mode string) (string, error) {
	// TODO: implement
	panic("SetJournalMode: not implemented")
}

// ----- TODO #2: BenchReadUnderWrite -----
//
// один пишущий goroutine + N читающих, собрать p50/p99 латентности чтения
func BenchReadUnderWrite(db *sql.DB, readers, iters int) (Latencies, error) {
	// TODO: implement
	panic("BenchReadUnderWrite: not implemented")
}

// ----- TODO #3: Checkpoint -----
//
// ручной PRAGMA wal_checkpoint(TRUNCATE)
func Checkpoint(db *sql.DB) error {
	// TODO: implement
	panic("Checkpoint: not implemented")
}

// _refs keeps the example domain types live while the TODO bodies are stubs.
// Удали эту переменную, когда реализуешь TODO выше.
var _refs = []any{
	Note{},
	RegionRevenue{},
	Latencies{},
}

// ----- main entry -----

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("Vibe Learn — sqlite lesson %s scaffold up", "l2_wal_mode")
	log.Printf("DATABASE_PATH: %s (driver: modernc.org/sqlite, pure-Go)", DatabasePath())
	log.Printf("Реализуй TODO-функции, затем `go test ./...`. README.md содержит задачу.")

	db, err := OpenDB(DatabasePath())
	if err != nil {
		log.Fatalf("OpenDB failed: %v", err)
	}
	defer db.Close()
	log.Printf("схема создана, БД готова")
}
