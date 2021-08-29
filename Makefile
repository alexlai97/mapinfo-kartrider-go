.PHONY: test
test:
	echo "WIP"

.PHONY: clean
clean:
	rm -r ./instance/db.sqlite3

.PHONY: build
build:
	go build
