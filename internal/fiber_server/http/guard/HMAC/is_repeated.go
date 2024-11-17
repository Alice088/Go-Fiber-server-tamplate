package HMAC

func IsUsed(nonce string, signature string, values *map[interface{}]interface{}) bool {
	if (*values)[signature] != nil {
		return true
	}
	(*values)[signature] = 1
	return false
}
