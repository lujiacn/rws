package rws

import (
	// . "./odms"
	"bytes"
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
	// replace <mdsol: to <mdsol_Query
	newBody := bytes.ReplaceAll(body, []byte("<mdsol:"), []byte("<"))
	reader := bytes.NewReader(newBody)
	//reader := strings.NewReader(string(newBody))
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
	var f func(map[string]interface{}, string, int)
	odmMap := map[string]string{}
	f = func(srcMap map[string]interface{}, pre string, rowNum int) {
		for k, v := range srcMap {
			switch reflect.ValueOf(v).Kind() {
			case reflect.Map:
				f(v.(map[string]interface{}), k, rowNum)
			case reflect.Slice:
				for i, tSlice := range v.([]interface{}) {
					newRowNum := rowNum + i
					f(tSlice.(map[string]interface{}), k, newRowNum)
				}
			case reflect.String:
				k := strings.Replace(k, "-", "_", 1)
				newK := pre + k
				colNameMap[k] = true
				if _, ok := rowSlice[rowNum]; ok {
					rowSlice[rowNum][newK] = v.(string)
				} else {
					rowSlice[rowNum] = map[string]string{newK: v.(string)}
				}
				if pre == "ODM" {
					odmMap[newK] = v.(string)
				}
			}
		}
	}
	// outMap\ := map[string]string{}
	f(tMap, "", 0)
	//colNames
	colNames := []string{}
	for k, _ := range colNameMap {
		colNames = append(colNames, k)
	}

	outPut := []map[string]string{}

	for _, row := range rowSlice {
		for k, v := range odmMap {
			row[k] = v
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
