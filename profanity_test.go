package profanity

import "testing"

func TestHasProfanity(t *testing.T) {

	exbl := []string{ // example blacklist
		"dang", "shoot",
	}

	var tests = []struct {
		input1   []string
		input2   string
		expected bool
	}{
		// case: no profanity
		{exbl, "no profanity here!", false},
		// case: contains profanity
		{exbl, "gosh dang it!", true},
		// case: pass empty string should return false
		{exbl, "", false},
		// case: pass empty blacklist should return false
		{[]string{}, "gosh dang it!", false},
	}

	for _, test := range tests {
		output := HasProfanity(test.input1, test.input2)
		if output != test.expected {
			t.Errorf(
				"Test failed:\ninput: %s, %s\nexpected: %t\nactual: %t\n",
				test.input1, test.input2, test.expected, output)
		}
	}
}

// Masking functions test cases:
// 		1. empty text string should return empty string
//		2. empty blacklist should return identicle text string
//		3. no profanity text should return identicle text string
//		4. profanity should be masked
//		5. multiple profanity should be masked
//		6. should work properly with utf8 unicode characters
//		7. function doesn't modify the original string

var (
	bl1 = []string{"dang", "shoot"}
	bl2 = []string{"绿茶婊"}
)

type maskingTest struct {
	blacklist []string
	text      string
	c         rune
	expected  string
}

func TestMaskAll(t *testing.T) {

	tests := []maskingTest{
		{bl1, "", '*', ""},
		{[]string{}, "some text", '*', "some text"},
		{bl1, "no profanity", '*', "no profanity"},
		{bl1, "dang it", '*', "**** it"},
		{bl1, "shoot dang", '*', "***** ****"},
		{bl2, "you 绿茶婊!", 'の', "you ののの!"},
	}

	for _, test := range tests {
		output := MaskAll(test.blacklist, test.text, test.c)
		if output != test.expected {
			t.Errorf(
				"Test failed:\nblacklist: %s\ntext: %s\nc: %c\n"+
					"expected: %s\nactual: %s\n",
				test.blacklist, test.text, test.c, test.expected, output)
		}
	}
}

func TestMaskMiddle(t *testing.T) {

	tests := []maskingTest{
		{bl1, "", '*', ""},
		{[]string{}, "some text", '*', "some text"},
		{bl1, "no profanity", '*', "no profanity"},
		{bl1, "dang it", '*', "d**g it"},
		{bl1, "shoot dang", '*', "s***t d**g"},
		{bl2, "you 绿茶婊!", 'の', "you 绿の婊!"},
		// Extra case: should work with one and two character blacklist words
		{[]string{"xy", "z"}, "xy z", '*', "xy z"},
	}

	for _, test := range tests {
		output := MaskMiddle(test.blacklist, test.text, test.c)
		if output != test.expected {
			t.Errorf(
				"Test failed:\nblacklist: %s\ntext: %s\nc: %c\n"+
					"expected: %s\nactual: %s\n",
				test.blacklist, test.text, test.c, test.expected, output)
		}
	}
}

// func TestMaskTail(t *testing.T) {

// 	tests := []maskingTest{
// 		{bl1, "", '*', ""},
// 		{[]string{}, "some text", '*', "some text"},
// 		{bl1, "no profanity", '*', "no profanity"},
// 		{bl1, "dang it", '*', "d*** it"},
// 		{bl1, "shoot dang", '*', "s**** d***"},
// 		{bl2, "you 绿茶婊!", 'の', "you 绿のの!"},
// 	}

// 	for _, test := range tests {
// 		output := MaskTail(test.blacklist, test.text, test.c)
// 		if output != test.expected {
// 			t.Errorf(
// 				"Test failed:\nblacklist: %s\ntext: %s\nc: %c\n"+
// 					"expected: %s\nactual: %s\n",
// 				test.blacklist, test.text, test.c, test.expected, output)
// 		}
// 	}
// }
