package watcher

func main() {

}

type Watcher struct {
	WatchList map[string]chan int
}

func NewWatcher() *Watcher {
	var things map[string](chan int)
	reader := kafka.NewReader()
	return &Watcher{WatchList: things}
}
