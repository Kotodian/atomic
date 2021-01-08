package proto_model

import (
	"errors"
	"reflect"
)

func ProtoToModel(src interface{}, dst interface{}) error {
	if src == nil || dst == nil {
		return errors.New("pb or model can't be nil")
	}
	srcElem := reflect.ValueOf(src).Elem()
	dstElem := reflect.ValueOf(dst).Elem()
	for i := 0; i < srcElem.NumField(); i++ {
		if srcElem.Type().Field(i).Name == "XXX_NoUnkeyedLiteral" ||
			srcElem.Type().Field(i).Name == "XXX_unrecognized" ||
			srcElem.Type().Field(i).Name == "XXX_sizecache" {

			continue
		}

		dv := dstElem.FieldByName(srcElem.Type().Field(i).Name)
		if dv.IsValid() && dv.CanSet() {
			dv.Set(srcElem.Field(i))
		}

	}
	return nil
}
