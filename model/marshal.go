package model

import (
	"errors"
	"time"
)

func unmarshalModel(m map[string]interface{}) error {
	if _, ok := m["initRun"]; !ok { // initialize data
		// 给model一个默认数据
	}
	if data, ok := m["data"]; ok { // read data
		if err := unmarshalData(data); err != nil {
			return err
		}
	}
	if logs, ok := m["logs"]; ok { // read log
		if err := unmarshalLog(logs); err != nil {
			return err
		}
	}
	return nil
}

func unmarshalData(data interface{}) error {
	dm, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("field data cannot be convert to map[string]interface{}")
	}
	if taskInc, ok := dm["task_inc"]; ok {
		taskIncVal, ok := taskInc.(float64)
		if !ok {
			return errors.New("field task_inc cannot be convert to float64")
		}
		DB.Data.TaskInc = int(taskIncVal)
	}
	if tasks, ok := dm["tasks"]; ok {
		if err := unmarshalTask(tasks); err != nil {
			return err
		}
	}
	if tagInc, ok := dm["tag_inc"]; ok {
		tagIncVal, ok := tagInc.(float64)
		if !ok {
			return errors.New("field tag_inc cannot be convert to float64")
		}
		DB.Data.TaskInc = int(tagIncVal)
	}
	if tags, ok := dm["tags"]; ok {
		if err := unmarshalTag(tags); err != nil {
			return err
		}
	}
	if taskTags, ok := dm["task_tags"]; ok {
		if err := unmarshalTaskTag(taskTags); err != nil {
			return err
		}
	}
	if projectInc, ok := dm["project_inc"]; ok {
		projectIncVal, ok := projectInc.(float64)
		if !ok {
			return errors.New("field project_inc cannot be convert to float64")
		}
		DB.Data.TaskInc = int(projectIncVal)
	}
	if projects, ok := dm["project"]; ok {
		if err := unmarshalProject(projects); err != nil {
			return err
		}
	}
	return nil
}

func unmarshalTask(tasks interface{}) error {
	var ok bool
	DB.Data.Tasks = make(map[int]Task)
	for _, task := range tasks.([]interface{}) {
		var taskMap []interface{}
		taskMap, ok = task.([]interface{})
		if !ok {
			return errors.New("field taskMap cannot be convert to []interface{}")
		}
		if len(taskMap) != 6 {
			return errors.New("count of taskMap fields is not matched")
		}

		var taskID, taskImportant, taskProjectID float64
		var taskContent, taskDue, taskLoop string

		if taskID, ok = taskMap[0].(float64); !ok {
			return errors.New("field taskMap[0] cannot be convert to float64")
		}
		if taskContent, ok = taskMap[1].(string); !ok {
			return errors.New("field taskMap[1] cannot be convert to string")
		}
		if taskImportant, ok = taskMap[2].(float64); !ok {
			return errors.New("field taskMap[2] cannot be convert to float64")
		}
		if taskProjectID, ok = taskMap[3].(float64); !ok {
			return errors.New("field taskMap[3] cannot be convert to float64")
		}
		if taskDue, ok = taskMap[4].(string); !ok {
			return errors.New("field taskMap[4] cannot be convert to time.Time")
		}
		loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return err
		}
		dueTime, err := time.ParseInLocation("2006-01-02T15:04:05Z", taskDue, loc)
		if err != nil {
			return err
		}
		if taskLoop, ok = taskMap[5].(string); !ok {
			return errors.New("field taskMap[5] cannot be convert to string")
		}

		task := Task{
			ID:        int(taskID),
			Content:   taskContent,
			Important: int(taskImportant),
			ProjectID: int(taskProjectID),
			Due:       dueTime,
			Loop:      taskLoop,
		}

		DB.Data.Tasks[task.ID] = task
		if task.ID <= DB.Data.TaskInc {
			DB.Data.TaskInc = task.ID - 1
		}
	}
	return nil
}

func unmarshalTag(tags interface{}) error {
	var ok bool
	DB.Data.Tags = make(map[int]Tag)
	for _, tag := range tags.([]interface{}) {
		var tagMap []interface{}
		if tagMap, ok = tag.([]interface{}); !ok {
			return errors.New("field tagMap cannot be convert to []interface{}")
		}

		if len(tagMap) != 4 {
			return errors.New("count of tagMap fields is not matched")
		}

		var tagID float64
		var tagContent, tagColor string
		var tagDeleted bool

		if tagID, ok = tagMap[0].(float64); !ok {
			return errors.New("field tagMap[0] cannot be convert to float64")
		}
		if tagContent, ok = tagMap[1].(string); !ok {
			return errors.New("field tagMap[1] cannot be convert to string")
		}
		if tagColor, ok = tagMap[2].(string); !ok {
			return errors.New("field tagMap[2] cannot be convert to string")
		}
		if tagDeleted, ok = tagMap[3].(bool); !ok {
			return errors.New("field tagMap[3] cannot be convert to bool")
		}

		tag := Tag{
			ID:      int(tagID),
			Content: tagContent,
			Color:   tagColor,
			Deleted: tagDeleted,
		}

		DB.Data.Tags[tag.ID] = tag
		if tag.ID <= DB.Data.TagInc {
			DB.Data.TagInc = tag.ID - 1
		}
	}
	return nil
}

