// Code generated by go generate; DO NOT EDIT.
// This file was generated by genenum.go

package main

import "fmt"

var namesHTTPCode = map[HTTPCode]string{
	400: "BadRequest",
	401: "Unauthorized",
	403: "Forbidden",
	404: "NotFound",
	429: "TooManyRequests", // DDOS alert, call admins!
	500: "Internal",
}

func (v HTTPCode) String() string {
	return namesHTTPCode[v]
}

type BadRequestError struct {
	Description string
}

func (v BadRequestError) Error() string {
	return fmt.Sprintf("HTTP 400 %s", v.Description)
}

type UnauthorizedError struct {
	Description string
}

func (v UnauthorizedError) Error() string {
	return fmt.Sprintf("HTTP 401 %s", v.Description)
}

type ForbiddenError struct {
	Description string
}

func (v ForbiddenError) Error() string {
	return fmt.Sprintf("HTTP 403 %s", v.Description)
}

type NotFoundError struct {
	Description string
}

func (v NotFoundError) Error() string {
	return fmt.Sprintf("HTTP 404 %s", v.Description)
}

type TooManyRequestsError struct {
	Description string
}

func (v TooManyRequestsError) Error() string {
	return fmt.Sprintf("HTTP 429 %s", v.Description)
}
func (v TooManyRequestsError) Alert() string {
	return "DDOS alert, call admins!"
}

type InternalError struct {
	Description string
}

func (v InternalError) Error() string {
	return fmt.Sprintf("HTTP 500 %s", v.Description)
}
