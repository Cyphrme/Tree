package tree

import (
	"encoding/json"
	"fmt"

	"github.com/cyphrme/coze"
)

func Example_IntToBytesLE() {
	// Byte one boundary
	fmt.Printf("0:%b\n", IntToBytesLE(0))
	fmt.Printf("1:%b\n", IntToBytesLE(1))
	fmt.Printf("2:%b\n", IntToBytesLE(2))
	// Byte two boundary
	fmt.Printf("255:%b\n", IntToBytesLE(255))
	fmt.Printf("256:%b\n", IntToBytesLE(256))
	fmt.Printf("257:%b\n", IntToBytesLE(257))
	// Byte three boundary
	fmt.Printf("65535:%b\n", IntToBytesLE(65535))
	fmt.Printf("65536:%b\n", IntToBytesLE(65536))
	fmt.Printf("65537:%b\n", IntToBytesLE(65537))
	// 8 byte boundary
	fmt.Printf("18446744073709551615:%b\n", IntToBytesLE(18446744073709551615))

	// Output:
	// 0:[0]
	// 1:[1]
	// 2:[10]
	// 255:[11111111]
	// 256:[0 1]
	// 257:[1 1]
	// 65535:[11111111 11111111]
	// 65536:[0 0 1]
	// 65537:[1 0 1]
	// 18446744073709551615:[11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111]
}

func Example_intToBytesBE() {
	// Byte one boundary
	fmt.Printf("0:%b\n", IntToBytesBE(0))
	fmt.Printf("1:%b\n", IntToBytesBE(1))
	fmt.Printf("2:%b\n", IntToBytesBE(2))
	// Byte two boundary
	fmt.Printf("255:%b\n", IntToBytesBE(255))
	fmt.Printf("256:%b\n", IntToBytesBE(256))
	fmt.Printf("257:%b\n", IntToBytesBE(257))
	// Byte three boundary
	fmt.Printf("65535:%b\n", IntToBytesBE(65535))
	fmt.Printf("65536:%b\n", IntToBytesBE(65536))
	fmt.Printf("65537:%b\n", IntToBytesBE(65537))
	// 8 byte boundary
	fmt.Printf("18446744073709551615:%b\n", IntToBytesBE(18446744073709551615))

	// Output:
	// 0:[0]
	// 1:[1]
	// 2:[10]
	// 255:[11111111]
	// 256:[1 0]
	// 257:[1 1]
	// 65535:[11111111 11111111]
	// 65536:[1 0 0]
	// 65537:[1 0 1]
	// 18446744073709551615:[11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111]
}

