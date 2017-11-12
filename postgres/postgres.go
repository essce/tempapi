package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/essce/tempapi"

	// Blank import of pq
	_ "github.com/lib/pq"
)

// Postgres type for database.
type Postgres struct {
	db *sql.DB
}

// Options contains the information to open a Postgres connection.
type Options struct {
	User     string
	Password string
	Name     string
}

// New returns an instance of a Postgres database.
func New() Postgres {
	opt := Options{
		User:     "postgres",
		Password: "",
		Name:     "postgres",
	}
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		opt.User, opt.Password, opt.Name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(fmt.Sprintf("unexpected error occured tryign to start a postgres connection: %s", err.Error()))
	}
	return Postgres{
		db: db,
	}
}

// InsertReading inserts the reading.
func (p *Postgres) InsertReading(ctx context.Context, temp, humidity float64) (string, error) {
	var id string
	err := p.db.QueryRowContext(ctx, `
		INSERT INTO reading
		(temperature, humidity, added_at)
		VALUES ($1, $2, now())
		RETURNING id;
	`).Scan(&id)

	return id, err
}

// ListReadings retrieves all readings from the database.
func (p *Postgres) ListReadings(ctx context.Context) ([]tempapi.Reading, error) {
	rows, err := p.db.QueryContext(ctx, `
		SELECT temperature, humidity, added_at
		FROM reading;
	`)
	if err != nil {
		return nil, err
	}

	var readings []tempapi.Reading
	for rows.Next() {
		var reading tempapi.Reading
		if err := rows.Scan(&reading.Temperature, &reading.Humidity, &reading.AddedAt); err != nil {
			return nil, err
		}
		readings = append(readings, reading)
	}
	return readings, rows.Err()
}
