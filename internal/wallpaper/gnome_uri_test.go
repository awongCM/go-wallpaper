package wallpaper

import (
	"path/filepath"
	"testing"
)

func TestPathToFileURI_and_FileURIToPath_roundTrip(t *testing.T) {
	path := filepath.Join(string(filepath.Separator), "usr", "share", "backgrounds", "photo with space.jpg")
	uri, err := PathToFileURI(path)
	if err != nil {
		t.Fatal(err)
	}
	if uri == path {
		t.Fatalf("expected file URI, got bare path %q", uri)
	}

	back, err := FileURIToPath(uri)
	if err != nil {
		t.Fatal(err)
	}
	if back != filepath.Clean(path) {
		t.Fatalf("round trip: want %q, got %q", filepath.Clean(path), back)
	}
}

func TestFileURIToPath_stripsGsettingsQuotes(t *testing.T) {
	path := filepath.Join(string(filepath.Separator), "home", "user", "wall.png")
	uri, err := PathToFileURI(path)
	if err != nil {
		t.Fatal(err)
	}
	quoted := "'" + uri + "'"
	got, err := FileURIToPath(quoted)
	if err != nil {
		t.Fatal(err)
	}
	if got != filepath.Clean(path) {
		t.Fatalf("want %q, got %q", filepath.Clean(path), got)
	}
}
