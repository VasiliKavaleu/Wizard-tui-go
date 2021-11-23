package utils


import (

	"io/ioutil"
    "log"
	// "os"
	"strconv"
	"strings"
	"fmt"

     "gopkg.in/yaml.v3"
	//  "github.com/rivo/tview"
)

func ReadConfigFile(config_path string, config_data interface{}) {
	yfile, err := ioutil.ReadFile(config_path)

	if err != nil {
		log.Fatal(err)
	}

	err2 := yaml.Unmarshal(yfile, &config_data)

	if err2 != nil {
		 log.Fatal(err2)
	}
}

func SaveConfigToFile(config_data interface{}, config_path string) {
	data, err := yaml.Marshal(&config_data)
	if err != nil {
		log.Fatal(err)
   }
   err2 := ioutil.WriteFile(config_path, data, 0644)
   if err2 != nil {
		log.Fatal(err2)
   }
}

func StrToList(value string) (values []string) {
	values = strings.Split(value, ",")
	return
}

func ListToStr(values []string) (value string) {
	value = strings.Join(values, ",")
	return
}

func ValidInt(value string, lastChar rune) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func ValidStr(value string, lastChar rune) bool {
	return len(value) != 0
}

func ValidReqField(value string) bool {
	return len(value) != 0
}

func GetReqFieldMsg(fieldName string) (msg string) {
	msg = fmt.Sprintf("Field '%s' must not be empty.\n", fieldName)
	return
}

func GetIndexFromVal(values []string, value string) int {
	for ind, val := range values {
        if value == val {
            return ind
        }
    }
    return len(values)
}

func BoolToIndexDisableAnable(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}
