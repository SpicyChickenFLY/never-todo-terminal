package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var dbFilePath = ""
var DB *Model

// Init DB with specified file
func Init(filePath string) error {
	dbFilePath = filePath
	if dbFilePath == "" {
		homePath, err := getHome()
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

	} else {
		fmt.Println("[INFO] init db successfully")
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
	return json.Unmarshal(b, DB)
}

// EndTransaction record model data into file
func EndTransaction() error {
	data, err := json.Marshal(DB)
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

func getHome() (string, error) {
	// user, err := user.Current()
	// if nil == err {
	// 	return user.HomeDir, nil
	// }

	if runtime.GOOS == "windows" {
		return homeDataWindows()
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
	// drive := os.Getenv("HOMEDRIVE")
	// path := os.Getenv("HOMEPATH")
	// home := drive + path
	// if drive == "" || path == "" {
	// 	home = os.Getenv("USERPROFILE")
	// }
	// if home == "" {
	// 	return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	// }
	home := os.Getenv("APPDATA")
	if home == "" {
		return "", errors.New("APPDATA are blank")
	}
	return home, nil
}
