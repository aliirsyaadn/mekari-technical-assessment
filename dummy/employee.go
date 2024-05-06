package dummy

import (
	"time"

	"github.com/aliirsyaadn/mekari-technical-assessment/model"
)

var DummyEmployees = []model.Employee{
	{
		ID:        1,
		FirstName: "Ali",
		LastName:  "Irsyaad",
		Email:     "aliirsyaadn@gmail.com",
		HireDate: model.Date{
			Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		ID:        2,
		FirstName: "Ali2",
		LastName:  "Irsyaad",
		Email:     "aliirsyaadn2@gmail.com",
		HireDate: model.Date{
			Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	},
}
