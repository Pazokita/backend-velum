package maps

import (
	"database/sql"
	"log"
)

type Repository interface {
	GetAllMaps() ([]MapMetadata, error)
	GetMapByID(id int) (*MapMetadata, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllMaps() ([]MapMetadata, error) {
	rows, err := r.db.Query("SELECT id, title, imageUrl, bbox, opacity FROM maps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var maps []MapMetadata
	for rows.Next() {
		var m MapMetadata
		if err := rows.Scan(&m.ID, &m.Title, &m.ImageURL, &m.BBox, &m.Opacity); err != nil {
			log.Println("Erreur de scan :", err)
			continue
		}
		maps = append(maps, m)
	}

	return maps, nil
}

func (r *repository) GetMapByID(id int) (*MapMetadata, error) {
	row := r.db.QueryRow("SELECT id, title, imageUrl, bbox, opacity FROM maps WHERE id = ?", id)

	var m MapMetadata
	err := row.Scan(&m.ID, &m.Title, &m.ImageURL, &m.BBox, &m.Opacity)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
