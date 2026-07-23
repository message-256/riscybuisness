all:
	gofmt -w vm.go assembler.go
	go build vm.go
	go build assembler.go

