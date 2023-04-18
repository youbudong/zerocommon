package converters

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
)

func TimeInt64() []copier.TypeConverter {
	return []copier.TypeConverter{
		TimeToInt64(),
		Int64ToTime(),
		TimepToInt64p(),
		Int64pToTimep(),
		TimepToint64(),
		Int64pToTime(),
	}
}

func TimeToInt64() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: time.Time{},
		DstType: int64(0),
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return s.UnixMilli(), nil
		},
	}
}

func Int64ToTime() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: int64(0),
		DstType: time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return time.UnixMilli(s), nil
		},
	}
}

func Int64pToTimep() copier.TypeConverter {
	var int64p *int64
	var timep *time.Time
	return copier.TypeConverter{
		SrcType: int64p,
		DstType: timep,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			t := time.UnixMilli(*s)
			return &t, nil
		},
	}
}

func TimepToInt64p() copier.TypeConverter {
	var int64p *int64
	var timep *time.Time
	return copier.TypeConverter{
		SrcType: timep,
		DstType: int64p,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			t := s.UnixMilli()
			return &t, nil
		},
	}
}

func TimepToint64() copier.TypeConverter {
	var int64p int64
	var timep *time.Time
	return copier.TypeConverter{
		SrcType: timep,
		DstType: int64p,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			t := s.UnixMilli()
			return t, nil
		},
	}
}

func Int64pToTime() copier.TypeConverter {
	var int64p *int64
	var timep time.Time
	return copier.TypeConverter{
		SrcType: int64p,
		DstType: timep,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			t := time.UnixMilli(*s)
			return t, nil
		},
	}
}
