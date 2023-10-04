package maps

import (
	"fmt"
	"reflect"
)

const (
	sep  string = "----"
	ind  string = "    |"
	tSep string = "****"
)

// print a map
func Print[M ~map[K]V, K comparable, V any](m *M) {
	fmt.Printf("%v", *m)

	fmt.Printf("\n%s\nType:   %T\nLength: %d\n%s\n", tSep, *m, len(*m), tSep)
}

// print a map with a newline
func Println[M ~map[K]V, K comparable, V any](m *M) {
	if IsEmpty(m) {
		fmt.Println(sep)
		fmt.Println("{}")
		fmt.Printf("***\nType: %T\nLength: %d\n", *m, len(*m))
		fmt.Println(sep)
		return
	}

	fmt.Println(sep)
	for k, v := range *m {
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Map:
			fmt.Printf("%v ", k)
			temp := reflect.ValueOf(v)
			n := make(map[interface{}]interface{})
			i := true
			var kType reflect.Kind
			var vType reflect.Kind
			for _, k := range temp.MapKeys() {
				if i {
					i = false
					kType = k.Kind()
					vType = temp.MapIndex(k).Kind()
				}
				n[k] = temp.MapIndex(k).Interface()
			}
			mPrintln(&n, kType, vType, "")
		case reflect.Slice:
			fmt.Printf("%v ", k)
			temp := reflect.ValueOf(v)
			n := make([]interface{}, temp.Len())
			var eType reflect.Kind
			for i := 0; i < temp.Len(); i++ {
				if i == 0 {
					eType = temp.Index(i).Kind()
				}
				n[i] = temp.Index(i).Interface()
			}
			sPrintln(&n, eType, "")
		default:
			fmt.Printf("%v: %#v\n", k, v)
		}
	}

	fmt.Printf("\n***\nType: %T\nLength: %d\n", *m, len(*m))
	fmt.Println(sep)

}

func mPrintln(m *map[interface{}]interface{}, kType reflect.Kind, vType reflect.Kind, indent string) {
	if len(indent) > 1000 {
		return
	}
	indent += ind
	var kTypeStr string
	var vTypeStr string
	if kType == reflect.Interface {
		kTypeStr = "interface{}"
	} else {
		kTypeStr = kType.String()
	}
	if vType == reflect.Interface {
		vTypeStr = "interface{}"
	} else {
		vTypeStr = vType.String()
	}

	length := len(*m)

	if length == 0 {
		fmt.Printf(">> empty  0\n")
		return
	}

	fmt.Printf(">> map[%s]%s  %d\n",
		kTypeStr, vTypeStr, length)

	for k, v := range *m {
		t := reflect.TypeOf(v)
		// fmt.Println(t.Kind().String())
		switch t.Kind() {
		case reflect.Map:
			fmt.Printf("%v%v ", indent, k)
			temp := reflect.ValueOf(v)
			n := make(map[interface{}]interface{})
			i := true
			var kType reflect.Kind
			var vType reflect.Kind
			for _, k := range temp.MapKeys() {
				if i {
					i = false
					kType = k.Kind()
					vType = temp.MapIndex(k).Kind()
				}
				n[k] = temp.MapIndex(k).Interface()
			}
			mPrintln(&n, kType, vType, indent)
		case reflect.Slice:
			fmt.Printf("%v%v ", indent, k)
			temp := reflect.ValueOf(v)
			n := make([]interface{}, temp.Len())
			var eType reflect.Kind
			for i := 0; i < temp.Len(); i++ {
				if i == 0 {
					eType = temp.Index(i).Kind()
				}
				n[i] = temp.Index(i).Interface()
			}
			sPrintln(&n, eType, indent)
		default:
			fmt.Print(indent)

			fmt.Printf("%v: %#v\n", k, v)
		}

	}

}

func sPrintln(s *[]interface{}, eType reflect.Kind, indent string) {
	if len(indent) > 1000 {
		return
	}

	indent += ind

	var eTypeStr string
	if eType == reflect.Interface {
		eTypeStr = "interface{}"
	} else {
		eTypeStr = eType.String()
	}

	length := len(*s)
	if length == 0 {

		fmt.Printf(">> []empty 0\n")

		return
	}

	fmt.Printf(">> []%s  %d\n",
		eTypeStr, length)

	for _, e := range *s {
		t := reflect.TypeOf(e)
		switch t.Kind() {
		case reflect.Map:
			fmt.Print(indent)
			temp := reflect.ValueOf(e)
			n := make(map[interface{}]interface{})
			i := true
			var kType reflect.Kind
			var vType reflect.Kind
			for _, k := range temp.MapKeys() {
				if i {
					i = false
					kType = k.Kind()
					vType = temp.MapIndex(k).Kind()
				}
				n[k] = temp.MapIndex(k).Interface()
			}
			mPrintln(&n, kType, vType, indent)
		case reflect.Slice:
			fmt.Print(indent)
			temp := reflect.ValueOf(e)
			n := make([]interface{}, temp.Len())
			var eType reflect.Kind
			for i := 0; i < temp.Len(); i++ {
				if i == 0 {
					eType = temp.Index(i).Kind()
				}
				n[i] = temp.Index(i).Interface()
			}
			sPrintln(&n, eType, indent)
		default:
			fmt.Print(indent)
			fmt.Printf("%#v\n", e)
		}
	}

}
