package pinyin

import (
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("assert failed: %v != %v", a, b)
	}
}

func TestEasy(t *testing.T) {
	assertEquals(t, Easy("套马的汉子你威武雄壮"), "taomadehanziniweiwuxiongzhuang")
	assertEquals(t, Easy("奔驰的骏马像疾风一样"), "benchidejunmaxiangjifengyiyang")
}

func TestGetPinyin(t *testing.T) {
	assertEquals(t, Get("知之为知之").String(), "zhizhiweizhizhi")
	assertEquals(t, Get("知之为知之").StringStyle(FullUpper), "ZHIZHIWEIZHIZHI")
	assertEquals(t, Get("知之为知之").StringStyle(Initial), "zhzhwzhzh")
	assertEquals(t, Get("知之为知之").StringStyle(InitialUpper), "ZHZHWZHZH")
	assertEquals(t, Get("知之为知之").StringStyle(First), "zzwzz")
	assertEquals(t, Get("知之为知之").StringStyle(FirstUpper), "ZZWZZ")

	py := Get("单和地")
	// 单: dan chan shan
	// 和: he hu huo
	// 地: de di
	assertEquals(t, py.Possibility(), 3*3*2)

	results := make(map[string]int)
	poss := py.Possibility()
	for i := 0; i < poss; i++ {
		results[py.StringIndex(i, DefaultStyle)] = 1
	}

	assertEquals(t, poss, len(results))
}
