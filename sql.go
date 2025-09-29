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

func (j SqlJSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

func (SqlJSON) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
