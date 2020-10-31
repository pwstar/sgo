package main

import (
	"github.com/ssgo/s"
	"github.com/ssgo/u"
	"path"
)

func main() {
	s.Static("/", "www/")
	s.Restful(0, "POST", "/data/{name}", save)
	s.Restful(0, "GET", "/data/{name}", load)
	s.Restful(0, "PUT", "/data/{name}", push)
	s.Restful(0, "GET", "/data/{name}/{num}", getList)
	s.Restful(0, "GET", "/data/{name}/{num}:{pos}", getList)
	s.Config.HttpVersion = 1
	s.Config.Listen = "80"
	s.Start()
}

func save(in struct{ Name, Data string }) bool {
	return u.WriteFile(path.Join("data", in.Name), in.Data) == nil
}

func load(in struct{ Name string }) string {
	data, _ := u.ReadFile(path.Join("data", in.Name), 1024*1024)
	return data
}

func push(in struct{ Name, Data string }) bool {
	list := make([]string, 0)
	u.Load(path.Join("data", in.Name), &list)
	list = append(list, in.Data)
	return u.Save(path.Join("data", in.Name), list) == nil
}

func getList(in struct {
	Name string
	Pos  uint
	Num  uint
}) []string {
	list := make([]string, 0)
	u.Load(path.Join("data", in.Name), &list)

	// pos=0, num=100, len(list)=500
	//   500-100-0=400 : 500-0=500
	// pos=100, num=100, len(list)=500
	//   500-100-100=300 : 500-100=400
	// pos=200, num=100, len(list)=500
	// 	 500-100-200=200 : 500-200=300

	// pos=600, num=100, len(list)=500 []
	// pos=100, num=100, len(list)=50 []
	// pos=0, num=100, len(list)=150  51~150
	// pos=100, num=100, len(list)=150 1~50
	listLen := uint(len(list))
	if in.Pos < listLen {
		if in.Num+in.Pos < listLen {
			return list[listLen-in.Num-in.Pos : listLen-in.Pos]
		} else {
			return list[0 : listLen-in.Pos]
		}
	} else {
		return make([]string, 0)
	}
}
