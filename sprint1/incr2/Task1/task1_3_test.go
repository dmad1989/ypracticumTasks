package main

import (
	"testing"
)

func TestAddNew(t *testing.T) {
	tests := []struct {
		name       string
		initFamily Family
		role       Relationship
		person     Person
		waitFamily Family
		waitError  error
	}{
		{
			name:       "Add new person in empty Family",
			initFamily: Family{},
			role:       Mother,
			person: Person{
				FirstName: "pupa",
				LastName:  "lupa",
				Age:       25,
			},
			waitFamily: Family{
				Members: map[Relationship]Person{Mother: {FirstName: "pupa",
					LastName: "lupa",
					Age:      25}},
			},
			waitError: nil,
		},
		{
			name: "Add new person existing role",
			initFamily: Family{
				Members: map[Relationship]Person{Mother: {FirstName: "pupa",
					LastName: "lupa",
					Age:      25}},
			},
			role: Mother,
			person: Person{
				FirstName: "pupa",
				LastName:  "lupa",
				Age:       25,
			},
			waitFamily: Family{
				Members: map[Relationship]Person{Mother: {FirstName: "pupa",
					LastName: "lupa",
					Age:      25}},
			},
			waitError: ErrRelationshipAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.initFamily.AddNew(test.role, test.person)
			t.Log(test.name)
			if err != test.waitError {
				t.Errorf("Error got = %s, want %s", err, test.waitError)
			}

			for k, val := range test.initFamily.Members {
				if waitVal := test.waitFamily.Members[k]; waitVal != (Person{}) {
					if val.Age != waitVal.Age || val.FirstName != waitVal.FirstName || val.LastName != waitVal.LastName {
						t.Errorf("Wrong person was created: %+v, want %+v", val, waitVal)
					}
				} else {
					t.Errorf("No with such Role accepted %s", k)
				}

			}

		})
	}

}
