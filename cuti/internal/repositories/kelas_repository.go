package repositories

import "database/sql"

type KelasRepository struct {
	db *sql.DB
}

func NewKelasRepository(db *sql.DB) *KelasRepository {
	return &KelasRepository{db}
}

type KelasSimple struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}

func (r *KelasRepository) GetKelas() (result []KelasSimple, err error) {
	query := `SELECT id, nama, kode FROM kelas WHERE deleted_at IS NULL`
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}

	var kelas []KelasSimple
	for rows.Next() {
		var k KelasSimple
		if err := rows.Scan(&k.ID, &k.Nama, &k.Kode); err != nil {
			return nil, err
		}
		kelas = append(kelas, k)
	}

	return
}
