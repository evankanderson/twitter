module github.com/vaikas/twitter

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/dghubble/go-twitter v0.0.0-20200725221434-4bc8ad7ad1b4
	github.com/dghubble/oauth1 v0.6.0
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/zap v1.16.0
	knative.dev/pkg v0.0.0-20200929022728-4efcf05498a5
)

replace github.com/dghubble/go-twitter => github.com/drswork/go-twitter v0.0.0-20190721142740-110a39637298
