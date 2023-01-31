package converters

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
)

func TimeToInt64() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.Time,
		DstType: copier.Int64,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return s.Unix(), nil
		},
	}
}

func Int64ToTime() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.Int64,
		DstType: copier.Time,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return time.UnixMilli(s), nil
		},
	}
}
