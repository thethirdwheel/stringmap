package stringmap

import (
	"fmt"
	"strings"
)

//StringPair is a struct holding two strings, Key and Val.
type StringPair struct {
	Key string
	Val string
}

//StringPairs is a slice of *StringPair.
type StringPairs []*StringPair

//Len returns the lenght of a StringPairs slice.
func (s StringPairs) Len() int { return len(s) }

//Swap exchanges to items in a StringPairs slice.
func (s StringPairs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

//Keys returns a slice of strings containing the Key of each StringPair
func (s StringPairs) Keys() []string {
	keys := []string{}
	for _, o := range s {
		keys = append(keys, o.Key)
	}
	return keys
}

//Vals returns a slice of strings containing the Val of each StringPair
func (s StringPairs) Vals() []string {
	vals := []string{}
	for _, o := range s {
		vals = append(vals, o.Val)
	}
	return vals
}

type StringMap map[string]interface{}

//String writes a StringMap's contents, omitting keys for maps with only one key/value pair.
func (m StringMap) String() string {
	vals := []string{}
	for k, v := range m {
		if len(m) == 1 {
			return fmt.Sprintf("%v", v)
		} else {
			vals = append(vals, fmt.Sprintf("%+v:%+v", k, v))
		}
	}
	return strings.Join(vals, " ")
}

type ByKey struct{ StringPairs }

func (s ByKey) Less(i, j int) bool { return s.StringPairs[i].Key < s.StringPairs[j].Key }

type ByVal struct{ StringPairs }

func (s ByVal) Less(i, j int) bool { return s.StringPairs[i].Val < s.StringPairs[j].Val }


