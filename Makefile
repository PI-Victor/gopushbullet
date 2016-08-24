build:
	cd cmd/gunner && go build -o ../../_out/bin/gunner -v .

clean:
	rm -r ./_out

install:
	cp -p ./_out/bin/gunner $(GOPATH)/bin
