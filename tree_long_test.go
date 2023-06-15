package tree

// import (
// 	"fmt"

// 	"github.com/cyphrme/coze"
// )

var tb = []byte(`{
		"alg": "SHA-256",
		"seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
		"id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
		"branch_level_sizes": [
				1,
				2,
				2
		],
		"branches": [
				"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
		],
		"path_calc": true,
		"paths": {
				"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
						"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
						"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
				],
				"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
						"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
						"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
				],
				"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
						"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
						"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
				],
				"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
						"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
						"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
				],
				"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
						"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
				],
				"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
						"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
				],
				"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4": [
						"RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"
				]
		},
		"leaves": [
				"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
				"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0",
				"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
				"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
		],
		"id_paths": {
				"Vq2PbysO0FInGOevXDeU5EHzX9UR5bkPdKs34W2UB4w": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
						"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
						"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
				],
				"i_1NQbnNVgPKIXnqN0639erzFeTT5TNUNDiyO4OH3e8": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
						"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
						"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
				],
				"RSpO8MpHPuFucs2L043_vTEeyIpiNTKaVbqJ6eAmr-A": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
						"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
						"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
				],
				"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
						"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
				],
				"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
						"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
				],
				"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94"
				],
				"4k28iyuwyRykWqPnUjqDPjr7-Z_qNH7eSFuisrpZ24I": [
						"JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
						"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
						"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
				]
		},
		"children": [
				{
						"alg": "SHA-256",
						"seed": "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
						"id": "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
						"branch_level_sizes": [
								2,
								2
						],
						"branches": [
								"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
								"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
						],
						"path_calc": true,
						"paths": {
								"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA": [
										"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
								],
								"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI": [
										"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
										"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
								],
								"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0": [
										"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
										"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
								],
								"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI": [
										"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
										"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
								],
								"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490": [
										"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
										"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
								],
								"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM": [
										"JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
								]
						},
						"leaves": [
								"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
								"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0",
								"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
								"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
						],
						"id_paths": {
								"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc": [
										"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
								],
								"4k28iyuwyRykWqPnUjqDPjr7-Z_qNH7eSFuisrpZ24I": [
										"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
										"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
								],
								"Vq2PbysO0FInGOevXDeU5EHzX9UR5bkPdKs34W2UB4w": [
										"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
										"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
								],
								"i_1NQbnNVgPKIXnqN0639erzFeTT5TNUNDiyO4OH3e8": [
										"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
										"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
								],
								"RSpO8MpHPuFucs2L043_vTEeyIpiNTKaVbqJ6eAmr-A": [
										"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
										"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
								],
								"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs": [
										"FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
								]
						},
						"children": [
								{
										"alg": "SHA-256",
										"seed": "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
										"id": "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs",
										"branch_level_sizes": [
												2
										],
										"branches": [
												"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
												"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0"
										],
										"path_calc": true,
										"paths": {
												"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI": [
														"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
												],
												"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0": [
														"0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
												]
										},
										"leaves": [
												"77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
												"9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0"
										],
										"id_paths": {
												"4k28iyuwyRykWqPnUjqDPjr7-Z_qNH7eSFuisrpZ24I": [
														"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
												],
												"Vq2PbysO0FInGOevXDeU5EHzX9UR5bkPdKs34W2UB4w": [
														"Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
												]
										}
								},
								{
										"alg": "SHA-256",
										"seed": "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA",
										"id": "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc",
										"branch_level_sizes": [
												2
										],
										"branches": [
												"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
												"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
										],
										"path_calc": true,
										"paths": {
												"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI": [
														"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
												],
												"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490": [
														"yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
												]
										},
										"leaves": [
												"IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
												"UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
										],
										"id_paths": {
												"i_1NQbnNVgPKIXnqN0639erzFeTT5TNUNDiyO4OH3e8": [
														"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
												],
												"RSpO8MpHPuFucs2L043_vTEeyIpiNTKaVbqJ6eAmr-A": [
														"tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
												]
										}
								}
						]
				}
		]
	}`)

// // TODO
// // ExampleTree_PathCalc tests calculating paths.  Note that this test cannot use
// // print the struct for output testing since maps are unordered(?!).  I'm not
// // sure why https://tip.golang.org/doc/go1.12#fmt isn't applying.
// func ExampleTree_PathCalc() {
// 	t := Tree{
// 		Alg:        coze.SHA256,
// 		Seed:       coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
// 		BLS: []int{1, 2, 2},
// 		PathCalc:   true,
// 	}
// 	t.Populate()
// 	//fmt.Println(t)

