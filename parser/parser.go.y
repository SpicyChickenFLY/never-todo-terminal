%{
package parser

import (
  "fmt"
  "strconv"
  "strings"
  "errors"
  "github.com/SpicyChickenFLY/never-todo-cmd/ast"
  "github.com/SpicyChickenFLY/never-todo-cmd/utils"
  "github.com/SpicyChickenFLY/never-todo-cmd/model"
)
%}

%union {
  str string
  num int
  root *ast.RootNode
  stmt ast.StmtNode
  helpNode *ast.HelpNode
  taskListNode *ast.TaskListNode
  taskListFilterNode *ast.TaskListFilterNode
  indefiniteTaskListFilterNode *ast.IndefiniteTaskListFilterNode
  taskListOptionNode *ast.TaskListOptionNode
  taskAddNode *ast.TaskAddNode  
  taskAddOptionNode *ast.TaskAddOptionNode
  taskTodoNode *ast.TaskTodoNode
  taskDeleteNode *ast.TaskDeleteNode
  taskDoneNode *ast.TaskDoneNode
  taskUpdateNode *ast.TaskUpdateNode
  taskUpdateOptionNode *ast.TaskUpdateOptionNode
  tagListNode *ast.TagListNode
  tagListFilterNode *ast.TagListFilterNode
  tagAddNode *ast.TagAddNode
  tagUpdateNode *ast.TagUpdateNode
  tagUpdateOptionNode *ast.TagUpdateOptionNode
  tagDeleteNode *ast.TagDeleteNode
  idGroupNode *ast.IDGroupNode
  contentGroupNode *ast.ContentGroupNode
  assignGroupNode *ast.AssignGroupNode
  timeFilterNode *ast.TimeFilterNode
  timeNode *ast.TimeNode
}

%left <str> PLUS MINUS
%left <str> NOT
%left <str> AND OR 

%token <str> LBRACK RBRACK MULTI
%token <str> NUM IDENT SETENCE DATE TIME WEEK
%token <str> UI EXPLAIN LOG UNDO
%token <str> LIST TODO TAG ADD DELETE UPDATE DONE ALL
%token <str> AGE DUE LOOP IMPORTANCE COLOR SORT
%token <str> HELP

%type <root> root
%type <stmt> stmt

%type <helpNode> help

%type <taskListNode> task_list
%type <taskListFilterNode> task_list_filter
%type <indefiniteTaskListFilterNode> indefinite_task_list_filter 
%type <taskListOptionNode> task_list_option
%type <taskAddNode> task_add
%type <taskAddOptionNode> task_add_option
%type <taskDeleteNode> task_delete 
%type <taskTodoNode> task_todo
%type <taskDoneNode> task_done
%type <taskUpdateNode> task_update
%type <taskUpdateOptionNode> task_update_option task_update_option_first

%type <tagListNode> tag_list
%type <tagListFilterNode> tag_list_filter
%type <tagAddNode> tag_add
%type <tagUpdateNode> tag_update
%type <tagUpdateOptionNode> tag_update_option tag_update_option_first
%type <tagDeleteNode> tag_delete
%type <num> id importance

%type <idGroupNode> id_group
%type <str> content definite_content indefinite_content color project sort
%type <contentGroupNode> content_group content_logic_p3 content_logic_p2 content_logic_p1
%type <assignGroupNode> assign_group positive_assign_group
%type <str> assign_tag unassign_tag
%type <timeFilterNode> time_filter
%type <timeNode> time

%start root

%%

root:
      { ast.Result = ast.NewRootNode(ast.CMDSummary, nil) }
    | UI { ast.Result = ast.NewRootNode(ast.CMDUI, nil) }
    | EXPLAIN stmt { ast.Result = ast.NewRootNode(ast.CMDExplain, $2) }
    | stmt { ast.Result = ast.NewRootNode(ast.CMDStmt, $1) }
    | help { ast.Result = ast.NewRootNode(ast.CMDHelp, $1) }
    ;

stmt:
      log_list  {}
    | undo_log {}
    | task_list { $$ = $1 }
    | task_add { $$ = $1 }
    | task_delete { $$ = $1 }
    | task_update { $$ = $1 }
    | task_done { $$ = $1 }
    | tag_list { $$ = $1 }
    | tag_add { $$ = $1 }
    | tag_delete { $$ = $1 }
    | tag_update { $$ = $1}
    ;


