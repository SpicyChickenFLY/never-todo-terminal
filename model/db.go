package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/constant"
)

var dbFilePath = ""

// DB exported for dao
var DB *Model

// Init DB with specified file
func Init(filePath string) error {
	dbFilePath = filePath
	if dbFilePath == "" {
		homePath, err := getDatePath()
		if err != nil {
			return err
		}
		dbFilePath = path.Join(homePath, ".nevertodo/data.json")
	}
	if _, err := os.Stat(dbFilePath); err != nil {
		// 创建文件
		_, err := os.Create(dbFilePath)
		if err != nil {
			return err
		}
		// TODO: 写入默认数据
		fp, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer fp.Close()
		_, err = fp.Write([]byte(constant.DBInitData))
	}
	return nil
}

// Import from other file
func Import(filePath string) {

}

// Export to local file
func Export(filePath string) {

}

// Begin return a model data
func Begin() error {
	fp, err := os.OpenFile(dbFilePath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	b, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	DB = &Model{}
	// var m map[string]interface{}
	// if err := json.Unmarshal(b, &m); err != nil {
	// 	return err
	// }
	// mapstructure.Decode(m, DB)
	// return nil
	return json.Unmarshal(b, &DB)

}

// Commit record model data into file
func Commit() error {
	data, err := json.Marshal(DB)
	if err != nil {
		return err
	}
	// fmt.Println(string(data))
	fp, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(data)
	return err
}

// RollBack wrap begin
func RollBack() error {
	return Begin()
}

func getDatePath() (string, error) {
	if runtime.GOOS == "windows" {
		return homeDataWindows()
	}

	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeDataWindows() (string, error) {
	home := os.Getenv("APPDATA")
	if home == "" {
		return "", errors.New("APPDATA are blank")
	}
	return home, nil
}