func ExampleTree_Populate() {
	t := Tree{
		Alg:  coze.SHA256,
		Seed: coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:  []int{3},
	}
	t.Populate()
	fmt.Println(t)
	// Output:
	// 	{
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "branch_level_sizes": [
	//         3
	//     ],
	//     "branches": [
	//         "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
	//         "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg",
	//         "nO49t9F-tf96Os0Eab4bqhnVywHbqJkaGlNgAACoRRI"
	//     ]
	// }
}

// ExampleTree_Skip tests skipping in the tree. With a skip of 1, the branch
// value should match ExampleTree_Populate starting at the second value.
func ExampleTree_Skip() {
	t := Tree{
		Alg:  coze.SHA256,
		Seed: coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:  []int{3},
		Skip: 1,
	}
	t.Populate()
	fmt.Println(t)

	// Output:
	// {
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "branch_level_sizes": [
	//         3
	//     ],
	//     "skip": 1,
	//     "branches": [
	//         "sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg",
	//         "nO49t9F-tf96Os0Eab4bqhnVywHbqJkaGlNgAACoRRI"
	//     ]
	// }
}

// ExampleTree_Populate_zero tests a zero sized tree.
func ExampleTree_Populate_zero() {
	t := Tree{
		Alg:      coze.SHA256,
		Seed:     coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:      []int{0},
		PathCalc: true,
	}
	t.Populate()
	fmt.Println(t)
	// Output:
	// {
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "branch_level_sizes": [
	//         0
	//     ],
	//     "path_calc": true
	// }

}

// Tests the nonce 255/256, one byte to two bytes, boundary
func ExampleTree_Populate_nonceBoundary() {
	t := Tree{
		Alg:  coze.SHA256,
		Seed: coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:  []int{257},
		Skip: 255,
	}
	t.Populate()
	fmt.Println(t)
	// Output:
	// {
	//     "alg": "SHA-256",
	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
	//     "branch_level_sizes": [
	//         257
	//     ],
	//     "skip": 255,
	//     "branches": [
	//         "1ILX9r1f7c3hK3Gwv_i5rmV-T7E4Sq0vjJsJE1J7SwM",
	//         "DLbAu75AbsA7GV2cxe_KRJiJvdOu3Jkbxt-Xbow5-LU"
	//     ]
	// }
}

// ExampleTree_PathCalc_len tests calculating paths.
func ExampleTree_PathCalc_len() {
	t := Tree{
		Alg:      coze.SHA256,
		Seed:     coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:      []int{1, 2, 2},
		PathCalc: true,
	}
	t.Populate()
	fmt.Printf("Paths: %d, PathsID %d, Leaves: %d, LeavesID: %d\n", len(t.Paths), len(t.PathsID), len(t.Leaves), len(t.LeavesID))
	t = Tree{
		Alg:      coze.SHA256,
		Seed:     coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:      []int{2},
		PathCalc: true,
	}
	t.Populate()
	fmt.Printf("Paths: %d, PathsID %d, Leaves: %d, LeavesID: %d\n", len(t.Paths), len(t.PathsID), len(t.Leaves), len(t.LeavesID))
	//fmt.Println(t)
	// Output:
	// Paths: 7, PathsID 7, Leaves: 4, LeavesID: 4
	// Paths: 2, PathsID 2, Leaves: 2, LeavesID: 2
}

func Example_big() {
	t := Tree{
		Alg:  coze.SHA256,
		Seed: coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
		BLS:  []int{2, 2},
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
	//     "branch_level_sizes": [
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
	//             "branch_level_sizes": [
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
	//             "branch_level_sizes": [
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
		BLS:            []int{2, 2},
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
	//     "branch_level_sizes": [
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
	//             "branch_level_sizes": [
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
	//             "branch_level_sizes": [
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
		SB64(coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno")): []coze.B64{
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

// ExampleB64Map_jsonUnmarshal tests B64Map JSON unmarshalling.  Only one key is
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

// ExampleTree_jsonUnmarshal tests tree JSON unmarshalling.
func ExampleTree_jsonUnmarshal() {
	// Generate the following using:
	// ts, _ := NewTreePopulated(coze.SHA256, coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"), []int{1, 2})
	tb := []byte(`{
    "alg": "SHA-256",
    "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
    "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
    "branch_level_sizes": [
        1,
        2
    ],
    "branches": [
        "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
    ],
    "children": [
        {
            "alg": "SHA-256",
            "seed": "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
            "id": "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
            "branch_level_sizes": [
                2
            ],
            "branches": [
                "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
                "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
            ]
        }
    ]
}`)
	t := new(Tree)
	err := json.Unmarshal(tb, t)
	if err != nil {
		panic(err)
	}
	ts := fmt.Sprintf("%s", t)
	fmt.Println(string(tb) == ts)

	// Output:
	//true
}

func ExampleNewTreePopulated_branches() {
	t, _ := NewTreePopulated(coze.SHA256, coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"), []int{3})
	fmt.Println(t.Branches)
	// Output:
	// [JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4 sGYg2-Wl8uE4OSaYWk2v7pBj7pYNHrl0K-pxWG6_YDg nO49t9F-tf96Os0Eab4bqhnVywHbqJkaGlNgAACoRRI]
}

// ExampleTree_printSB64 ensures SB64 and B64Map/B64MapP prints out correctly.
func ExampleTree_printSB64() {
	t2 := Tree{
		Paths: B64MapP{
			SB64(coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno")): &[]coze.B64{
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
