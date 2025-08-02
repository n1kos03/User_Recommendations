module github.com/n1kos03/User_Recommendations/services/recommendations

go 1.24.1

require github.com/n1kos03/User_Recommendations/common v0.0.0

require (
	github.com/golang-migrate/migrate/v4 v4.18.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace github.com/n1kos03/User_Recommendations/common => ../../common
