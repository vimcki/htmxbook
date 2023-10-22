package archiver

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
			status: StatusWaiting,
		}
		a.archives[id] = archive
	}
	return archive
}

type Archive struct {
	status   string
	progrest float64
}

func (a *Archive) Status() string {
	return a.status
}

func (a *Archive) Progress() float64 {
	return a.progrest
}

func (a *Archive) Run() {
}

func (a *Archive) Reset() {
}

func (a *Archive) ArchiveFile() {
}
