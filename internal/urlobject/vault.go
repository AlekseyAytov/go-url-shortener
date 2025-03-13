package urlobject

type PersistentStorage interface {
	SaveObject(URLObject) error
	ReadObjects() ([]URLObject, error)
}

type Vault struct {
	urls []URLObject
	// fileName string
	storage PersistentStorage
}

func GetVault(storage PersistentStorage) *Vault {
	v := Vault{storage: storage}
	objs, err := v.storage.ReadObjects()
	if err != nil {
		// TODO: обработать ошибку
		return &Vault{}
	}
	v.urls = objs
	return &v
}

func (v *Vault) Add(u URLObject) {
	v.urls = append(v.urls, u)
	// TODO: обработать ошибку
	_ = v.storage.SaveObject(u)
}

func (v *Vault) Find(search string, checker func(URLObject, string) bool) (*URLObject, bool) {
	for _, u := range v.urls {
		if checker(u, search) {
			return &u, true
		}
	}
	return &URLObject{}, false
}
