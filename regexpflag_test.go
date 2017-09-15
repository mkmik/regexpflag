package regexpflag_test

import (
	"flag"
	"testing"

	"github.com/mmikulicic/regexpflag"
)

func TestFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	defer func(saved *flag.FlagSet) {
		flag.CommandLine = saved
	}(flag.CommandLine)
	flag.CommandLine = fs

	reg := regexpflag.Flag("dummy", "a.*b", "test")

	if !reg.MatchString("aab") {
		t.Error("should match default")
	}

	testCases := [][]string{
		{"--dummy=a.*c"},
		{"--dummy", "a.*c"},
	}
	for _, tc := range testCases {
		err := fs.Parse(tc)
		if err != nil {
			t.Error(err)
		}
		if !reg.MatchString("abc") {
			t.Errorf("expecting to match")
		}
	}
}

func TestBadFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	defer func(saved *flag.FlagSet) {
		flag.CommandLine = saved
	}(flag.CommandLine)
	flag.CommandLine = fs

	reg := regexpflag.Flag("dummy", "a.*b", "test")

	testCases := [][]string{
		{"--dummy=("},
	}
	for _, tc := range testCases {
		err := fs.Parse(tc)
		if err == nil {
			t.Error("expecting error, got no error")
		}

		if !reg.MatchString("aab") {
			t.Error("should match default")
		}
	}
}
