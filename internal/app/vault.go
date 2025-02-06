package app

type Vault struct {
	urls []URLObject
}

func GetVault() *Vault {
	return &Vault{}
}

func (v *Vault) Add(u URLObject) {
	v.urls = append(v.urls, u)
}

func (v *Vault) Find(search string, checker func(URLObject, string) bool) (*URLObject, bool) {
	for _, u := range v.urls {
		if checker(u, search) {
			return &u, true
		}
	}
	return &URLObject{}, false
}
