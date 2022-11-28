package api

import (
	"database/sql"
	"log"
	"marketeer/config"
	"marketeer/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func LogInUser(c *gin.Context) {

	var loginPayload models.InputLogIn
	var retrieveResponse models.RetrieveResponse
	err := c.ShouldBindJSON(&loginPayload)

	cfg := mysql.Config{
		User:   config.GetEnvConfig("DBUSER"),
		Passwd: config.GetEnvConfig("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "marketeer_database1",
	}

	var dberr error
	db, dberr = sql.Open("mysql", cfg.FormatDSN())
	if dberr != nil {
		log.Fatal(err)
	}

	loginStatement := " SELECT `UserID`, `FirstName`, `LastName`, `EMail`, `ContactNum`, `BirthDate`, `Address`, `UserName`, `Password` FROM marketeer_database1.Users WHERE `UserName` = ? AND `Password` = ?; "
	row := db.QueryRow(loginStatement, loginPayload.UserName, loginPayload.Password)
	err2 := row.Scan()

	if err2 != nil && err2 != sql.ErrNoRows {
		print("\nTHERE IS THAT USER \n")

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
		rows, err := db.Query(loginStatement, loginPayload.UserName, loginPayload.Password)
		for rows.Next() {
			err := rows.Scan(&id, &firstName, &lastName, &eMail, &contactNum, &birthDate, &address, &userName, &passWord)
			//err := rows.Scan(&userName)
			if err != nil && err != sql.ErrNoRows {
				print("\nError 2\n")
				log.Fatal(err)
			}
		}

		if err == nil && err != sql.ErrNoRows {
			retrieveResponse.Status = "Success"
			retrieveResponse.Code = 200
			retrieveResponse.Message = "--USER DATA---"
			retrieveResponse.ID = id
			retrieveResponse.FirstName = firstName
			retrieveResponse.LastName = lastName
			retrieveResponse.BirthDate = birthDate
			retrieveResponse.Address = address
			retrieveResponse.EMail = eMail
			retrieveResponse.ContactNum = contactNum
			retrieveResponse.Password = passWord
			retrieveResponse.UserName = userName

			print("\nUser Log in Details \n")
			println(userName)
			println(passWord)
			print("\n")
		}

		c.JSON(http.StatusOK, retrieveResponse)

	} else {
		var genereicResponse models.GenericResponse
		print("\nTHERE IS NO SUCH USER \n")
		genereicResponse.Code = 500
		genereicResponse.Status = "Error"
		genereicResponse.Message = "No Such Credentials Is found"

		c.JSON(http.StatusOK, genereicResponse)
	}

}

/*
choco@5892
Elson R. Galaroza Jr.9:47 PM
CREATE SCHEMA marketeer_database1
CREATE TABLE marketeer_database1.Users
Elson R. Galaroza Jr.9:49 PM
CREATE TABLE marketeer_database1.Users
(
UserID int NOT NULL,
PRIMARY KEY(UserID),
FirstName varchar(255),
LastName varchar(255),
EMail varchar(255),
ContactNum char(14),
BirthDate date,
Address varchar(255),
UserName varchar (255),
Password varchar (255)
);
Elson R. Galaroza Jr.9:50 PM
select * from marketeer_database1.Users;
INSERT INTO marketeer_database1.Users
VALUES (1,
'TestingFirstName',
'TestingLastName',
'TestingEmail',
09914326973,
'2000-1-01',
'TestingAddress',
'TestingUserName',
'TestingPassword'
);
Elson R. Galaroza Jr.9:51 PM
http://localhost:/marketeer/systemcheck
Elson R. Galaroza Jr.9:54 PM
http://localhost/marketeer/registerUser
{
	"FirstName" : "Elson",
	"LastName" : "Galaroza",
	"E-Mail" : "egal@jobtarget.com",
	"ContactNumber" : "09914326860",
	"BirthDate" : "2000-3-01",
	"Address" : "Lapu - Lapu City",
	"Username" : "elson_G",
	"Password" : "e123"
}
*/
