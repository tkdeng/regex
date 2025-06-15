package regex

// Match returns true if a []byte matches a regex
func (reg *Regexp) Match(buf []byte) bool {
	return reg.RE.Match(buf)
}

// Split splits a string, and keeps capture groups
//
// Similar to JavaScript .split(/re/)
func (reg *Regexp) Split(buf []byte) [][]byte {
	ind := reg.RE.FindAllIndex(buf, -1)

	res := [][]byte{}
	trim := 0
	for _, pos := range ind {
		v := buf[pos[0]:pos[1]]
		m := reg.RE.FindSubmatch(v)

		if trim == 0 {
			res = append(res, buf[:pos[0]])
		} else {
			res = append(res, buf[trim:pos[0]])
		}
		trim = pos[1]

		for i := 1; i <= len(m)-1; i++ {
			res = append(res, m[i])
		}
	}

	res = append(res, buf[trim:])

	return res
}

// SplitStr splits a string, and keeps capture groups
//
// Similar to JavaScript .split(/re/)
func (reg *Regexp) SplitStr(str string) []string {
	buf := []byte(str)

	ind := reg.RE.FindAllIndex(buf, -1)

	res := []string{}
	trim := 0
	for _, pos := range ind {
		v := buf[pos[0]:pos[1]]
		m := reg.RE.FindSubmatch(v)

		if trim == 0 {
			res = append(res, string(buf[:pos[0]]))
		} else {
			res = append(res, string(buf[trim:pos[0]]))
		}
		trim = pos[1]

		for i := 1; i <= len(m)-1; i++ {
			res = append(res, string(m[i]))
		}
	}

	res = append(res, string(buf[trim:]))

	return res
}
