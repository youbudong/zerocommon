package converters

import (
	"errors"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectID() []copier.TypeConverter {
	return []copier.TypeConverter{
		ObjectIdToStringConverter(),
		StringToObjectIdConverter(),
		StringptrToObjectIdConverter(),
		ObjectIdptrToStringConverter(),
		StringptrToObjectIdptrConverter(),
		ObjectIdptrToStringptrConverter(),
	}
}

func ObjectIdToStringConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: primitive.ObjectID{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(primitive.ObjectID)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return s.Hex(), nil
		},
	}
}

func StringToObjectIdConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: primitive.ObjectID{},
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(string)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return primitive.ObjectIDFromHex(s)
		},
	}
}

func StringptrToObjectIdConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: *new(string),
		DstType: primitive.ObjectID{},
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*string)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return primitive.ObjectIDFromHex(*s)
		},
	}
}

func ObjectIdptrToStringConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: *new(primitive.ObjectID),
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*primitive.ObjectID)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			return s.Hex(), nil
		},
	}
}

func StringptrToObjectIdptrConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: *new(string),
		DstType: *new(primitive.ObjectID),
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*string)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			id, err := primitive.ObjectIDFromHex(*s)
			if err != nil {
				return nil, err
			}

			return &id, nil
		},
	}
}

func ObjectIdptrToStringptrConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: *new(primitive.ObjectID),
		DstType: *new(string),
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*primitive.ObjectID)
			if !ok {
				return nil, errors.New("src type not matching")
			}

			id := s.Hex()
			return &id, nil
		},
	}
}
