package ProjectPublicPage
/*根据大小写区分权限
大写方法名外部包可以调用变量，变量命也要大写*/
type UserInfo struct {
	Name   string
	Age    int
	Sex    int
}
type userAddress struct {
	city    string
	country string
}