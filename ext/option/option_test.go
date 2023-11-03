package option

import (
	"encoding/json"
	"testing"
)

func Test_OptionInt(t *testing.T) {
	i := Some(42)

	if i.IsNone() {
		t.Errorf("OptionInt should not be None")
	}

	if i.Value() != 42 {
		t.Errorf("OptionInt should be 42")
	}

	data := struct {
		Data Option[int] `json:"d"`
	}{
		Data: i,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionInt should marshal")
	}

	if string(bytes) != `{"d":42}` {
		t.Errorf("Expected %s, got %s", `{"d":42}`, string(bytes))
	}
}

func Test_OptionIntNone(t *testing.T) {
	i := None[int]()

	if !i.IsNone() {
		t.Errorf("OptionInt should be None")
	}

	data := struct {
		Data Option[int] `json:"d"`
	}{
		Data: i,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionInt should marshal")
	}

	if string(bytes) != `{"d":null}` {
		t.Errorf("Expected %s, got %s", `{"d":null}`, string(bytes))
	}
}

func Test_OptionString(t *testing.T) {
	s := Some("foo")

	if s.IsNone() {
		t.Errorf("OptionString should not be None")
	}

	if s.Value() != "foo" {
		t.Errorf("OptionString should be foo")
	}

	data := struct {
		Data Option[string] `json:"d"`
	}{
		Data: s,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionString should marshal")
	}

	if string(bytes) != `{"d":"foo"}` {
		t.Errorf("Expected %s, got %s", `{"d":"foo"}`, string(bytes))
	}
}

func Test_OptionStringNone(t *testing.T) {
	s := None[string]()

	if !s.IsNone() {
		t.Errorf("OptionString should be None")
	}

	data := struct {
		Data Option[string] `json:"d"`
	}{
		Data: s,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionString should marshal")
	}

	if string(bytes) != `{"d":null}` {
		t.Errorf("Expected %s, got %s", `{"d":null}`, string(bytes))
	}
}

func Test_OptionBool(t *testing.T) {
	b := Some(true)

	if b.IsNone() {
		t.Errorf("OptionBool should not be None")
	}

	if b.Value() != true {
		t.Errorf("OptionBool should be true")
	}

	data := struct {
		Data Option[bool] `json:"d"`
	}{
		Data: b,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionBool should marshal")
	}

	if string(bytes) != `{"d":true}` {
		t.Errorf("Expected %s, got %s", `{"d":true}`, string(bytes))
	}
}

func Test_OptionBoolNone(t *testing.T) {
	b := None[bool]()

	if !b.IsNone() {
		t.Errorf("OptionBool should be None")
	}

	data := struct {
		Data Option[bool] `json:"d"`
	}{
		Data: b,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionBool should marshal")
	}

	if string(bytes) != `{"d":null}` {
		t.Errorf("Expected %s, got %s", `{"d":null}`, string(bytes))
	}
}

func Test_OptionStruct(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}

	f := Some(Foo{Bar: "baz"})

	if f.IsNone() {
		t.Errorf("OptionStruct should not be None")
	}

	if f.Value().Bar != "baz" {
		t.Errorf("OptionStruct should be baz")
	}

	data := struct {
		Data Option[Foo] `json:"d"`
	}{
		Data: f,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionStruct should marshal")
	}

	if string(bytes) != `{"d":{"bar":"baz"}}` {
		t.Errorf("Expected %s, got %s", `{"d":{"bar":"baz"}}`, string(bytes))
	}
}

func Test_OptionStructNone(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}

	f := None[Foo]()

	if !f.IsNone() {
		t.Errorf("OptionStruct should be None")
	}

	data := struct {
		Data Option[Foo] `json:"d"`
	}{
		Data: f,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionStruct should marshal")
	}

	if string(bytes) != `{"d":null}` {
		t.Errorf("Expected %s, got %s", `{"d":null}`, string(bytes))
	}
}

func Test_OptionSlice(t *testing.T) {
	s := Some([]string{"foo", "bar"})

	if s.IsNone() {
		t.Errorf("OptionSlice should not be None")
	}

	if s.Value()[0] != "foo" {
		t.Errorf("OptionSlice should be foo")
	}

	data := struct {
		Data Option[[]string] `json:"d"`
	}{
		Data: s,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionSlice should marshal")
	}

	if string(bytes) != `{"d":["foo","bar"]}` {
		t.Errorf("Expected %s, got %s", `{"d":["foo","bar"]}`, string(bytes))
	}
}

func Test_OptionSliceNone(t *testing.T) {
	s := None[[]string]()

	if !s.IsNone() {
		t.Errorf("OptionSlice should be None")
	}

	data := struct {
		Data Option[[]string] `json:"d"`
	}{
		Data: s,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("OptionSlice should marshal")
	}

	// Technically, you might expect this to be `{"d":[]}`, but that's not the case here, fork and fix if you need to.
	if string(bytes) != `{"d":null}` {
		t.Errorf("Expected %s, got %s", `{"d":null}`, string(bytes))
	}
}
