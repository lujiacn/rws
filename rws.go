package rws

import (
	// . "./odms"
	// "encoding/xml"
	// "bufio"
	// "fmt"

	"io/ioutil"
	"net/http"
	"os"
	"rws/rwsStruct"
	"strings"

	"github.com/clbanning/x2j"
)

func RwsRead(api_url, user, pwd string) ([]byte, error) {
	//read xml string
	client := &http.Client{}
	req, err := http.NewRequest("GET", api_url, nil)
	// req.SetBasicAuth("ALMAC_VAL", "Ea1?ZSX3")
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

func RwsToMap(api_url, user, pwd string) (map[string]interface{}, error) {
	body, err := RwsRead(api_url,
		user, pwd)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(body))

	output, err := x2j.ToMap(reader)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func UnmashalData(xml_data, type_str string) {
	switch type_str {
	case "StudyList":
		v := &rwsStruct.StudyList{}
		// fmt.Println(v)

	}

}
