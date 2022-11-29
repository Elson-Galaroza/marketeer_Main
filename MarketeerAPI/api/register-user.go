package api

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"log"
	"marketeer/config"
	"marketeer/models"
	"net/http"

	//"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)

var db *sql.DB

var registrationPayload models.RegistrationInput
var genereicResponse models.GenericResponse
var retrieveResponse models.RetrieveResponse

func RegisterUser(c *gin.Context) {

	errJson := c.ShouldBindJSON(&registrationPayload)

	if errJson == nil {
		genereicResponse.Status = "Successfull Input"
		genereicResponse.Code = 200
		genereicResponse.Message = "Hi, " + registrationPayload.FirstName + ". Your User Details are Below: "
		retrieveResponse.Code = 200
		retrieveResponse.Status = "Succesfull Retrieve"
		retrieveResponse.FirstName = registrationPayload.FirstName
		retrieveResponse.LastName = registrationPayload.LastName
		retrieveResponse.EMail = registrationPayload.EMail
		retrieveResponse.ContactNum = registrationPayload.ContactNum
		retrieveResponse.BirthDate = registrationPayload.BirthDate
		retrieveResponse.Address = registrationPayload.Address
		retrieveResponse.UserName = registrationPayload.UserName
		retrieveResponse.Password = registrationPayload.Password
	} else {
		genereicResponse.Status = "Failed"
		genereicResponse.Code = 500
		genereicResponse.Message = errJson.Error()

	}

	c.JSON(http.StatusOK, genereicResponse)
	c.JSON(http.StatusOK, retrieveResponse)

	cfg := mysql.Config{
		User:   config.GetEnvConfig("DBUSER"),
		Passwd: config.GetEnvConfig("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "marketeer_database1",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	FirstName := retrieveResponse.FirstName
	LastName := retrieveResponse.LastName
	EMail := retrieveResponse.EMail
	ContactNum := retrieveResponse.ContactNum
	BirthDate := retrieveResponse.BirthDate
	Address := retrieveResponse.Address
	UserName := retrieveResponse.UserName
	Password := retrieveResponse.Password

	insertStatement := " INSERT INTO `marketeer_database1`.`Users` (  `FirstName`, `LastName`, `EMail`, `ContactNum`, `BirthDate`, `Address`, `UserName`, `Password`) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?); "
	insert, err := db.Query(insertStatement, FirstName, LastName, EMail, ContactNum, BirthDate, Address, UserName, Password)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Print("Successfull Connection")

	var (
		id         int
		firstName  string
		lastName   string
		eMail      string
		contactNum string
		birthDate  string
		address    string
		userName   string
		passWord   string
	)
	getUserIDStatement := "SELECT * FROM marketeer_database1.Users ORDER BY UserID DESC LIMIT 1"

	row := db.QueryRow(getUserIDStatement)
	err2 := row.Scan()

	if err2 != nil && err2 != sql.ErrNoRows {
		rows, err := db.Query(getUserIDStatement)
		for rows.Next() {
			err := rows.Scan(&id, &firstName, &lastName, &eMail, &contactNum, &birthDate, &address, &userName, &passWord)
			if err != nil && err != sql.ErrNoRows {
				print("\nError\n")
				log.Fatal(err)
			}
		}
		if err == nil && err != sql.ErrNoRows {
			println("\n", id)
		}
	}

	idnum := id

	retrieveResponse.ID = idnum + 1
}

func getUserID() {

	var (
		id         int
		firstName  string
		lastName   string
		eMail      string
		contactNum string
		birthDate  string
		address    string
		userName   string
		passWord   string
	)
	getUserIDStatement := "SELECT * FROM marketeer_database1.Users ORDER BY UserID DESC LIMIT 1"

	row := db.QueryRow(getUserIDStatement)
	err2 := row.Scan()

	if err2 != nil && err2 != sql.ErrNoRows {
		rows, err := db.Query(getUserIDStatement)
		for rows.Next() {
			err := rows.Scan(&id, &firstName, &lastName, &eMail, &contactNum, &birthDate, &address, &userName, &passWord)
			if err != nil && err != sql.ErrNoRows {
				print("\nError\n")
				log.Fatal(err)
			}
		}
		if err == nil && err != sql.ErrNoRows {
			println("\n", id)
		}
	}

}

// sample input
// {
// 	"FirstName" : "John",
// 	"LastName" : "Dalisay",
// 	"E-Mail" : "cards@jobtarget.com",
// 	"ContactNumber" : "099912345",
// 	"BirthDate" : "2000-1-01",
// 	"Address" : "Cebu City",
// 	"Username" : "cards_D",
// 	"Password" : "1234"
// }
