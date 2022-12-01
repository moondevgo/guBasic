package guBasic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRootFolder(t *testing.T) {
	defaultFolder := GetRootFolder()
	if defaultFolder != `C:\MoonDev\_setup\_config\` {
		t.Error("GetRootFolder() at defaultFolder is not working")
	}

	rootFolder := GetRootFolder("moon", "config", "root")
	if rootFolder != `C:\Dev\_mm_configs` {
		t.Error("GetRootFolder() is not working")
	}
}

func TestGetConfigMap(t *testing.T) {
	// yaml config
	path := `C:\MoonDev\withLang\inGo\goUtils\_config\database_conn.yaml`
	config := GetConfigMap(path, "mysql_Oracle1_root_ex")

	if val, ok := config["host"]; ok {
		assert.Equal(t, val, "152.67.230.230")
	}

	// json config
	path = `C:\MoonDev\withLang\inGo\goUtils\_config\google_bot_moonsats.json`
	config2 := GetConfigMap(path)
	if val2, ok2 := config2["token_uri"]; ok2 {
		assert.Contains(t, val2, "token") // "token_uri": "https://oauth2.googleapis.com/token"
	}
}

func TestGetConfigFromExcel(t *testing.T) {
	// path := `C:\MoonDev\_setup\configs\onStock\specs_ebest\XingApiDllSpec.xlsx`
	// configs := GetConfigFromExcel(path, "_brief", []string{"sheetName", "desc", "content"})
	// if val, ok := configs[0]["sheetName"]; ok {
	// 	assert.Equal(t, val, "functions")
	// }

	path := `C:\MoonDev\_setup\configs\onStock\test.xlsx`
	configs := GetConfigFromExcel(path, "sheet_1", []string{"name", "count"})

	if val, ok := configs[0]["name"]; ok {
		assert.Equal(t, val, "moon")
	}
}
