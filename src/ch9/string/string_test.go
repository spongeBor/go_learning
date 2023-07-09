package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	t.Log(len(s))
	// s[1] = '3' // string 是不可变的byte slice
	s = "\xE4\xB8\xA5"
	t.Log("s1:")
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xA5\xFF"
	t.Log("s2:")
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log("s3:")
	t.Log(len(s))
	c := []rune(s) // 取出unicode
	t.Log("c:")
	t.Log(len(c))
	// Unicode 中一种字符集， UTF8是Unicode 的储存实现
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s { // range 输出的是rune
		t.Logf("%[1]c %[1]x", c)
	}
}
