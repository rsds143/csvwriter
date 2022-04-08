package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jaswdr/faker"
)

func main() {
	f, err := os.Create("./records.csv")
	if err != nil {
		log.Fatalf("unexpected error opening file for writing %v", f)
	}
	w := csv.NewWriter(f)
	fake := faker.New()
	header := []string{"id", "first_name", "last_name", "postal_code", "state", "created_on", "weight", "ranking", "ext_id"}
	if err := w.Write(header); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
    for i := 0; i < 400000; i++ {
		firstName := fake.Person().FirstName()
		lastName := fake.Person().LastName()
		postalCode := fake.Address().PostCode()
		state := fake.Address().State()
		createdOn := fake.Time().Unix(time.Now()) * 1000
		weight := fake.RandomFloat(2, 0, 3)
		ranking := fake.Int32()
		extId := fake.UUID().V4()
		record := []string{fmt.Sprint(i), firstName, lastName, postalCode, state, fmt.Sprint(createdOn), fmt.Sprint(weight), fmt.Sprint(ranking), extId}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
