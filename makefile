watch:
ifndef $(shell command -v CompileDaemon)
    $(shell go get github.com/githubnemo/CompileDaemon)
endif
    $(shell CompileDaemon -build="go build ./cmd/tagonapi/..." -command="./tagonapi" -color=true -log-prefix=false)
