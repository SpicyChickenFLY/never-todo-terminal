<div align=center><img src="./static/logo.png" width = "200" height = "200" /><img src="./static/icon-cmd.png" width = "200" height = "200" /><h1>never-todo-cmd</h1></div>


> NeverToDo待办列表——命令行端

[ Document in English ](./README_EN.md)

## 总览
这个应用被分为五部分
* [后端数据库（开发中）](https://github.com/SpicyChickenFLY/never-todo-backend) - 使用Golang进行开发
* [前端Web页面（开发中）](https://github.com/bluepongo/never-todo-frontend) - 使用Vue进行开发，可能会用dart写Vue
* [PC端（Win/Linux/Mac）（发布v0.0.1）](https://github.com/bluepongo/never-todo-client)- 使用Electron-Vue框架搭建
* [命令行端（Win/Linux/Mac）（开发中）](https://github.com/SpicyChickenFLY/never-todo-cmd) - 使用Golang进行开发
* [移动端（Android/IOS）（尚未开发）](https://github.com/SpicyChickenFLY/never-todo-mobile) - 使用Dart/Flutter搭建


本项目为跨平台的never-todo系列产品提供了一个在线同步的后端服务功能，您可以在自己的服务器中部署该服务并在客户端中配置相应信息来实现同步功能，本项目由[SpicyChickenFLY](https://github.com/SpicyChickenFLY)与[bluepongo](https://github.com/bluepongo)合作开发

## 实现功能
* [x] 实现通过输入指令实现待办、标签的增删改查

## 需要注意的特性
* 任务的删除是完全删除，数据中不会再记录这一条，但是日志中保存着这个操作，因此在下个版本中可以通过undo撤销来恢复
* 标签作为归档任务分类的重要依据，删除的方式为软删除，即删除之后你仍可以在分配了这个标签的待办任务中看见这个标签，但是查询不到这一个标签
* 被服务器通过

## 目前已知的Bug
* [ ] 在通过nt tag add命令创建标签的时候设置标签颜色不生效

## 下个版本实现的功能
* [ ] 实现通过文字终端UI的界面交互完成相应操作
* [ ] 把对于待办和标签的增删改操作记录在操作日志中便于多端同步 
* [ ] 当操作日志超过指定数目时（默认为200条），系统会在新指令执行完毕后对超出限制的旧日志进行合并压缩，已同步的日志不会和尚未同步的日志合并
* [ ] 实现日程相关功能
* [ ] 实现命令行端和桌面端或者后端服务器的同步机制

#### 项目搭建

```bash
# 查看总览
never
# 查看帮助(当前版本未实现)
#never -h

# 通过UI交互界面进行操作(当前版本未实现）
#never ui

# 查看历史操作
#never log [<num>]

# 撤销(当前版本未实现）
#never undo [<log_id>]

# 解释指令的解析结果和执行计划
never explain [<log_id>]

# 查看、搜索待办任务
never [list] [todo]|done|all [<FILTER_TODO_LIST>]
    # FILTER_TODO_LIST
    <id>[-<id>] [<id>]             # 通过ID直接定位
    <content> [and|or <content>]   # 通过内容模糊搜索
    +<tag>|-<tag> [+<tag>|-<tag>]  # 通过标签筛选
    age:<age>|[<age>]-[<age>]      # 通过创建时间筛选(未完成)
    due:<due>|[<due>]-[<due>]      # 通过截止时间筛选(未完成)

# 新增待办任务
never add <content> [<FILTER_TODO_ADD>]
    # FILTER_TODO_ADD
    +<tag> [+<tag>]         # 分配标签
    due:<due>               # 设置截止时间(未完成)
    loop: y|m|w[-SMTWTFS]|d # 设置重复提醒(每周日，一，四：w-SM...T..)(未完成)

# 查看已完成、已删除任务
never done

# 完成、删除任务
never done|del <id>[-<id>] [<id>]

# 修改任务
never [set] <id> [<content>] [<FILTER_TODO_UPDATE>]
    # FILTER_TODO_UPDATE
    +<tag>|-<tag> [+<tag>|-<tag>]   # 分配标签
    due:<due>                       # 设置截止时间  (未完成)
    loop: y|m|w[-SMTWTFS]|d         # 设置重复提醒(每周日，一，四：w-SM...T..)(未完成)

# 查看某一时段内的统计(当前版本未实现）
# never stat year|month|week|day # 默认为day

# 查看所有标签、修改标签
never tag [<FILTER_TAG_LIST>]
    # FILTER_TAG_LIST
    <id>[~<id>] [<id>]                  # 通过ID直接定位
    like <content> [and|or <content>]   # 通过内容模糊搜索
never tag [set] <id> <content>

```
