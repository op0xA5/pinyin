package pinyin

import "testing"

func TestEasy(t *testing.T) {
	t.Log(Easy("套马的汉子你威武雄壮"))
}

func TestGetPinyin(t *testing.T) {
	t.Log(Get("天安门").String())
	t.Log(Get("天安门").StringStyle(FullUpper))
	t.Log(Get("天安门").StringStyle(Initial))
	t.Log(Get("天安门").StringStyle(InitialUpper))
	t.Log(Get("天安门").StringStyle(First))
	t.Log(Get("天安门").StringStyle(FirstUpper))

	py := Get("单和地")
	poss := py.Possibility()
	t.Log("Possibility", poss)
	for i := 0; i < poss; i++ {
		t.Log(py.StringIndex(i, DefaultStyle))
	}
}
