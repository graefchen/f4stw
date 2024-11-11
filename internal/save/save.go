package save

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"os"
	"strings"
	"syscall"
)

type Statistic struct {
	name string
	t    uint8
	n    uint32
}

type F4Save struct {
	// internel fields
	fileName string

	// The important info stuff (file header)
	engineVersion              uint32
	saveNumber                 uint32
	characterName              string
	characterLevel             uint32
	characterLocation          string
	playtime                   string
	characterRace              string
	characterSex               uint16
	currentCharacterExperience float32
	requiredExperience         float32
	fileTime                   syscall.Filetime
	snapshot                   *image.RGBA
	formatVersion              uint8
	gameVersion                string
	plugins                    []string
	plugins2                   []string

	// Following Stats(?)
	statistic []Statistic
}

// Checks if the given file of filename is a Fallout 4 savefile
func IsF4Save(filename string) (bool, error) {
	bytecode, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}

	// cheking if the magicid is the correct magicid of "FO4_SAVEGAME"
	if strings.Compare(string(bytecode[0:11]), "FO4_SAVEGAME") != 0 {
		return true, nil
	}

	return false, nil
}

func ReadF4Save(filename string) (F4Save, error) {
	save := F4Save{fileName: filename}
	bytecode, err := os.ReadFile(filename)
	if err != nil {
		return F4Save{}, err
	}
	reader := bytes.NewReader(bytecode)
	mid := make([]byte, 12)
	var u32 uint32
	var sz uint16
	binary.Read(reader, binary.BigEndian, &mid)
	binary.Read(reader, binary.LittleEndian, &u32)
	binary.Read(reader, binary.LittleEndian, &save.engineVersion)
	binary.Read(reader, binary.LittleEndian, &save.saveNumber)

	binary.Read(reader, binary.LittleEndian, &sz)
	charName := make([]byte, sz)
	binary.Read(reader, binary.LittleEndian, &charName)
	save.characterName = string(charName)

	binary.Read(reader, binary.LittleEndian, &save.characterLevel)

	binary.Read(reader, binary.LittleEndian, &sz)
	charLocation := make([]byte, sz)
	binary.Read(reader, binary.LittleEndian, &charLocation)
	save.characterLocation = string(charLocation)

	binary.Read(reader, binary.LittleEndian, &sz)
	charPlaytime := make([]byte, sz)
	binary.Read(reader, binary.LittleEndian, &charPlaytime)
	save.playtime = string(charPlaytime)

	binary.Read(reader, binary.LittleEndian, &sz)
	charRace := make([]byte, sz)
	binary.Read(reader, binary.LittleEndian, &charRace)
	save.characterRace = string(charRace)

	binary.Read(reader, binary.LittleEndian, &save.characterSex)
	binary.Read(reader, binary.LittleEndian, &save.currentCharacterExperience)
	binary.Read(reader, binary.LittleEndian, &save.requiredExperience)

	var ldt, hdt uint32
	binary.Read(reader, binary.LittleEndian, &ldt)
	binary.Read(reader, binary.LittleEndian, &hdt)
	save.fileTime.LowDateTime = ldt
	save.fileTime.HighDateTime = hdt

	var snapshotWidth, snapshotHeight uint32
	binary.Read(reader, binary.LittleEndian, &snapshotWidth)
	binary.Read(reader, binary.LittleEndian, &snapshotHeight)
	snapshot := make([]uint8, snapshotWidth*snapshotHeight*4)
	binary.Read(reader, binary.LittleEndian, &snapshot)
	rect := image.Rect(0, 0, int(snapshotWidth), int(snapshotHeight))
	save.snapshot = &image.RGBA{Pix: snapshot, Rect: rect, Stride: int(snapshotWidth) * 4}

	binary.Read(reader, binary.LittleEndian, &save.formatVersion)

	binary.Read(reader, binary.LittleEndian, &sz)
	gameVersion := make([]byte, sz)
	binary.Read(reader, binary.LittleEndian, &gameVersion)
	save.gameVersion = string(gameVersion)

	binary.Read(reader, binary.LittleEndian, &u32)

	var pluginCount uint8
	binary.Read(reader, binary.LittleEndian, &pluginCount)
	save.plugins = make([]string, pluginCount)
	for i := 0; i < int(pluginCount); i++ {
		binary.Read(reader, binary.LittleEndian, &sz)
		plugin := make([]byte, sz)
		binary.Read(reader, binary.LittleEndian, &plugin)
		save.plugins[i] = string(plugin)
	}

	binary.Read(reader, binary.LittleEndian, &pluginCount)
	save.plugins2 = make([]string, pluginCount)
	for i := 0; i < int(pluginCount); i++ {
		binary.Read(reader, binary.LittleEndian, &sz)
		plugin := make([]byte, sz)
		binary.Read(reader, binary.LittleEndian, &plugin)
		save.plugins2[i] = string(plugin)
	}

	// TODO: Reverse engineer the whole rest of the file (possible skyrim?)
	// fmt.Println(reader)

	// Unkown ... some data garbage? (last 8 byte might be interesting)
	// after that follwoing are the stats
	// unknown := make([]byte, 113)
	// unknown := make([]byte, 109)
	// possible 10 uint32 important data type, probably no offsets
	// unknown := make([]byte, 105)
	// fmt.Println(reader)
	for i := 0; i < 10; i++ {
		binary.Read(reader, binary.LittleEndian, &u32)
		// fmt.Println(i, ":", u32)
	}
	unknown := make([]byte, 105-40)
	binary.Read(reader, binary.LittleEndian, &unknown)
	// the  size of the statistic block
	binary.Read(reader, binary.LittleEndian, &u32)
	// fmt.Println(u32)

	var dataSize uint32
	binary.Read(reader, binary.LittleEndian, &dataSize)
	save.statistic = make([]Statistic, dataSize)

	for i := 0; i < int(dataSize); i++ {
		var statType byte
		var number uint32
		binary.Read(reader, binary.LittleEndian, &sz)
		name := make([]byte, sz)
		binary.Read(reader, binary.LittleEndian, &name)
		// Possible way to describe data?
		binary.Read(reader, binary.LittleEndian, &statType)
		binary.Read(reader, binary.LittleEndian, &number)
		save.statistic[i] = Statistic{string(name), statType, number}
	}
	fmt.Println(reader)

	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println(u32)
	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println(u32)
	number := u32
	for i := 0; i < int(number); i++ {
		binary.Read(reader, binary.LittleEndian, &u32)
	}

	return save, nil
}

