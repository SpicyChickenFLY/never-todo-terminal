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
	"time"
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
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return unmarshallModel(m)
}

// Commit record model data into file
func Commit() error {
	m, err := marshallModel()
	if err != nil {
		return err
	}
	data, err := json.Marshal(m)
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

func unmarshallModel(m map[string]interface{}) error {
	if _, ok := m["initRun"]; !ok { // initialize data
		// 给model一个默认数据
	}
	if data, ok := m["data"]; ok { // read data
		dm := data.(map[string]interface{})
		if taskInc, ok := dm["task_inc"]; ok {
			DB.Data.TaskInc = taskInc.(int)
		}
		if tasks, ok := dm["tasks"]; ok {
			for _, taskMap := range tasks.([]interface{}) {
				if len(taskMap.([]interface{})) != 6 {
					return errors.New("count of task fields is not matched")
				}
				task := Task{
					ID:        taskMap.([]interface{})[0].(int),
					Content:   taskMap.([]interface{})[1].(string),
					Important: taskMap.([]interface{})[2].(int),
					ProjectID: taskMap.([]interface{})[3].(int),
					Due:       taskMap.([]interface{})[4].(time.Time),
					Loop:      taskMap.([]interface{})[5].(string),
				}
				DB.Data.Tasks[task.ID] = task
				if task.ID <= DB.Data.TaskInc {
					DB.Data.TaskInc = task.ID - 1
				}
			}
		}
		if tagInc, ok := dm["tag_inc"]; ok {
			DB.Data.TaskInc = tagInc.(int)
		}
		if tags, ok := dm["tags"]; ok {
			for _, tagMap := range tags.([]interface{}) {
				if len(tagMap.([]interface{})) != 4 {
					return errors.New("count of tag fields is not matched")
				}
				tag := Tag{
					ID:      tagMap.([]interface{})[0].(int),
					Content: tagMap.([]interface{})[1].(string),
					Color:   tagMap.([]interface{})[2].(string),
					Deleted: tagMap.([]interface{})[3].(bool),
				}
				DB.Data.Tags[tag.ID] = tag
				if tag.ID <= DB.Data.TagInc {
					DB.Data.TagInc = tag.ID - 1
				}
			}
		}
		if taskTags, ok := dm["task_tags"]; ok {
			for _, taskTagMap := range taskTags.([]interface{}) {
				if len(taskTagMap.([]interface{})) != 2 {
					return errors.New("count of task_tag fields is not matched")
				}
				taskID := taskTagMap.([]interface{})[0].(int)
				tagID := taskTagMap.([]interface{})[1].(int)
				DB.Data.TaskTags[taskID] = append(DB.Data.TaskTags[taskID], tagID)
			}
		}
		if projectInc, ok := dm["project_inc"]; ok {
			DB.Data.TaskInc = projectInc.(int)
		}
		if projects, ok := dm["project"]; ok {
			for _, projectMap := range projects.([]interface{}) {
				if len(projectMap.([]interface{})) != 2 {
					return errors.New("count of project fields is not matched")
				}
				project := Project{
					ID:      projectMap.([]interface{})[0].(int),
					Content: projectMap.([]interface{})[1].(string),
				}
				DB.Data.Projects[project.ID] = project
				if project.ID <= DB.Data.ProjectInc {
					DB.Data.ProjectInc = project.ID - 1
				}
			}
		}
	}
	if logs, ok := m["logs"]; ok { // read log
		for _, logMap := range logs.([]interface{}) {
			if len(logMap.([]interface{})) != 4 {
				return errors.New("count of log fields is not matched")
			}
			log := Log{
				Table: logMap.([]interface{})[0].(string),
				ID:    logMap.([]interface{})[0].(int),
				Type:  logMap.([]interface{})[0].(int),
				Data:  logMap.([]interface{})[0].(map[string]interface{}),
			}
			DB.Logs = append(DB.Logs, log)
		}
	}
	return nil
}

func marshallModel() (m map[string]interface{}, err error) {
	var dm map[string]interface{}
	var tasks []interface{}
	for _, task := range DB.Data.Tasks {
		var taskMap []interface{}
		taskMap = append(taskMap,
			task.ID,
			task.Content,
			task.Important,
			task.ProjectID,
			task.Due,
			task.Loop,
		)
		tasks = append(tasks, taskMap)
	}
	dm["tasks"] = tasks
	var tags []interface{}
	for _, tag := range DB.Data.Tags {
		var tagMap []interface{}
		tagMap = append(tagMap,
			tag.ID,
			tag.Content,
			tag.Color,
		)
		tags = append(tags, tagMap)
	}
	dm["tags"] = tags
	var taskTags []interface{}
	for taskID, tagIDsOfTask := range DB.Data.TaskTags {
		var taskTagMap []interface{}
		for _, tagID := range tagIDsOfTask {
			taskTagMap = append(taskTagMap, []int{taskID, tagID})
			taskTags = append(taskTags, taskTagMap)
		}
	}
	dm["tas_tags"] = taskTags
	dm["task_inc"] = DB.Data.TaskInc
	dm["tag_inc"] = DB.Data.TagInc
	dm["project_inc"] = DB.Data.ProjectInc
	m["data"] = dm
	var logs []interface{}
	for _, log := range DB.Logs {
		var logMap []interface{}
		logMap = append(logMap,
			log.Table,
			log.ID,
			log.Type,
			log.Data,
		)
		logs = append(logs, logMap)
	}
	m["logs"] = logs
	return m, nil
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
