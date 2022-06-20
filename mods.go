package osugo

type Mods int32

const (
	Nomod            Mods = 0
	NoFail           Mods = 1 << 0
	Easy             Mods = 1 << 1
	TouchDevice      Mods = 1 << 2
	Hidden           Mods = 1 << 3
	HardRock         Mods = 1 << 4
	SuddenDeath      Mods = 1 << 5
	DoubleTime       Mods = 1 << 6
	Relax            Mods = 1 << 7
	HalfTime         Mods = 1 << 8
	Nightcore        Mods = 1 << 9
	Flashlight       Mods = 1 << 10
	Autoplay         Mods = 1 << 11
	SpunOut          Mods = 1 << 12
	Relax2           Mods = 1 << 13
	Perfect          Mods = 1 << 14
	Key4             Mods = 1 << 15
	Key5             Mods = 1 << 16
	Key6             Mods = 1 << 17
	Key7             Mods = 1 << 18
	Key8             Mods = 1 << 19
	FadeIn           Mods = 1 << 20
	Random           Mods = 1 << 21
	Cinema           Mods = 1 << 22
	Target           Mods = 1 << 23
	Key9             Mods = 1 << 24
	Key10            Mods = 1 << 25
	Key1             Mods = 1 << 26
	Key3             Mods = 1 << 27
	Key2             Mods = 1 << 28
	LastMod          Mods = 1 << 29
	KeyMod           Mods = Key1 | Key2 | Key3 | Key4 | Key5 | Key6 | Key7 | Key8 | Key9 | Key10
	KeyModUnranked   Mods = Key1 | Key2 | Key3 | Key9 | Key10
	FreeModAllowed   Mods = NoFail | Easy | Hidden | HardRock | SuddenDeath | Flashlight | FadeIn | Relax | Relax2 | SpunOut | KeyMod | KeyMod
	ScoreIncreasMods Mods = Hidden | HardRock | DoubleTime | Flashlight | FadeIn
)
