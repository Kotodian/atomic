package proto_model

import (
	"errors"
	"reflect"
)

func ProtoToModel(pb interface{}, model interface{}) error {
	if pb == nil || model == nil {
		return errors.New("pb or model can't be nil")
	}
	src := reflect.ValueOf(pb).Elem()
	dst := reflect.ValueOf(model).Elem()
	for i := 0; i < src.NumField(); i++ {
		if src.Type().Field(i).Name == "XXX_NoUnkeyedLiteral" ||
			src.Type().Field(i).Name == "XXX_unrecognized" ||
			src.Type().Field(i).Name == "XXX_sizecache" {

			continue
		}

		dv := dst.FieldByName(src.Type().Field(i).Name)
		if dv.IsValid() && dv.CanSet() {
			dv.Set(src.Field(i))
		}

	}
	return nil
}
