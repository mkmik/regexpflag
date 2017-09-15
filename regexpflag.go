/*
Package regexpflag implements a custom "regexp" flag

Usage:

	var (
		myFlag = regexpflag.Flag("myflag", "a.*b", "a regular expression that matches some stuff")
	)

	func main() {
		flag.Parse()

		fmt.Println(myFlag.MatchString("ab"))
	}

If you invoke it with:

	--myflag 'a.*c'

myFlag will contain a *regexp.Regexp
*/
package regexpflag

import (
	"flag"
	"regexp"
)

// RegexpValue implements flag.Value
type regexpValue regexp.Regexp

// Set implements the flag.Value interface
func (v *regexpValue) Set(arg string) error {
	r, err := regexp.Compile(arg)
	if err != nil {
		return err
	}
	*(*regexp.Regexp)(v) = *r.Copy()
	return nil
}

func (v *regexpValue) String() string {
	return (*regexp.Regexp)(v).String()
}

// Flag returns a regexp flag. Panics if the default value doesn't compile.
func Flag(name string, def string, usage string) *regexp.Regexp {
	r := regexp.MustCompile(def)
	flag.Var((*regexpValue)(r), name, usage)
	return r
}
