run: build
	./conway data/glider.cells

build: clean
	go build conway/cmd/conway

clean:
	rm -f conway

test:
	go test conway/internal/universe
