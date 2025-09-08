module github.com/n1kos03/User_Recommendations/common

go 1.24.1

require (
	github.com/confluentinc/confluent-kafka-go/v2 v2.11.0
	github.com/golang-migrate/migrate/v4 v4.18.3
	github.com/lib/pq v1.10.9
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

// replace github.com/n1kos03/User_Recommendations/services/recommendations => ../services/recommendations
