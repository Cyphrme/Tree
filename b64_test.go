package tree

import (
	"fmt"

	"github.com/cyphrme/coze"
)

// ExampleSB64 demonstrates using SB64 as a map key and that fmt prints "RFC
// 4648 base 64 URI canonical with padding truncated" properly.
func ExampleSB64() {
	b := coze.MustDecode("zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0")
	b2 := coze.MustDecode("vZIAk8rjcSIKZKokGylCtVoI3DXvFYJn4XNWzf_C_FA")

	lp := make(map[SB64]coze.B64)
	lp[SB64(b)] = coze.B64(b2)

	fmt.Printf("%s\n", SB64(b))
	fmt.Printf("%s\n", lp)
	fmt.Printf("%+v\n", lp)
	fmt.Printf("%#v\n", lp)

	// Output:
	// zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0
	// map[zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0:vZIAk8rjcSIKZKokGylCtVoI3DXvFYJn4XNWzf_C_FA]
	// map[zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0:vZIAk8rjcSIKZKokGylCtVoI3DXvFYJn4XNWzf_C_FA]
	// map[tree.SB64]coze.B64{zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0:vZIAk8rjcSIKZKokGylCtVoI3DXvFYJn4XNWzf_C_FA}
}

// TODO write SB64 unmarshal test.
