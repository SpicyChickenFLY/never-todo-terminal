%{
package parser

import (
    "fmt"
)
%}

%union {
   str string
   num int
   cmd *Command
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

%type <cmd> command
%type <str> id id_group

%start command

%%
command:
      command_stmt {/*fmt.Println("command")*/}

command_stmt:
      {fmt.Println("command_stmt-summary")}
    | HELP {fmt.Println("command_stmt-HELP")}
    | UI {fmt.Println("command_stmt-UI")}
    | GUI {fmt.Println("command_stmt-GUI")}
    | EXPLAIN {fmt.Println("command_stmt-EXPLAIN")}

    | log_list  {fmt.Println("command_stmt-log_list")}
    | undo_log {fmt.Println("command_stmt-undo_log")}

    | task_help {fmt.Println("command_stmt-task_help")}
    | task_list {fmt.Println("command_stmt-task_list")}
    | task_add {fmt.Println("command_stmt-task_add")}
    | task_delete {fmt.Println("command_stmt-task_delete")}
    | task_set {fmt.Println("command_stmt-task_set")}
    | task_done {fmt.Println("command_stmt-task_done")}
    
    | tag_help {fmt.Println("command_stmt-tag_help")}
    | tag_list {fmt.Println("command_stmt-tag_list")}
    | tag_set {fmt.Println("command_stmt-tag_set")}
    ;

// ========== LOG =============
log_list:
      LOG {/*fmt.Println("log")*/}
    | LOG NUM {/*fmt.Println("log")*/}
    ;

undo_log:
      UNDO {/*fmt.Println("log")*/}
    | UNDO id {/*fmt.Println("log")*/}
    ;

// ========== TASK COMMAND =============
task_help:
      TASK HELP {/*fmt.Println("task_help")*/}
    | task_list HELP {/*fmt.Println("task_help")*/}
    | task_add HELP {/*fmt.Println("task_help")*/}
    | task_delete HELP {/*fmt.Println("task_help")*/}
    | task_set HELP {/*fmt.Println("task_help")*/}
    ;

task_list:
      TASK task_list_filter {/*fmt.Println("task_list")*/}
    | task_list_filter {/*fmt.Println("task_list")*/}
    ;

task_add:
      TASK ADD content task_add_filter {/*fmt.Println("task_add")*/}
    | ADD content task_add_filter {/*fmt.Println("task_add")*/}
    | content ADD task_add_filter {/*fmt.Println("task_add")*/}
    ;

task_done:
      TASK DONE id_group {/*fmt.Println("task_done")*/}
    | DONE id_group {/*fmt.Println("task_done")*/}
    | id_group DONE {/*fmt.Println("task_done")*/}
    ;

task_delete:
      TASK DELETE id_group {/*fmt.Println("task_delete")*/}
    | DELETE id_group {/*fmt.Println("task_delete")*/}
    | id_group DELETE {/*fmt.Println("task_delete")*/}
    ;

task_set:
      TASK SET id content task_update_filter {/*fmt.Println("task_set")*/}
    | SET id content task_update_filter {/*fmt.Println("task_set")*/}
    | TASK id SET content task_update_filter {/*fmt.Println("task_set")*/}
    | id SET content task_update_filter {/*fmt.Println("task_set")*/}
    | TASK id content task_update_filter {/*fmt.Println("task_set")*/}
    | id content task_update_filter {/*fmt.Println("task_set")*/}
    ;


// ========== TAG COMMAND =============
tag_help:HELP
      TAG HELP {/*fmt.Println("tag_help")*/}
    | tag_list HELP {/*fmt.Println("tag_help")*/}
    | tag_set HELP  {/*fmt.Println("tag_help")*/}
    ;

tag_list:
      TAG tag_list_filter {/*fmt.Println("tag_list")*/}
    ;

tag_set:
      TAG SET id content {/*fmt.Println("task_set")*/}
    ;



// ========== TASK FILTER =============
task_list_filter:
     {/*fmt.Println("task_list_filter-end")*/}
    | id_group task_list_filter {/*fmt.Println("task_list_filter")*/}
    | LIKE content_group task_list_filter {/*fmt.Println("task_list_filter")*/}
    | content_group task_list_filter {/*fmt.Println("task_list_filter")*/}
    | assign_group task_list_filter {/*fmt.Println("task_list_filter")*/}
    | AGE COLON time_list_filter task_list_filter {/*fmt.Println("task_list_filter")*/}
    | DUE COLON time_list_filter task_list_filter {/*fmt.Println("task_list_filter")*/}
    ;

task_add_filter:
      {/*fmt.Println("task_add_filter")*/}
    | positive_assign_group task_add_filter {/*fmt.Println("task_add_filter")*/}
    | DUE COLON time_single {/*fmt.Println("task_add_filter")*/}
    | LOOP COLON loop_time {/*fmt.Println("task_add_filter")*/}
    ;

task_update_filter:
      {/*fmt.Println("task_update_filter")*/}
    | assign_group task_update_filter {/*fmt.Println("task_update_filter")*/}
    | DUE COLON time_single {/*fmt.Println("task_update_filter")*/}
    | LOOP COLON loop_time {/*fmt.Println("task_update_filter")*/}
    ;

// ========== TAG FILTER =============
tag_list_filter:
      {/*fmt.Println("task_list_filter")*/}
    | id_group tag_list_filter {/*fmt.Println("task_list_filter")*/}
    | LIKE content_group tag_list_filter {/*fmt.Println("task_list_filter")*/}
    ;

// ========== COMMON =============
id_group:
      id id_group {
        $$ =  $1 + $2
        // fmt.Println("id_group:", $1, $2)
      }
    | id MINUS id {
        $$ = $1 + $2
        // fmt.Println("id_group:",$1, $3)
      }
    | id {
        $$ = $1
        // fmt.Println("id_group:", $1)
      }
    ;

id: 
      NUM {
        $$ = $1
        // fmt.Printf("id:%s\n", $1) 
      }
    ; 

content_group:
      content_logic {/*fmt.Println("content_group")*/}
    | content {/*fmt.Println("content_group")*/}
    ;

content_logic:
      content AND content {/*fmt.Println("content_logic")*/}
    | content OR content {/*fmt.Println("content_logic")*/}
    | content XOR content {/*fmt.Println("content_logic")*/}
    | NOT content {/*fmt.Println("content_logic")*/}
    ;

assign_group:
      assign_tag assign_group {/*fmt.Println("assign_group")*/}
    | unassign_tag assign_group {/*fmt.Println("assign_group")*/}
    | {/*fmt.Println("assign_group")*/}
    ;

positive_assign_group:
      assign_tag positive_assign_group {/*fmt.Println("positive_assign_group")*/}
    | {/*fmt.Println("positive_assign_group")*/}
    ;

assign_tag: 
      PLUS IDENT  {/*fmt.Println("assign_tag")*/}
    ;

unassign_tag: 
      MINUS IDENT  {/*fmt.Println("unassign_tag")*/}
    ;

content:
      DQUOTE shard_content DQUOTE {}
    | QUOTE shard_content QUOTE {}
    | DQUOTE content DQUOTE {}
    | QUOTE content QUOTE {}
    | shard_content {}

// 转义所有关键字
shard_content:
      {fmt.Println("content-end")}
    | IDENT shard_content {fmt.Println("content-merge")}
    | id_group shard_content {fmt.Println("content-merge")}
    | ADD shard_content {fmt.Println("content-merge")}
    | DELETE shard_content {fmt.Println("content-merge")}
    | SET shard_content {fmt.Println("content-merge")}
    | DONE shard_content {fmt.Println("content-merge")}

    | AGE shard_content {fmt.Println("content-merge")}
    | DUE shard_content {fmt.Println("content-merge")}
    | LIKE shard_content {fmt.Println("content-merge")}
    | LOOP shard_content {fmt.Println("content-merge")}

    | COLON shard_content {fmt.Println("content-merge")}
    | PLUS shard_content {fmt.Println("content-merge")}
    | MINUS shard_content {fmt.Println("content-merge")}
    | AND shard_content {fmt.Println("content-merge")}
    | OR shard_content {fmt.Println("content-merge")}
    | XOR shard_content {fmt.Println("content-merge")}
    | NOT shard_content {fmt.Println("content-merge")}
    ;

time_list_filter:
      time_single {/*fmt.Println("time_list_filter")*/}
    | time_range {/*fmt.Println("time_list_filter")*/}
    ;

time_single:
      IDENT {/*fmt.Println("time_single")*/}
    | {/*fmt.Println("time_single")*/}
    ;

time_range:
      time_single MINUS time_single {/*fmt.Println("time_range")*/}
    ;

loop_time:
      IDENT {/*fmt.Println("loop_time")*/}
     ;

%%