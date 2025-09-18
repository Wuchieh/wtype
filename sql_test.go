package wtype_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/wuchieh/wtype"
	"gorm.io/gorm/migrator"
)

func TestSqlJson(t *testing.T) {
	interfaceCheck := func(a any) error {
		_, ok := a.(migrator.GormDataTypeInterface)
		if !ok {
			return errors.New("not a GormDataTypeInterface")
		}

		_, ok = a.(sql.Scanner)
		if !ok {
			return errors.New("not a SqlScanner")
		}

		_, ok = a.(driver.Valuer)
		if !ok {
			return errors.New("not a Valuer")
		}
		return nil
	}

	j := wtype.SqlJSON{}
	if err := interfaceCheck(&j); err != nil {
		t.Error(err)
	}
}

func TestSqlJSON2_Unmarshal(t *testing.T) {
	type data struct {
		Name string `json:"name"`
	}

	j := wtype.SqlJSON2[data]{}

	err := j.Scan([]byte(`{"name":123}`))
	if err == nil {
		t.Error("should error")
		return
	}

	err = j.Scan([]byte(`{"name":"test"}`))
	if err != nil {
		t.Error(err)
		return
	}

	var d data
	err = j.Unmarshal(&d)
	if err != nil {
		t.Error(err)
		return
	}

	if d.Name != "test" {
		t.Error("name error")
	}
}

func TestSqlJSON2(t *testing.T) {
	interfaceCheck := func(a any) error {
		_, ok := a.(migrator.GormDataTypeInterface)
		if !ok {
			return errors.New("not a GormDataTypeInterface")
		}

		_, ok = a.(sql.Scanner)
		if !ok {
			return errors.New("not a SqlScanner")
		}

		_, ok = a.(driver.Valuer)
		if !ok {
			return errors.New("not a Valuer")
		}
		return nil
	}

	j := wtype.SqlJSON2[string]{}
	if err := interfaceCheck(&j); err != nil {
		t.Error(err)
	}
}
