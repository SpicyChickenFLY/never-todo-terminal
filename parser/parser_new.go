package parser

import (
	"strconv"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/ast"
	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/utils"
)

// command alias
var (
	cmdAliasExplain = []string{"explain"}
	cmdAliasUI      = []string{"ui"}
	cmdAliasHelp    = []string{"-h", "--help"}
	cmdAliasList    = []string{"ls", "list"}
	cmdAliasAdd     = []string{"add", "create"}
	cmdAliasDel     = []string{"del", "delete", "rm", "remove"}
	cmdAliasSet     = []string{"set", "update", "upd"}
	cmdAliasTodo    = []string{"todo", "td"}
	cmdAliasDone    = []string{"done", "complete"}
	cmdAliasTag     = []string{"tag"}
)

func ParseRoot(cmdParts []string) *ast.RootNode {
	partLen := len(cmdParts)
	needExplain := utils.MatchWord(cmdAliasExplain, cmdParts[0])
	if needExplain {
		if partLen == 1 {
			return ast.NewRootNode(ast.CMDExplain, nil)
		}
		result := parseStmt(cmdParts[1:])
		if result.CmdType == ast.CMDStmt {
			result.CmdType = ast.CMDExplain
			return result
		}
		return ast.NewRootNode(ast.CMDExplain, nil)
	}
	return parseStmt(cmdParts)
}

func parseStmt(cmdParts []string) (result *ast.RootNode) {
	partLen := len(cmdParts)
	if partLen == 0 {
		return ast.NewRootNode(ast.CMDSummary, nil)
	}
	needHelp := utils.MatchWord(cmdAliasHelp, cmdParts[partLen-1])
	switch {
	case utils.MatchWord(cmdAliasHelp, cmdParts[0]):
		return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpRoot))
	case utils.MatchWord(cmdAliasList, cmdParts[0]):

	case utils.MatchWord(cmdAliasAdd, cmdParts[0]):
		if partLen == 1 || needHelp {
			return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTaskAdd))
		}
		return ast.NewRootNode(ast.CMDStmt, parseTaskAddCmd(cmdParts[1:]))
	case utils.MatchWord(cmdAliasDel, cmdParts[0]):
		if partLen == 1 || needHelp {
			return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTaskDelete))
		}
		return ast.NewRootNode(ast.CMDStmt, parseTaskDeleteCmd(cmdParts[1:]))
	case utils.MatchWord(cmdAliasSet, cmdParts[0]):
		if partLen == 1 || needHelp {
			return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTaskUpdate))
		}
		return ast.NewRootNode(ast.CMDStmt, parseTaskUpdateCmd(cmdParts[1:]))
	case utils.MatchWord(cmdAliasTodo, cmdParts[0]):
		if partLen == 1 || needHelp {
			return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTaskTodo))
		}
		return ast.NewRootNode(ast.CMDStmt, parseTaskTodoCmd(cmdParts[1:]))
	case utils.MatchWord(cmdAliasDone, cmdParts[0]):
		if partLen == 1 || needHelp {
			return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTaskDone))
		}
		return ast.NewRootNode(ast.CMDStmt, parseTaskDoneCmd(cmdParts[1:]))
	case utils.MatchWord(cmdAliasTag, cmdParts[0]):
		if needHelp {
			return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTag))
		}
		if partLen == 1 { // list tag

		}
		switch {
		case utils.MatchWord(cmdAliasAdd, cmdParts[1]):
			if partLen == 2 || needHelp {
				return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTagAdd))
			}
			return ast.NewRootNode(ast.CMDStmt, parseTagAddCmd(cmdParts[2:]))
		case utils.MatchWord(cmdAliasDel, cmdParts[1]):
			if partLen == 2 || needHelp {
				return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTagDelete))
			}
			return ast.NewRootNode(ast.CMDStmt, parseTagDeleteCmd(cmdParts[2:]))
		case utils.MatchWord(cmdAliasSet, cmdParts[1]):
			if partLen == 2 || needHelp {
				return ast.NewRootNode(ast.CMDHelp, ast.NewHelpNode(ast.HelpTagUpdate))
			}
			return ast.NewRootNode(ast.CMDStmt, parseTagUpdateCmd(cmdParts[2:]))
		default: // List
		}
	default:

	}
	return result
}

