package her

import (
	"github.com/go-code/her/notify"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
)

type Watcher struct {
	watchers    []*notify.Watcher
	watcherPath string
	notifyMutex sync.Mutex
}

func NewWatcher() *Watcher {
	return &Watcher{}
}

func (w *Watcher) Listen(p string) {
	w.watcherPath = p
	if w.watcherPath != "" {
		watcher, err := notify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		watcher.Event = make(chan *notify.FileEvent, 100)
		watcher.Error = make(chan error, 10)

		filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Error walking path:", err)
				return nil
			}

			if info.IsDir() {
				err = watcher.Watch(path)
				if err != nil {
					log.Println("Failed to watch", path, ":", err)
				}
			}
			return nil
		})

		w.watchers = append(w.watchers, watcher)
	}
}

func (w *Watcher) Notify() {
	w.notifyMutex.Lock()
	defer w.notifyMutex.Unlock()

	for _, watcher := range w.watchers {
		for {
			select {
			case ev := <-watcher.Event:
				if path.Ext(ev.Name) == Config.GetString("TemplateExt") {
					templates = loadTemplate()
					log.Println("event:", ev)
				}
				continue
			case err := <-watcher.Error:
				log.Println("error:", err)
				continue
			default:
			}
			break
		}
	}
}
