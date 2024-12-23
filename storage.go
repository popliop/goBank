package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
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
	//postgres://postgres:gobank@localhost:5432/postgres
	//connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

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
func (s *PostgreStorage) CreateAccount(acc *Account) error {

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
	query := `DELETE FROM account where id = $1`

	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreStorage) GetAccountByID(id int) (*Account, error) {
	query := `SELECT * FROM account where id = $1`
	fmt.Println("id", id)
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanRowAccount(rows)
	}

	return nil, fmt.Errorf("Account %d not found", id)
}

func (s *PostgreStorage) GetAccounts() ([]*Account, error) {

	query := `SELECT * FROM account ORDER BY id`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanRowAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanRowAccount(rows *sql.Rows) (*Account, error) {

	account := Account{}
	err := rows.Scan(
		&account.ID,
		&account.Firstname,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.Createdtime)

	return &account, err
}
