BINS=webby

$(BINS): webby.go
	GOOS=linux GOARCH=amd64 go build .

clean:
	rm -rf $(BINS)

all: clean $(BINS)
