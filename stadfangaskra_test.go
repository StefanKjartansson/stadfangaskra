package stadfangaskra

import (
	"testing"
)

func BenchmarkGeocode(b *testing.B) {

	f, _ := ByLocationString("Furugrund 40, 200 Kópavogur")

	if f == nil {
		return
	}

	for i := 0; i < b.N; i++ {
		_, _ = Stadfangaskra.Find(f)
	}

}
