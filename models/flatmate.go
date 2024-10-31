package models

import (
	"database/sql"
	"log"
	"time"
)

type Flatmate struct {
	Id         int64     `json:"id"`
	FlatId     *int64     `json:"flat_id"`
	Name       string    `json:"name"`
	Created_At time.Time `json:"created_at"`
}

func CheckIfUserJoined(db *sql.DB, userId int64)(bool,error){
	log.Printf("Checking if user with ID %d is joined", userId)
	var count int
	err:=db.QueryRow("SELECT COUNT(*) FROM flatmates WHERE id=$1",userId).Scan(&count)
	if err!=nil{
		return false,err
	}

	return count>0,nil
}

func AddFlatmateToFlat(db *sql.DB, flatmate Flatmate) error{

	_,err:=db.Exec("INSERT INTO flatmates (flat_id,name,created_at) VALUES ($1,$2,$3)",flatmate.FlatId,flatmate.Name,time.Now())
	if err!=nil{
      return err
	}

	return nil
}
