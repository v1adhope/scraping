.SILENT:

build:
	go build -o .bin/scraping main.go

run: build
	./.bin/scraping
