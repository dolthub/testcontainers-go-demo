package customer

import (
	"context"
	"database/sql"
	"errors"
	// Import mysql into the scope of this package (required)
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (r Repository) CreateCustomer(ctx context.Context, customer *Customer) (*Customer, error) {
	res, err := r.db.ExecContext(ctx, "INSERT INTO customers (name, email) VALUES (?, ?);", customer.Name, customer.Email)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	customer.Id = int(id)
	return customer, nil
}

func (r Repository) GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	customer := Customer{}
	row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM customers WHERE email = ?;", email)
	err := row.Scan(&customer.Id, &customer.Name, &customer.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (r Repository) UpdateCustomer(ctx context.Context, customer *Customer) error {
	_, err := r.db.ExecContext(ctx, "UPDATE customers SET name = ?, email = ? WHERE id = ?;", customer.Name, customer.Email, customer.Id)
	return err
}

func (r Repository) DeleteCustomer(ctx context.Context, customer *Customer) error {
	_, err := r.db.ExecContext(ctx, "DELETE from customers WHERE id = ?;", customer.Id)
	return err
}