// ========== HELP =============
help:
      HELP { $$ = ast.NewHelpNode(ast.HelpRoot) }
    | task_list HELP { $$ = ast.NewHelpNode(ast.HelpTaskList) }
    | task_add HELP { $$ = ast.NewHelpNode(ast.HelpTaskList) }
    | ADD HELP { $$ = ast.NewHelpNode(ast.HelpTaskAdd) }
    | ADD { $$ = ast.NewHelpNode(ast.HelpTaskAdd) }
    | task_delete HELP { $$ = ast.NewHelpNode(ast.HelpTaskDelete) }
    | DELETE HELP { $$ = ast.NewHelpNode(ast.HelpTaskDelete) }
    | DELETE { $$ = ast.NewHelpNode(ast.HelpTaskDelete) }
    | task_update HELP { $$ = ast.NewHelpNode(ast.HelpTaskUpdate) }
    | UPDATE HELP { $$ = ast.NewHelpNode(ast.HelpTaskUpdate) }
    | UPDATE { $$ = ast.NewHelpNode(ast.HelpTaskUpdate) }
    | tag_list HELP { $$ = ast.NewHelpNode(ast.HelpTagList) }
    | TAG LIST HELP { $$ = ast.NewHelpNode(ast.HelpTagList) }
    | tag_add HELP { $$ = ast.NewHelpNode(ast.HelpTagAdd) }
    | TAG ADD HELP { $$ = ast.NewHelpNode(ast.HelpTagAdd) }
    | TAG ADD { $$ = ast.NewHelpNode(ast.HelpTagAdd) }
    | tag_delete HELP { $$ = ast.NewHelpNode(ast.HelpTagDelete) }
    | TAG DELETE HELP { $$ = ast.NewHelpNode(ast.HelpTagDelete) }
    | TAG DELETE { $$ = ast.NewHelpNode(ast.HelpTagDelete) }
    | tag_update HELP { $$ = ast.NewHelpNode(ast.HelpTagUpdate) }
    | TAG UPDATE HELP { $$ = ast.NewHelpNode(ast.HelpTagUpdate) }
    | TAG UPDATE { $$ = ast.NewHelpNode(ast.HelpTagUpdate) }
    ;

// ========== LOG =============
log_list:
      LOG {}
    | LOG NUM {}
    ;

undo_log:
      UNDO {}
    | UNDO id {}
    ;

// ========== TASK COMMAND ==============
task_list:
      LIST task_list_filter task_list_option { $$ = ast.NewTaskListNode(model.TaskTodo, $2, $3) } 
    | LIST TODO task_list_filter task_list_option { $$ = ast.NewTaskListNode(model.TaskTodo, $3, $4) }
    | LIST DONE task_list_filter task_list_option { $$ = ast.NewTaskListNode(model.TaskDone, $3, $4) }
    | LIST ALL task_list_filter task_list_option { $$ = ast.NewTaskListNode(model.TaskAll, $3, $4) }
    ;

task_add:
      ADD content task_add_option { $$ = ast.NewTaskAddNode($2, $3) }
    ;

task_todo:
      id_group TODO { $$ = ast.NewTaskTodoNode($1) }
    | TODO id_group { $$ = ast.NewTaskTodoNode($2) }
    ;

task_done:
      DONE id_group { $$= ast.NewTaskDoneNode($2) }
    | id_group DONE { $$= ast.NewTaskDoneNode($1) }
    ;

task_delete:
      DELETE id_group {$$ = ast.NewTaskDeleteNode($2)}
    | id_group DELETE {$$ = ast.NewTaskDeleteNode($1)}
    ;

task_update:
      id task_update_option {$$ = ast.NewTaskUpdateNode($1, $2)}
    | UPDATE id task_update_option {$$ = ast.NewTaskUpdateNode($2, $3)} 
    ;

// ========== TASK LIST FILTER =============
task_list_filter:
      { $$ = ast.NewTaskListFilterNode(nil, nil) }
    | id_group { $$ = ast.NewTaskListFilterNode($1, nil) }
    | indefinite_task_list_filter { $$ = ast.NewTaskListFilterNode(nil, $1) }
    ;

