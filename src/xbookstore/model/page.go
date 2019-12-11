package model

type Page struct {
	Books       []*Book //每页查询出来的图书存放的切片
	PageNo      int64   //当前页
	PageSize    int64   //每页显示的条数
	TotalPageNo int64   //总页数，通过计算得到
	TotalRecord int64   //总记录数，通过查询数据库得到
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

//IsHasPrev 是否存在上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 是否存在下一页
func (p *Page) IsHasNext() bool {
	return p.TotalPageNo > p.PageNo
}

//GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		//存在上一页，上一页页码 = 当前页 - 1
		return p.PageNo - 1
	} else {
		//不存在上一页，就是首页咯
		return 1
	}
}

//GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		//存在下一页，下一页页码 = 当前页 + 1
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
