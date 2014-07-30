package regex

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

var (
	testFileName      = "test.dat"
	testFileFalseName = "test_false.dat"
)

func TestRegexTrue(t *testing.T) {
	f, err := os.OpenFile(testFileName, os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("Can't read test file: %s", testFileName)
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Fatalf("read file error: %s", err)
			}
		}
		splits := strings.SplitN(line, " ", 2)
		if len(splits) != 2 {
			t.Fatalf("INVALID file format, line = %s", line)
		}
		str := strings.TrimLeft(splits[0], " \t\r\n")
		match := strings.TrimLeft(splits[1], " \t\r\n")
		match = strings.TrimRight(match, " \t\r\n")
		reg, err := NewRegex(str)
		if err != nil {
			t.Fatalf("build regex %s error: %s", str, err)
		}
		if !reg.Match(match) {
			t.Fatalf("Regex:%s doesn't match %s", str, match)
		}
	}
}

func TestRegexFalse(t *testing.T) {
	f, err := os.OpenFile(testFileFalseName, os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("Can't read test file: %s", testFileFalseName)
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Fatalf("read file error: %s", err)
			}
		}
		splits := strings.SplitN(line, " ", 2)
		if len(splits) != 2 {
			t.Fatalf("INVALID file format, line = %s", line)
		}
		str := strings.TrimLeft(splits[0], " \t\r\n")
		match := strings.TrimLeft(splits[1], " \t\r\n")
		match = strings.TrimRight(match, " \t\r\n")
		reg, err := NewRegex(str)
		if err != nil {
			t.Fatalf("build regex %s error: %s", str, err)
		}
		if reg.Match(match) {
			t.Fatalf("Regex:%s match %s", str, match)
		}
	}
}
