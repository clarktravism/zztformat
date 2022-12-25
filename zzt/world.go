package zzt

import (
	"io"
	"bytes"
	"encoding/binary"
)

func ReadWorld(b []byte) (World, error) {
	var w World
	r := bytes.NewReader(b)
	err := w.Read(r)
	return w, err
} 


// World
//
type World struct {
	Header WorldHeader
	Boards []Board
}

func (w *World) Read(r io.Reader) error {
	if err := w.Header.Read(r); err != nil {
		return err
	}
	w.Boards = make([]Board, w.Header.NumBoards + 1)
	for i := range w.Boards {
		if err := w.Boards[i].Read(r); err != nil {
			return err
		}
	}
	return nil
}

func (w *World) Write(wr io.Writer) error {
	w.UpdateNumBoards()
	if err := w.Header.Write(wr); err != nil {
		return err
	}

	for i := range w.Boards {
		if err := w.Boards[i].Write(wr); err != nil {
			return err
		}
	}
	return nil
}

func (w *World) UpdateNumBoards() {
	w.Header.NumBoards = int16(len(w.Boards)) - 1
}

//WorldHeader
//
type WorldHeader struct {
	WorldType int16
	NumBoards int16
	PlayerAmmo int16
	PlayerGems int16
	PlayerKeys [7]byte
	PlayerHealth int16
	PlayerBoard int16
	PlayerTorches int16
	TorchCycles int16
	EnergyCycles int16
	Unused25 int16
	PlayerScore int16
	WorldName String20
	Flag0 String20
	Flag1 String20
	Flag2 String20
	Flag3 String20
	Flag4 String20
	Flag5 String20
	Flag6 String20
	Flag7 String20
	Flag8 String20
	Flag9 String20
	TimePassed int16
	TimePassedTicks int16
	Locked byte
	Unused265 [14]byte
	_ [233]byte
}

func (h *WorldHeader) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, h)
}

func (h *WorldHeader) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, h)
}