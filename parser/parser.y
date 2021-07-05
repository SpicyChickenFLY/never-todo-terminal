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

%token EOF

%token <num> NUM 
%token <str> IDENT

%token <str> PLUS MINUS
%token <str> COLON

%token <str> NOT
%left <str> AND OR 
%left <str> XOR

%token <str> HELP UI GUI EXPLAIN LOG UNDO 
%token <str> TASK TAG ADD DELETE SET DONE
%token <str> AGE DUE LIKE LOOP 

%right PLUS MINUS

%type <cmd> command

%%
command:
      command_stmt EOF {}

command_stmt:
      HELP {fmt.Println("command")}
    | UI {fmt.Println("command")}
    | GUI {fmt.Println("command")}
    | EXPLAIN {fmt.Println("command")}

    | log_find  {fmt.Println("command")}
    | undo_log {fmt.Println("command")}

    | task_help {fmt.Println("command")}
    | task_find {fmt.Println("command")}
    | task_add {fmt.Println("command")}
    | task_delete {fmt.Println("command")}
    | task_set {fmt.Println("command")}
    | task_done {fmt.Println("command")}
    
    | tag_help {fmt.Println("command")}
    | tag_find {fmt.Println("command")}
    | tag_set {fmt.Println("command")}
    ;

// ========== LOG =============
log_find:
      LOG {fmt.Println("log")}
    | LOG NUM {fmt.Println("log")}
    ;

undo_log:
      UNDO {fmt.Println("log")}
    | UNDO id {fmt.Println("log")}
    ;

// ========== TASK COMMAND =============
task_help:
      TASK HELP {fmt.Println("task_help")}
    | task_find HELP {fmt.Println("task_help")}
    | task_add HELP {fmt.Println("task_help")}
    | task_delete HELP {fmt.Println("task_help")}
    | task_set HELP {fmt.Println("task_help")}
    ;

task_find:
      TASK task_find_filter {fmt.Println("task_find")}
    | task_find_filter {fmt.Println("task_find")}
    ;

task_add:
      TASK ADD content task_add_filter {fmt.Println("task_add")}
    | content ADD task_add_filter {fmt.Println("task_add")}
    ;

task_done:
      TASK DONE id_group {fmt.Println("task_done")}
    | DONE id_group {fmt.Println("task_done")}
    | id_group DONE {fmt.Println("task_done")}
    ;

task_delete:
      TASK DELETE id_group {fmt.Println("task_delete")}
    | DELETE id_group {fmt.Println("task_delete")}
    | id_group DELETE {fmt.Println("task_delete")}
    ;

task_set:
      TASK SET id content task_update_filter {fmt.Println("task_set")}
    | SET id content task_update_filter {fmt.Println("task_set")}
    | TASK id SET content task_update_filter {fmt.Println("task_set")}
    | id SET content task_update_filter {fmt.Println("task_set")}
    | TASK id content task_update_filter {fmt.Println("task_set")}
    | id content task_update_filter {fmt.Println("task_set")}
    ;

// ========== TAG COMMAND =============
tag_help:HELP
      TAG HELP {fmt.Println("tag_help")}
    | tag_find HELP {fmt.Println("tag_help")}
    | tag_set HELP  {fmt.Println("tag_help")}
    ;

tag_find:
      TAG tag_find_filter {fmt.Println("tag_find")}
    ;

tag_set:
      TAG SET id content {fmt.Println("task_set")}
    ;

// ========== TASK FILTER =============
task_find_filter:
     {fmt.Println("task_find_filter")}
    | id_group task_find_filter {fmt.Println("task_find_filter")}
    | LIKE content_group task_find_filter {fmt.Println("task_find_filter")}
    | assign_group task_find_filter {fmt.Println("task_find_filter")}
    | AGE COLON time_find_filter task_find_filter {fmt.Println("task_find_filter")}
    | DUE COLON time_find_filter task_find_filter {fmt.Println("task_find_filter")}
    ;

task_add_filter:
      {fmt.Println("task_add_filter")}
    | positive_assign_group task_add_filter {fmt.Println("task_add_filter")}
    | DUE COLON time_single {fmt.Println("task_add_filter")}
    | LOOP COLON loop_time {fmt.Println("task_add_filter")}
    ;

task_update_filter:
      {fmt.Println("task_update_filter")}
    | assign_group task_update_filter {fmt.Println("task_update_filter")}
    | DUE COLON time_single {fmt.Println("task_update_filter")}
    | LOOP COLON loop_time {fmt.Println("task_update_filter")}
    ;

// ========== TAG FILTER =============
tag_find_filter:
      {fmt.Println("task_find_filter")}
    | id_group tag_find_filter {fmt.Println("task_find_filter")}
    | LIKE content_group tag_find_filter {fmt.Println("task_find_filter")}
    ;

// ========== COMMON =============
id_group:
      id id_group {fmt.Println("id_group")}
    | id MINUS id {fmt.Println("id_group")}
    | id {fmt.Println("id_group")}
    ;

id: 
      NUM {fmt.Println("id")}
    ; 

content_group:
      content_logic {fmt.Println("content_group")}
    | content {fmt.Println("content_group")}
    ;

content_logic:
      content AND content {fmt.Println("content_logic")}
    | content OR content {fmt.Println("content_logic")}
    | content XOR content {fmt.Println("content_logic")}
    | NOT content {fmt.Println("content_logic")}
    ;

assign_group:
      assign_tag assign_group {fmt.Println("assign_group")}
    | unassign_tag assign_group {fmt.Println("assign_group")}
    | {fmt.Println("assign_group")}
    ;

positive_assign_group:
      assign_tag positive_assign_group {fmt.Println("positive_assign_group")}
    | {fmt.Println("positive_assign_group")}
    ;

assign_tag: 
      PLUS IDENT  {fmt.Println("assign_tag")}
    ;

unassign_tag: 
      MINUS IDENT  {fmt.Println("unassign_tag")}
    ;

content:
      {fmt.Println("content")}
    | IDENT content {fmt.Println("content")}
    ;

time_find_filter:
      time_single {fmt.Println("time_find_filter")}
    | time_range {fmt.Println("time_find_filter")}
    ;

time_single:
      IDENT {fmt.Println("time_single")}
    | {fmt.Println("time_single")}
    ;

time_range:
      time_single MINUS time_single {fmt.Println("time_range")}
    ;

loop_time:
      IDENT {fmt.Println("loop_time")}
     ;

%%