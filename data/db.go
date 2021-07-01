package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	defaultFilePath = "./data.json"
)

type DB struct {
	filePath string
}

func NewDB() *DB {
	if _, err := os.Stat(defaultFilePath); err != nil {
		// 创建文件
	}
	return &DB{
		filePath: defaultFilePath,
	}
}

func (db *DB) Read(model *Model) error {
	fp, err := os.OpenFile(db.filePath, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, model)
}

func (db *DB) Write(model *Model) error {
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(db.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(data)
	return err
}
