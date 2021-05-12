package pinyin

func Easy(s string) string {
	return Get(s).String()
}

func Get(s string) Pinyin {
	n := 0
	for range s {
		n++
	}
	py := make(Pinyin, 0, n)
	for _, c := range s {
		if n, ok := lookupMapping(c); ok {
			code := int(lookup[n])
			if code != 0 {
				py = append(py, pinyin{
					lookup: int(code >> 3),
					poss:   int(code&0x07) + 1,
				})
			}
		}
	}
	return py
}

type Style int

const (
	Full = Style(iota)
	FullUpper
	Initial
	InitialUpper
	First
	FirstUpper

	DefaultStyle = Full
)

type Pinyin []pinyin
type pinyin struct {
	lookup int
	poss   int
}

func (py Pinyin) Possibility() int {
	if len(py) == 0 {
		return 0
	}
	p := 1
	for _, pyi := range py {
		p *= pyi.poss
	}
	return p
}
func (py Pinyin) String() string {
	return py.StringStyle(DefaultStyle)
}
func (py Pinyin) BywordStyle(style Style) []string {
	return py.BywordIndex(0, DefaultStyle)
}
func (py Pinyin) BywordIndex(n int, style Style) []string {
	s := make([]string, len(py))
	for i, v := range py {
		p := int(pyTable[v.lookup+(n%v.poss)])
		n = n / v.poss

		switch style {
		case Full:
			s[i] = pinyinCombines[p]
		case FullUpper:
			s[i] = pinyinCombinesU[p]
		case Initial:
			s[i] = pinyinInitial[p]
		case InitialUpper:
			s[i] = pinyinInitialU[p]
		case First:
			s[i] = pinyinFirst[p]
		case FirstUpper:
			s[i] = pinyinFirstU[p]
		}
	}
	return s
}
func (py Pinyin) StringStyle(style Style) string {
	return py.StringIndex(0, style)
}
func (py Pinyin) StringIndex(n int, style Style) string {
	w, n := py.BywordIndex(n, style), 0
	for _, c := range w {
		n += len(c)
	}
	b, n := make([]byte, n), 0
	for _, c := range w {
		copy(b[n:], c)
		n += len(c)
	}
	return string(b)
}
