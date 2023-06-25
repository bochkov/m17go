package services

import (
	"database/sql"
	"log"
)

type Member struct {
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

func (members Members) FindAll() []Member {
	queryStr := `SELECT me.id, me.name, inst.text, me.actual 
				 FROM members me, instrument inst 
				 WHERE me.instrument = inst.id 
				 ORDER BY me.weight`
	result := make([]Member, 0)
	rows, err := members.db.Query(queryStr)
	if err != nil {
		log.Println(err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		var m Member
		rows.Scan(&m.Id, &m.Name, &m.Instrument, &m.Actual)
		result = append(result, m)
	}
	return result
}

func (members Members) FindActual() []Member {
	var result []Member
	for _, mem := range members.FindAll() {
		if mem.Actual {
			result = append(result, mem)
		}
	}
	return result
}
