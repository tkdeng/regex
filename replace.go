package regex

// RepFunc replaces a string with the result of a function
//
// similar to JavaScript .replace(/re/, function(b){})
func (reg *Regexp) RepFunc[T ~string | ~[]byte](buf T, rep func(b func(int) []byte) []byte) T {
	return T(reg.RE.ReplaceAllFunc([]byte(buf), func(b []byte) []byte {
		m := reg.RE.FindSubmatch(b)

		r := rep(func(g int) []byte {
			if g < 0 || g >= len(m) {
				return []byte{}
			}
			return m[g]
		})

		if r == nil {
			return []byte{}
		}
		return r
	}))
}

// RepFuncBreak replaces a string with the result of a function
// and gives you the option to break the loop
//
// similar to JavaScript .replace(/re/, function(b){})
//
// return true to continue loop, false to break loop
func (reg *Regexp) RepFuncBreak[T ~string | ~[]byte](buf T, rep func(b func(int) []byte) ([]byte, bool)) T {
	stop := false
	return T(reg.RE.ReplaceAllFunc([]byte(buf), func(b []byte) []byte {
		if stop {
			return b
		}

		m := reg.RE.FindSubmatch(b)

		r, next := rep(func(g int) []byte {
			if g < 0 || g >= len(m) {
				return []byte{}
			}
			return m[g]
		})

		if !next {
			stop = true
		}

		if r == nil {
			return []byte{}
		}
		return r
	}))
}

// Rep replaces a string with another string
//
// this function will replace things in the result like $1 with your capture groups
//
// use $0 to use the full regex capture group
//
// use ${123} to use numbers with more than one digit
func (reg *Regexp) Rep[T ~string | ~[]byte](buf T, rep T) T {
	return T(reg.RE.ReplaceAll([]byte(buf), []byte(rep)))
}

// RepLit replaces a string with another string literal
//
// note: this function does not accept replacements like $1
func (reg *Regexp) RepLit[T ~string | ~[]byte](buf T, rep T) T {
	return T(reg.RE.ReplaceAllLiteral([]byte(buf), []byte(rep)))
}