func (s F4Save) GetFileName() string {
	return s.fileName
}

func (s F4Save) GetEngineVersion() uint32 {
	return s.engineVersion
}

func (s F4Save) GetSaveNumber() uint32 {
	return s.saveNumber
}

func (s F4Save) GetCharacterName() string {
	return s.characterName
}

func (s F4Save) GetCharacterLevel() uint32 {
	return s.characterLevel
}

func (s F4Save) GetCharacterLocation() string {
	return s.characterLocation
}

func (s F4Save) GetPlaytime() string {
	return s.playtime
}

func (s F4Save) GetCharacterRace() string {
	return s.characterRace
}

func (s F4Save) GetCharacterSex() uint16 {
	return s.characterSex
}

func (s F4Save) GetCurrentCharacterExperience() float32 {
	return s.currentCharacterExperience
}

func (s F4Save) GetRequiredExperience() float32 {
	return s.requiredExperience
}

func (s F4Save) GetFileTime() syscall.Filetime {
	return s.fileTime
}

func (s F4Save) GetSnapshot() *image.RGBA {
	return s.snapshot
}

func (s F4Save) GetFormatVersion() uint8 {
	return s.formatVersion
}

func (s F4Save) GetGameVersion() string {
	return s.gameVersion
}

func (s F4Save) GetPlugins() []string {
	return s.plugins
}

func (s F4Save) GetPlugins2() []string {
	return s.plugins2
}

func (s F4Save) GetStatisatics() []Statistic {
	return s.statistic
}
