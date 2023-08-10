package tree

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/cyphrme/coze"
)

// SB64 is useful for B64 map keys. Idiomatically, map key type should be `B64`,
// but currently in Go map keys are only type `string`, not `[]byte`.  Since
// B64's underlying type is `[]byte` it cannot be used as a map key. See
// https://github.com/golang/go/issues/283 and
// https://github.com/google/go-cmp/issues/67.  SB64 will be deprecated if/when
// Go supports []byte keys.
//
// This is an acceptable hack because (from https://go.dev/blog/strings)
//
//	>[A] string holds arbitrary bytes. It is not required to hold Unicode text,
//	> UTF-8 text, or any other predefined format. As far as the content of a
//	> string is concerned, it is exactly equivalent to a slice of bytes.
type SB64 string

// String implements fmt.Stringer
func (b SB64) String() string {
	return coze.B64(b).String()
}

// GoString implements fmt.GoString
func (b SB64) GoString() string {
	return b.String()
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *SB64) UnmarshalJSON(b []byte) error {
	// JSON.Unmarshal returns b encapsulated in quotes which is invalid base64 characters.
	bb, err := base64.URLEncoding.Strict().WithPadding(base64.NoPadding).DecodeString(strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	b64 := coze.B64(bb)
	*t = SB64(b64.String())

	return nil
}

// B64Map is for maps with B64 as keys.  See notes on SB64.  Without this,
// defining struct values as `map[SB64][]B64` will undesirably print map keys as
// []byte and not B64 since json.Marshal looks for the key's primitive type if a
// defined custom type for the top level struct type is not present.
type B64Map map[SB64][]coze.B64

// MarshalJSON implements json.Marshaler.
func (t B64Map) MarshalJSON() ([]byte, error) {
	i := 0
	l := len(t)
	s := "{"
	for k, v := range t {
		i++
		vj, err := coze.Marshal(v)
		if err != nil {
			return nil, err
		}
		s += `"` + fmt.Sprint(k) + `":` + string(vj)
		if i != l {
			s += ","
		}
	}
	return []byte(s + "}"), nil
}

// TODO
// func (t B64Map) UnmarshalJSON(b []byte) error {
// 	fmt.Println(b)
// 	// JSON.Unmarshal returns b encapsulated in quotes which is invalid base64 characters.
// 	s, err := base64.URLEncoding.Strict().WithPadding(base64.NoPadding).DecodeString(strings.Trim(string(b), "\""))
// 	if err != nil {
// 		return err
// 	}
// 	*t = B64(s)
// 	return nil
// }

// B64Map is for maps with B64 as keys.  See notes on SB64.  Without this,
// defining struct values as `map[tree.SB64][]coze.B64` will undesirably print
// map keys as []byte and not B64 since json.Marshal looks for the key's
// primitive type if a defined custom type for the top level struct type is not
// present.
type B64MapP map[SB64]*[]coze.B64

// MarshalJSON implements json.Marshaler.
func (t B64MapP) MarshalJSON() ([]byte, error) {
	var s = "{"
	var i = 0
	var l = len(t)
	for k, v := range t {
		i++
		vj, err := coze.Marshal(v)
		if err != nil {
			return nil, err
		}
		s += "\"" + fmt.Sprint(k) + "\":" + string(vj)
		if i != l {
			s += ","
		}
	}
	return []byte(s + "}"), nil
}
