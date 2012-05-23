all:config.6 main.6
	go tool 6l -o test main.6
config.6:config.go
	go tool 6g config.go
main.6:config.go
	go tool 6g main.go

clean:
	rm -f *.6 test
