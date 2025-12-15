package turn

import "time"

type Phase string

const (
	PhasePVE  Phase = "PVE"
	PhasePVP  Phase = "PVP"
	PhaseBOSS Phase = "BOSS"
)

type Player struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	HP    int    `json:"hp"`
	Alive bool   `json:"alive"`
}

type Match struct {
	ID          string    `json:"id"`
	Round       int       `json:"round"`
	Phase       Phase     `json:"phase"`
	ActiveIndex int       `json:"activeIndex"`
	Players     []Player  `json:"players"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	pveLeft  int `json:"-"`
	pvpLeft  int `json:"-"`
	bossLeft int `json:"-"`
}
