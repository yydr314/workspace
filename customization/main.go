package customization

import "github.com/yydr314/workspace"

func TestFunc(s string) string {
	return s + " new"
}

func InCusFunc(num int) string {
	return workspace.InMainFunc(num)
}
