go:
	go tool asm 3.func.s
	go build -o asd
	./asd
	sudo perf stat -r 100 ./asd
notgo:
	go build -gcflags '-N -l' ../3.1.go # вырубаем оптимизацию, по дефолту она включена, и разницы не видно
	./3.1
	sudo perf stat -r 100 ./3.1
