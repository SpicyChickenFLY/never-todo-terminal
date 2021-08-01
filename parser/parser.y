%{
package parser

import (
  "fmt"
  "strconv"
  "github.com/SpicyChickenFLY/never-todo-cmd/parser/ast"
)
%}

%union {
  str string
  num int
  root ast.Node
  stmt ast.StmtNode

  taskListNode ast.TaskListNode
  taskListFilterNode *ast.TaskListFilterNode
  indefiniteTaskListFilterNode *ast.IndefiniteTaskListFilterNode
  taskDeleteNode ast.TaskDeleteNode
  taskDoneNode ast.TaskDoneNode
  taskUpdateNode ast.TaskUpdateNode
  taskUpdateOptionNode *ast.TaskUpdateOptionNode

  tagListFilterNode *ast.TagListFilterNode

  idGroupNode *ast.IDGroupNode
  contentGroupNode *ast.ContentGroupNode
  assignGroupNode *ast.AssignGroupNode
}

%token <str> NUM IDENT

%right <str> PLUS MINUS
%token <str> COLON QUOTE DQUOTE
%token <str> LBRACK RBRACK

%token <str> NOT
%left <str> AND OR 
%left <str> XOR

%token <str> UI GUI EXPLAIN LOG UNDO 
%token <str> TASK TAG ADD DELETE SET DONE
%token <str> AGE DUE LIKE LOOP 
%token <str> HELP


%type <root> root
%type <stmt> stmt

%type <taskListNode> task_list
%type <taskListFilterNode> task_list_filter
%type <indefiniteTaskListFilterNode> indefinite_task_list_filter itlf_1 itlf_2 itlf_3
%type <taskDeleteNode> task_delete
%type <taskDoneNode> task_done
%type <taskUpdateNode> task_update
%type <taskUpdateOptionNode> task_update_option

%type <tagListFilterNode> tag_list_filter

%type <num> id 
%type <idGroupNode> id_group
%type <str> definite_content indefinite_content shard_content
%type <contentGroupNode> content_group content_filter content_logic_p3 content_logic_p2 content_logic_p1
%type <assignGroupNode> assign_group positive_assign_group
%type <str> assign_tag unassign_tag
%type <str> time_list_filter

%start root

%%

root:
      { result = ast.NewRootNode(ast.CMDSummary, nil) }
    | UI { result = ast.NewRootNode(ast.CMDUI, nil) }
    | GUI { result = ast.NewRootNode(ast.CMDGUI, nil) }
    | EXPLAIN stmt { result = ast.NewRootNode(ast.CMDExplain, $2) }
    | stmt { result = ast.NewRootNode(ast.CMDStmt, $1) }
    | help { result = ast.NewRootNode(ast.CMDHelp, nil) }
    ;

stmt:
      log_list  {if debug {fmt.Println("stmt_log_list")}}
    | undo_log {if debug {fmt.Println("stmt_undo_log")}}

    | task_help {if debug {fmt.Println("stmt_task_help")}}
    | task_list { $$ = &$1 }
    | task_add {if debug {fmt.Println("stmt_task_add")}}
    | task_delete { $$ = &$1 }
    | task_update { $$ = &$1 }
    | task_done { $$ = &$1 }
    
    | tag_help {if debug {fmt.Println("stmt_tag_help")}}
    | tag_list {if debug {fmt.Println("stmt_tag_list")}}
    | tag_set {if debug {fmt.Println("stmt_tag_set")}}
    ;


// ========== HELP =============
help:
      HELP {}
    | task_help {}
    | tag_help {}

task_help:
      TASK HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_list HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_add HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_delete HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_update HELP {/*if debug {fmt.Println("task_help")}*/}
    ;

tag_help:
      TAG HELP {/*if debug {fmt.Println("tag_help")}*/}
    | tag_list HELP {/*if debug {fmt.Println("tag_help")}*/}
    | tag_set HELP  {/*if debug {fmt.Println("tag_help")}*/}
    ;

// ========== LOG =============
log_list:
      LOG {/*if debug {fmt.Println("log")}*/}
    | LOG NUM {/*if debug {fmt.Println("log")}*/}
    ;

undo_log:
      UNDO {/*if debug {fmt.Println("log")}*/}
    | UNDO id {/*if debug {fmt.Println("log")}*/}
    ;

