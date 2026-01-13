package mongodb

// MongoShellParseError represents a parse error with position information.
type MongoShellParseError struct {
	Line    int
	Column  int
	Message string
}