indefinite_task_list_filter:
      { $$ = ast.NewIndefiniteTaskListFilterNode() }
    | content_group indefinite_task_list_filter { $$ = $2.SetContentGroup($1) }
    | indefinite_task_list_filter content_group { $$ = $1.SetContentGroup($2) }
    | importance indefinite_task_list_filter { $$ = $2.SetImportance($1) }
    | indefinite_task_list_filter importance { $$ = $1.SetImportance($2) }
    | assign_group indefinite_task_list_filter { $$ = $2.MergeAssignGroup($1) }
    | indefinite_task_list_filter assign_group { $$ = $1.MergeAssignGroup($2) }
    | AGE time_filter indefinite_task_list_filter { $$ = $3.SetAge($2) }
    | indefinite_task_list_filter AGE time_filter { $$ = $1.SetAge($3) }
    | DUE time_filter indefinite_task_list_filter { $$ = $3.SetDue($2) }
    | indefinite_task_list_filter DUE time_filter { $$ = $1.SetDue($3) }
    | project indefinite_task_list_filter { $$=$2.SetProject($1) }
    | indefinite_task_list_filter project { $$=$1.SetProject($2) }
    ;

// ========== TASK LIST OPTION =============

task_list_option:
      { $$ = ast.NewTaskListOptionNode() }
    | sort task_list_option { $$ = $2.SetSortMetric($1) }
    | task_list_option sort { $$ = $1.SetSortMetric($2) } 
    /* | LOOP loop_time {} */
    ;

// ========== TASK ADD OPTION =============

task_add_option:
      { $$ = ast.NewTaskAddOptionNode() }
    | positive_assign_group task_add_option { $$ = $2.SetAssignGroup($1) }
    | task_add_option positive_assign_group { $$ = $1.SetAssignGroup($2) }
    | task_add_option importance { $$ = $1.SetImportance($2) }
    | importance task_add_option { $$ = $2.SetImportance($1) }
    | task_add_option DUE time_filter { $$ = $1.SetDue($3) }
    | DUE time_filter task_add_option { $$ = $3.SetDue($2) }    
    /* | LOOP loop_time {} */
    ;

// ========== TASK UPDATE OPTION =============

task_update_option:
      task_update_option_first { $$ = $1 }
    | content task_update_option { $$ = $2.SetContent($1) }
    | task_update_option content { $$ = $1.SetContent($2) }
    | assign_group task_update_option { $$ = $2.SetAssignGroup($1) }
    | task_update_option assign_group { $$ = $1.SetAssignGroup($2) }
    | task_update_option importance { $$ = $1.SetImportance($2) }
    | importance task_update_option { $$ = $2.SetImportance($1) }
    | task_update_option DUE time { $$ = $1.SetDue($3) }
    | DUE time task_update_option { $$ = $3.SetDue($2) }
    ;

task_update_option_first:
      content { $$ = ast.NewTaskUpdateOptionNode().SetContent($1) }
    | assign_group { $$ = ast.NewTaskUpdateOptionNode().SetAssignGroup($1) }
    | importance { $$ = ast.NewTaskUpdateOptionNode().SetImportance($1) }
    | DUE time { $$ = ast.NewTaskUpdateOptionNode().SetDue($2) }
    ;

// ========== TAG COMMAND =============

tag_list:
      TAG tag_list_filter { $$ = ast.NewTagListNode($2) }
    ;

tag_add:
      TAG ADD content { $$ = ast.NewTagAddNode($3, "") }
    | TAG ADD content color { $$ = ast.NewTagAddNode($3, $4) }
    ;

tag_delete:
      TAG DELETE id_group { $$ = ast.NewTagDeleteNode($3) }
    ;

tag_update:
      TAG id tag_update_option { $$ = ast.NewTagUpdateNode($2, $3) }
    ;

// ========== TAG FILTER =============
tag_list_filter:
      { $$ = ast.NewTagListFilterNode(nil, "") }
    | id_group { $$ = ast.NewTagListFilterNode($1, "") }
    | content { $$ = ast.NewTagListFilterNode(nil, $1) }
    ;

// ========== TASK UPDATE OPTION =============

tag_update_option:
      tag_update_option_first { $$ = $1 }
    | content tag_update_option { $$ = $2.SetContent($1) }
    | tag_update_option content { $$ = $1.SetContent($2) }
    | color tag_update_option { $$ = $2.SetColor($1) }
    | tag_update_option color { $$ = $1.SetColor($2) }
    ;

tag_update_option_first:
      content { $$ = ast.NewTagUpdateOptionNode().SetContent($1) }
    | color { $$ = ast.NewTagUpdateOptionNode().SetColor($1) }
    ;

// ========== UTILS =============
sort:
      SORT content { $$ = $2 }
    ;

project:
      MULTI content { $$ = $2 }
    ;

