package archiver

import "time"

const (
	StatusWaiting  = "waiting"
	StatusRunning  = "running"
	StatusComplete = "complete"
)

type Archiver struct {
	archives map[string]*Archive
}

func New() *Archiver {
	return &Archiver{
		archives: map[string]*Archive{},
	}
}

func (a *Archiver) Get(id string) *Archive {
	archive, found := a.archives[id]
	if !found {
		archive = &Archive{
			Status: StatusWaiting,
		}
		a.archives[id] = archive
	}
	return archive
}

type Archive struct {
	Status   string
	Progress float64
}

func (a *Archive) Run() {
	a.Status = StatusRunning
	go func() {
		for a.Progress < 1 {
			a.Progress = a.Progress + 0.01
			time.Sleep(100 * time.Millisecond)
		}
		a.Status = StatusComplete
	}()
}

func (a *Archive) Reset() {
	a.Status = StatusWaiting
	a.Progress = 0
}

func (a *Archive) ArchiveFile() []byte {
	return []byte(`{"hello": "world"}`)
}
