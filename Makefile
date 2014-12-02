cmd/get-git-ignore/get-git-ignore: *.go
	cd cmd/get-git-ignore && go build

~/bin/get-git-ignore: cmd/get-git-ignore/get-git-ignore
	ln -s $GOPATH/src/github.com/handlename/go-get-git-ignore/cmd/get-git-ignore/get-git-ignore ~/bin/get-git-ignore
