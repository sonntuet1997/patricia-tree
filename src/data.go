package src

import (
	"encoding/hex"
)

type MyData string

func (data *MyData) Serialize() string {
	return hex.EncodeToString([]byte(*data))
}

func (data *MyData) Deserialize(h string) error {
	bytes, err := hex.DecodeString(h)
	if err != nil {
		return err
	}
	*data = MyData(bytes)
	return nil
}

func Deserialize(h string) (*MyData, error) {
	bytes, err := hex.DecodeString(h)
	if err != nil {
		return nil, err
	}
	data := MyData(bytes)
	return &(data), nil
}
