package wallpaper

import (
	"runtime"
	"testing"
)

func TestNewBackendFor_unsupportedOS(t *testing.T) {
	_, err := newBackendFor("windows")
	if err == nil {
		t.Fatal("expected error for windows")
	}
}

func TestNew_supportedPlatforms(t *testing.T) {
	switch runtime.GOOS {
	case "darwin", "linux":
		b, err := New()
		if err != nil {
			t.Fatal(err)
		}
		if b.Name() == "" {
			t.Fatal("expected non-empty backend name")
		}
		if b.DefaultWallpaperDir() == "" {
			t.Fatal("expected default wallpaper dir")
		}
	default:
		t.Skipf("unsupported test OS: %s", runtime.GOOS)
	}
}

func TestEscapeAppleScriptString(t *testing.T) {
	got := escapeAppleScriptString(`C:\path "quote"`)
	want := `C:\\path \"quote\"`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
