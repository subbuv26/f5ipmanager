package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type repo struct {
	db *sql.DB
	dbName string
}
type Params struct {
	dbHost string
	dbPort string
	dbName string
	dbUser string
	dbPassword string
}

func NewRepo(params Params) *repo {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", params.dbUser, params.dbPassword, params.dbName, params.dbHost, params.dbPort)
	store, err := sql.Open("postgres", dbinfo)
		if err != nil {
			panic(err)
		}
	return &repo{
		db: store,
		dbName: params.dbName,
	}
}

 func (r repo) AddNewIP(ctx context.Context, ipamlabel string, ip string) error{
 	//Add ip address record to database
	 _, err := r.db.Exec(`INSERT INTO IPRecord (key, ipamlabel, ip) VALUES ("", $2, $3)`)
	 if err != nil {
		 return err
	 }
	 return nil
 }

 func (r repo) RemoveIPRange(ctx context.Context, ipamlabel string) error{
 	//Remove all records related to ipamlabel
	 _, err := r.db.Exec(`DELETE FROM IPRecord WHERE ipamlabel = $2`)
	 if err != nil {
		 return err
	 }
 	return nil
 }

