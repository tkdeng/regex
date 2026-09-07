package regex

// Match returns true if a []byte matches a regex
func (reg *Regexp) Match[T ~string | ~[]byte](buf T) bool {
	return reg.RE.Match([]byte(buf))
}

// Split splits a string, and keeps capture groups
//
// Similar to JavaScript .split(/re/)
func (reg *Regexp) Split[T ~string | ~[]byte](buf T) []T {
	b := []byte(buf)
	
	ind := reg.RE.FindAllIndex(b, -1)

	

	res := []T{}
	trim := 0
	for _, pos := range ind {
		v := b[pos[0]:pos[1]]
		m := reg.RE.FindSubmatch(v)

		if trim == 0 {
			res = append(res, T(b[:pos[0]]))
		} else {
			res = append(res, T(b[trim:pos[0]]))
		}
		trim = pos[1]

		for i := 1; i <= len(m)-1; i++ {
			res = append(res, T(m[i]))
		}
	}

	res = append(res, T(b[trim:]))

	return res
}
