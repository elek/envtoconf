package app

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToYaml(t *testing.T) {
	input := map[string]string{
		"test.something.asd":"value1",
		"test.something.bsd":"value2",
		"test.other.bsd":"value3",
	}
	result := ToYaml(input)
	expected := "test:\n  other:\n    bsd: value3\n  something:\n    asd: value1\n    bsd: value2\n"

	assert.Equal(t, expected, result, "Results are not the expected")

	input = map[string]string{
		"test.something.1":"value1",
		"test.something.2":"value2",
	}
	expected = "test:\n  something:\n  - value1\n  - value2\n"
	assert.Equal(t, expected, ToYaml(input), "Results are not the expected\n")

	input = map[string]string{
		"test.something.1.kev":"value1",
		"test.something.1.xxx":"value2",
		"test.something.2.kev":"value3",
	}
	expected =  "test:\n  something:\n  - kev: value1\n    xxx: value2\n  - kev: value3\n"
	assert.Equal(t, expected, ToYaml(input), "Results are not the expected\n")

}

func TestToSh(t *testing.T) {
	input := map[string]string{
		"test.something.asd":"value1",
		"test.something.bsd":"value2",
		"test.other.bsd":"value3",
	}
	result := ToSh(input)
	expected := "export test.something.asd=value1\nexport test.something.bsd=value2\nexport test.other.bsd=value3\n"

	assert.Equal(t, expected, result, "Results are not the expected")

}

func TestToProperties(t *testing.T) {
	input := map[string]string{
		"test.something.asd":"value1",
		"test.something.bsd":"value2",
		"test.other.bsd":"value3",
	}
	result := ToProperties(input)
	expected := "test.something.asd: value1\ntest.something.bsd: value2\ntest.other.bsd: value3\n"

	assert.Equal(t, expected, result, "Results are not the expected")

}
