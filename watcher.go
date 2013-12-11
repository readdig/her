package handy

import (
	"github.com/go-web-framework/handy/notify"
	"log"
)

type Watcher struct {
	watcher     *notify.Watcher
	watcherPath string
}

func NewWatcher() *Watcher {
	return &Watcher{watcher: nil}
}

func (w *Watcher) Listen(path string) {
	w.watcherPath = path
	if w.watcherPath != "" {
		watcher, err := notify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		watcher.Event = make(chan *notify.FileEvent, 100)
		watcher.Error = make(chan error, 10)

		err = watcher.Watch(path)
		if err != nil {
			log.Fatal(err)
		}
		w.watcher = watcher
	}
}

func (w *Watcher) Notify() {
	if w.watcherPath != "" {
		watcher := w.watcher
		go func() {
			for {
				select {
				case ev := <-watcher.Event:
					templates = loadTemplate()
					log.Println("event:", ev)
					continue
				case err := <-watcher.Error:
					log.Println("error:", err)
					continue
				}
			}
		}()
	}
}
