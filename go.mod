module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
)
