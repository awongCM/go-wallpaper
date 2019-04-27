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
func TestGetCurrentWallpaper(t *testing.T) {
	osEnvironment := CheckOSEnviroment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
func TestGetDefaultLocation(t *testing.T) {
	osEnvironment := CheckOSEnviroment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
func TestGetListOfPictures(t *testing.T) {
	osEnvironment := CheckOSEnviroment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
func TestAlternateWallPapers(t *testing.T) {
	osEnvironment := CheckOSEnviroment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
func TestCurrentlySetWallpaper(t *testing.T) {
	osEnvironment := CheckOSEnviroment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
