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

	// Following Stuff
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
	// var i32 int32
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

	var lightPluginCount uint16
	binary.Read(reader, binary.LittleEndian, &lightPluginCount)
	save.plugins2 = make([]string, lightPluginCount)
	for i := 0; i < int(lightPluginCount); i++ {
		binary.Read(reader, binary.LittleEndian, &sz)
		plugin := make([]byte, sz)
		binary.Read(reader, binary.LittleEndian, &plugin)
		save.plugins2[i] = string(plugin)
	}

	// fmt.Println(reader)

	// File location Table
	var globalDataTable3Offset, globalDataTable1Count, globalDataTable2Count, globalDataTable3Count, changeFormCount uint32
	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println("formIDArrayCountOffset:", u32)
	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println("unknownTable3Offset:", u32)
	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println("globalDataTable1Offset:", u32)
	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println("globalDataTable2Offset:", u32)
	binary.Read(reader, binary.LittleEndian, &u32)
	fmt.Println("changeFormsOffset:", u32)
	binary.Read(reader, binary.LittleEndian, &globalDataTable3Offset)
	fmt.Println("globalDataTable3Offset:", globalDataTable3Offset)
	binary.Read(reader, binary.LittleEndian, &globalDataTable1Count)
	fmt.Println("globalDataTable1Count:", globalDataTable1Count)
	binary.Read(reader, binary.LittleEndian, &globalDataTable2Count)
	fmt.Println("globalDataTable2Count:", globalDataTable2Count)
	binary.Read(reader, binary.LittleEndian, &globalDataTable3Count)
	fmt.Println("globalDataTable3Count:", globalDataTable3Count)
	binary.Read(reader, binary.LittleEndian, &changeFormCount)
	fmt.Println("changeFormCount:", changeFormCount)

	unused := make([]uint32, 15)
	binary.Read(reader, binary.LittleEndian, &unused)

	// Misc Stats
	for i := 0; i < int(globalDataTable1Count); i++ {
		fmt.Printf("=== Global Data Table 1[%d] ===\n", i)
		binary.Read(reader, binary.LittleEndian, &u32)
		fmt.Println("type:", u32)
		var length uint32
		binary.Read(reader, binary.LittleEndian, &length)
		fmt.Println("length:", length)
		u := make([]uint8, length)
		binary.Read(reader, binary.LittleEndian, u)
	}

	for i := 0; i < int(globalDataTable2Count); i++ {
		fmt.Printf("=== Global Data Table 2[%d] ===\n", i)
		binary.Read(reader, binary.LittleEndian, &u32)
		fmt.Println("type:", u32)
		var length uint32
		binary.Read(reader, binary.LittleEndian, &length)
		fmt.Println("length:", length)
		u := make([]uint8, length)
		binary.Read(reader, binary.LittleEndian, u)
	}

	// Change Form ...

	// fmt.Println("globalDataTable3Offset:", globalDataTable3Offset)
	reader = bytes.NewReader(bytecode[globalDataTable3Offset:])
	for i := 0; i < int(globalDataTable3Count); i++ {
		fmt.Printf("=== Global Data Table 3[%d] ===\n", i)
		binary.Read(reader, binary.LittleEndian, &u32)
		fmt.Println("type:", u32)
		var length uint32
		binary.Read(reader, binary.LittleEndian, &length)
		fmt.Println("length:", length)
		u := make([]uint8, length)
		binary.Read(reader, binary.LittleEndian, u)
	}

	// Old...
	// var count uint32
	// binary.Read(reader, binary.LittleEndian, &count)
	// // fmt.Println("count:", count)
	// save.plugins2 = make([]string, count)
	// for i := 0; i < int(count); i++ {
	// 	var category uint8
	// 	binary.Read(reader, binary.LittleEndian, &sz)
	// 	str := make([]byte, sz)
	// 	binary.Read(reader, binary.LittleEndian, &str)
	// 	// fmt.Println("name:", string(str))
	// 	binary.Read(reader, binary.LittleEndian, &category)
	// 	// fmt.Println("categoty:", category)
	// 	binary.Read(reader, binary.LittleEndian, &u32)
	// 	// fmt.Println("value:", u32)
	// }

	// // Player Location

	// fmt.Println("=== Global Data Table 1[1] ===")
	// binary.Read(reader, binary.LittleEndian, &u32)
	// fmt.Println("type:", u32)
	// binary.Read(reader, binary.LittleEndian, &u32)
	// fmt.Println("length:", u32)

	// binary.Read(reader, binary.LittleEndian, &u32)
	// // fmt.Println("nextObjectid:", u32)
	// var u8 uint8
	// // fmt.Print("worldSpace1: ")
	// binary.Read(reader, binary.LittleEndian, &u8)
	// // fmt.Print(u8)
	// binary.Read(reader, binary.LittleEndian, &u8)
	// // fmt.Print(u8)
	// binary.Read(reader, binary.LittleEndian, &u8)
	// // fmt.Print(u8, "\n")
	// binary.Read(reader, binary.LittleEndian, &i32)
	// // fmt.Println("coorX:", i32)
	// binary.Read(reader, binary.LittleEndian, &i32)
	// // fmt.Println("coorY:", i32)
	// // fmt.Print("worldSpace2: ")
	// binary.Read(reader, binary.LittleEndian, &u8)
	// // fmt.Print(u8)
	// binary.Read(reader, binary.LittleEndian, &u8)
	// // fmt.Print(u8)
	// binary.Read(reader, binary.LittleEndian, &u8)
	// // fmt.Print(u8, "\n")
	// var f32 float32
	// binary.Read(reader, binary.LittleEndian, &f32)
	// // fmt.Println("posX:", f32)
	// binary.Read(reader, binary.LittleEndian, &f32)
	// // fmt.Println("posY:", f32)
	// binary.Read(reader, binary.LittleEndian, &f32)
	// // fmt.Println("posZ:", f32)

	// // TES(?)
	// fmt.Println("=== Global Data Table 1[2] ===")
	// binary.Read(reader, binary.LittleEndian, &u32)
	// fmt.Println("type:", u32)
	// binary.Read(reader, binary.LittleEndian, &u32)
	// fmt.Println("length:", u32)

	// binary.Read(reader, binary.LittleEndian, &u8)
	// vsval := u8
	// fmt.Println("vsval?:", vsval&3)
	// fmt.Println("vsval is uint8?:", vsval <= 0x40)
	// binary.Read(reader, binary.LittleEndian, &u8)
	// u16vsval := (uint16(vsval) | (uint16(u8) << 8)) >> 2
	// fmt.Println("vsval is uint16?:", u16vsval <= 0x4000)

	// fmt.Println(reader)

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
