package assert

func Equal(message string, expected, actual interface{}) error {
	if expected != actual {
		return exception(message, expected, actual)
	}
	return nil
}

func NotEqual(message string, expected, actual interface{}) error {
	if expected == actual {
		return exception(message, expected, actual)
	}
	return nil
}
