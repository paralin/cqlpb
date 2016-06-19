package marshal

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type fieldType struct {
	Field    reflect.Value
	Property *proto.Properties
}

var protoFieldName = "proto"

func buildFieldMap(protoEnum reflect.Value, structProps *proto.StructProperties) map[string]fieldType {
	fieldByName := make(map[string]fieldType)
	for _, prop := range structProps.Prop {
		field := protoEnum.FieldByName(prop.Name)
		fieldByName[prop.OrigName] = fieldType{
			Field:    field,
			Property: prop,
		}
	}
	return fieldByName
}

// Marshal marshals a protocol buffer into a cassandra map.
// Serialize using camelCase
// It's not checked, but cassandra template should contain primitives only
func Marshal(messageo proto.Message, template map[string]interface{}) (map[string]interface{}, error) {
	message := proto.Clone(messageo)
	result := make(map[string]interface{})
	protoEnum := reflect.ValueOf(message).Elem()
	structProps := proto.GetProperties(protoEnum.Type())
	fieldByName := buildFieldMap(protoEnum, structProps)

	// todo: check field is bytes
	if _, ok := template[protoFieldName]; !ok {
		return nil, fmt.Errorf("Template must contain 'proto' field.")
	}

	// we expect ALL of the template vals to exist.
	for k, v := range template {
		if k == protoFieldName {
			continue
		}
		prop, ok := fieldByName[k]
		if !ok {
			return nil, fmt.Errorf("Cannot match template field '%s' to proto field.", k)
		}
		rtyp := reflect.TypeOf(v)
		if prop.Field.Type() != rtyp {
			return nil, fmt.Errorf("Type of template '%s' doesn't match proto type '%s'.", rtyp.Name(), prop.Field.Type().Name())
		}
		result[k] = prop.Field.Interface()
		fmt.Printf("Result %s is kind %s.\n", k, prop.Field.Kind().String())
		prop.Field.Set(reflect.Zero(rtyp))
	}

	// Serialize the rest as proto
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	result["proto"] = data

	return result, nil
}

// Unmarshal marshals a cassandra map to a protocol buffer.
// It's not checked, but cassandra template should contain primitives only
func Unmarshal(message proto.Message, row map[string]interface{}) error {
	protoEnum := reflect.ValueOf(message).Elem()
	structProps := proto.GetProperties(protoEnum.Type())
	fieldByName := buildFieldMap(protoEnum, structProps)

	// todo: check field is bytes
	protoRow, ok := row[protoFieldName]
	if ok && protoRow != nil {
		// return fmt.Errorf("Template must contain 'proto' field.")
		protoArr, ok := protoRow.([]byte)
		if !ok {
			foundProtoType := reflect.ValueOf(protoRow).Type()
			return fmt.Errorf("Expected []byte for 'proto' field, found %s.", foundProtoType.String())
		}

		// unmarshal proto
		if err := proto.Unmarshal(protoArr, message); err != nil {
			return err
		}
	}

	// we expect ALL of the row vals to exist.
	for k, v := range row {
		if k == protoFieldName {
			continue
		}
		prop, ok := fieldByName[k]
		if !ok {
			return fmt.Errorf("Cannot match row field '%s' to proto field.", k)
		}
		rtyp := reflect.TypeOf(v)
		if prop.Field.Type() != rtyp {
			return fmt.Errorf("Type of row '%s' doesn't match proto type '%s'.", rtyp.Name(), prop.Field.Type().Name())
		}
		prop.Field.Set(reflect.ValueOf(v))
	}

	// Serialize the rest as proto
	return nil
}
