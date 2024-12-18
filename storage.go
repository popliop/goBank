package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAcount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgreStorage struct {
	db *sql.DB
}

var (
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("SERVER_HOST")
	port     = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_NAME")
	sslmode  = os.Getenv("DB_SSLMODE")
)

func NewPostgressStorage() (*PostgreStorage, error) {

	//connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	//postgres://postgres:gobank@localhost:5432/postgres

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStorage{
		db: db,
	}, nil

}

func (s *PostgreStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgreStorage) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account(
			id serial PRIMARY KEY,
			first_name VARCHAR(50),
			last_name VARCHAR(50),
			number SERIAL,
			balance SERIAL,
			created_at timestamp DEFAULT now()
	)`

	_, err := s.db.Exec(query)
	return err
}

// Handlers
func (s *PostgreStorage) CreateAcount(acc *Account) error {

	query := `INSERT INTO account 
	(first_name, last_name, number, balance, created_at)
	VALUES
	($1,$2,$3,$4,$5)`

	_, err := s.db.Query(query,
		acc.Firstname,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.Createdtime)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreStorage) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgreStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgreStorage) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgreStorage) GetAccount(id int) ([]*Account, error) {

	query := `SELECT * FROM acounts`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.Firstname,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.Createdtime)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}
