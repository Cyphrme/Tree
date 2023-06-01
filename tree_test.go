package tree

import (
	"encoding/json"
	"fmt"

	"github.com/cyphrme/coze"
)

func Example_intToBytesLE() {
	// Byte one boundary
	fmt.Printf("0:%b\n", intToBytesLE(0))
	fmt.Printf("1:%b\n", intToBytesLE(1))
	fmt.Printf("2:%b\n", intToBytesLE(2))
	// Byte two boundary
	fmt.Printf("255:%b\n", intToBytesLE(255))
	fmt.Printf("256:%b\n", intToBytesLE(256))
	fmt.Printf("257:%b\n", intToBytesLE(257))
	fmt.Printf("257:%b\n", intToBytesLE(258))
	// Byte three boundary
	fmt.Printf("65535:%b\n", intToBytesLE(65535))
	fmt.Printf("65536:%b\n", intToBytesLE(65536))
	fmt.Printf("65537:%b\n", intToBytesLE(65537))
	// 8 byte boundry
	fmt.Printf("18446744073709551615:%b\n", intToBytesLE(18446744073709551615))

	// Output:
	// 0:[0]
	// 1:[1]
	// 2:[10]
	// 255:[11111111]
	// 256:[0 1]
	// 257:[1 1]
	// 257:[10 1]
	// 65535:[11111111 11111111]
	// 65536:[0 0 1]
	// 65537:[1 0 1]
	// 18446744073709551615:[11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111]
}

func ExampleTree_Populate() {
	t := Tree{
		Alg:        coze.SHA256,
		Seed:       coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		DepthSizes: []int{3},
	}
	t.Populate()
	fmt.Println(t)
	// Output:
	// 	{
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "depth_sizes": [
	//         3
	//     ],
	//     "branches": [
	//         "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
	//         "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg",
	//         "nO49t9F-tf96Os0Eab4bqhnVywHbqJkaGlNgAACoRRI"
	//     ]
	// }
}

// ExampleTree_PathCalc tests calculating paths.  Note that this test cannot use
// print the struct for output testing since maps are unordered.  I'm not sure
// why https://tip.golang.org/doc/go1.12#fmt isn't applying.
func ExampleTree_PathCalc() {
	t := Tree{
		Alg:        coze.SHA256,
		Seed:       coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		DepthSizes: []int{1, 2, 2},
		PathCalc:   true,
	}
	t.Populate()

	fmt.Printf("Paths: %d, Leaves: %d", len(t.Paths), len(t.Leaves))

	// Output:
	// Paths: 7, Leaves: 4
}

func Example_big() {
	t := Tree{
		Alg:        coze.SHA256,
		Seed:       coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		DepthSizes: []int{2, 2},
	}
	err := t.Populate()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("%+v\n", t)
	fmt.Printf("%s\n", t)

	// Output:
	// {
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "depth_sizes": [
	//         2,
	//         2
	//     ],
	//     "branches": [
	//         "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
	//         "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg"
	//     ],
	//     "children": [
	//         {
	//             "alg": "SHA-256",
	//             "seed": "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
	//             "id": "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
	//             "depth_sizes": [
	//                 2
	//             ],
	//             "branches": [
	//                 "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
	//                 "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
	//             ]
	//         },
	//         {
	//             "alg": "SHA-256",
	//             "seed": "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg",
	//             "id": "DG6CXAv4yPV4Ebs1igPW1uZspBO5e5QW9BmY7dJ_lv8",
	//             "depth_sizes": [
	//                 2
	//             ],
	//             "branches": [
	//                 "IVhtm0XLC2K_pp0P6PYESAfQ8qJJuerxVR9DHu5mozA",
	//                 "JokEZgl_qDf7Nkf4z5AQdfeUOBSW0DqbWdRdmE4afl0"
	//             ]
	//         }
	//     ]
	// }

}

func Example_max() {
	max := 3

	t := Tree{
		Alg:            coze.SHA256,
		Seed:           coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		DepthSizes:     []int{2, 2},
		MaxTotalLeaves: &max,
	}
	err := t.Populate()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("%+v\n", t)
	fmt.Printf("%s\n", t)

	// Output:
	// {
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "depth_sizes": [
	//         2,
	//         2
	//     ],
	//     "branches": [
	//         "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
	//         "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg"
	//     ],
	//     "children": [
	//         {
	//             "alg": "SHA-256",
	//             "seed": "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
	//             "id": "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
	//             "depth_sizes": [
	//                 2
	//             ],
	//             "branches": [
	//                 "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
	//                 "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
	//             ]
	//         },
	//         {
	//             "alg": "SHA-256",
	//             "seed": "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg",
	//             "id": "DG6CXAv4yPV4Ebs1igPW1uZspBO5e5QW9BmY7dJ_lv8",
	//             "depth_sizes": [
	//                 2
	//             ],
	//             "branches": [
	//                 "IVhtm0XLC2K_pp0P6PYESAfQ8qJJuerxVR9DHu5mozA",
	//                 ""
	//             ]
	//         }
	//     ]
	// }

}

// ExampleB64Map_jsonMarshal does a round trip Marshal and Unmarshal.  Only one
// key is used for testing since since Go maps are unordered and this test uses
// a static string.
func ExampleB64Map_jsonMarshal() {
	lp := B64Map{
		coze.SB64(coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno")): []coze.B64{
			coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
			coze.MustDecode("zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0")},
	}

	b, err := coze.MarshalPretty(lp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

	// Output:
	// {
	//     "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno": [
	//         "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//         "zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0"
	//     ]
	// }
}

// ExampleB64Map_jsonUnmarshal tests JSON unmarshalling.  Only one key is
// unmarshalled since Go maps are unordered and the test uses a static string.
func ExampleB64Map_jsonUnmarshal() {
	sB64Map := `{
    "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno": [
        "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
        "zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0"
    ]
}`

	lp := new(B64Map)
	err := json.Unmarshal([]byte(sB64Map), lp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", lp)

	// Output:
	// &map[UnBNTTRfbFU2akNqM2FzWkV0SUZ5WXFQakMyTDZtbHVjbDdWR012QXVubw:[RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0]]
}

func ExampleNewTreePopulated_branches() {
	t, _ := NewTreePopulated(coze.SHA256, coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"), []int{3})
	fmt.Println(t.Branches)
	// Output:
	// [JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4 sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg nO49t9F-tf96Os0Eab4bqhnVywHbqJkaGlNgAACoRRI]
}

// ExampleTree_printSB64 ensures SB64 and B64Map prints out correctly.
func ExampleTree_printSB64() {
	t2 := Tree{
		Paths: B64MapP{
			coze.SB64(coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno")): &[]coze.B64{
				coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
				coze.MustDecode("zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0")},
		},
	}
	fmt.Printf("%s\n", t2)

	// Output:
	// {
	//     "alg": "",
	//     "seed": "",
	//     "id": "",
	//     "paths": {
	//         "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno": [
	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//             "zVzgRU3WFpnrlVJAnI4ZU1Od4Agl5Zd4jIP79oubOW0"
	//         ]
	//     }
	// }
}