// 	t2 := new(Tree)
// 	err := json.Unmarshal(tb, t2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(reflect.DeepEqual(t, t2))
// 	// Output:
// 	// true
// }

// // ExampleTree_PathCalc tests calculating paths.  Note that this test cannot use
// // print the struct for output testing since maps are unordered(?!).  I'm not
// // sure why https://tip.golang.org/doc/go1.12#fmt isn't applying.
// func ExampleTree_PathCalc2() {
// 	t := Tree{
// 		Alg:        coze.SHA256,
// 		Seed:       coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
// 		BLS: []int{1, 2, 2},
// 		PathCalc:   true,
// 	}
// 	t.Populate()

// 	t2 := new(Tree)

// 	err := json.Unmarshal(tb, t2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(t2)

// 	fmt.Println(reflect.DeepEqual(t, t2))
// 	// Output:
// 	// true
// }

// // // ExampleTree_PathCalc tests calculating paths.  Note that this test cannot use
// // // print the struct for output testing since maps are unordered(?!).  I'm not
// // // sure why https://tip.golang.org/doc/go1.12#fmt isn't applying.
// // func ExampleTree_PathCalc2() {
// // 	t := Tree{
// // 		Alg:        coze.SHA256,
// // 		Seed:       coze.MustDecode("RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"),
// // 		BLS: []int{1, 2, 2},
// // 		PathCalc:   true,
// // 	}
// // 	t.Populate()

// // 	fmt.Printf("%s", t)

