package regex

// RepFuncStr replaces a string with the result of a function
//
// similar to JavaScript .replace(/re/, function(b){})
func (reg *Regexp) RepFuncStr(buf string, rep func(b func(int) []byte) []byte) string {
	return string(reg.RE.ReplaceAllFunc([]byte(buf), func(b []byte) []byte {
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

// RepFuncBreakStr replaces a string with the result of a function
// and gives you the option to break the loop
//
// similar to JavaScript .replace(/re/, function(b){})
//
// return true to continue loop, false to break loop
func (reg *Regexp) RepFuncBreakStr(buf string, rep func(b func(int) []byte) ([]byte, bool)) string {
	stop := false
	return string(reg.RE.ReplaceAllFunc([]byte(buf), func(b []byte) []byte {
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

// RepStr replaces a string with another string
//
// this function will replace things in the result like $1 with your capture groups
//
// use $0 to use the full regex capture group
//
// use ${123} to use numbers with more than one digit
func (reg *Regexp) RepStr(buf string, rep string) string {
	return reg.RE.ReplaceAllString(buf, rep)
}

// RepLitStr replaces a string with another string literal
//
// note: this function does not accept replacements like $1
func (reg *Regexp) RepLitStr(buf string, rep string) string {
	return reg.RE.ReplaceAllLiteralString(buf, rep)
}
