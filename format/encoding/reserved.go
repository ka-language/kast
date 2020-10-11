package oatenc

//MAGIC is the oat magic number
const MAGIC = "OAT BETA"

var reserved = map[string]rune{
	"next action":                    'అ',
	"escaper":                        0,
	"var":                            1,
	"declare":                        362,
	"log":                            2,
	"print":                          3,
	"if":                             4,
	"elif":                           5,
	"else":                           6,
	"try":                            363,
	"catch":                          364,
	"condition":                      7,
	"while":                          8,
	"each":                           9,
	"function":                       10,
	"return":                         11,
	"await":                          12,
	"proto":                          13,
	"ovld":                           14,
	"let":                            15,
	"cast":                           16,
	"::":                             17,
	":":                              360,
	"?":                              361,
	"+":                              18,
	"-":                              19,
	"*":                              20,
	"/":                              21,
	"%":                              22,
	"^":                              23,
	"==":                             24,
	"!=":                             25,
	">":                              26,
	"<":                              27,
	">=":                             28,
	"<=":                             29,
	"!":                              30,
	"&":                              31,
	"|":                              32,
	"<-":                             290,
	"<~":                             291,
	"++":                             292,
	"--":                             293,
	"+=":                             294,
	"-=":                             295,
	"*=":                             296,
	"/=":                             297,
	"%=":                             298,
	"^=":                             299,
	"break":                          300,
	"continue":                       301,
	"{":                              302,
	"(":                              303,
	"c-hash":                         304,
	"r-hash":                         305,
	"c-array":                        306,
	"r-array":                        307,
	"string":                         308,
	"rune":                           309,
	"bool":                           310,
	"undef":                          311,
	"number":                         312,
	"variable":                       313,
	"varname start":                  314,
	"start multi action":             315,
	"end multi action":               316,
	"hash key seperator":             317,
	"value seperator":                318,
	"make bool":                      319,
	"make undef":                     320,
	"make rune":                      321,
	"make string":                    322,
	"start number":                   323,
	"end number":                     324,
	"decimal spot":                   325,
	"make c-array":                   326,
	"make c-hash":                    327,
	"start proto":                    328,
	"end proto":                      329,
	"seperate proto name":            330,
	"seperate proto static instance": 335,
	"start function":                 336,
	"end function":                   337,
	"seperate type-param":            338,
	"start params":                   339,
	"new global":                     340,
	"set global":                     341,
	"start r-hash":                   342, //how's "life"
	"end r-hash":                     343,
	"start r-array":                  344,
	"end r-array":                    345,
	"param body split":               346,
	"seperate file":                  347,
	"seperate line":                  348,
	"seperate type":                  349,
	"seperate name":                  350,
	"seperate value":                 351,
	"seperate expact":                352,
	"seperate first":                 353,
	"seperate second":                354,
	"seperate array":                 355,
	"seperate hash":                  356,
	"seperate overload":              357,
	"sub value seperator":            359,
}
