#json-diff

Small utility for comparing json documents. Prints out a json document containing those properties in the second document that differ from the first.

##Building & running
```sh
go get github.com/mahe-go/json-diff
cd $GOPATH/github.com/mahe-go/json-diff
go build
./json-diff file1 file2
```

##Running tests
```sh
go test ./...
```