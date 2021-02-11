.PHONY: build clean invoke

build: clean
	env GOOS=linux GOARCH=amd64 go build -o bin/main ./main.go

clean:
	rm -rf ./bin

invoke: clean build
	sam local generate-event kinesis get-records --data "$$(jq -c < ./event_data.json)" | sam local invoke -e - MyLambda