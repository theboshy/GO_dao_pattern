package mysql

import (
	"../../models"
	//"fmt"
)

type RadicacionImplMysql struct {
}

func (dao RadicacionImplMysql) Create(radicaci *models.Radicacion) error {
	/*query := "INSERT INTO users (,,,) VALUES(?, ?, ?)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(radicaci.IdRemitente, radicaci.NumeroRadicacion, radicaci.FechaRadicacion)
	if err != nil {
		return err
	}*/

	/*id, err := result.LastInsertId()
	if err != nil {
		return err
	}*/

	//radicaci.NumeroRadicacion = int(id)
	return nil
}

func (dao RadicacionImplMysql) GetAll() ([]models.Radicacion, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := get()
	defer db.Close()
	var radci []models.Radicacion
	error :=db.Find(&radci)
	if error != nil {
		return radci,error
	}
		return radci,nil
}