assign_group:
      assign_tag { $$ = ast.NewAssignGroupNode($1, "") }
    | unassign_tag { $$ = ast.NewAssignGroupNode("", $1) }
    | assign_tag assign_group { $$ = $2.AssignTag($1) }
    | unassign_tag assign_group { $$ = $2.UnassignTag($1) }
    ;

positive_assign_group:
      assign_tag { $$ = ast.NewAssignGroupNode($1, "") }
    | assign_tag positive_assign_group { $$ = $2.AssignTag($1) }
    ;

assign_tag: 
      PLUS content  { $$ = $2 } 
    ;

unassign_tag: 
      MINUS content  { $$ = $2 } 
    ;

time_filter:
      time { $$ = ast.NewTimeFilterNode ($1, nil) }
    | time MINUS time { $$ = ast.NewTimeFilterNode ($1, $3) }
    ;

time:
      DATE { $$ = ast.NewTimeNode($1, ast.TimeFormatDate) }
    | TIME { $$ = ast.NewTimeNode($1, ast.TimeFormatTime) }
    | DATE TIME { $$ = ast.NewTimeNode($1 + " " + $2, ast.TimeFormatDateTime) }
    ;

importance:
      IMPORTANCE { $$, _ = strconv.Atoi($1) }
    ;

color:
      COLOR content { $$ = $2 }
    ;

/* loop_time:
      IDENT {}
     ; */

id_group:
      id_group id_group { $$ = $1.MergeIDNode($2) }
    | id MINUS id { $$ = ast.NewIDGroupNode($1, $3) }
    | id { $$ = ast.NewIDGroupNode($1) }
    ;

id: 
      NUM { $$, _ = strconv.Atoi($1) }
    | MINUS NUM { $$, _ = strconv.Atoi("-"+$2) }
    ; 

content_group:
      content_logic_p3 { $$ = $1 }
    ;

content_logic_p3:
      content_logic_p2 { $$ = $1 }
    | content_logic_p3 AND content_logic_p3 { $$ = ast.NewContentGroupNode("", ast.OPAND, []*ast.ContentGroupNode{$1, $3}) }
    ;

content_logic_p2:
      content_logic_p1 { $$ = $1 }
    | content_logic_p2 OR content_logic_p2 { $$ = ast.NewContentGroupNode("", ast.OPOR, []*ast.ContentGroupNode{$1, $3}) }
    ;

content_logic_p1:
      content { $$ = ast.NewContentGroupNode($1, ast.OPNone, []*ast.ContentGroupNode{}) }
    | LBRACK content_logic_p3 RBRACK { $$ = $2 }
    | NOT content_logic_p1 { $$ = ast.NewContentGroupNode("", ast.OPNOT, []*ast.ContentGroupNode{$2}) }
    ;

content:
      definite_content 
      { 
        var err error
          $$, err = utils.DecodeCmd($1)
        if err != nil {
            ast.ErrorList = append(ast.ErrorList, errors.New("Illegal character in CMD"))
        }	
      }
    | indefinite_content 
      { 
        var err error
        $$, err = utils.DecodeCmd(strings.Trim($1, " "))
        if err != nil {
            ast.ErrorList = append(ast.ErrorList, errors.New("Illegal character in CMD"))
        }	
      }
    ;

definite_content:
      SETENCE { $$ = ast.SearchVarMap($1) }
    ;

indefinite_content:
      NUM  { $$ = fmt.Sprint($1)}
    | IDENT {$$ = $1}
    | TODO { $$ = $1 }
    | TAG { $$ = $1 }
    | ADD { $$ = $1 }
    | DELETE { $$ = $1 }
    | DONE { $$ = $1 }
    | DATE { $$ = $1 }
    | TIME { $$ = $1 }
    | indefinite_content NUM  { $$ = $1 + " " + fmt.Sprint($2) }
    | indefinite_content IDENT {$$ = $1 + " " + $2}
    | indefinite_content TODO { $$ = $1 + " " + $2 }
    | indefinite_content TAG { $$ = $1 + " " + $2 }
    | indefinite_content ADD { $$ = $1 + " " + $2 }
    | indefinite_content DELETE { $$ = $1 + " " + $2 }
    | indefinite_content DONE { $$ = $1 + " " + $2 }
    | indefinite_content DATE { $$ = $1 + " " + $2 }
    | indefinite_content TIME { $$ = $1 + " " + $2 }
    ;
%%
