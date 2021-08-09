package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	defaultFilePath = "%APPDATA%\\.nevertodo\\data.json"
)

var dbFilePath = defaultFilePath
var M *Model

// Init DB with specified file
func Init() error {
	if _, err := os.Stat(dbFilePath); err != nil {
		// 创建文件
		_, err := os.Create(dbFilePath)
		if err != nil {
			return err
		}
		// TODO: 写入默认数据

	}
	return nil
}

// Import from other file
func Import(filePath string) {

}

// Export to local file
func Export(filePath string) {

}

// StartTransaction return a model data
func StartTransaction() error {
	fp, err := os.OpenFile(dbFilePath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	b, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, M)
}

// EndTransaction record model data into file
func EndTransaction() error {
	data, err := json.Marshal(M)
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(data)
	return err
}
