package oatenc

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/tusklang/tools"

	. "github.com/tusklang/tusk/lang/types"
)

//export OatEncode
func OatEncode(filename string, data map[string]*TuskType) error {

	f, e := os.Create(filename)

	if e != nil {
		return e
	}

	//versioning and magic #
	fmt.Fprint(f, MAGIC)
	fmt.Fprintf(f, "%d.%d.%d\n", tools.TuskMajor, tools.TuskMinor, tools.TuskBug)
	////////////////////////

	for k, v := range data {

		var name = EncodeStr([]rune(k))
		var nameinter = make([]interface{}, len(name))
		var format = ""
		for k, v := range name {
			format += "%c"
			nameinter[k] = v
		}

		fmt.Fprintf(f, format, nameinter...)
		fmt.Fprintf(f, "%c", reserved["set global"])

		var encoded = EncodeValue(*v)

		for _, v := range encoded {
			fmt.Fprintf(f, "%c", v)
		}

		fmt.Fprintf(f, "%c", reserved["new global"])
	}

	return nil
}

func EncodeValue(v TuskType) []rune {

	var final []rune

	switch v.(type) {

	case TuskArray:

		final = append(final, reserved["make c-array"])

		for _, v := range v.(TuskArray).Array {
			final = append(final, EncodeValue(*v)...)
			final = append(final, reserved["value seperator"])
		}

	case TuskBool:

		final = append(final, reserved["make bool"], reserved["escaper"])
		if v.(TuskBool).ToGoType() {
			final = append(final, 1)
		} else {
			final = append(final, 0)
		}

	case TuskFunc:

		final = append(final, reserved["start function"])

		for _, v := range v.(TuskFunc).Overloads {
			for k := range v.Params {
				final = append(final, EncodeStr([]rune(v.Types[k]))...)
				final = append(final, reserved["seperate type-param"])
				final = append(final, EncodeStr([]rune(v.Params[k]))...)
				final = append(final, reserved["value seperator"])
			}
			final = append(final, reserved["param body split"])
			final = append(final, EncodeActions(v.Body)...)

			final = append(final, reserved["seperate overload"])
		}

		final = append(final, reserved["end function"])

	case TuskHash:

		final = append(final, reserved["make c-hash"])

		for k, v := range v.(TuskHash).Hash {
			final = append(final, EncodeStr([]rune(k))...)
			final = append(final, reserved["hash key seperator"])
			final = append(final, EncodeValue(*v)...)
			final = append(final, reserved["value seperator"])
		}

	case TuskNumber:

		final = append(final, reserved["start number"])

		if v.(TuskNumber).Integer != nil && len(*v.(TuskNumber).Integer) != 0 {
			for _, v := range *v.(TuskNumber).Integer {
				final = append(final, reserved["escaper"], rune(v))
			}
		}

		final = append(final, reserved["decimal spot"])

		if v.(TuskNumber).Decimal != nil && len(*v.(TuskNumber).Decimal) != 0 {
			for _, v := range *v.(TuskNumber).Decimal {
				final = append(final, reserved["escaper"], rune(v))
			}
		}

		final = append(final, reserved["end number"])

	case TuskProto:

		final = append(final, reserved["start proto"])

		//put the name
		final = append(final, EncodeStr([]rune(v.(TuskProto).ProtoName))...)
		final = append(final, reserved["seperate proto name"])
		//////////////

		for k, v := range v.(TuskProto).Static {
			final = append(final, EncodeStr([]rune(k))...)
			final = append(final, reserved["hash key seperator"])
			final = append(final, EncodeValue(*v)...)
			final = append(final, reserved["value seperator"])
		}

		final = append(final, reserved["seperate proto static instance"])

		for k, v := range v.(TuskProto).Instance {
			final = append(final, EncodeStr([]rune(k))...)
			final = append(final, reserved["hash key seperator"])
			final = append(final, EncodeValue(*v)...)
			final = append(final, reserved["value seperator"])
		}

		final = append(final, reserved["seperate proto static instance"])

		final = append(final, reserved["end proto"])

	case TuskRune:

		final = append(final, reserved["make rune"], reserved["escaper"])
		final = append(final, v.(TuskRune).ToGoType())

	case TuskString:

		final = append(final, reserved["make string"])
		final = append(final, v.(TuskString).ToRuneList()...)

	case TuskUndef:

		final = append(final, reserved["make undef"])

	}

	return final
}

func EncodeActions(data []Action) []rune {

	var final []rune

	for _, v := range data {
		fieldt := reflect.TypeOf(v)

		for i := 0; i < fieldt.NumField(); i++ {

			switch fieldt.Field(i).Name {

			case "File":

				final = append(final, EncodeStr([]rune(v.File))...)

			case "Line":

				final = append(final, reserved["escaper"], rune(v.Line))

			case "Type":

				final = append(final, reserved[v.Type])

			case "Name":

				final = append(final, EncodeStr([]rune(v.Name))...)

			case "Value":

				final = append(final, EncodeValue(v.Value)...)

			case "ExpAct":

				if len(v.ExpAct) != 0 {
					final = append(final, reserved["start multi action"])
					final = append(final, EncodeActions(v.ExpAct)...)
					final = append(final, reserved["end multi action"])
				}

			case "First":

				if len(v.First) != 0 {
					final = append(final, reserved["start multi action"])
					final = append(final, EncodeActions(v.First)...)
					final = append(final, reserved["end multi action"])
				}

			case "Second":

				if len(v.Second) != 0 {
					final = append(final, reserved["start multi action"])
					final = append(final, EncodeActions(v.Second)...)
					final = append(final, reserved["end multi action"])
				}

			case "Array":

				final = append(final, reserved["start r-array"])

				for _, v := range v.Array {
					final = append(final, EncodeActions(v)...)
					final = append(final, reserved["value seperator"])
				}

				final = append(final, reserved["end r-array"])

			case "Hash":

				final = append(final, reserved["start r-hash"])

				for _, v := range v.Hash {
					final = append(final, reserved["start multi action"])
					final = append(final, EncodeActions(v[0])...)
					final = append(final, reserved["end multi action"])
					final = append(final, reserved["hash key seperator"])
					final = append(final, EncodeActions(v[1])...)
					final = append(final, reserved["value seperator"])
				}

				final = append(final, reserved["end r-hash"])

			}

			final = append(final, reserved["seperate "+strings.ToLower(fieldt.Field(i).Name)])

		}

		final = append(final, reserved["next action"])
	}

	return final
}

//export EncodeStr
func EncodeStr(splitted []rune) []rune {
	var slc []rune
	for _, v := range splitted {

		for _, vv := range reserved {
			if vv == v {
				slc = append(slc, reserved["escaper"])
				break
			}
		}

		slc = append(slc, v)
	}
	return slc
}
