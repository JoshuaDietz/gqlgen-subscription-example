// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// `Time` is a simple type only containing the current time as
// a unix epoch timestamp and a string timestamp.
type Time struct {
	UnixTime  int    `json:"unixTime"`
	TimeStamp string `json:"timeStamp"`
}
