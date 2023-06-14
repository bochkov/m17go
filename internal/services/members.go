package services

import (
	"database/sql"
	"log"
)

type member struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Instrument string `json:"instrument"`
	Actual     bool   `json:"actual"`
}

type Members struct {
	db *sql.DB
}

func NewMembers(db *sql.DB) *Members {
	return &Members{db: db}
}

func (members Members) FindAll() []member {
	result := make([]member, 0)
	rows, err := members.db.
		Query(`SELECT m.id, m.name, i.text, m.actual 
			FROM members m, instrument i 
			WHERE m.instrument = i.id 
			ORDER BY m.weight`)
	if err != nil {
		log.Println(err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var m member
		rows.Scan(&m.Id, &m.Name, &m.Instrument, &m.Actual)
		result = append(result, m)
	}
	return result
}

func (members Members) FindActual() []member {
	var result []member
	for _, mem := range members.FindAll() {
		if mem.Actual {
			result = append(result, mem)
		}
	}
	return result
}
