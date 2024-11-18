package system

type System struct {
	App    App
	Logger Logger
	Mysql  Mysql
	Jwt    Jwt
	Upload Upload
	Redis  Redis
	Email  Email
	Github Github
	ES     ES
}
