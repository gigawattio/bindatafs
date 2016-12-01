package bindatafs

//go:generate bash -c "( command -v go-bindata || go get -u github.com/jteeuwen/go-bindata/... ) && echo 'Generating bindatafs package from templates/...' && cd templates && go-bindata -pkg bindatafs -o ../templates_bindatafs.go ./... && gofmt -w ../templates_bindatafs.go || ( echo 'oops.. there was a problem, you may need to install go-bindata or fix path issues, see README.md' && exit 1 )"
