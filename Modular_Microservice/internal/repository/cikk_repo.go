package repository

import (
	"database/sql"
	"fmt"

	"modular_microservice/internal/model"
)

type CikkRepository struct {
	db *sql.DB
}

func NewCikkRepository(db *sql.DB) *CikkRepository {
	return &CikkRepository{db: db}
}

func (r *CikkRepository) GetAllCikk() ([]model.Cikk, error) {

	query := `
	SELECT id, nev, cikkszam, cikk_tipus_id
	FROM atlas.cikk
	LIMIT 10
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var items []model.Cikk

	for rows.Next() {
		var item model.Cikk
		err := rows.Scan(&item.ID, &item.Nev, &item.Cikkszam, &item.CikkTipusId)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rws error: %w", err)
	}

	return items, nil
}
