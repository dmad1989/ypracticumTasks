package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	tests := []struct {
		name string
		user User
		want string
	}{
		{
			name: "empty strs",
			user: User{
				FirstName: "",
				LastName:  ""},
			want: " ",
		},
		{
			name: "empty fname",
			user: User{
				FirstName: "",
				LastName:  "Rururr"},
			want: " Rururr",
		},
		{
			name: "empty lname",
			user: User{
				FirstName: "Rururr",
				LastName:  ""},
			want: "Rururr ",
		},
		{
			name: "notempty strs",
			user: User{
				FirstName: "Tut",
				LastName:  "Tam"},
			want: "Tut Tam",
		},
		{
			name: "no lname",
			user: User{
				FirstName: "Tut"},
			want: "Tut ",
		},
		{
			name: "no fname",
			user: User{
				LastName: "Tam"},
			want: " Tam",
		},
		{
			name: "no vals",
			user: User{},
			want: " ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := test.user.FullName()
			assert.Equal(t, test.want, res, "FullName() = %s, want %s", res, test.want)

		})
	}
}
