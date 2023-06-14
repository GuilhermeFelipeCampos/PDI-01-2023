package mock

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
)

type MockRows struct {
	FillInterfaceOnScan func(dest *[]interface{}) error
	RowCounter          int
}

func (rows *MockRows) Close() {
}

func (rows *MockRows) Err() error {
	return nil
}

func (rows *MockRows) CommandTag() pgconn.CommandTag {
	return nil
}

func (rows *MockRows) FieldDescriptions() []pgproto3.FieldDescription {
	return nil
}

func (rows *MockRows) Scan(dest ...interface{}) error {
	return rows.FillInterfaceOnScan(&dest)
}

func (rows *MockRows) Next() bool {
	ret := false

	if rows.RowCounter > 0 {
		ret = true
		rows.RowCounter--
	}

	return ret
}

func (rows *MockRows) Values() ([]interface{}, error) {
	return nil, nil
}

func (rows *MockRows) RawValues() [][]byte {
	return nil
}