func unmarshalTaskTag(taskTags interface{}) error {
	var ok bool
	DB.Data.TagTasks = make(map[int]map[int]bool)
	DB.Data.TaskTags = make(map[int]map[int]bool)
	for _, taskTag := range taskTags.([]interface{}) {
		var taskTagMap []interface{}
		if taskTagMap, ok = taskTag.([]interface{}); !ok {
			return errors.New("field taskTagMap cannot be convert to []interface{}")
		}

		if len(taskTagMap) != 2 {
			return errors.New("count of taskTagMap fields is not matched")
		}
		taskIDVal, ok := taskTagMap[0].(float64)
		if !ok {
			return errors.New("field taskTagMap[0] cannot be convert to float64")
		}
		taskID := int(taskIDVal)
		tagIDVal, ok := taskTagMap[1].(float64)
		if !ok {
			return errors.New("field taskTagMap[1] cannot be convert to float64")
		}
		tagID := int(tagIDVal)
		if tagMap, ok := DB.Data.TaskTags[taskID]; ok {
			tagMap[tagID] = true
			DB.Data.TaskTags[taskID] = tagMap
		} else {
			tagMap = map[int]bool{tagID: true}
			DB.Data.TaskTags[taskID] = tagMap
		}
		if taskMap, ok := DB.Data.TagTasks[tagID]; ok {
			taskMap[taskID] = true
			DB.Data.TagTasks[tagID] = taskMap
		} else {
			taskMap = map[int]bool{tagID: true}
			DB.Data.TagTasks[tagID] = taskMap
		}
	}
	return nil
}

func unmarshalProject(projects interface{}) error {
	var ok bool
	DB.Data.Projects = make(map[int]Project)
	for _, project := range projects.([]interface{}) {
		var projectMap []interface{}
		if projectMap, ok = project.([]interface{}); !ok {
			return errors.New("field projectMap cannot be convert to []interface{}")
		}

		if len(projectMap) != 2 {
			return errors.New("count of projectMap fields is not matched")
		}

		var projectID float64
		var projectContent string

		projectID, ok := projectMap[0].(float64)
		if !ok {
			return errors.New("field projectMap[0] cannot be convert to float64")
		}
		if projectContent, ok = projectMap[1].(string); !ok {
			return errors.New("field projectMap[1] cannot be convert to string")
		}

		project := Project{
			ID:      int(projectID),
			Content: projectContent,
		}

		DB.Data.Projects[project.ID] = project
		if project.ID <= DB.Data.ProjectInc {
			DB.Data.ProjectInc = project.ID - 1
		}
	}
	return nil
}

func unmarshalLog(logs interface{}) error {
	var ok bool
	for _, log := range logs.([]interface{}) {
		var logMap []interface{}
		if logMap, ok = log.([]interface{}); !ok {
			return errors.New("field logMap cannot be convert to []interface{}")
		}

		if len(logMap) != 4 {
			return errors.New("count of logMap fields is not matched")
		}

		var logTable string
		var logID, logType float64
		var logData []interface{}
		var ok bool

		if logTable, ok = logMap[0].(string); !ok {
			return errors.New("field logMap[0] cannot be convert to string")
		}
		if logID, ok = logMap[1].(float64); !ok {
			return errors.New("field logMap[1] cannot be convert to float64")
		}
		if logType, ok = logMap[2].(float64); !ok {
			return errors.New("field logMap[2] cannot be convert to float64")
		}
		if logData, ok = logMap[3].([]interface{}); !ok {
			return errors.New("field logMap[3] cannot be convert to map[string]interface{}")
		}

		log := Log{
			Table: logTable,
			ID:    int(logID),
			Type:  int(logType),
			Data:  logData,
		}

		DB.Logs = append(DB.Logs, log)
	}
	return nil
}

func marshalModel() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	dm := make(map[string]interface{})
	var tasks []interface{}
	for _, task := range DB.Data.Tasks {
		var taskMap []interface{}
		dueTime := task.Due.Format("2006-01-02T15:04:05Z")
		taskMap = append(taskMap,
			task.ID,
			task.Content,
			task.Important,
			task.ProjectID,
			dueTime,
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
			tag.Deleted,
		)
		tags = append(tags, tagMap)
	}
	dm["tags"] = tags
	var taskTags []interface{}
	for taskID, tagIDsOfTask := range DB.Data.TaskTags {
		var taskTagMap []interface{}
		for tagID := range tagIDsOfTask {
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
