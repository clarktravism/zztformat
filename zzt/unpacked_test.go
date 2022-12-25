package zzt

import (
	"os"
	"testing"
)

func TestUnpackPack(t *testing.T) {

	b, err := os.ReadFile("./testdata/WEAVEDMO.zzt")
	if err != nil {
		t.Fatal(err)
	}

	w, err := ReadWorld(b)
	if err != nil {
		t.Fatal(err)
	}

	for _, brd := range w.Boards {
		brd2 := brd.Unpack().Pack()
		BoardHeadersAreEqual(t, brd.Header, brd2.Header)
		BoardPropertiesAreEqual(t, brd.Properties, brd2.Properties, brd.Header)
		StatusElementsAreEqual(t, brd.StatusElements, brd2.StatusElements, brd.Header)
		TilesAreEqual(t, brd.Tiles, brd2.Tiles, brd.Header)
	}
}

func BoardHeadersAreEqual(t *testing.T, a, b BoardHeader) {
	if a != b {
		t.Fail()
		t.Logf("board headers are not equal at %v", a.BoardName)
	}
}

func BoardPropertiesAreEqual(t *testing.T, a, b BoardProperties, h BoardHeader) {
	if a != b {
		t.Fail()
		t.Logf("board properties are not equal at %v", h.BoardName)
	}
}

func StatusElementsAreEqual(t *testing.T, a, b []StatusElement, h BoardHeader) {
	if len(a) != len(b) {
		t.Fail()
		t.Log("status elements are different len()")
		return
	}
	
	for _, sea := range a {
		for _, seb := range b {
			if sea.Properties == seb.Properties {
				if string(sea.Code) != string(seb.Code) {
					t.Fail()
					t.Logf("code not equal at %v, %v, %v", h.BoardName.String(), sea.Properties.LocationX, sea.Properties.LocationY)
				}
				goto next
			}
		}
		t.Fail()
		t.Logf("could not match status element at %v, %v, %v", h.BoardName.String(), sea.Properties.LocationX, sea.Properties.LocationY)
		next:
	}
}

func TilesAreEqual(t *testing.T, a, b []Tile, h BoardHeader) {
	if len(a) != len(b) {
		t.Fail()
		t.Log("tile are different len()")
		return
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fail()
			t.Logf("tiles not matching at %v, %v", i, h.BoardName.String())
		} 
	}
}