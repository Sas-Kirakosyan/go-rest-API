package models

import (
	"log"
	"time"

	"github.com/Sas-Kirakosyan/my-go-project/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	          VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		log.Println("Error executing statement:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	// Iterate over the rows returned by the query
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		events = append(events, event)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		// log.Println("Error scanning row:", err)
		// continue
		return nil, err
	}
	return &event, nil
}
