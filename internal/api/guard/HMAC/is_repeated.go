package HMAC

func IsRepeated(signature string, values *map[interface{}]interface{}) bool {
	if (*values)[signature] != nil {
		return true
	} else {
		(*values)[signature] = 1
		return false
	}
}
