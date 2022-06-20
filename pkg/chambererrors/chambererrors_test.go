package chambererrors_test

import (
	"testing"

	"github.com/graytonio/chamber/pkg/chambererrors"
)

type error_test struct {
	Err  *chambererrors.ChamberError
	Want string
}

var error_tests = []error_test{
	{&chambererrors.BadConfigError, "Error 1: Error Reading Config File. Check Syntax"},
	{&chambererrors.ChamberError{StatusCode: 23, Message: "Foo Bar Error"}, "Error 23: Foo Bar Error"},
}

func TestError(t *testing.T) {
	for _, test := range error_tests {
		if output := test.Err.Error(); output != test.Want {
			t.Errorf("FAILED\nWanted:\n\"%s\"\nGot:\n\"%s\"", test.Want, output)
		}
	}
}
