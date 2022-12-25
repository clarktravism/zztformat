package zzt

import (
	"encoding/binary"
	"io"
)

//Tile
//
const TileSize = 3
type Tile struct {
	Count   byte
	Element byte
	Color   byte
}

func (t *Tile) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, t)
}

func (t *Tile) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, t)
}

//StatusElement
//
type StatusElement struct {
	Properties StatusElementProperties
	Code       []byte
}

func (e *StatusElement) Read(r io.Reader) error {
	if err := e.Properties.Read(r); err != nil {
		return err
	}
	if e.Properties.Length > 0 {
		e.Code = make([]byte, e.Properties.Length)
		if _, err := r.Read(e.Code); err != nil {
			return err
		}
	}
	return nil
}

func (e *StatusElement) Write(w io.Writer) error {
	if err := e.Properties.Write(w); err != nil {
		return err
	}
	if _, err := w.Write(e.Code); err != nil {
		return err
	}
	return nil
}

//UpdateLength updates the length of code and returns the size of the StatusElement
func (se *StatusElement) UpdateLength() int16 {
	if len(se.Code) > 0 {
		se.Properties.Length = int16(len(se.Code))
	} else if se.Properties.Length > 0 {
		se.Properties.Length = 0
	}
	return StatusElementPropertiesSize + int16(len(se.Code))
}

//StatusElementProperties
//
const StatusElementPropertiesSize = 33
type StatusElementProperties struct {
	LocationX          byte
	LocationY          byte
	StepX              int16
	StepY              int16
	Cycle              int16
	P1                 byte
	P2                 byte
	P3                 byte
	Follower           int16
	Leader             int16
	UnderID            byte
	UnderColor         byte
	Pointer            int32
	CurrentInstruction int16
	Length             int16
	_                  [8]byte
}

func (p *StatusElementProperties) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, p)
}

func (p *StatusElementProperties) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, p)
}

// Elements
//
const (
	Empty       byte = 0x00
	BoardEdge   byte = 0x01
	Messenger   byte = 0x02
	Monitor     byte = 0x03
	Player      byte = 0x04
	Ammo        byte = 0x05
	Torch       byte = 0x06
	Gem         byte = 0x07
	Key         byte = 0x08
	Door        byte = 0x09
	Scroll      byte = 0x0A
	Passage     byte = 0x0B
	Duplicator  byte = 0x0C
	Bomb        byte = 0x0D
	Energizer   byte = 0x0E
	Star        byte = 0x0F
	Clockwise   byte = 0x10
	Counter     byte = 0x11
	Bullet      byte = 0x12
	Water       byte = 0x13
	Forest      byte = 0x14
	Solid       byte = 0x15
	Normal      byte = 0x16
	Breakable   byte = 0x17
	Boulder     byte = 0x18
	SliderNS    byte = 0x19
	SliderEW    byte = 0x1A
	Fake        byte = 0x1B
	Invisible   byte = 0x1C
	BlinkWall   byte = 0x1D
	Transporter byte = 0x1E
	Line        byte = 0x1F
	Ricochet    byte = 0x20
	BlinkRayH   byte = 0x21
	Bear        byte = 0x22
	Ruffian     byte = 0x23
	Object      byte = 0x24
	Slime       byte = 0x25
	Shark       byte = 0x26
	SpinningGun byte = 0x27
	Pusher      byte = 0x28
	Lion        byte = 0x29
	Tiger       byte = 0x2A
	BlickRayV   byte = 0x2B
	Head        byte = 0x2C
	Segment     byte = 0x2D
	Element46   byte = 0x2E
	TextBlue    byte = 0x2F
	TextGreen   byte = 0x30
	TextCyan    byte = 0x31
	TextRed     byte = 0x32
	TextPurple  byte = 0x33
	TextBrown   byte = 0x34
	TextBlack   byte = 0x35
)

const (
	Black      byte = 0x00
	DarkBlue   byte = 0x01
	DarkGreen  byte = 0x02
	DarkCyan   byte = 0x03
	DarkRed    byte = 0x04
	DarkPurple byte = 0x05
	Brown      byte = 0x06
	Gray       byte = 0x07
	DarkGray   byte = 0x08
	Blue       byte = 0x09
	Green      byte = 0x0A
	Cyan       byte = 0x0B
	Red        byte = 0x0C
	Purple     byte = 0x0D
	Yellow     byte = 0x0E
	White      byte = 0x0F
)

func Color(fg, bg byte) byte {
	return (bg << 4) + fg
}

func GetColors(c byte) (byte, byte) {
	return c & 0x0F, (c & 0xF0) >> 4
} 
