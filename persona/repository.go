package persona

import (
	"database/sql"
	"time"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	GetPersonByID(personaID string) (*Person, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository crear el nuevo repositorio y retorna con la BD conectada*/
func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}

func (repo *repository) GetPersonByID(personaID string) (*Person, error) {
	const sql = `SELECT * FROM PERSONA_01 WHERE PERSONA_ID = :1 AND ESTADO <> 0`
	row := repo.db.QueryRow(sql, personaID)
	var fechaNac time.Time
	persona := &Person{}
	err := row.Scan(
		&persona.ID,
		&persona.Nombre,
		&persona.ApellidoPaterno,
		&persona.ApellidoMaterno,
		&persona.Genero,
		&persona.Dni,
		&fechaNac,
		&persona.Estado,
	)
	// year, month, day := fecha_nac.Date()
	// fmt.Printf("Date : [%d]year : [%d]month : [%d]day \n", year, month, day)
	persona.FechaNacimiento = fechaNac.Format("02/01/2006")
	return persona, err
}
