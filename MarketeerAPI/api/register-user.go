package api

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"log"
	"marketeer/config"
	"marketeer/models"
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)

var db *sql.DB

func RegisterUser(c *gin.Context) {

	var registrationPayload models.RegistrationInput
	var genereicResponse models.GenericResponse
	var retrieveResponse models.RetrieveResponse

	errJson := c.ShouldBindJSON(&registrationPayload)

	// Generate random UserID
	UserID := rand.Intn(10)
	if errJson == nil {
		genereicResponse.Status = "Successfull Input"
		genereicResponse.Code = 200
		genereicResponse.Message = "Hi, " + registrationPayload.FirstName + ". Your User Details are Below: "

		retrieveResponse.Code = 200
		retrieveResponse.Status = "Succesfull Retrieve"
		retrieveResponse.ID = UserID
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

	insertStatement := " INSERT INTO `marketeer_database1`.`Users` (`UserID`, `FirstName`, `LastName`, `EMail`, `ContactNum`, `BirthDate`, `Address`, `UserName`, `Password`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?); "
	insert, err := db.Query(insertStatement, UserID, FirstName, LastName, EMail, ContactNum, BirthDate, Address, UserName, Password)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Print("Successfull Connection")

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
