package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
func main(){
	db, err := sql.Open("postgres",
	"user=postgres password=7zDXiDUkYct8nZV dbname=postgres sslmode=disable host=database-1.cip6dyjtjole.us-east-1.rds.amazonaws.com port=5432" )
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	rowsId, rowsIdErr := db.Query("SELECT id FROM fmembers")
	//var member DynamoMember
	var ids []int
	if rowsIdErr != nil {
		log.Fatal(rowsIdErr)
	}
	defer rowsId.Close()

	for rowsId.Next(){

		var id int
		rowsId.Scan(&id)
		ids = append(ids, id)


	}
	rowsRes, rowsResErr := db.Query("SELECT results FROM fmembers")
	var results []Result
	if rowsResErr != nil {
		log.Fatal(rowsResErr)
	}
	defer rowsRes.Close()
	
	for rowsRes.Next(){
		var resJson []byte
		var result []Result
		rowsRes.Scan(&resJson)
		if err:= json.Unmarshal(resJson, &result); err != nil {
			log.Fatal(err)
		}
		for _, r := range result{
			results = append(results, r)
		}

		
		
		

	}
	
	
	var members []DynamoMember
	for i := 0; i < len(ids); i++ {
		members = append(members, DynamoMember{
			Id: ids[i],
			Result: []Result{results[i]},
		})
	}
	log.Printf("Id lenght %v results length %s", fmt.Sprint(len(ids)),fmt.Sprint(len(results)))
	//log.Print(members)
}
type DynamoMember struct {
	Id int `json:"id"`
	Result []Result `json:"results"`

}
type Result struct{
	Link string `json:"link"`
	Headline string `json:"headline"`
}