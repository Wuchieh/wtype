package wtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// SqlJSON implements sql.Scanner interface
type (
	//revive:disable-next-line var-naming
	SqlJSON json.RawMessage
	SQLJSON = SqlJSON
)

// Scan implements sql.Scanner interface
func (j *SqlJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = SqlJSON(result)
	return err
}

// Value implements driver.Valuer interface
func (j SqlJSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// GormDBDataType implements migrator.GormDataTypeInterface interface
func (SqlJSON) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

//revive:disable-next-line var-naming
type SqlJSON2[T any] struct {
	SqlJSON
}
type SQLJSON2[T any] = SqlJSON2[T]

func (j SqlJSON2[T]) Unmarshal(a *T) error {
	return json.Unmarshal(j.SqlJSON, a)
}

func (j *SqlJSON2[T]) Scan(value interface{}) error {
	if value == nil {
		j.SqlJSON = SqlJSON("null")
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot scan %T into SqlJSON2", value)
	}

	// 驗證 JSON 格式並檢查類型
	var t T
	if err := json.Unmarshal(bytes, &t); err != nil {
		return fmt.Errorf("invalid JSON for type %T: %w", t, err)
	}

	// 只需要一次解析
	j.SqlJSON = bytes
	return nil
}
