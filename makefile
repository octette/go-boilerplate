watch:
ifndef $(shell command -v CompileDaemon)
    $(shell go get github.com/githubnemo/CompileDaemon)
endif
    $(shell CompileDaemon -build="go build ./cmd/goboilerplate/..." -command="./goboilerplate" -color=true -log-prefix=false)