// // 	// Output:
// // 	// {
// // 	//     "alg": "SHA-256",
// // 	//     "seed": "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//     "id": "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//     "branch_level_sizes": [
// // 	//         1,
// // 	//         2,
// // 	//         2
// // 	//     ],
// // 	//     "branches": [
// // 	//         "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
// // 	//     ],
// // 	//     "path_calc": true,
// // 	//     "paths": {
// // 	//         "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//             "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//             "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//         ],
// // 	//         "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno"
// // 	//         ],
// // 	//         "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//             "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
// // 	//         ],
// // 	//         "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//             "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
// // 	//         ],
// // 	//         "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//             "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//             "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
// // 	//         ],
// // 	//         "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//             "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//             "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
// // 	//         ],
// // 	//         "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI": [
// // 	//             "RpMM4_lU6jCj3asZEtIFyYqPjC2L6mlucl7VGMvAuno",
// // 	//             "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//             "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//         ]
// // 	//     },
// // 	//     "leaves": [
// // 	//         "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
// // 	//         "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0",
// // 	//         "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
// // 	//         "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
// // 	//     ],
// // 	//     "id_paths": {
// // 	//         "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94"
// // 	//         ],
// // 	//         "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//             "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
// // 	//         ],
// // 	//         "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//             "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
// // 	//         ],
// // 	//         "4k28iyuwyRykWqPnUjqDPjr7-Z_qNH7eSFuisrpZ24I": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//             "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//             "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
// // 	//         ],
// // 	//         "Vq2PbysO0FInGOevXDeU5EHzX9UR5bkPdKs34W2UB4w": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//             "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//             "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
// // 	//         ],
// // 	//         "i_1NQbnNVgPKIXnqN0639erzFeTT5TNUNDiyO4OH3e8": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//             "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//             "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
// // 	//         ],
// // 	//         "RSpO8MpHPuFucs2L043_vTEeyIpiNTKaVbqJ6eAmr-A": [
// // 	//             "JArIzKVsB7CBjQ9zqB-Y1RbDdET2Bi-oyWADoVIUJ94",
// // 	//             "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//             "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
// // 	//         ]
// // 	//     },
// // 	//     "children": [
// // 	//         {
// // 	//             "alg": "SHA-256",
// // 	//             "seed": "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//             "id": "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//             "branch_level_sizes": [
// // 	//                 2,
// // 	//                 2
// // 	//             ],
// // 	//             "branches": [
// // 	//                 "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
// // 	//                 "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//             ],
// // 	//             "path_calc": true,
// // 	//             "paths": {
// // 	//                 "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM": [
// // 	//                     "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
// // 	//                 ],
// // 	//                 "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA": [
// // 	//                     "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4"
// // 	//                 ],
// // 	//                 "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI": [
// // 	//                     "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//                     "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
// // 	//                 ],
// // 	//                 "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0": [
// // 	//                     "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//                     "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
// // 	//                 ],
// // 	//                 "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI": [
// // 	//                     "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//                     "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//                 ],
// // 	//                 "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490": [
// // 	//                     "JmXqHaRJWmRD8qLK6yPyAfH-jiJ9EDJXQsUnAwO5ot4",
// // 	//                     "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//                 ]
// // 	//             },
// // 	//             "leaves": [
// // 	//                 "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
// // 	//                 "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0",
// // 	//                 "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
// // 	//                 "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
// // 	//             ],
// // 	//             "id_paths": {
// // 	//                 "Vq2PbysO0FInGOevXDeU5EHzX9UR5bkPdKs34W2UB4w": [
// // 	//                     "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//                     "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
// // 	//                 ],
// // 	//                 "i_1NQbnNVgPKIXnqN0639erzFeTT5TNUNDiyO4OH3e8": [
// // 	//                     "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//                     "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
// // 	//                 ],
// // 	//                 "RSpO8MpHPuFucs2L043_vTEeyIpiNTKaVbqJ6eAmr-A": [
// // 	//                     "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//                     "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
// // 	//                 ],
// // 	//                 "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs": [
// // 	//                     "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
// // 	//                 ],
// // 	//                 "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc": [
// // 	//                     "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0"
// // 	//                 ],
// // 	//                 "4k28iyuwyRykWqPnUjqDPjr7-Z_qNH7eSFuisrpZ24I": [
// // 	//                     "FBrxbr4wZPHuGyM2Nl7gzoZTfZ-oLQP4pwKBpDxIIb0",
// // 	//                     "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
// // 	//                 ]
// // 	//             },
// // 	//             "children": [
// // 	//                 {
// // 	//                     "alg": "SHA-256",
// // 	//                     "seed": "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM",
// // 	//                     "id": "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs",
// // 	//                     "branch_level_sizes": [
// // 	//                         2
// // 	//                     ],
// // 	//                     "branches": [
// // 	//                         "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
// // 	//                         "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0"
// // 	//                     ],
// // 	//                     "path_calc": true,
// // 	//                     "paths": {
// // 	//                         "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI": [
// // 	//                             "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
// // 	//                         ],
// // 	//                         "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0": [
// // 	//                             "0RJw1SSIXkrW3SID6eoANHq7Y1gB6ZNKo0nuC8pL_qM"
// // 	//                         ]
// // 	//                     },
// // 	//                     "leaves": [
// // 	//                         "77Job9cusXWaUDsy4qpP7qyTwHfP_3fHFGPuZyTaTVI",
// // 	//                         "9xVTD9gwrTFFigX8-fw9wHNoDXIW3yZUh2zesft-gN0"
// // 	//                     ],
// // 	//                     "id_paths": {
// // 	//                         "4k28iyuwyRykWqPnUjqDPjr7-Z_qNH7eSFuisrpZ24I": [
// // 	//                             "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
// // 	//                         ],
// // 	//                         "Vq2PbysO0FInGOevXDeU5EHzX9UR5bkPdKs34W2UB4w": [
// // 	//                             "Nk6_zWOgHW7apqKoSIwZNst_yVryqQ--stm1jCKy6Cs"
// // 	//                         ]
// // 	//                     }
// // 	//                 },
// // 	//                 {
// // 	//                     "alg": "SHA-256",
// // 	//                     "seed": "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA",
// // 	//                     "id": "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc",
// // 	//                     "branch_level_sizes": [
// // 	//                         2
// // 	//                     ],
// // 	//                     "branches": [
// // 	//                         "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
// // 	//                         "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
// // 	//                     ],
// // 	//                     "path_calc": true,
// // 	//                     "paths": {
// // 	//                         "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI": [
// // 	//                             "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//                         ],
// // 	//                         "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490": [
// // 	//                             "yZJ4i_WzjbPHrbczNt2AW-zk1mDJE77ps98jN9fYaFA"
// // 	//                         ]
// // 	//                     },
// // 	//                     "leaves": [
// // 	//                         "IegtglAF2dvixDP5ApZsQ5AWHNYbEt577whrEZp6mnI",
// // 	//                         "UzCJNRuN5SD84w8x7mfvriZNkltv6mI5A5i4r5q2490"
// // 	//                     ],
// // 	//                     "id_paths": {
// // 	//                         "RSpO8MpHPuFucs2L043_vTEeyIpiNTKaVbqJ6eAmr-A": [
// // 	//                             "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
// // 	//                         ],
// // 	//                         "i_1NQbnNVgPKIXnqN0639erzFeTT5TNUNDiyO4OH3e8": [
// // 	//                             "tsdM7o8IiUvtp-S3pmKzg-9iH18C8f9vFuHgmlDFJrc"
// // 	//                         ]
// // 	//                     }
// // 	//                 }
// // 	//             ]
// // 	//         }
// // 	//     ]
// // 	// }
// // }
