package regex

// RepFunc replaces a string with the result of a function
//
// similar to JavaScript .replace(/re/, function(b){})
func (reg *Regexp) RepFunc(buf []byte, rep func(b func(int) []byte) []byte) []byte {
	stop := false
	return reg.RE.ReplaceAllFunc(buf, func(b []byte) []byte {
		if stop {
			return b
		}

		m := reg.RE.FindSubmatch(b)

		r := rep(func(g int) []byte {
			if g < 0 || g >= len(m) {
				return []byte{}
			}
			return m[g]
		})

		if r == nil {
			stop = true
		}

		return r
	})
}

// Rep replaces a string with another string
//
// this function will replace things in the result like $1 with your capture groups
//
// use $0 to use the full regex capture group
//
// use ${123} to use numbers with more than one digit
func (reg *Regexp) Rep(buf []byte, rep []byte) []byte {
	return reg.RE.ReplaceAll(buf, rep)
}

// RepLit replaces a string with another string literal
//
// note: this function does not accept replacements like $1
func (reg *Regexp) RepLit(str []byte, rep []byte) []byte {
	return reg.RE.ReplaceAllLiteral(str, rep)
}
