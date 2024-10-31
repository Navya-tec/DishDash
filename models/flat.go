package models

import (
	"database/sql"
	"time"
)

type Flat struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Created_At time.Time `json:"created_at"`
	Created_By int64     `json:"created_by"`  //flatmate id
}

func CreateFlat(db *sql.DB,flat Flat)(int64,error){

	var flatId int64
	err:=db.QueryRow(`INSERT INTO flats (name,created_at,created_by) VALUES ($1,$2,$3) RETURNING id`,flat.Name,time.Now(),flat.Created_By).Scan(&flatId)
	if err!=nil{
		return 0,err
	}

	return flatId,nil
}




