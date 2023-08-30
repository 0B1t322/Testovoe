## Launch app
type into terminal:
```shell
go run cmd/http/main.go
```

this command start http server that listen port `:8080`

## endpoints

| endpoint       | descriptions                                                                  |
|----------------|-------------------------------------------------------------------------------|
| /api/get-items | get items. you can filter items by id.<br/>example `api/get-items?id=120,116` |
