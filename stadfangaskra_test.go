package stadfangaskra

import (
	"testing"
)

func BenchmarkGeocode(b *testing.B) {

	content := "Furugrund 40, 200 Kópavogur"

	for i := 0; i < b.N; i++ {
		_, _ = Geocode(content)
	}

}
