# Welcome to Star's Go World.

You can find some go source and use them.

## Server

Server is a web server and data server.

Please put web content to ```www/``` 

#### Data Api

```go
s.Restful(0, "POST", "/data/{name}", save)
s.Restful(0, "GET", "/data/{name}", load)
s.Restful(0, "PUT", "/data/{name}", push)
s.Restful(0, "GET", "/data/{name}/{num}", getList)
s.Restful(0, "GET", "/data/{name}/{num}:{pos}", getList)

func save(in struct{ Name, Data string }) bool
func load(in struct{ Name string }) string 
func push(in struct{ Name, Data string }) bool 
func getList(in struct {
	Name string
	Pos  uint
	Num  uint
}) []string
```

## Wall

Is a demo for put content to the wall.

```
1、download wall via https://github.com/pwstar/sgo/archive/main.zip 
2、put bin/server.xxx to wall directory
3、run wall/server.xxx
4、open http://localhost
```

