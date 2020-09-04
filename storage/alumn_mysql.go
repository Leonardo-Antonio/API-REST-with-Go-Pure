package storage

import (
	"database/sql"

	"github.com/Leonardo-Antonio/api-rest-mysql/storage/validations"

	"github.com/Leonardo-Antonio/api-rest-mysql/model"
)

const (
	sqlCreate  = "INSERT INTO tb_alumnos VALUES(null, ?, ?, ? ,?)"
	sqlGetAll  = "SELECT id, name, lastname, age, dni FROM tb_alumnos"
	sqlGetByID = "SELECT id, name, lastname, age, dni FROM tb_alumnos WHERE id = ?"
	sqlDelete  = "DELETE FROM tb_alumnos WHERE id = ?"
	sqlUpdate  = "UPDATE tb_alumnos SET name = ?, lastname = ?, age = ?, dni = ? WHERE id = ?"
)

// Alumn . (CRUD)
type Alumn struct {
	db *sql.DB
}

// NewAlumn method contructor of the class Alumno
func NewAlumn(db *sql.DB) *Alumn {
	return &Alumn{db}
}

// Create .
func (a *Alumn) Create(alumn model.Alumn) error {
	stmt, err := a.db.Prepare(sqlCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		validations.StringNull(alumn.Name),
		validations.StringNull(alumn.LastName),
		validations.IntNull(int32(alumn.Age)),
		alumn.Dni,
	)

	if err != nil {
		return err
	}

	if rA, _ := rs.RowsAffected(); rA != 1 {
		return validations.ErrorRowsAffected
	}

	return nil
}

// GetAll .
func (a *Alumn) GetAll() (alumns []model.Alumn, err error) {
	stmt, err := a.db.Prepare(sqlGetAll)
	if err != nil {
		return
	}
	defer stmt.Close()

	// valid null
	nullName := sql.NullString{}
	nullLastName := sql.NullString{}
	nullAge := sql.NullInt32{}

	rows, err := stmt.Query()
	for rows.Next() {
		alumn := model.Alumn{}
		err := rows.Scan(
			&alumn.ID,
			&nullName,
			&nullLastName,
			&nullAge,
			&alumn.Dni,
		)
		if err != nil {
			return nil, err
		}
		alumn.Name = nullName.String
		alumn.LastName = nullLastName.String
		alumn.Age = uint8(nullAge.Int32)

		alumns = append(alumns, alumn)
	}

	return alumns, nil
}

// Delete .
func (a *Alumn) Delete(id int) error {
	stmt, err := a.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if rA, _ := rs.RowsAffected(); rA != 1 {
		return validations.ErrorRowsAffected
	}
	return nil
}

// GetByID .
func (a *Alumn) GetByID(id int) (alumn model.Alumn, err error) {

	stmt, err := a.db.Prepare(sqlGetByID)
	if err != nil {
		return
	}
	defer stmt.Close()

	// valid null
	nullName := sql.NullString{}
	nullLastName := sql.NullString{}
	nullAge := sql.NullInt32{}

	err = stmt.QueryRow(id).Scan(
		&alumn.ID,
		&nullName,
		&nullLastName,
		&nullAge,
		&alumn.Dni,
	)

	alumn.Name = nullName.String
	alumn.LastName = nullLastName.String
	alumn.Age = uint8(nullAge.Int32)

	return alumn, nil
}

// Update .
func (a *Alumn) Update(id int, alumn model.Alumn) error {
	stmt, err := a.db.Prepare(sqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		validations.StringNull(alumn.Name),
		validations.StringNull(alumn.LastName),
		validations.IntNull(int32(alumn.Age)),
		alumn.Dni,
		id,
	)
	if err != nil {
		return err
	}
	if rA, _ := rs.RowsAffected(); rA != 1 {
		return validations.ErrorRowsAffected
	}
	return nil
}
