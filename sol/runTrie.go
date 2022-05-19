package sol

func Run(commands *[]string, val *[][]string) []string {
	c := *commands
	values := *val
	result := []string{"null"}
	obj := Constructor()
	for idx := 1; idx < len(c); idx++ {
		command := c[idx]
		v := values[idx][0]
		switch command {
		case "insert":
			obj.Insert(v)
			result = append(result, "null")
		case "search":
			if obj.Search(v) {
				result = append(result, "true")
			} else {
				result = append(result, "false")
			}
		case "startsWith":
			if obj.StartsWith(v) {
				result = append(result, "true")
			} else {
				result = append(result, "false")
			}
		}
	}
	return result
}
