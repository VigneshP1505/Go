package occupation

import (
	"fmt"
	"time"
)

type office struct {
	address string
}

type Occupation struct {
	Title     string
	salary    string
	isSenior  bool
	office    office
	updatedAt time.Time
}

func CreateOccupation() *Occupation {
	return &Occupation{}
}

func SpawnOccupation(title string, salary string, isSenior bool) (*Occupation, error) {
	return &Occupation{
		Title:    title,
		salary:   salary,
		isSenior: isSenior,
		office: office{
			address: "Toronto",
		},
		updatedAt: time.Now(),
	}, nil
}

func (o Occupation) PrintOccupation() {
	fmt.Println(o.Title, o.salary, o.isSenior, o.updatedAt.Local(), o.office.printOffice())
}

func (o office) printOffice() string {
	return o.address
}
