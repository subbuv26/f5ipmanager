package postgres

import (
	"context"
	"database/sql"
	spec "f5ipmanager/internal/core/spec/manager"
	"fmt"
)

// import (
// 	ports "f5ipmanager/internal/core/ports/manager"
// )
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

func (r repo) GetIPAddress(ctx context.Context, label, key string) (spec.IPAddress, error) {
	var ip spec.IPAddress
	queryString := fmt.Sprintf("SELECT ip FROM IPRecord where key=\"%s\" AND ipamlabel=\"%s\" order by ip ASC limit 1",
		key,
		label,
	)
	err := r.db.QueryRow(queryString).Scan(&ip)
	if err != nil {
		return "", err
	}
	return ip, nil
}

func (r repo) AllocateIPAddress(ctx context.Context, ipamlabel, key string) (spec.IPAddress, error) {
	var ip spec.IPAddress
	queryString := fmt.Sprintf(
		"SELECT ip FROM IPRecord WHERE key is null AND ipam_label=\"%s\" order by ip ASC limit 1",
		ipamlabel,
	)
	err := r.db.QueryRow(queryString).Scan(&ip)
	if err != nil {
		return " ", spec.ErrorResourcesExhausted
	}
	allocateIPSql := fmt.Sprintf("UPDATE IPRecord set key=\"%s\" WHERE ip=?",
		key,
	)
	_, err = r.db.Exec(allocateIPSql, ip)
	if err != nil {
		return " ", err
	}
	return ip, nil

}
func (r repo) FreeIPAddress(ctx context.Context, label, key string) error {
	deallocateIPSql := fmt.Sprintf("UPDATE IPRecord set key=\"\" where ipamlabel=?")
	_, err := r.db.Exec(deallocateIPSql, label)
	if err != nil {
		return err
	}
	return nil
}
