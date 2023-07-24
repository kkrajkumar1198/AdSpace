// supply_side_service/main.go
package main

import (
	"fmt"
	"time"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type AdSpace struct {
	ID          int       `json:"ad_space_id"`
	Name        string    `json:"ad_space_name"`
	BasePrice   float64   `json:"ad_space_price"`
	IsAdSpaceAvailable   bool   `json:"is_ad_available"`
	AuctionTime string `json:"auction_time"`
}

type Bid struct {
	ID         int       `json:"id"`
	BidderID   int       `json:"bidder_id"`
	AdSpaceID  int       `json:"ad_space_id"`
	Amount     float64   `json:"amount"`
	BidTime    time.Time `json:"bid_time"`
}

// handleListAdSpaces fetches all ad spaces from the database and returns them as JSON response.
func handleListAdSpaces() []*AdSpace {
	db, err := sql.Open("mysql", "rk:rk@tcp(db:3306)/adspace")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ad_space")
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()

	var adSpaces []*AdSpace
	for rows.Next() {
		var adSpace AdSpace
		if err := rows.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &adSpace.IsAdSpaceAvailable, &adSpace.AuctionTime); err != nil {
			panic(err.Error())
		}
		adSpaces = append(adSpaces, &adSpace)
	}

	return adSpaces
}

// adSpacePage handles the HTTP request to list all ad spaces.
func adSpacePage(w http.ResponseWriter, r *http.Request) {
	adspaces := handleListAdSpaces()
	fmt.Println("Endpoint Hit: adSpacePage")
	json.NewEncoder(w).Encode(adspaces)
}

func main() {
	http.HandleFunc("/ad_spaces", adSpacePage)
	fmt.Println("Supply side service started at :8080")
	http.ListenAndServe(":8080", nil)
}
