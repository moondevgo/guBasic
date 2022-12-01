package guBasic

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_ROOT_FOLDER = `C:\MoonDev\_setup\_config\`
	DEFAULT_ROOT_NAME   = "DEV_CONFIG_ROOT"
)

// 환경변수를 먼저 찾고, 없으면 기본값을 사용한다.
func GetRootFolder(names ...string) string {
	root_folder := DEFAULT_ROOT_FOLDER
	root_name := DEFAULT_ROOT_NAME

	if len(names) > 0 {
		root_name = strings.ToUpper(strings.Join(names, "_"))
	}

	if os.Getenv(root_name) != "" {
		root_folder = os.Getenv(root_name)
	}
	return root_folder
}

// 파일을 읽어서 []byte 형태로 리턴한다.
func GetConfigBuf(path string) []byte {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return buf
}

// unmarshal
func UnmarshalBuf(buf []byte, configs map[string]interface{}, path string) map[string]interface{} {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(buf, &configs); err != nil {
			return nil
		}
	case ".json":
		if err := json.Unmarshal(buf, &configs); err != nil {
			return nil
		}
	}

	return configs
}

// org Map에서 key에 해당하는 Sub Map을 리턴한다.
// TODO: result 를 임의의 type으로 리턴할 수 있도록 수정
func GetSubMap(org map[string]interface{}, keys ...string) map[string]interface{} {
	for _, key := range keys {
		org = org[key].(map[string]interface{})
	}

	return org
}

// path에 해당하는 파일을 읽어서 map[string]interface{} 형태로 리턴한다.
func GetConfigMap(path string, keys ...string) (configs map[string]interface{}) {
	return GetSubMap(UnmarshalBuf(GetConfigBuf(path), configs, path), keys...)
}

// Get Config From Excel(Xlsx) File
// keys {<file path>, <sheet name>, ...}
func GetConfigFromExcel(path, sheet string, fields []string) (configs []map[string]string) {
	config := map[string]string{}
	for _, row := range InitExcel(path, sheet).Read() {
		for _, k := range fields {
			config[k] = row[k]
		}
		configs = append(configs, config)
	}
	return configs
}
