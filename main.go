package main

import (
	"fmt"
	"main/internal/save"
)

func main() {
	// save.ReadF4Save("F4Save_long.fos")
	save, err := save.ReadF4Save("F4Save_long.fos")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Filename: ", save.GetFileName())
	// fmt.Println("Engine Version: ", save.GetEngineVersion())
	// fmt.Println("Save Number: ", save.GetSaveNumber())
	// fmt.Println("Character Name: ", save.GetCharacterName())
	// fmt.Println("Character Level: ", save.GetCharacterLevel())
	// fmt.Println("Character Location: ", save.GetCharacterLocation())
	// fmt.Println("Playtime: ", save.GetPlaytime())
	// fmt.Println("Character Race: ", save.GetCharacterRace())
	// fmt.Println("Character Sex: ", save.GetCharacterSex())
	// fmt.Println("Character Experience: ", save.GetCurrentCharacterExperience())
	// fmt.Println("Required Experience: ", save.GetRequiredExperience())
	// fmt.Println("File Time: ", save.GetFileTime())
	// fmt.Println("Here comes the snapshot...")

	// f, err := os.Create("snapshot.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// png.Encode(f, save.GetSnapshot())

	// fmt.Println("Format Version: ", save.GetFormatVersion())
	fmt.Println("Game Version: ", save.GetGameVersion())
	// fmt.Println(save.GetPlugins())
	// fmt.Println(save.GetPlugins2())
	// fmt.Println(save.GetStatisatics())
}
