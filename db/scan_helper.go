package db

import "database/sql"

func GetLong(number sql.NullInt64) *int64 {
	if number.Valid {
		return &number.Int64
	}
	return nil
}

func GetString(str sql.NullString) *string {
	if str.Valid {
		return &str.String
	}
	return nil
}

func GetFloat(float sql.NullFloat64) *float64 {
	if float.Valid {
		return &float.Float64
	}
	return nil
}
