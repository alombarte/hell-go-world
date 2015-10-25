# hell-go-world
Stupid go code to learn the syntax

Code organization:

    mkdir -p $HOME/git/goworkspace/{src,pkg,bin}
    git clone git@github.com:alombarte/hell-go-world.git $HOME/git/goworkspace/src/github.com/alombarte/hell-go-world
    export GOPATH=$HOME/git/goworkspace
    export PATH=$PATH:$GOPATH/bin

Install the package under `bin/hell-go-world`

    go install github.com/alombarte/hell-go-world
