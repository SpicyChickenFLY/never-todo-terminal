%{
package parser

import (
    "fmt"
)
%}

%union {
   str string
   num int
   cmd *RootNode
   stmt *StmtNode
}

%token <str> NUM 
%token <str> IDENT

%token <str> PLUS MINUS
%token <str> COLON

%token <str> NOT
%left <str> AND OR 
%left <str> XOR

%token <str> UI GUI EXPLAIN LOG UNDO 
%token <str> TASK TAG ADD DELETE SET DONE
%token <str> AGE DUE LIKE LOOP 

%token <str> HELP

%right PLUS MINUS

%token <str> QUOTE DQUOTE

%type <cmd> cmd
%type <str> id id_group

%start root

%%
root:
      cmd {/*if debug {fmt.Println("cmd")}*/}
      stmt

cmd:
      {if debug {fmt.Println("cmd-summary")}}
    | HELP {if debug {fmt.Println("cmd-HELP")}}
    | UI {if debug {fmt.Println("cmd-UI")}}
    | GUI {if debug {fmt.Println("cmd-GUI")}}
    | EXPLAIN stmt {if debug {fmt.Println("cmd-EXPLAIN")}}

stmt:
    | log_list  {if debug {fmt.Println("stmt_log_list")}}
    | undo_log {if debug {fmt.Println("stmt_undo_log")}}

    | task_help {if debug {fmt.Println("stmt_task_help")}}
    | task_list {if debug {fmt.Println("stmt_task_list")}}
    | task_add {if debug {fmt.Println("stmt_task_add")}}
    | task_delete {if debug {fmt.Println("stmt_task_delete")}}
    | task_set {if debug {fmt.Println("stmt_task_set")}}
    | task_done {if debug {fmt.Println("stmt_task_done")}}
    
    | tag_help {if debug {fmt.Println("stmt_tag_help")}}
    | tag_list {if debug {fmt.Println("stmt_tag_list")}}
    | tag_set {if debug {fmt.Println("stmt_tag_set")}}
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

// ========== TASK COMMAND =============
task_help:
      TASK HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_list HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_add HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_delete HELP {/*if debug {fmt.Println("task_help")}*/}
    | task_set HELP {/*if debug {fmt.Println("task_help")}*/}
    ;

task_list:
      TASK task_list_filter {/*if debug {fmt.Println("task_list")}*/}
    | task_list_filter {/*if debug {fmt.Println("task_list")}*/}
    ;

task_add:
      TASK ADD content task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    | ADD content task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    | content ADD task_add_filter {/*if debug {fmt.Println("task_add")}*/}
    ;

task_done:
      TASK DONE id_group {/*if debug {fmt.Println("task_done")}*/}
    | DONE id_group {/*if debug {fmt.Println("task_done")}*/}
    | id_group DONE {/*if debug {fmt.Println("task_done")}*/}
    ;

task_delete:
      TASK DELETE id_group {/*if debug {fmt.Println("task_delete")}*/}
    | DELETE id_group {/*if debug {fmt.Println("task_delete")}*/}
    | id_group DELETE {/*if debug {fmt.Println("task_delete")}*/}
    ;

task_set:
      TASK SET id content task_update_filter {/*if debug {fmt.Println("task_set")}*/}
    | SET id content task_update_filter {/*if debug {fmt.Println("task_set")}*/}
    | TASK id SET content task_update_filter {/*if debug {fmt.Println("task_set")}*/}
    | id SET content task_update_filter {/*if debug {fmt.Println("task_set")}*/}
    | TASK id content task_update_filter {/*if debug {fmt.Println("task_set")}*/}
    | id content task_update_filter {/*if debug {fmt.Println("task_set")}*/}
    ;


// ========== TAG COMMAND =============
tag_help:HELP
      TAG HELP {/*if debug {fmt.Println("tag_help")}*/}
    | tag_list HELP {/*if debug {fmt.Println("tag_help")}*/}
    | tag_set HELP  {/*if debug {fmt.Println("tag_help")}*/}
    ;

tag_list:
      TAG tag_list_filter {/*if debug {fmt.Println("tag_list")}*/}
    ;

tag_set:
      TAG SET id content {/*if debug {fmt.Println("task_set")}*/}
    ;

// ========== TASK FILTER =============
task_list_filter:
     {/*if debug {fmt.Println("task_list_filter-end")}*/}
    | id_group task_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    | LIKE content_group task_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    | content_group task_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    | assign_group task_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    | AGE COLON time_list_filter task_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    | DUE COLON time_list_filter task_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    ;

