package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Bidder struct {
	ID         int       `json:"bidder_id"`
	Name       string    `json:"bidder_name"`
	Budget     int       `json:"bidder_budget"`
	BidTime    string  `json:"bid_time"`
}

type Bid struct {
	ID         int       `json:"id"`
	BidderID   int       `json:"bidder_id"`
	AdSpaceID  int       `json:"ad_space_id"`
	Amount     float64   `json:"amount"`
	BidTime    string `json:"bid_time"`
}

type AdSpaceID struct {
	ID         int      `json:"ad_space_id"`    
}

var db *sql.DB

// BiddersPage handles the HTTP request to list all bidders.
func BiddersPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering BiddersPage")
	biddersList := getBiddersFromDB()
	fmt.Println("Endpoint Hit: biddersPage")
	json.NewEncoder(w).Encode(biddersList)
}

// getBiddersFromDB fetches all bidders from the database.
func getBiddersFromDB() []*Bidder {
	// fmt.Println("entering db fetch")
	db, err := sql.Open("mysql", "rk:rk@tcp(db:3306)/adspace")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM bidders")
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()

	var bidders []*Bidder
	for rows.Next() {
		var bidder Bidder
		err := rows.Scan(&bidder.ID, &bidder.Name, &bidder.Budget, &bidder.BidTime)
		if err != nil {
			log.Print(err.Error())
		}
		bidders = append(bidders, &bidder)
	}

	return bidders
}

// handleGetBidderDetails handles the HTTP POST request to add a new bidder.
func createNewBids(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var bids Bid

	// Set the current time as the bid_time for the new bidder
	bids.BidTime = time.Now().Format("2006-01-02 15:04:05")
	
	if err := json.NewDecoder(r.Body).Decode(&bids); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to parse the request data", http.StatusBadRequest)
		return
	}
	
	err := insertBidsIntoDB(bids)
	if err != nil {
		http.Error(w, "Failed to store bidder in the database", http.StatusInternalServerError)
		return
	}
}

// insertBidsIntoDB inserts a new bids into the database.
func insertBidsIntoDB(bid Bid) error {

	db, err := sql.Open("mysql", "rk:rk@tcp(db:3306)/adspace")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	fmt.Println(db)
	fmt.Println("entered")
	// Insert the bidder into the database
	query := "INSERT INTO bids (bidder_id, ad_space_id, amount, bid_time) VALUES (?, ?, ?, ?)"
	if err != nil {
		return err
	}
	// Execute the query with the data from the newBid struct
	result, err := db.Exec(query, bid.BidderID, bid.AdSpaceID, bid.Amount, bid.BidTime)
	if err != nil {
		log.Fatal(err)
	}

	// Get the ID of the newly inserted record
	newID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New bid with ID %d inserted successfully.\n", newID)

	return nil
}

// GetSpecificAdBidderDetails handles the HTTP POST request to get bidders for a specific ad space.
func GetSpecificAdBidDetails(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var adSpaceID AdSpaceID
	
	if err := json.NewDecoder(r.Body).Decode(&adSpaceID); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to parse the request data", http.StatusBadRequest)
		return
	}
	fmt.Println(adSpaceID.ID)
	
	bids := getSpecificAdBidsFromDB(adSpaceID)
	
	if err := json.NewEncoder(w).Encode(bids); err != nil {
		http.Error(w, "Failed to encode bidder details", http.StatusInternalServerError)
		return
	}
}

// getSpecificAdBiddersFromDB fetches bidders for a specific ad space from the database.
func getSpecificAdBidsFromDB(adSpaceID AdSpaceID) []*Bid {
	fmt.Println("entering db fetch")
	db, err := sql.Open("mysql", "rk:rk@tcp(db:3306)/adspace")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM bids where ad_space_id = ?", adSpaceID.ID)
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()
	
	var bids []*Bid
	for rows.Next() {
		var bid Bid
		if err := rows.Scan(&bid.ID, &bid.BidderID, &bid.AdSpaceID, &bid.Amount, &bid.BidTime); err != nil {
			log.Fatal(err)
		}
		bids = append(bids, &bid)
	}

	return bids
}

func main() {
	http.HandleFunc("/list_bidders", BiddersPage)
	http.HandleFunc("/get_ad_space_bids", GetSpecificAdBidDetails)
	http.HandleFunc("/createnewbids", createNewBids)
	http.ListenAndServe(":8081", nil)
}
