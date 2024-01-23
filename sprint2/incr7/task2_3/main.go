package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Slice []byte

// MarshalJSON реализует интерфейс json.Marshaler.
func (s Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(s))
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler.
func (s *Slice) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	slice, error := hex.DecodeString(tmp)
	if error != nil {
		return error
	}
	*s = slice
	// s = slice
	// используйте hex.DecodeString, чтобы получить из hex-строки слайс,
	// и присвойте этот слайс значению указателя s
	// ...
	return nil
}

type MySlice struct {
	ID    int
	Slice Slice
}

func main() {
	ret, err := json.Marshal(MySlice{ID: 7, Slice: []byte{1, 2, 3, 10, 11, 255}})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
	var result MySlice
	if err = json.Unmarshal(ret, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
