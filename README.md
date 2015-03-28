# martini sample

# setup db

sqlite3 todo.db

create table todo(id integer, content text);

insert into todo(id,content) values (1,'study English');


# to start server

export GOPATH=\`pwd\`

go get github.com/go-martini/martini

go get github.com/mattn/go-sqlite3

go src

go run server.go

