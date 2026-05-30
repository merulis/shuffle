package entity

type Locator struct {
	Path string
	Ref  string
}

func NewLocator(path, ref string) Locator {
	return Locator{
		Path: path,
		Ref:  ref,
	}
}
