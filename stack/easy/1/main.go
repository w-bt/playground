package main

func ValidParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	stackList := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "{" || string(s[i]) == "(" || string(s[i]) == "[" {
			stackList = append(stackList, string(s[i]))
		} else {
			if len(stackList) == 0 {
				return false
			}

			if (string(s[i]) == "}" && stackList[len(stackList)-1] == "{") ||
				(string(s[i]) == "]" && stackList[len(stackList)-1] == "[") ||
				(string(s[i]) == ")" && stackList[len(stackList)-1] == "(") {
				stackList = stackList[:len(stackList)-1]
			}
		}
	}

	return len(stackList) == 0
}
