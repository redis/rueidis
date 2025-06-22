package rueidis

import "testing"

func TestBitmap(t *testing.T) {
	var bm bitmap
	bm.Init(2953)
	if bm.Len() < 2953 {
		t.Errorf("bitmap length = %d, expected at least %d", bm.Len(), 2953)
	}
	for i := 0; i < bm.Len(); i++ {
		bm.Set(i)
		for j := 0; j < bm.Len(); j++ {
			if j <= i {
				if !bm.Get(j) {
					t.Errorf("bitmap[%d] should be set after iteration %d, but is not", j, i)
				}
			} else {
				if bm.Get(j) {
					t.Errorf("bitmap[%d] should not be set after iteration %d, but is", j, i)
				}
			}
		}
	}
}
