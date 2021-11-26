package utils


import (

	"io/ioutil"
    "log"
	// "os"
	"strconv"
	"strings"
	"fmt"

     "gopkg.in/yaml.v3"
	 "errors"
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

func GetWrongFormatMsg(fieldName, exampleFormat string) (msg string) {
	msg = fmt.Sprintf("Wrong format! Field '%s' should be like this: '%s'. \n", fieldName, exampleFormat)
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

func ConvToListOfIntStrDict(inputData string) ([]map[int]string, error) {
	result := []map[int]string{}
	err := errors.New("wrong value format")
	
	values := strings.Split(inputData, ",")
	for _, pairValues := range values {
		mapValues := strings.Split(pairValues, ":")
		if len(mapValues) < 2 {
			return result, err
		} else {
			key := strings.TrimSpace(mapValues[0])
			keyPrepared, convErr := strconv.Atoi(key)
			if convErr != nil {
				return result, convErr
			}
			valuePrepared := strings.TrimSpace(mapValues[1])
			if len(valuePrepared) == 0 {
				return result, err
			}

			result = append(result, map[int]string{
				keyPrepared: valuePrepared,
			})
		}

	}
	return result, nil
}

func FromListOfDictToStr(values []map[int]string) string {
	interResult := []string{}
	var pairVal string
	for _, baseVal := range values {
		for key, value := range baseVal {
			keyPrepared := strconv.Itoa(key)
			pairVal = fmt.Sprintf("%s:%s", keyPrepared, value)
			interResult = append(interResult, pairVal)
		}
	}
	return strings.Join(interResult, ",")
}
