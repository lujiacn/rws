package rws

import (
	// . "./odms"
	"crypto/tls"
	"encoding/xml"
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
	reader := strings.NewReader(string(body))
	output, err := x2j.ToMap(reader)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func RwsToFlatMap(body []byte) ([]map[string]string, error) {
	tMap, err := RwsToMap(body)
	if err != nil {
		return nil, err
	}

	rowSlice := map[int]map[string]string{}
	var f func(map[string]interface{}, string, int)
	f = func(srcMap map[string]interface{}, pre string, rowNum int) {
		for k, v := range srcMap {
			switch reflect.ValueOf(v).Kind() {
			case reflect.Map:
				f(v.(map[string]interface{}), k, rowNum)
			case reflect.Slice:
				for i, tSlice := range v.([]interface{}) {
					// fmt.Println(k, i)
					newRowNum := rowNum + i
					f(tSlice.(map[string]interface{}), k, newRowNum)
				}
			case reflect.String:
				k := strings.Replace(k, "-", "", 1)
				if _, ok := rowSlice[rowNum]; ok {
					rowSlice[rowNum][k] = v.(string)
				} else {
					rowSlice[rowNum] = map[string]string{k: v.(string)}
				}
			}
		}
	}
	// outMap\ := map[string]string{}
	f(tMap, "", 0)

	//flatten
	colNames := []string{}
	outPut := []map[string]string{}
	for k, _ := range rowSlice[0] {
		colNames = append(colNames, k)
	}
	outPut = append(outPut, rowSlice[0])

	for i := 1; i < len(rowSlice); i++ {
		for _, col := range colNames {
			if _, ok := rowSlice[i][col]; !ok {
				rowSlice[i][col] = rowSlice[0][col]
			}
		}

		outPut = append(outPut, rowSlice[i])
	}

	return outPut, nil
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
