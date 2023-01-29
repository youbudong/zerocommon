package converters

import (
	"errors"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
