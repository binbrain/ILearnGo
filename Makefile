# Run tests on submodule
#

build:
	docker-compose build

# make app=fib run
run:
	docker-compose run -w /go learngo go run $(app)

# make app=binarytrees/bst test
test:
	docker-compose run -w /go learngo go test $(app)

# make test app=binarytrees/bst t=Find
t1:
	docker-compose run -w /go learngo go test -run $(t) $(app)

bench:
	docker-compose run -w /go learngo go test -bench -benchmem $(app)
