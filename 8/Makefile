go:
	go tool asm 8.func.s
	go build -o asd
	./asd < input.txt
	sudo perf stat -r 100 ./asd
notgo:
	go build ../8.1.go
	./8.1 < input.txt
	sudo perf stat -r 100 ./8.1 < input.txt