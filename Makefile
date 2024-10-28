all: configure compile run

configure:
	mkdir -p bin/

compile:
	go build -o ./bin/pdfgener ./cmd/pdfgener/main.go

run:
	bin/pdfgener

