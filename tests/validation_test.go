package tests

import (
	"get-tube/pkg"
	"log"
	"os"
	"testing"
)

func TestCheckParameters(t *testing.T) {

	result, _ := pkg.CheckParameters("https://www.youtube.com/watch?v=YphOsmaYIfE")
	if result != "YphOsmaYIfE" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestValidDIr(t *testing.T) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	result, err := pkg.CheckValidDir(dirname)
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: %s.", result, "TRUE")
	}
}
