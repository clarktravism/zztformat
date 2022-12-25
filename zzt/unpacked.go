package zzt

//UnpackedBoard
//
type UnpackedBoard struct {
	Header     BoardHeader
	Tiles      [1500]UnpackedTile
	Properties BoardProperties
}

func (b Board) Unpack() UnpackedBoard {
	ub := UnpackedBoard{Header: b.Header, Properties: b.Properties}

	i := 0
	for _, t := range b.Tiles {
		ct := int(t.Count)
		if ct == 0 {
			ct = 256
		}
		for j := 0; j < ct; j++ {
			ub.Tiles[i] = UnpackedTile{ Element: t.Element, Color: t.Color}
			i++
		}
	}

	for i := range b.StatusElements {
		se := b.StatusElements[i]
		ub.Tiles[Index(se.Properties.LocationX, se.Properties.LocationY)].StatusElement = &se
	}

	return ub
}

func (ub UnpackedBoard) Pack() Board {
	b := Board{Header: ub.Header, Properties: ub.Properties, StatusElements: make([]StatusElement, 1)}

	//place player at 0 status element
	//pse := ub.Tiles[Index(ub.Properties.PlayerEnterX, ub.Properties.PlayerEnterY)].StatusElement
	//pse.Properties.LocationX, pse.Properties.LocationY = ub.Properties.PlayerEnterX, ub.Properties.PlayerEnterY
	//b.StatusElements = []StatusElement{ *pse }
	
	tile := Tile{ Element: ub.Tiles[0].Element, Color: ub.Tiles[0].Color }
	for i, t := range ub.Tiles {
		if t.StatusElement != nil {
			se := *t.StatusElement
			se.Properties.LocationX, se.Properties.LocationY = XY(i)
			if t.Element == Player {
				b.StatusElements[0] = se
			} else {
				b.StatusElements = append(b.StatusElements, se)
			}
		}

		if t.Color == tile.Color && t.Element == tile.Element {
			if tile.Count == 255 {
				b.Tiles = append(b.Tiles, tile)
				tile.Count = 0
			}
			tile.Count++
		} else {
			b.Tiles = append(b.Tiles, tile)
			tile = Tile{ Element: t.Element, Color: t.Color, Count: 1 }
		}
	}
	b.Tiles = append(b.Tiles, tile)

	return b
}

//UnpackedTile
//
//todo; track negative length 
type UnpackedTile struct {
	Element byte
	Color   byte
	StatusElement *StatusElement
}

func Index(x, y byte) int {
	if x < 1 || y < 1 || x > 60 || y > 25 {
		panic("x, y out of range")
	}
	return (int(y) - 1) * 60 + int(x) - 1
}

func XY(i int) (x, y byte) {
	if i < 0 || i > 1499 {
		panic("i out of range")
	}
	return byte(i % 60 + 1), byte(i / 60 + 1)
}