func parseTaskAddCmd(cmdPart []string) *ast.TaskAddNode {
	content := ""
	opt := ast.NewTaskAddOptionNode()
	agn := ast.NewAssignGroupNode("", "")
	for i := len(cmdPart); i > 0; i-- {
		lexType, lexVal := lex(cmdPart[i])
		switch lexType {
		case Assign:
			agn.AssignTag(lexVal.(string))
		case Unassign: // FIXME: no unassgin in task add
			content = strings.Join(cmdPart[0:i], " ") + content
			i = 0
		case Importance:
			opt.SetImportance(lexVal.(int))
		case Due:
			// opt.SetDue()
		case Loop:
			opt.SetLoop(ast.NewLoopNode(lexVal.(string)))
		default: // no more option, rest cmd are all content
			content = strings.Join(cmdPart[0:i], " ") + content
			i = 0
		}
	}
	opt.SetAssignGroup(agn)
	return ast.NewTaskAddNode(content, opt)
}

func parseTaskUpdateCmd(cmdParts []string) *ast.TaskUpdateNode {
	var id int
	ugn := ast.NewTaskUpdateOptionNode()
	// TODO:
	return ast.NewTaskUpdateNode(id, ugn)
}

func parseTaskDeleteCmd(cmdParts []string) *ast.TaskDeleteNode {
	ign := ast.NewIDGroupNode()
	for i := 0; i < len(cmdParts); i++ {
		ignSub := parseIDGroup(cmdParts[i])
		if ignSub != nil {
			ast.WarnList = append(ast.WarnList, "Parameter cannot be parsed as ID group")
			break
		}
		ign.MergeIDNode(ignSub)
	}
	return ast.NewTaskDeleteNode(ign)
}

func parseTaskTodoCmd(cmdParts []string) *ast.TaskTodoNode {
	ign := ast.NewIDGroupNode()
	for i := 0; i < len(cmdParts); i++ {
		ignSub := parseIDGroup(cmdParts[i])
		if ignSub != nil {
			ast.WarnList = append(ast.WarnList, "Parameter cannot be parsed as ID group")
			break
		}
		ign.MergeIDNode(ignSub)
	}
	return ast.NewTaskTodoNode(ign)
}

func parseTaskDoneCmd(cmdParts []string) *ast.TaskDoneNode {
	ign := ast.NewIDGroupNode()
	for i := 0; i < len(cmdParts); i++ {
		ignSub := parseIDGroup(cmdParts[i])
		if ignSub != nil {
			ast.WarnList = append(ast.WarnList, "Parameter cannot be parsed as ID group")
			break
		}
		ign.MergeIDNode(ignSub)
	}
	return ast.NewTaskDoneNode(ign)
}

func parseTagAddCmd(cmdParts []string) *ast.TagAddNode    { return nil }
func parseTagDeleteCmd(cmdParts []string) *ast.TagAddNode { return nil }
func parseTagUpdateCmd(cmdParts []string) *ast.TagAddNode { return nil }

func parseContentGroup(contentParts []string) *ast.ContentGroupNode { return nil }
func parseTimeFilter(timeStr string) *ast.TimeFilterNode            { return nil }

func parseIDGroup(idGroupStr string) *ast.IDGroupNode {
	ign := ast.NewIDGroupNode()
	idRanges := strings.Split(idGroupStr, ",")
	for _, idRange := range idRanges {
		ids := []int{}
		var err error
		start, end := 0, 0
		minusIdx := strings.Index(idRange, "-")
		switch minusIdx {
		case -1: // No range
			start, err = strconv.Atoi(idRange)
			end = start
		case 0: // only end
			end, err = strconv.Atoi(idRange[minusIdx+1:])
		case len(idRange): // only start
			start, err = strconv.Atoi(idRange[:minusIdx])
		default:
			start, err = strconv.Atoi(idRange[:minusIdx])
			end, err = strconv.Atoi(idRange[minusIdx+1:])
		}
		if err != nil {
			// FIXME: handle Atoi error
		}
		for num := start; num <= end; num++ {
			ids = append(ids, num)
		}
		ign.MergeIDNode(ast.NewIDGroupNode(ids...))
	}
	return ign
}
