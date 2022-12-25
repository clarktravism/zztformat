package zzt

import (
	"bytes"
	"encoding/binary"
	"io"
)

//Board
//
type Board struct {
	Header         BoardHeader
	Tiles          []Tile
	Properties     BoardProperties
	StatusElements []StatusElement
}

func (b *Board) Read(r io.Reader) error {
	if err := b.Header.Read(r); err != nil {
		return err
	}

	bb := make([]byte, b.Header.BoardSize - BoardHeaderSize)
	if _, err := r.Read(bb); err != nil {
		return err
	}
	br := bytes.NewReader(bb)

	n := 0
	for n < 1500 {
		var t Tile
		if err := t.Read(br); err != nil {
			return err
		}
		b.Tiles = append(b.Tiles, t)
		if t.Count == 0 {
			n += 256
		}
		n += int(t.Count)
	}

	if err := b.Properties.Read(br); err != nil {
		return err
	}

	b.StatusElements = make([]StatusElement, b.Properties.StatElementCount+1)
	for i := range b.StatusElements {
		if err := b.StatusElements[i].Read(br); err != nil {
			return err
		}
	}

	return nil
}

func (b *Board) Write(w io.Writer) error {
	b.UpdateBoardSize()
	if err := b.Header.Write(w); err != nil {
		return err
	}

	for _, t := range b.Tiles {
		if err := t.Write(w); err != nil {
			return err
		}
	}

	b.UpdateStatElementCount()
	if err := b.Properties.Write(w); err != nil {
		return err
	}

	for _, se := range b.StatusElements {
		if err := se.Write(w); err != nil {
			return err
		}
	}

	return nil
}

func (b *Board) UpdateBoardSize() {
	b.Header.BoardSize = BoardHeaderSize + BoardPropertiesSize + int16(len(b.Tiles))*TileSize
	for i := range b.StatusElements {
		b.Header.BoardSize += b.StatusElements[i].UpdateLength()
	}
}

func (b *Board) UpdateStatElementCount() {
	b.Properties.StatElementCount = int16(len(b.StatusElements)) - 1
}

//BoardHeader
//
const BoardHeaderSize = 51 //exclude BoardSize
type BoardHeader struct {
	BoardSize int16
	BoardName String50
}

func (h *BoardHeader) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, h)
}

func (h *BoardHeader) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, h)
}

//BoardProperties
//
const BoardPropertiesSize = 88

type BoardProperties struct {
	MaxPlayerShots   byte
	IsDark           byte
	ExitNorth        byte
	ExitSouth        byte
	ExitWest         byte
	ExitEast         byte
	RestartOnZap     byte
	Message          String58
	PlayerEnterX     byte
	PlayerEnterY     byte
	TimeLimit        int16
	Unused           [16]byte
	StatElementCount int16
}

func (p *BoardProperties) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, p)
}

func (p *BoardProperties) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, p)
}
