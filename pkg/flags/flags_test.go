package flags

import (
	"fmt"
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	expected := "./../../configs/config.yml"
	os.Args[1] = fmt.Sprintf("-config=%s", expected)

	path, err := ParseFlags()
	if err != nil {
		t.Errorf("Error parsing flags:\n%s", err)
	}

	if path != expected {
		t.Errorf("The flag read is not the same as the expected: %s", path)
	}
}