// ========== TASK COMMAND ==============
task_list:
      TASK task_list_filter { $$ = ast.NewTaskListNode($2) }
    | task_list_filter { $$ = ast.NewTaskListNode($1) }
    ;

task_add:
      TASK ADD indefinite_content {/*if debug {fmt.Println("task_add")}*/}
    | ADD indefinite_content {/*if debug {fmt.Println("task_add")}*/}
    | TASK ADD definite_content task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    | ADD definite_content task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    | definite_content ADD task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    ;

task_done:
      TASK DONE id_group { $$ = ast.NewTaskDoneNode($3) }
    | DONE id_group { $$= ast.NewTaskDoneNode($2) }
    | id_group DONE { $$= ast.NewTaskDoneNode($1) }
    ;

task_delete:
      TASK DELETE id_group {$$ = ast.NewTaskDeleteNode($3)}
    | DELETE id_group {$$ = ast.NewTaskDeleteNode($2)}
    | id_group DELETE {$$ = ast.NewTaskDeleteNode($1)}
    ;

task_update:
      TASK SET id definite_content task_update_option { $$ = ast.NewTaskUpdateNode($3, $4, $5) }
    | SET id definite_content task_update_option { $$ = ast.NewTaskUpdateNode($2, $3, $4) }
    | TASK id SET definite_content task_update_option { $$ = ast.NewTaskUpdateNode($2, $4, $5) }
    | id SET definite_content task_update_option { $$ = ast.NewTaskUpdateNode($1, $3, $4) }
    | TASK id definite_content task_update_option { $$ = ast.NewTaskUpdateNode($2, $3, $4) }
    | id definite_content task_update_option { $$ = ast.NewTaskUpdateNode($1, $2, $3) }
    ;

// ========== TASK FILTER =============
task_list_filter:
      { $$ = ast.NewTaskListFilterNode(nil, nil) }
    | id_group { $$ = ast.NewTaskListFilterNode($1, nil) }
    | indefinite_task_list_filter { $$ = ast.NewTaskListFilterNode(nil, $1) }
    ;

indefinite_task_list_filter:
      itlf_1 { $$ = $1 }
    |  content_filter itlf_1 {
        $$ = $2
        $$.SetContentFilter($1)
      }
    | itlf_1 content_filter {
        $$ = $1
        $$.SetContentFilter($2)
      }
    ;

content_filter:
      LIKE content_group { $$ = $2 }
    | content_group { $$ = $1}
    ;

itlf_1:
      itlf_2 { $$ = $1 }
    | assign_group itlf_2 {
        $$ = $2
        $$.SetAssignFilter($1)
      }
    | itlf_2 assign_group {
        $$ = $1
        $$.SetAssignFilter($2)
      }
    ;

itlf_2:
      itlf_3 { $$ = $1 }
    | AGE COLON time_list_filter itlf_3 {
        $$ = $4
        $$.SetAgeFilter($3)
      }
    | itlf_3 AGE COLON time_list_filter {
        $$ = $1
        $$.SetAgeFilter($4)
      }
    ;

itlf_3:
      { $$ = ast.NewIndefiniteTaskListFilterNode() }
    | DUE COLON time_list_filter { $$.SetDueFilter($3) }
    ;

task_add_filter:
      {}
    | positive_assign_group task_add_filter {}
    /* | DUE COLON time_single {}
    | LOOP COLON loop_time {} */
    ;

task_update_option:
      { $$ = ast.NewTaskUpdateOptionNode() }
    | assign_group task_update_option { 
        $2.AssignTag($1)
        $$ = $2 
      }
    /* | DUE COLON time_single {}
    | LOOP COLON loop_time {} */
    ;

// ========== TAG COMMAND =============

tag_list:
      TAG tag_list_filter {  }
    ;

tag_set:
      TAG SET id definite_content {}
    ;

// ========== TAG FILTER =============
tag_list_filter:
      { $$ = ast.NewTagListFilterNode(nil, nil) }
    | id_group { $$ = ast.NewTagListFilterNode($1, nil) }
    | LIKE content_group { $$ = ast.NewTagListFilterNode(nil, $2) }
    | content_group { $$ = ast.NewTagListFilterNode(nil, $1) }
    ;

// ========== UTILS =============
id_group:
      id_group id_group { 
        $1.MergeIDNode($2)
        $$ =  $1 
      }
    | id MINUS id { $$ = ast.NewIDGroupNode($1, $3) }
    | id { $$ = ast.NewIDGroupNode($1) }
    ;

