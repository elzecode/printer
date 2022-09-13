ifeq ($(BINARY_PATH),)
BINARY_PATH := dist
endif

clean:
	rm -rf ./${BINARY_PATH}/*

build:
	docker exec -ti go bash -c "cd /var/server/src/cmd/main && go run main.go"