task_add_filter:
      {/*if debug {fmt.Println("task_add_filter")}*/}
    | positive_assign_group task_add_filter {/*if debug {fmt.Println("task_add_filter")}*/}
    | DUE COLON time_single {/*if debug {fmt.Println("task_add_filter")}*/}
    | LOOP COLON loop_time {/*if debug {fmt.Println("task_add_filter")}*/}
    ;

task_update_filter:
      {/*if debug {fmt.Println("task_update_filter")}*/}
    | assign_group task_update_filter {/*if debug {fmt.Println("task_update_filter")}*/}
    | DUE COLON time_single {/*if debug {fmt.Println("task_update_filter")}*/}
    | LOOP COLON loop_time {/*if debug {fmt.Println("task_update_filter")}*/}
    ;

// ========== TAG FILTER =============
tag_list_filter:
      {/*if debug {fmt.Println("task_list_filter")}*/}
    | id_group tag_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    | LIKE content_group tag_list_filter {/*if debug {fmt.Println("task_list_filter")}*/}
    ;

// ========== COMMON =============
id_group:
      id id_group {
        $$ =  $1 + $2
        // if debug {fmt.Println("id_group:", $1, $2)
      }
    | id MINUS id {
        $$ = $1 + $2
        // if debug {fmt.Println("id_group:",$1, $3)
      }
    | id {
        $$ = $1
        // if debug {fmt.Println("id_group:", $1)
      }
    ;

id: 
      NUM {
        $$ = $1
        // fmt.Printf("id:%s\n", $1) 
      }
    ; 

content_group:
      content_logic {/*if debug {fmt.Println("content_group")}*/}
    | content {/*if debug {fmt.Println("content_group")}*/}
    ;

content_logic:
      content AND content {/*if debug {fmt.Println("content_logic")}*/}
    | content OR content {/*if debug {fmt.Println("content_logic")}*/}
    | content XOR content {/*if debug {fmt.Println("content_logic")}*/}
    | NOT content {/*if debug {fmt.Println("content_logic")}*/}
    ;

assign_group:
      assign_tag assign_group {/*if debug {fmt.Println("assign_group")}*/}
    | unassign_tag assign_group {/*if debug {fmt.Println("assign_group")}*/}
    | {/*if debug {fmt.Println("assign_group")}*/}
    ;

positive_assign_group:
      assign_tag positive_assign_group {/*if debug {fmt.Println("positive_assign_group")}*/}
    | {/*if debug {fmt.Println("positive_assign_group")}*/}
    ;

assign_tag: 
      PLUS IDENT  {/*if debug {fmt.Println("assign_tag")}*/}
    ;

unassign_tag: 
      MINUS IDENT  {/*if debug {fmt.Println("unassign_tag")}*/}
    ;

content:
      DQUOTE shard_content DQUOTE {}
    | QUOTE shard_content QUOTE {}
    | DQUOTE content DQUOTE {}
    | QUOTE content QUOTE {}
    | shard_content {}

// 转义所有关键字
shard_content:
      {if debug {fmt.Println("content-end")}}
    | IDENT shard_content {if debug {fmt.Println("content-merge")}}
    | id_group shard_content {if debug {fmt.Println("content-merge")}}
    | ADD shard_content {if debug {fmt.Println("content-merge")}}
    | DELETE shard_content {if debug {fmt.Println("content-merge")}}
    | SET shard_content {if debug {fmt.Println("content-merge")}}
    | DONE shard_content {if debug {fmt.Println("content-merge")}}

    | AGE shard_content {if debug {fmt.Println("content-merge")}}
    | DUE shard_content {if debug {fmt.Println("content-merge")}}
    | LIKE shard_content {if debug {fmt.Println("content-merge")}}
    | LOOP shard_content {if debug {fmt.Println("content-merge")}}

    | COLON shard_content {if debug {fmt.Println("content-merge")}}
    | PLUS shard_content {if debug {fmt.Println("content-merge")}}
    | MINUS shard_content {if debug {fmt.Println("content-merge")}}
    | AND shard_content {if debug {fmt.Println("content-merge")}}
    | OR shard_content {if debug {fmt.Println("content-merge")}}
    | XOR shard_content {if debug {fmt.Println("content-merge")}}
    | NOT shard_content {if debug {fmt.Println("content-merge")}}
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

loop_time:
      IDENT {/*if debug {fmt.Println("loop_time")}*/}
     ;

%%