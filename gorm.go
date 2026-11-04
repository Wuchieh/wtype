package wtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// GormSlice is a slice type that is stored as a JSON column by GORM.
type GormSlice[T any] []T

// ToSlice convert to slice
//
//	use for gorm where
func (g GormSlice[T]) ToSlice() []T {
	return g
}

// Scan implements the sql.Scanner interface.
func (g *GormSlice[T]) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &g)
}

// Value implements the driver.Valuer interface.
func (g GormSlice[T]) Value() (driver.Value, error) {
	if len(g) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(g)
}

// GormDBDataType returns the database column type for the current dialect.
func (GormSlice[T]) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
