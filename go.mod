module github.com/mrsufgi/todo_api

go 1.15

require (
	github.com/golang/mock v1.4.4
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.0.0
	github.com/prometheus/client_golang v1.7.0
	github.com/sirupsen/logrus v1.6.0
	github.com/uber/athenadriver v1.1.12
	google.golang.org/appengine v1.6.7
)

replace github.com/mrsufgi/todo_api => ./
