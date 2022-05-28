package conv

// padding
func padding(bytes []byte, length int) []byte {
	if len(bytes) < length {
		return append(make([]byte, length-len(bytes)), bytes...)
	}
	return bytes
}
