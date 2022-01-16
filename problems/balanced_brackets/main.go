package main

import "fmt"

func main() {
	// s := "{[()]}"
	// s := "{[(])}"
	// s := "{{[[(())]]}}"
	// s := "{(([])[])[]}"
	// s := "{{)[](}}"
	// s := "{(([])[])[]}"
	// s := "{(([])[])[]]}"
	s := "{(([])[])[]}[]"

	fmt.Println(isBalanced(s))
}

func isBalanced(s string) string {
	if len(s)%2 != 0 {
		return "NO"
	}

	mapBrackets := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	bracketStack := []string{string(s[len(s)-1])}
	for i := len(s) - 2; i >= 0; i-- {
		bracket := string(s[i])
		topBracketEnqueued := ""
		if len(bracketStack) > 0 {
			topBracketEnqueued = bracketStack[len(bracketStack)-1]
		}

		if topBracketEnqueued != "" && mapBrackets[bracket] == topBracketEnqueued {

			bracketStack = bracketStack[:len(bracketStack)-1]
			continue
		}

		bracketStack = append(bracketStack, bracket)
	}

	if len(bracketStack) == 0 {
		return "YES"
	}

	return "NO"
}
