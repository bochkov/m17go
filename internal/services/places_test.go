package services_test

import (
	"database/sql"
	"testing"

	"github.com/bochkov/m17go/internal/services"
	_ "github.com/lib/pq"
)

func TestFindById(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://m17:m17@10.10.10.10:5432/m17?sslmode=disable")
	if err != nil {
		t.Error("cannot connect test database", err)
	}
	places := services.NewPlaces(db)
	expected := services.Place{
		Id:           1,
		Name:         "Дом печати",
		Address:      "Екатеринбург, ул. Ленина, 49",
		Link:         "https://tele-club.ru/dompechati",
		Slug:         "dompechati",
		InvertedLogo: true,
	}
	v := places.FindById(1)
	if expected != v {
		t.Error(
			"Expected", expected,
			"Got", v,
		)
	}

}
