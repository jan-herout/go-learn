package cc

// Metadata about DDL scripts.
type Metadata struct {
	Path      string // path to the DDL of the table
	TableName string // the table this DDL scripts creates
	Content   []byte // full content of the file in questions
}
