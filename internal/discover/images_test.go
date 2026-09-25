package discover

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListImages_filtersExtensionsAndDirs(t *testing.T) {
	root := t.TempDir()

	if err := os.Mkdir(filepath.Join(root, ".thumbnails"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".thumbnails", "hidden.jpg"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}

	want := map[string]struct{}{
		filepath.Join(root, "a.jpg"):         {},
		filepath.Join(root, "nested", "b.PNG"): {},
	}
	for path := range want {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("img"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "readme.txt"), []byte("nope"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := ListImages(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(want) {
		t.Fatalf("expected %d images, got %d: %v", len(want), len(got), got)
	}
	for _, path := range got {
		if _, ok := want[path]; !ok {
			t.Fatalf("unexpected path %q", path)
		}
	}
}

func TestListImages_missingRoot(t *testing.T) {
	_, err := ListImages(filepath.Join(t.TempDir(), "does-not-exist"))
	if err == nil {
		t.Fatal("expected error for missing root")
	}
}
