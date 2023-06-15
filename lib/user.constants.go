package lib

const Password = "Welcome1!"

const (
	DefaultEmail = "api_default@gmail.com"
	ActiveUser   = "api_active_user@gmail.com"
	ReadUser     = "api_read_user@gmail.com"
	WriteUser    = "api_write_user@yopmail.com"
	AdminUser    = "api_admin_user@gmail.com"
)

const (
	ReadAcl  = "read"
	WriteAcl = "write"
	AdminAcl = "admin"
)

type KeyType int

const (
	Custom KeyType = iota
	Teacher
	Student
	Project
)

var KeyTypes = map[KeyType]string{
	Custom:  "CustomKey",
	Teacher: "TeacherKey",
	Student: "StudentKey",
	Project: "ProjectKey",
}

const (
	AlphabeticScaleType = "Alphabetic"
	NumericScaleType    = "Numeric"
)
