package id

import (
	"database/sql/driver"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	NullID = ID{0, false}
	ZeroID = ID{0, true}
)

type ID struct {
	Uint64 uint64
	Valid  bool
}

//Generate new ID
func NewID() (ID, error) {
	var id ID
	idUint, err := strconv.ParseUint(fmt.Sprintf("%d", time.Now().UnixNano()+rand.Int63n(100)), 10, 64)
	if err != nil {
		return id, err
	}

	id.Uint64 = idUint
	id.Valid = true
	return id, nil
}

// Parses the s string to create a ID.
func Parse(s string) (ID, error) {
	var id ID
	if len(s) == 0 {
		return id, fmt.Errorf("empty string")
	}
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return id, err
	}
	id.Uint64 = value
	id.Valid = true
	return id, nil
}

// Scan implements the Scanner interface.
func (id *ID) Scan(value interface{}) error {
	if value == nil {
		id.Uint64 = 0
		id.Valid = false
		return nil
	}

	switch vc := value.(type) {
	case int64:
		id.Uint64 = 0 + uint64(vc)
		id.Valid = true
		return nil
	case []byte:
		i64, err := strconv.ParseInt(string(vc), 10, 64)
		if err != nil {
			return err
		}
		id.Uint64 = 0 + uint64(i64)
		id.Valid = true
		return nil
	default:
		return fmt.Errorf("interface {} is not int64")
	}
}

// Value implements the driver Valuer interface.
func (id ID) Value() (driver.Value, error) {
	if !id.Valid {
		return nil, nil
	}
	// uint64 isn't a valid driver.Value value, so casting to int64
	return int64(id.Uint64), nil
}