id: 
      NUM {
        val, err := strconv.Atoi($1)
        if err != nil {
          panic(err)
        }
        $$ = val
      }
    ; 

content_group:
      content_logic_p3 { 
        $$ = ast.NewContentGroupNode("", ast.OPNOT, []*ast.ContentGroupNode{$1})
      }
      | indefinite_content {
        $$ = ast.NewContentGroupNode($1, ast.OPNone, []*ast.ContentGroupNode{})
      }
    ;

content_logic_p3:
      content_logic_p2 {}
    | content_logic_p2 AND content_logic_p2 {
        $$ = ast.NewContentGroupNode("", ast.OPAND, []*ast.ContentGroupNode{$1, $3})
      }
    ;

content_logic_p2:
      content_logic_p1 {}
    | content_logic_p2 OR content_logic_p2 { 
        $$ = ast.NewContentGroupNode("", ast.OPOR, []*ast.ContentGroupNode{$1, $3})
      }
    ;

content_logic_p1:
      definite_content { 
        $$ = ast.NewContentGroupNode($1, ast.OPNone, []*ast.ContentGroupNode{}) 
      }
    | LBRACK content_logic_p3 RBRACK { $$ = $2 }
    | NOT content_logic_p1 { 
        $$ = ast.NewContentGroupNode("", ast.OPNOT, []*ast.ContentGroupNode{$2}) 
      }
    ;

definite_content:
      DQUOTE shard_content DQUOTE { $$ = $2 }
    | QUOTE shard_content QUOTE { $$ = $2 }
    | DQUOTE definite_content DQUOTE { $$ = $2 }
    | QUOTE definite_content QUOTE { $$ = $2 }
    | DQUOTE indefinite_content DQUOTE {}
    | QUOTE indefinite_content QUOTE {}
    ;

indefinite_content:
      indefinite_content TASK { $$ = $1 + $2 }
    | indefinite_content TAG { $$ = $1 + $2 }

    | indefinite_content ADD { $$ = $1 + $2 }
    | indefinite_content DELETE { $$ = $1 + $2 }
    | indefinite_content SET { $$ = $1 + $2 }
    | indefinite_content DONE { $$ = $1 + $2 }

    | indefinite_content AGE { $$ = $1 + $2 }
    | indefinite_content DUE { $$ = $1 + $2 }
    | indefinite_content LIKE { $$ = $1 + $2 }
    | indefinite_content LOOP { $$ = $1 + $2 }

    | indefinite_content COLON { $$ = $1 + $2 }
    | indefinite_content PLUS { $$ = $1 + $2 }
    | indefinite_content MINUS { $$ = $1 + $2 }
    | indefinite_content AND { $$ = $1 + $2 }
    | indefinite_content OR { $$ = $1 + $2 }
    | indefinite_content XOR { $$ = $1 + $2 }
    | indefinite_content NOT { $$ = $1 + $2 }
    | indefinite_content NUM  { $$ = $1 + fmt.Sprint($2) }

    | shard_content { $$ = $1 }
    ;


// 转义所有关键字
shard_content:
      { $$ = "" }
    | id_group shard_content { $$ = $1.Restore() + $2}
    | IDENT shard_content { $$ = $1 + $2 }
    
    ;

assign_group:
      { $$ = ast.NewAssignGroupNode() }
    | assign_tag assign_group {
        $$ = $2
        $$.AssignTag($1)
      }
    | unassign_tag assign_group {
        $$ = $2
        $$.UnassignTag($1)
      }
    ;

positive_assign_group:
      { $$ = ast.NewAssignGroupNode() }
    | assign_tag positive_assign_group {
        $$ = $2
        $$.AssignTag($1)}
    ;

assign_tag: 
      PLUS IDENT  { $$ = $2 }
    ;

unassign_tag: 
      MINUS IDENT  { $$ = $2}
    ;


time_list_filter:
      time_single {}
    | time_range {}
    ;

time_single:
      IDENT {/*if debug {fmt.Println("time_single")}*/}
    | {/*if debug {fmt.Println("time_single")}*/}
    ;

time_range:
      time_single MINUS time_single {/*if debug {fmt.Println("time_range")}*/}
    ;

/* loop_time:
      IDENT {}
     ; */

%%