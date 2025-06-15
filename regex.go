package regex

import (
	"regexp"
	"strconv"
	"time"

	"github.com/tkdeng/regex/common"
)

type Regexp struct {
	RE  *regexp.Regexp
	len int64
}

var reg_comment = regexp.MustCompile(`\(\?#.*?\)`)
var reg_param_quote = regexp.MustCompile(`\%([0-9]|\{[0-9]+\})|\\[\\']`)
var reg_escapeParam = regexp.MustCompile(`[\\%]`)

var cache common.CacheMap[*Regexp] = common.NewCache[*Regexp]()

func init() {
	go func() {
		for {
			time.Sleep(10 * time.Minute)

			// default: remove cache items have not been accessed in over 2 hours
			cacheTime := 2 * time.Hour

			// SysFreeMemory returns the total free system memory in megabytes
			mb := common.SysFreeMemory()
			if mb < 200 && mb != 0 {
				// low memory: remove cache items have not been accessed in over 10 minutes
				cacheTime = 10 * time.Minute
			} else if mb < 500 && mb != 0 {
				// low memory: remove cache items have not been accessed in over 30 minutes
				cacheTime = 30 * time.Minute
			} else if mb < 2000 && mb != 0 {
				// low memory: remove cache items have not been accessed in over 1 hour
				cacheTime = 1 * time.Hour
			} else if mb > 64000 {
				// high memory: remove cache items have not been accessed in over 12 hour
				cacheTime = 12 * time.Hour
			} else if mb > 32000 {
				// high memory: remove cache items have not been accessed in over 6 hour
				cacheTime = 6 * time.Hour
			} else if mb > 16000 {
				// high memory: remove cache items have not been accessed in over 3 hour
				cacheTime = 3 * time.Hour
			}

			cache.DelOld(cacheTime)

			time.Sleep(10 * time.Second)

			// clear cache if were still critically low on available memory
			if mb := common.SysFreeMemory(); mb < 10 && mb != 0 {
				cache.DelOld(0)
			}
		}
	}()
}

//* regex compile methods

// this method compiles the RE string to add more functionality to it
func compRE(re string, params []string) string {
	reB := []byte(re)

	reB = reg_comment.ReplaceAllLiteral(reB, []byte{})

	reB = reg_param_quote.ReplaceAllFunc(reB, func(b []byte) []byte {
		if b[0] == '\\' {
			if b[1] == '\'' {
				return []byte{'`'}
			}
			return b
		}

		if b[0] == '%' {
			b = b[1:]
			if len(b) > 1 {
				b = b[1 : len(b)-1]
			}

			if i, err := strconv.Atoi(string(b)); err == nil {
				if i < 1 || i > len(params) {
					return []byte{}
				}

				return []byte(regexp.QuoteMeta(params[i-1]))
			}
		}

		return []byte{}
	})

	return string(reB)
}

// Comp compiles a regular expression and store it in the cache
func Comp(re string, params ...string) *Regexp {
	re = compRE(re, params)

	if val, err := cache.Get(re); val != nil || err != nil {
		if err != nil {
			panic(err)
		}

		return val
	}

	reg := regexp.MustCompile(re)

	compReg := Regexp{RE: reg, len: int64(len(re))}

	cache.Set(re, &compReg, nil)
	return &compReg
}

// CompTry tries to compile or returns an error
func CompTry(re string, params ...string) (*Regexp, error) {
	re = compRE(re, params)

	if val, err := cache.Get(re); val != nil || err != nil {
		if err != nil {
			return &Regexp{}, err
		}

		return val, nil
	}

	reg, err := regexp.Compile(re)
	if err != nil {
		cache.Set(re, nil, err)
		return &Regexp{}, err
	}

	compReg := Regexp{RE: reg, len: int64(len(re))}

	cache.Set(re, &compReg, nil)
	return &compReg, nil
}

//* other regex methods

// Escape will escape regex special chars
func Escape(re string) string {
	re = regexp.QuoteMeta(re)
	re = reg_escapeParam.ReplaceAllString(re, `\$0`)
	return re
}

// IsValid will return true if a regex is valid and can be compiled by this module
func IsValid(re string) bool {
	re = compRE(re, []string{})
	if _, err := regexp.Compile(re); err == nil {
		return true
	}
	return false
}

// IsValidRE2 will return true if a regex is valid and can be compiled by the builtin RE2 module
func IsValidRE2(re string) bool {
	if _, err := regexp.Compile(re); err == nil {
		return true
	}
	return false
}

// JoinBytes is an easy way to join multiple values into a single []byte
func JoinBytes(bytes ...any) []byte {
	return common.JoinBytes(bytes...)
}
