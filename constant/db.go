package constant

// DBInitData init database
var DBInitData = `
	{
		"data":{
			"tasks":[
				{"id":-1,"content":"单击查看任务详情","deleted":false,"completed":false,"important":false,"due":"0001-01-01T00:00:00Z"},
				{"id":-2,"content":"双击快捷完成任务","deleted":false,"completed":false,"important":false,"due":"0001-01-01T00:00:00Z"},
				{"id":-3,"content":"点击加号新增对应项目","deleted":false,"completed":false,"important":false,"due":"0001-01-01T00:00:00Z"},
				{"id":-4,"content":"可以随意调整待办对应标签","deleted":false,"completed":false,"important":false,"due":"0001-01-01T00:00:00Z"},
				{"id":-5,"content":"同样，双击已完成会撤销","deleted":false,"completed":true,"important":false,"due":"0001-01-01T00:00:00Z"}
			],
			"tags":[
				{"id":-1,"content":"同左","deleted":false,"color":"#FF0000"},
				{"id":-2,"content":"待办","deleted":false,"color":"#00FF80"}
			],
			"task_tags":[
				{"task_id":-1,"tag_id":-1},
				{"task_id":-2,"tag_id":-1},
				{"task_id":-3,"tag_id":-1},
				{"task_id":-1,"tag_id":-2},
				{"task_id":-2,"tag_id":-2},
				{"task_id":-4,"tag_id":-2},
				{"task_id":-5,"tag_id":-2}
			],
			"taskAutoIncVal":-6,
			"tagAutoIncVal":-3},
		"log":[]
	}`
