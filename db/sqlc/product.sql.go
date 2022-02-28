// Code generated by sqlc. DO NOT EDIT.
// source: product.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product (
	account_id,
  title,
  content,
	product_tag
) VALUES (
  $1, $2, $3,$4
) RETURNING id, account_id, title, content, product_tag, created_at, last_update
`

type CreateProductParams struct {
	AccountID  int64    `json:"account_id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	ProductTag []string `json:"product_tag"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.AccountID,
		arg.Title,
		arg.Content,
		pq.Array(arg.ProductTag),
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Title,
		&i.Content,
		pq.Array(&i.ProductTag),
		&i.CreatedAt,
		&i.LastUpdate,
	)
	return i, err
}

const deleteAccountProduct = `-- name: DeleteAccountProduct :exec
DELETE FROM product
WHERE id = $1
`

func (q *Queries) DeleteAccountProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, account_id, title, content, product_tag, created_at, last_update FROM product
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Title,
		&i.Content,
		pq.Array(&i.ProductTag),
		&i.CreatedAt,
		&i.LastUpdate,
	)
	return i, err
}

const listMyProduct = `-- name: ListMyProduct :many
SELECT id, account_id, title, content, product_tag, created_at, last_update FROM product
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListMyProductParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListMyProduct(ctx context.Context, arg ListMyProductParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listMyProduct, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Title,
			&i.Content,
			pq.Array(&i.ProductTag),
			&i.CreatedAt,
			&i.LastUpdate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductDetail = `-- name: UpdateProductDetail :one
UPDATE product
SET title = $2, content = $3, product_tag = $4, last_update = now()
WHERE id = $1
RETURNING id, account_id, title, content, product_tag, created_at, last_update
`

type UpdateProductDetailParams struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	ProductTag []string `json:"product_tag"`
}

func (q *Queries) UpdateProductDetail(ctx context.Context, arg UpdateProductDetailParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProductDetail,
		arg.ID,
		arg.Title,
		arg.Content,
		pq.Array(arg.ProductTag),
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Title,
		&i.Content,
		pq.Array(&i.ProductTag),
		&i.CreatedAt,
		&i.LastUpdate,
	)
	return i, err
}
