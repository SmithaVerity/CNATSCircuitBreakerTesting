cd gobreaker/

go mod init client-go
go get github.com/sony/gobreaker
go mod tidy
go run client-go/main.go
