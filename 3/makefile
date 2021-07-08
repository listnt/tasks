go:
	go tool asm 3.func.s
	go build -o asd
	./asd
	sudo perf stat -r 100 ./asd
notgo:
	go build ../3.1.go
	./3.1
	sudo perf stat -r 100 ./3.1