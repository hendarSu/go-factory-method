package main

type handGun struct {
	gun
}

func newHandGun() iGun {
	return &handGun{
		gun: gun{"Handgun", 1},
	}
}
