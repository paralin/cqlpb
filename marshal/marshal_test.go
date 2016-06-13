package marshal

import (
	"encoding/json"
	"reflect"
	"testing"

	cqlpb_test "github.com/paralin/cqlpb/marshal/marshal_test"
)

func TestMarshalTest1(t *testing.T) {
	val := &cqlpb_test.TestOne{
		StringFoo:  "foos",
		IntFoo:     4,
		StringFooB: "kappa",
	}
	temp := make(map[string]interface{})
	temp["stringFoo"] = ""
	temp["intFoo"] = int32(0)
	temp["proto"] = make([]byte, 1)

	result, err := Marshal(val, temp)
	if err != nil {
		t.Error("Expected no err, got ", err)
	}

	// debug output
	jval, err := json.Marshal(result)
	if err != nil {
		t.Error("Expected no err for json serialize, got ", err)
	}
	t.Log(string(jval))

	// now deserialize
	// intentionally put oldVal here, this will make sure it doesn't destroy the data
	oldVal := *val
	val = &cqlpb_test.TestOne{}
	err = Unmarshal(val, result)

	// assert the same
	jval, err = json.Marshal(*val)
	if err != nil {
		t.Error("Expected no err for json serialize, got ", err)
	}
	t.Log(string(jval))
	jval, err = json.Marshal(oldVal)
	if err != nil {
		t.Error("Expected no err for json serialize, got ", err)
	}
	t.Log(string(jval))
	if !reflect.DeepEqual(oldVal, *val) {
		t.Error("Unmarshalled value is not identical.")
	}
}
