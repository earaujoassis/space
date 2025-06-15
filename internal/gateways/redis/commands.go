package redis

// Do is used to send commands to a Redis datastore, throguh an active connection
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return memoryStore.Do(commandName, args...)
}
