package database

import (
	"fmt"
	cnfg "github.com/daria/PortMicroservice/data/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
)

type Port struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	IsActive   bool    `json:"isActive" db:"is_active"`
	Company    string  `json:"company"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address" db:"address"`
	About      string  `json:"about"`
	Registered string  `json:"registered"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
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
	if strconv.Itoa(item.ID) != "" {
		_, err := r.db.NamedExec(`INSERT INTO ports (id,name, is_active,company,email,phone,address,about,registered,latitude,longitude)
        VALUES (:id, :name, :is_active, :company, :email, :phone, :address, :about, :registered, :latitude, :longitude)`, item)
		if err != nil {
			panic(err)
		}
		return
	}
	_, err := r.db.NamedExec(`INSERT INTO ports (name, is_active,company,email,phone, address,about,registered,latitude,longitude)
        VALUES (:name, :is_active, :company, :email, :phone, :address :about, :registered, :latitude, :longitude)`, item)
	if err != nil {
		panic(err)
	}
}
func (r *Repo) UpdatePort(item Port) {
	_, err := r.db.NamedExec(`UPDATE ports SET name=:name, is_active=:is_active, company= :company, email=:email, phone= :phone, address= :address, about= :about, registered=:registered, latitude=:latitude, longitude=:longitude WHERE id =:id`, item)
	if err != nil {
		panic(err)
	}
}
func (r *Repo) Close() error {
	err := r.db.Close()
	return err
}

func Init(config *cnfg.Config) *Repo {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host =%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName))
	if err != nil {
		return nil
	}
	repo := New(db)

	return repo
}
