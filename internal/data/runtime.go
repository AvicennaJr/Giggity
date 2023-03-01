package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotesJSONValue := strconv.Quote(jsonValue)

	return []byte(quotesJSONValue), nil
}
