package main

import (
	"reflect"
	"testing"
)

func TestCheckOSEnvironment(t *testing.T) {
	// assuming we're using Mac OSX is default OS
	expectedOS := "darwin"

	osEnvironment := CheckOSEnvironment()
	if osEnvironment != expectedOS {
		t.Fatalf("Expected %s, got %s", expectedOS, osEnvironment)
	}
}

// TODO - mock this exec command
func TestGetCurrentWallpaper(t *testing.T) {
	osEnvironment := CheckOSEnvironment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
func TestGetDefaultLocation(t *testing.T) {
	// assuming we're using Mac OSX is default OS
	expectedDefaultLocation := "/Library/Desktop Pictures/"

	defaultLocation := GetDefaultLocation()
	if defaultLocation != expectedDefaultLocation {
		t.Fatalf("Expected %s, got %s", expectedDefaultLocation, defaultLocation)
	}
}

// TODO - mock this exec command
func TestGetListOfPictures(t *testing.T) {
	osEnvironment := CheckOSEnvironment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}

// TODO - mock this exec command
func TestAlternateWallPapers(t *testing.T) {
	osEnvironment := CheckOSEnvironment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}

// TODO - mock this exec command
func TestCurrentlySetWallpaper(t *testing.T) {
	osEnvironment := CheckOSEnvironment()
	if reflect.TypeOf(osEnvironment).String() != "string" {
		t.Error()
	}
}
