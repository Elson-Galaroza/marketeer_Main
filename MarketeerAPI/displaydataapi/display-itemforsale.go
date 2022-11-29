package displaydataapi

import (
	"database/sql"
	"fmt"
	"log"
	"marketeer/config"
	"marketeer/itemmodels"
	"marketeer/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func RetrieveForSaleItems(c *gin.Context) {

	var retrieveAllItemsForSale itemmodels.ItemRetrieveResponse
	var genereicResponse models.GenericResponse

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

	var (
		itemID       int
		sellerID     int
		itemImageURL string
		itemName     string
		itemPrice    int
		itemQty      int
		itemDesc     string
	)

	if err == nil {
		// retrieveAllItemsForSale.SellerID = sellerID
		// retrieveAllItemsForSale.ItemImageLinkURL = itemImageURL
		// retrieveAllItemsForSale.ItemID = itemID
		// retrieveAllItemsForSale.ItemName = itemName
		// retrieveAllItemsForSale.ItemPrice = itemPrice
		// retrieveAllItemsForSale.ItemQty = itemQty
		// retrieveAllItemsForSale.ItemDesc = itemDesc

	} else {
		genereicResponse.Status = "Failed to Retrieve"
		genereicResponse.Code = 500
		genereicResponse.Message = err.Error()

	}

	if err != nil {
		panic(err)
	}

	displayAllItemsStatement := "SELECT * FROM `marketeer_database1`.`ItemsForSale`"

	rows, err := db.Query(displayAllItemsStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&itemID, &sellerID, &itemImageURL, &itemName, &itemPrice, &itemQty, &itemDesc)
		if err != nil {
			log.Fatal(err)
		}
		retrieveAllItemsForSale.SellerID = sellerID
		retrieveAllItemsForSale.ItemImageLinkURL = itemImageURL
		retrieveAllItemsForSale.ItemID = itemID
		retrieveAllItemsForSale.ItemName = itemName
		retrieveAllItemsForSale.ItemPrice = itemPrice
		retrieveAllItemsForSale.ItemQty = itemQty
		retrieveAllItemsForSale.ItemDesc = itemDesc
		log.Println(itemID, itemName)
		c.JSON(http.StatusOK, retrieveAllItemsForSale)

		currentTime := time.Now()
		fmt.Println("The time is", currentTime)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func CheckErrorRetrieve(err error) {
	if err != nil {
		panic(err)
	}
}
