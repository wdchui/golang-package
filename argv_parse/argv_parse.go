package argv_parse

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ArgvItem struct {
	Name  string
	Value Value
	Usage string
}

type Value interface {
	Set(string) error
	String() string
}

type (
	intValue    int
	int64Value  int64
	boolValue   bool
	stringValue string
)

var ArgvItems = make(map[string]ArgvItem)

func newIntValue(value int, p *int) *intValue {
	*p = value
	return (*intValue)(p)
}

func (value *intValue) Set(s string) error {
	val, err := strconv.ParseInt(s, 0, 32)
	*value = intValue(val)
	return err
}

func (value *intValue) String() string {
	return fmt.Sprintf("%v", *value)
}

func newInt64Value(value int64, p *int64) *int64Value {
	*p = value
	return (*int64Value)(p)
}

func (value *int64Value) Set(s string) error {
	val, err := strconv.ParseInt(s, 0, 64)
	*value = int64Value(val)
	return err
}

func (value *int64Value) String() string {
	return fmt.Sprintf("%v", *value)
}

func newBoolValue(value bool, p *bool) *boolValue {
	*p = value
	return (*boolValue)(p)
}

func (value *boolValue) Set(s string) error {
	val, err := strconv.ParseBool(s)
	*value = boolValue(val)
	return err
}

func (value *boolValue) String() string {
	return fmt.Sprintf("%v", *value)
}

func newStringValue(value string, p *string) *stringValue {
	*p = value
	return (*stringValue)(p)
}

func (value *stringValue) Set(s string) error {
	*value = stringValue(s)
	return nil
}

func (value *stringValue) String() string {
	return fmt.Sprintf("%v", *value)
}

func IntVar(p *int, name string, value int, usage string) {
	ArgvItems[name] = ArgvItem{name, newIntValue(value, p), usage}
}

func Int64Var(p *int64, name string, value int64, usage string) {
	ArgvItems[name] = ArgvItem{name, newInt64Value(value, p), usage}
}

func BoolVar(p *bool, name string, value bool, usage string) {
	ArgvItems[name] = ArgvItem{name, newBoolValue(value, p), usage}
}

func StringVar(p *string, name string, value string, usage string) {
	ArgvItems[name] = ArgvItem{name, newStringValue(value, p), usage}
}

func Parse() {
	for _, param := range os.Args {
		if strings.HasPrefix(param, "-") && strings.Contains(param, "=") {
			splitIndex := strings.Index(param, "=")
			name := param[1:splitIndex]
			item, exist := ArgvItems[name]
			if exist {
				hasErr := item.Value.Set(param[splitIndex+1:])
				if hasErr != nil {
					fmt.Println("Can not get ", name, " param!! Error msg: ", hasErr, " Usage:", item.Usage)
					os.Exit(1)
				}
			}
		}
	}
}
