%{
    package main
	import (
		"fmt"
		"os"
		"io/ioutil"
		"flag"
		"bufio"
		)
	var fi *bufio.Reader

%}

%union
{
	vvar   string;
	numval float;
}

%token NUM STR

%token PLUS MINUS
%token COLON DASH

%token NOT
%token AND OR XOR

%token TASK TAG ADD DONE EDIT AGE DUE LIKE LOOP DELETE

%%

command:
    task_find
    | task_add
    | task_delete
    | task_edit
    
    | tag

// ========== TASK COMMAND =============
task_find:
    TASK task_find_filter
    | task_find_filter

task_add:
    TASK ADD task_content task_add_filter
    | task_content ADD task_add_filter

task_done:
    TASK DONE id_group
    | DONE id_group
    | id_group DONE

task_delete:
    TASK DELETE id_group
    | DELETE id_group
    | id_group DELETE

task_edit:
    TASK EDIT task_id task_content task_update_filter
    | EDIT task_id task_content task_update_filter
    | TASK task_id EDIT task_content task_update_filter
    | task_id EDIT task_content task_update_filter 

// ========== TASK FILTER =============
task_find_filter:
    task_find_filter
    | id_group task_find_filter
    | LIKE content_group task_find_filter
    | assign_group task_find_filter
    | AGE COLON time_find_filter task_find_filter
    | DUE COLON time_find_filter task_find_filter
    | {/*empty*/}

task_add_filter:
    task_add_filter
    | positive_assign_group task_add_filter
    | DUE COLON time_single
    | LOOP COLON loop_time
    | {/*empty*/}

task_update_filter:
    task_update_filter
    | assign_group task_update_filter
    | DUE COLON time_single
    | LOOP COLON loop_time
    | {/*empty*/}


// ========== COMMON =============
id_group:
    task_id id_group
    | task_id DASH task_id
    | task_id

task_id: NUM

content_group:
    task_content_logic
    | task_content

task_content_logic:
    task_content AND task_content
    | task_content OR task_content
    | task_content XOR task_content
    | NOT task_content

assign_group:
    assign_tag assign_group
    | unassign_tag assign_group

positive_assign_group:
    assign_group assign_group

assign_tag: 
    PLUS STR 

unassign_tag: 
    MINUS STR 

task_content:
    STR

time_find_filter:
    time_single
    | time_range

time_single:
    STR
    | {/*empty*/}

time_range:
    time_single DASH time_single

loop_time:


%%