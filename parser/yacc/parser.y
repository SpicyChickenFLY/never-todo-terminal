%{
package parser

import (
  "fmt"
  "strconv"
  "github.com/SpicyChickenFLY/never-todo-cmd/ast"
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
  taskAddNode ast.TaskAddNode
  taskDeleteNode ast.TaskDeleteNode
  taskDoneNode ast.TaskDoneNode
  taskUpdateNode ast.TaskUpdateNode
  taskUpdateOptionNode *ast.TaskUpdateOptionNode

  tagListNode ast.TagListNode
  tagListFilterNode *ast.TagListFilterNode

  idGroupNode *ast.IDGroupNode
  contentGroupNode *ast.ContentGroupNode
  assignGroupNode *ast.AssignGroupNode
}

%token <str> NUM IDENT SETENCE WHITE 

%right <str> PLUS MINUS
%token <str> LBRACK RBRACK

%token <str> NOT
%left <str> AND OR 

%token <str> UI EXPLAIN LOG UNDO 
%token <str> TASK TAG ADD DELETE SET DONE
%token <str> AGE DUE LIKE LOOP 
%token <str> HELP


%type <root> root
%type <stmt> stmt

%type <taskListNode> task_list
%type <taskListFilterNode> task_list_filter
%type <indefiniteTaskListFilterNode> indefinite_task_list_filter itlf_p3 itlf_p2 itlf_p1
%type <taskAddNode> task_add
%type <taskDeleteNode> task_delete
%type <taskDoneNode> task_done
%type <taskUpdateNode> task_update
%type <taskUpdateOptionNode> task_update_option

%type <tagListNode> tag_list
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
    | EXPLAIN stmt { result = ast.NewRootNode(ast.CMDExplain, $2) }
    | stmt { result = ast.NewRootNode(ast.CMDStmt, $1) }
    | help { result = ast.NewRootNode(ast.CMDHelp, nil) }
    ;

stmt:
      log_list  {if debug {fmt.Println("stmt_log_list")}}
    | undo_log {if debug {fmt.Println("stmt_undo_log")}}

    | task_list { $$ = &$1 }
    | task_add { $$ = &$1 }
    | task_delete { $$ = &$1 }
    | task_update { $$ = &$1 }
    | task_done { $$ = &$1 }
    
    | tag_list { $$ = &$1 }
    | tag_set {if debug {fmt.Println("stmt_tag_set")}}
    ;


// ========== HELP =============
help:
      HELP {}
    | task_help {}
    | tag_help {}

task_help:
      TASK WHITE HELP {}
    | task_list WHITE HELP {}
    | task_add WHITE HELP {}
    | task_delete WHITE HELP {}
    | task_update WHITE HELP {}
    ;

tag_help:
      TAG WHITE HELP {/*if debug {fmt.Println("tag_help")}*/}
    | tag_list WHITE HELP {/*if debug {fmt.Println("tag_help")}*/}
    | tag_set WHITE HELP  {/*if debug {fmt.Println("tag_help")}*/}
    ;

// ========== LOG =============
log_list:
      LOG {/*if debug {fmt.Println("log")}*/}
    | LOG WHITE NUM {/*if debug {fmt.Println("log")}*/}
    ;

undo_log:
      UNDO {/*if debug {fmt.Println("log")}*/}
    | UNDO WHITE id {/*if debug {fmt.Println("log")}*/}
    ;

// ========== TASK COMMAND ==============
task_list:
      TASK WHITE task_list_filter { $$ = ast.NewTaskListNode($3) }
    | task_list_filter { $$ = ast.NewTaskListNode($1) }
    ;

task_add:
      TASK ADD indefinite_content { $$ = ast.NewTaskAddNode($3) }
    | ADD indefinite_content { $$ = ast.NewTaskAddNode($2) }
    | TASK ADD definite_content task_add_option { $$ = ast.NewTaskAddNode($3) }
    | ADD definite_content task_add_option { $$ = ast.NewTaskAddNode($2) }
    ;

task_done:
       DONE id_group { $$= ast.NewTaskDoneNode($2) }
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
      itlf_p3 { 
        $$ = $1
      }
    |  content_filter itlf_p3 {
        $$ = $2
        $$.SetContentFilter($1)
      }
    | itlf_p3 content_filter {
        $$ = $1
        $$.SetContentFilter($2)
      }
    ;

itlf_p3:
      itlf_p2 { $$ = $1 }
    | assign_group itlf_p2 {
        $$ = $2
        $$.SetAssignFilter($1)
      }
    | itlf_p2 assign_group {
        $$ = $1
        $$.SetAssignFilter($2)
      }
    ;

itlf_p2:
      itlf_p1 { $$ = $1 }
    | AGE time_list_filter itlf_p1 {
        $$ = $4
        $$.SetAgeFilter($3)
      }
    | itlf_p1 AGE time_list_filter {
        $$ = $1
        $$.SetAgeFilter($4)
      }
    ;

itlf_p1:
      { $$ = ast.NewIndefiniteTaskListFilterNode() }
    | DUE time_list_filter { $$.SetDueFilter($3) }
    ;

task_add_option:
      {}
    | positive_assign_group task_add_option {}
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
      TAG tag_list_filter { $$ = ast.NewTagListNode($2) }
    ;

tag_set:
      TAG SET id definite_content {  }
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

content_filter:
        content_logic_p3 { $$ = $1 }
      | NOT content_logic_p3 { 
          $$ = ast.NewContentGroupNode("", ast.OPNOT, []*ast.ContentGroupNode{$2})
        }
      | indefinite_content {
        $$ = ast.NewContentGroupNode($1, ast.OPNone, []*ast.ContentGroupNode{})
      }
    ;

content_logic_p3:
      content_logic_p2 { $$ = $1 }
    | content_logic_p2 AND content_logic_p2 {
        $$ = ast.NewContentGroupNode("", ast.OPAND, []*ast.ContentGroupNode{$1, $3})
      }
    ;

content_logic_p2:
      content_logic_p1 { $$ = $1 }
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
      SETENCE { $$ = $2 }
    ;

indefinite_content:
      { $$ = "" }

    | indefinite_content NUM  { $$ = $1 + fmt.Sprint($2) }
    | indefinite_content IDENT {$$ = $1 + $2}
    | indefinite_content WHITE {$$ = $1 + $2}

    | indefinite_content TASK { $$ = $1 + $2 }
    | indefinite_content TAG { $$ = $1 + $2 }

    | indefinite_content ADD { $$ = $1 + $2 }
    | indefinite_content DELETE { $$ = $1 + $2 }
    | indefinite_content SET { $$ = $1 + $2 }
    | indefinite_content DONE { $$ = $1 + $2 }
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