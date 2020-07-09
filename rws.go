package rws

import (
	// . "./odms"
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"

	// "bufio"
	// "fmt"
	"reflect"

	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/lujiacn/rws/rwsStruct"

	"github.com/clbanning/x2j"
)

func RwsRead(apiUrl, user, pwd string, proxyUrl string) ([]byte, error) {

	var client = &http.Client{}
	if proxyUrl != "" {
		proxy, _ := url.Parse(proxyUrl)
		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{
			Transport: tr,
		}
	}

	//read body string
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.SetBasicAuth(user, pwd)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Retrieve the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}
	return []byte(body), nil
}

func ReadXml(file_name string) ([]byte, error) {
	file, err := os.Open(file_name) // For read access.
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RwsToMap(body []byte) (map[string]interface{}, error) {
	// replace <mdsol: to <mdsol_Query
	reg := regexp.MustCompile(`<(\/?)(\w+)(:)(\w+)>`)
	newBody := reg.ReplaceAll(body, []byte("<$1$2_$4>"))
	reader := bytes.NewReader(newBody)
	output, err := x2j.ToMap(reader)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// return rows, colNames, error
func RwsToFlatMap(body []byte) ([]map[string]string, []string, error) {
	tMap, err := RwsToMap(body)
	if err != nil {
		return nil, nil, err
	}

	rowSlice := map[int]map[string]string{}
	colNameMap := map[string]bool{}
	var f func(map[string]interface{}, int, int)
	//odmMap := map[string]string{}
	f = func(srcMap map[string]interface{}, pre int, rowNum int) {
		pre = pre + 1
		for k, v := range srcMap {
			k = strings.TrimPrefix(k, "-")

			switch reflect.ValueOf(v).Kind() {
			case reflect.Map:
				f(v.(map[string]interface{}), pre, rowNum)
			case reflect.Slice:
				strValue := ""

				for i, tSlice := range v.([]interface{}) {
					//f(tSlice.(map[string]interface{}), k, newRowNum)
					if reflect.TypeOf(tSlice).Kind() == reflect.String {
						strValue = strValue + " | " + tSlice.(string)
					} else {
						newRowNum := rowNum + i
						f(tSlice.(map[string]interface{}), pre, newRowNum)
					}

				}
				// as string output
				if strValue != "" {
					k := strings.Replace(k, "-", "_", 1)
					//newK := pre + k
					newK := k
					colNameMap[newK] = true
					if _, ok := rowSlice[rowNum]; ok {
						rowSlice[rowNum][newK] = strValue
					} else {
						rowSlice[rowNum] = map[string]string{newK: strValue}
					}
					rowSlice[rowNum]["level"] = fmt.Sprintf("%v", pre)
					//odmMap[newK] = strValue
					//if pre == "ODM" {
					//odmMap[newK] = strValue
					//}

				}
			case reflect.String:
				k := strings.Replace(k, "-", "_", 1)
				//newK := pre + k
				newK := k
				colNameMap[newK] = true
				if _, ok := rowSlice[rowNum]; ok {
					rowSlice[rowNum][newK] = v.(string)
				} else {
					rowSlice[rowNum] = map[string]string{newK: v.(string)}
				}

				rowSlice[rowNum]["level"] = fmt.Sprintf("%v", pre)
				//odmMap[newK] = v.(string)
				//if pre == "ODM" {
				//odmMap[newK] = v.(string)
				//}
			}
		}
	}
	// outMap\ := map[string]string{}
	f(tMap, 0, 0)
	//colNames
	colNames := []string{"level"}
	for k := range colNameMap {
		colNames = append(colNames, k)
	}

	outPut := []map[string]string{}
	//lagRow := map[string]string{}
	levelMap := map[int]map[string]string{}
	for i := 0; i < len(rowSlice); i++ {
		row := rowSlice[i]

		level, _ := strconv.Atoi(row["level"])

		if _, found := levelMap[level]; !found {
			levelMap[level] = row
		}

		preRow := map[string]string{}

		if i > 0 {
			for j := level; j < 0; j-- {
				if _, found := levelMap[j]; found {
					preRow = levelMap[j]
				}
			}
		}

		for _, c := range colNames {
			if v, ok := preRow[c]; ok {
				row[c] = v
			}
		}
		outPut = append(outPut, row)
	}

	return outPut, colNames, nil
}

func UnmashalData(xml_byte []byte, type_str string) interface{} {
	switch type_str {
	case "StudyList":
		v := &rwsStruct.StudyList{}
		xml.Unmarshal(xml_byte, &v)
		return v
	}
	return nil

}
