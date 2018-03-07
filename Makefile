k8s-safe-cronjob: main.go
	CGO_ENABLED=0 go build
	strip k8s-safe-cronjob
