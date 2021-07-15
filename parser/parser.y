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
   taskDeleteNode *ast.TaskDeleteNode
   taskDoneNode *ast.TaskDoneNode
   taskUpdateNode *ast.TaskUpdateNode
   taskUpdateOptionNode *ast.TaskUpdateOptionNode
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
%type <taskDeleteNode> task_delete
%type <taskDoneNode> task_done
%type <taskUpdateNode> task_update
%type <taskUpdateOptionNode> task_update_option

%type <num> id 
%type <idGroupNode> id_group
%type <str> content shard_content
%type <contentGroupNode> content_group
%type <assignGroupNode> assign_group positive_assign_group
%type <str> assign_tag unassign_tag

%start root

%%

root:
      { result = ast.NewRootNode(ast.CMDSummary, nil) }
    | UI { result = ast.NewRootNode(ast.CMDUI, nil) }
    | GUI { result = ast.NewRootNode(ast.CMDGUI, nil) }
    | EXPLAIN stmt { result = ast.NewRootNode(ast.CMDExplain, &$2) }
    | stmt { result = ast.NewRootNode(ast.CMDStmt, &$1) }
    | help { result = ast.NewRootNode(ast.CMDHelp, nil) }
    ;

stmt:
      log_list  {if debug {fmt.Println("stmt_log_list")}}
    | undo_log {if debug {fmt.Println("stmt_undo_log")}}

    | task_help {if debug {fmt.Println("stmt_task_help")}}
    | task_list {if debug {fmt.Println("stmt_task_list")}}
    | task_add {if debug {fmt.Println("stmt_task_add")}}
    | task_delete { $$ = $1 }
    | task_update { $$ = $1 }
    | task_done { $$ = $1 }
    
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
      TASK task_list_filter {}
    | task_list_filter {}
    ;

task_add:
      TASK ADD content task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    | ADD content task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    | content ADD task_add_filter {/*if debug {fmt.Println("task_add")}*/}
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
      TASK SET id content task_update_option { $$ = ast.NewTaskUpdateNode($3, $4, $5) }
    | SET id content task_update_option { $$ = ast.NewTaskUpdateNode($2, $3, $4) }
    | TASK id SET content task_update_option { $$ = ast.NewTaskUpdateNode($2, $4, $5) }
    | id SET content task_update_option { $$ = ast.NewTaskUpdateNode($1, $3, $4) }
    | TASK id content task_update_option { $$ = ast.NewTaskUpdateNode($2, $3, $4) }
    | id content task_update_option { $$ = ast.NewTaskUpdateNode($1, $2, $3) }
    ;

// ========== TASK FILTER =============
task_list_filter:
      definite_task_list_filter { $$ = ast.NewDefiniteTaskListFilterNode($1, nil) }
    | indefinite_task_list_filter { $$ = ast.NewIndefiniteTaskListFilterNode(nil, $1) }
    ;

definite_task_list_filter:
      id_group {}
    ;

indefinite_task_list_filter:
      LIKE content_group task_list_filter { 
        $$ = $3 
        $$.SetContentGroupNode($2)
      }
    | content_group task_list_filter { 
        $$ = $2
        $$.SetContentGroupNode($1)
      }
    | assign_group task_list_filter {  }
    | AGE COLON time_list_filter task_list_filter {  }
    | DUE COLON time_list_filter task_list_filter {  }
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
      TAG SET id content {}
    ;

// ========== TAG FILTER =============
tag_list_filter:
      { $$ = ast.NewTagListFilter() }
    | id_group { $$ = ast.NewTagListFilter($1) }
    | LIKE content_group { $$ = ast.NewTagListFilter($2) }
    | content_group { $$ = ast.NewTagListFilter($1) }
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
      content { $$ = NewContentGroupNode(ast.OPNone, []ContentGroupNode{$1}) }
    | LBRACK content_group RBRACK { $$ = $2 }
    | content_group AND content_group { $$ = NewContentGroupNode(ast.OPAND, []ContentGroupNode{$1, $3}) }
    | content_group OR content_group { $$ = NewContentGroupNode(ast.OPOR, []ContentGroupNode{$1, $3}) }
    | content_group XOR content_group { $$ = NewContentGroupNode(ast.OPXOR, []ContentGroupNode{$1, $3}) }
    | NOT content_group { $$ = NewContentGroupNode(ast.OPNOT, []ContentGroupNode{$2}) }
    ;

content:
      DQUOTE shard_content DQUOTE { $$ = $2 }
    | QUOTE shard_content QUOTE { $$ = $2 }
    | DQUOTE content DQUOTE { $$ = $2 }
    | QUOTE content QUOTE { $$ = $2 }
    | shard_content { $$ = $1 }
    ;

// 转义所有关键字
shard_content:
      { $$ = "" }
    | id_group shard_content { $$ = $1.Restore() + $2}
    | IDENT shard_content { $$ = $1 + $2 }
    
    | ADD shard_content { $$ = $1 + $2 }
    | DELETE shard_content { $$ = $1 + $2 }
    | SET shard_content { $$ = $1 + $2 }
    | DONE shard_content { $$ = $1 + $2 }

    | AGE shard_content { $$ = $1 + $2 }
    | DUE shard_content { $$ = $1 + $2 }
    | LIKE shard_content { $$ = $1 + $2 }
    | LOOP shard_content { $$ = $1 + $2 }

    | COLON shard_content { $$ = $1 + $2 }
    | PLUS shard_content { $$ = $1 + $2 }
    | MINUS shard_content { $$ = $1 + $2 }
    | AND shard_content { $$ = $1 + $2 }
    | OR shard_content { $$ = $1 + $2 }
    | XOR shard_content { $$ = $1 + $2 }
    | NOT shard_content { $$ = $1 + $2 }
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
      time_single {/*if debug {fmt.Println("time_list_filter")}*/}
    | time_range {/*if debug {fmt.Println("time_list_filter")}*/}
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