package modles

type UserInfo struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//c.ShouldBind(&user)基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象
