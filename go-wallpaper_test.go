// TODOS - obviously...
package main

import (
	"reflect"
	"testing"
)

func TestCheckOSEnviroment(t *testing.T) {
	osEnvironment := CheckOSEnviroment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
