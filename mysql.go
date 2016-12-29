package yphp

// mysql字符转接
func Mysql_escape_string(escapestr string) string {
	b := []byte(escapestr)
	ret := make([]byte, len(b)*2)

	j := 0
	for _, v := range b {
		switch v {
		case '\r', '\\', '\'', '\n', '"':
			ret[j] = '\\'
			ret[j+1] = v
			j = j + 2
		case '\032':
			ret[j] = '\\'
			ret[j+1] = 'Z'
			j = j + 2
		default:
			ret[j] = v
			j = j + 1
		}
	}
	return string(ret[0:j])
}
