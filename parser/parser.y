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

%token TODO TAG ADD DONE EDIT AGE DUE LIKE LOOP

%%

task_find:
    TODO task_find_filter
    | task_find_filter
    

task_add:
    TODO ADD task_content task_add_filter
    | task_content ADD task_add_filter 

task_find_filter:
    task_find_filter
    | id_group task_find_filter {}
    | LIKE content_group task_find_filter {}
    | assign_group task_find_filter {}
    | AGE COLON time_find_filter task_find_filter {}
    | DUE COLON time_find_filter task_find_filter {}
    | {/*empty*/}

task_add_filter:
    task_add_filter
    | positive_assign_group task_add_filter {}
    | DUE COLON time_single
    | AGE COLON time_single
    | LOOP COLON loop_time

id_group:
    task_id id_group
    | task_id DASH task_id
    | task_id

task_id: NUM {}

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
    STR {}

time_find_filter:
    time_single {}
    | time_range {}

time_single:
    STR {}
    | {/*empty*/}

time_range:
    time_single DASH time_single

loop_time:


%%