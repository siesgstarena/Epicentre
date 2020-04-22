run: bin/epicentre
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/epicentre: main.go
	go build -o bin/epicentre main.go

clean:
	rm -rf bin