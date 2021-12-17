package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Port struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetPorts() []Port {
	ports := []Port{}
	rows, err := r.db.Queryx("SELECT * FROM ports")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var p Port
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ports = append(ports, p)
	}
	return ports
}
func (r *Repo) AddPort(item Port) {
	if item.ID != "" {
		_, err := r.db.NamedExec(`INSERT INTO ports (id,name, city)
        VALUES (:id, :name, :city)`, item)
		if err != nil {
			panic(err)
		}
		return
	}
	_, err := r.db.NamedExec(`INSERT INTO ports (name, city)
        VALUES (:name, :city)`, item)
	if err != nil {
		panic(err)
	}
	//return ports
}
func (r *Repo) UpdatePort(item Port) {
	_, err := r.db.NamedExec(`UPDATE ports SET name=:name, city=:city WHERE id =:id`, item)
	if err != nil {
		panic(err)
	}
}
func (r *Repo) Close() error {
	err := r.db.Close()
	return err
}

func Init() *Repo {
	db, err := sqlx.Open("postgres", "user=postgres password=12345 dbname=portDb sslmode=disable")
	//db, err := sqlx.Open("postgres", "user=postgres password=12345"+databaseUrl)
	if err != nil {
		return nil
	}
	repo := New(db)

	return repo
}
