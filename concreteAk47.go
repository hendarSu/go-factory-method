package main

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			"AK47 gun",
			4,
		},
	}
}
