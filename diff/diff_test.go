package diff

import (
	"encoding/json"
	"testing"
	"reflect"
)

func TestPrimitives(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": "bar",
	"fooInt": 1,
	"fooDouble": 0.0,
	"fooBool": true 
	}`)
	jsonData2 := []byte(`{
	"foo": "baz",
	"fooInt": 2,
	"fooDouble": 0.1,
	"fooBool": false
	}`)
	
	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, jsonObject2)) {
		t.Fail()
	}
}

func TestEqual(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": "bar",
	"bar": {
		"baz": "xyzzy"
	},
	"xyzzy": [1,2]		
	}`)

	expectedData := []byte(`{}`)
	
	var jsonObject1 interface{}
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject1.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestEqual: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestAddedField(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": "bar"
	}`)
	jsonData2 := []byte(`{
	"foo": "bar",
	"bar": "baz"
	}`)
	expectedData := []byte(`{
		"bar": "baz"
	}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestAddedField: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestRemovedField(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": "bar",
	"bar": "baz"
	}`)
	jsonData2 := []byte(`{
	"foo": "bar"
	}`)
	expectedData := []byte(`{
		"bar": "undefined"
	}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestRemovedField: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestTypeMismatch(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": "bar"
	}`)
	jsonData2 := []byte(`{
	"foo": 1
	}`)
	expectedData := []byte(`{
		"foo": 1
	}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestTypeMismatch: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestNested(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": "bar",
	"bar": {
		"baz": "xyzzy"
	}}`)
	jsonData2 := []byte(`{
	"foo": "bar",
	"bar": {
		"baz": 1
	}}`)
	expectedData := []byte(`{
	"bar": {
		"baz": 1
	}}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestNested: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestChangedArray(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": [1,2]
	}`)
	jsonData2 := []byte(`{
	"foo": [1,2,3]
	}`)
	expectedData := []byte(`{
	"foo": [1,2,3]
	}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestArrays: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestArrayOfNestedObjects(t *testing.T) {
	jsonData1 := []byte(`{
	"foo": [{"bar":1},{"baz": 2}]
	}`)
	jsonData2 := []byte(`{
	"foo": [{"bar":2},{"baz":2}]
	}`)
	expectedData := []byte(`{
	"foo": [{"bar":2},{"baz":2}]
	}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestArrayOfNestedObjects: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}

func TestEmpty(t *testing.T) {
	jsonData1 := []byte(`{}`)
	jsonData2 := []byte(`{}`)
	expectedData := []byte(`{}`)

	var jsonObject1 interface{}
	var jsonObject2 interface{}	
	_ = json.Unmarshal(jsonData1, &jsonObject1)
	_ = json.Unmarshal(jsonData2, &jsonObject2)	
	var expectedObject interface{}
	_ = json.Unmarshal(expectedData, &expectedObject)
	
	result := Diff(jsonObject1.(map[string]interface{}), jsonObject2.(map[string]interface{}))
	
	if(!reflect.DeepEqual(result, expectedObject)) {
		resultBytes, _ := json.Marshal(result)
		expectedBytes, _ := json.Marshal(expectedObject)
		t.Errorf("TestEmpty: expected %s actual %s\n", expectedBytes, resultBytes)
	}
}
