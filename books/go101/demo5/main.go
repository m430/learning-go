package main

type L interface {
	Run() error
	Stop()
}

type M interface {
	L
	Step() error
}

type N interface {
	M
	interface{ Resume() }
	~map[int]bool
	~[]byte | string
}

type O interface {
	Pause()
	N
	string
	int64 | ~chan int | any
}
