package main

var root = string(os.PathSeparator)

func getExec() string {
	return "open"
}

func getArgs(dir string) []string {
	return []string{"-na", "Goland.app", "--args", dir}
}
