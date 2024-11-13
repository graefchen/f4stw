<h1 align="center">Fallout 4 Save File Format</h1>

initial file taken from [here](https://gist.github.com/SirTony/5832ad8a2b8fd4acb636)
and added with information from [here](https://en.uesp.net/wiki/Skyrim_Mod:Save_File_Format)

The binary format for Fallout 4 PC save files.
This document was created by reverse-engineering files from version 1.10.984.0 of the game.

_**Note**: This document is incomplete!_

## Table of Contents

- _[Table of Contents](#table-of-contents)_
- _[Types](#types)_
- _[Format](#format)_
  - _[Header](#header)_
  - _[File Location Table](#file-location-table)_
  - _[Global Data](#global-data)_
    - _[Misc Stats](#misc-stats)_

## Types

| Type Name  | Size (in bytes) | Remarks                                                                                                        |
| ---------- | :-------------- | -------------------------------------------------------------------------------------------------------------- |
| `char`     | 1               | An 8-bit character                                                                                             |
| `wstring`  | Variable        | A `wstring` is a string prefixed with a `uint16` denoting the length, followed by exactly that many characters |
| `uint8`    | 1               | An unsigned 8-bit integer                                                                                      |
| `uint16`   | 2               | An unsigned 16-bit integer                                                                                     |
| `uint32`   | 4               | An unsigned 32-bit integer                                                                                     |
| `float32`  | 4               | A single-precision, 32-bit, floating-point number                                                              |
| `FILETIME` | 8               | _**See**: https://msdn.microsoft.com/en-us/library/windows/desktop/ms724284(v=vs.85).aspx_                     |

## Format

| Field Name          | Type                                                    | Remarks                                                                                              |
| ------------------- | ------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| Magic ID            | `char[12]`                                              | Always `FO4_SAVEGAME`                                                                                |
| Header Size         | `uint32`                                                | The total size (in bytes) of the header                                                              |
| Header              | `Header`                                                | See: [Header](#header)                                                                               |
| Snapshot            | `uint8[Width * Height * 4]`                             | An array containing raw pixel data for the thumbnail. The array is stored as 32-bits-per-pixel ARGB. |
| Format Version      | `uint8`                                                 | The save file format version (?). Current value is 61                                                |
| Game Version        | `wstring`                                               | The game's patch version when the save was created in dot-notation (ex `1.2.37.0`)                   |
| Plugin Info Size    | `uint32`                                                | The total size (in bytes) of the plugin information                                                  |
| Plugins Count       | `uint8`                                                 | The number of plugins used by this save                                                              |
| Plugins             | `wstring[Plugins Count]`                                | Each string is a file name for a `.esm` or `.esp` file in the `Data` directory.                      |
| Light Plugins Count | `uint16`                                                | The number of light plugins used by this save                                                        |
| Light Plugins       | `wstring[Plugins Count]`                                | Each string is a file name for a `.esm` or `.esp` file in the `Data` directory.                      |
| File Location Table | `File Location Table`                                   |                                                                                                      |
| Global Date Table 1 | `Global Data [fileLocationTable.globalDataTable1Count]` |                                                                                                      |
| Global Date Table 2 | `Global Data [fileLocationTable.globalDataTable2Count]` |                                                                                                      |
| ???                 | `???`                                                   |                                                                                                      |
| Global Date Table 3 | `Global Data [fileLocationTable.globalDataTable2Count]` |                                                                                                      |

### Header

| Field Name                   | Type       | Remarks                                                                                                                                                                                                                                                |
| ---------------------------- | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Engine Version               | `uint32`   | The version of Creation Engine that created this file (?). Current value is 11                                                                                                                                                                         |
| Save Number                  | `uint32`   | Incremented by 1 each time a game is saved                                                                                                                                                                                                             |
| Character Name               | `wstring`  |                                                                                                                                                                                                                                                        |
| Character Level              | `uint32`   |                                                                                                                                                                                                                                                        |
| Character Location           | `wstring`  | Name of the player's current location                                                                                                                                                                                                                  |
| Play Time                    | `wstring`  | The amount of time played. Stored as `xd.yh.zm.x days.y hours.z minutes` where `x`, `y`, and `z` are any arbitrary integers. An example of for a character that has played 2 days, 2 hours, and 3 minutes would be `2d.2h.3m.2 days.2 hours.3 minutes` |
| Character Race               | `wstring`  | The internal editor ID of the player's race. Probably always `HumanRace`                                                                                                                                                                               |
| Character Sex                | `uint16`   | The sex (gender) of the player character. `0` = male, `1` = female                                                                                                                                                                                     |
| Current Character Experience | `float32`  | The current amount of experience the player has attained for progressing to the next level                                                                                                                                                             |
| Required Experience          | `float32`  | The amount of experience needed to progress to the next level                                                                                                                                                                                          |
| Filetime                     | `FILETIME` | The real-world time the save file was created                                                                                                                                                                                                          |
| Snapshot Width               | `uint32`   | The width (in pixels) of the save thumbnail                                                                                                                                                                                                            |
| Snapshot Height              | `uint32`   | The height (in pixels) of the save thumbnail                                                                                                                                                                                                           |

### File Location Table

| Field Name                 | Type         | Remarks |
| -------------------------- | ------------ | ------- |
| Form ID Array Count Offset | `uint32`     |         |
| Unknown Table 3 Offset     | `uint32`     |         |
| Global Data Table 1 Offset | `uint32`     |         |
| Global Data Table 2 Offset | `uint32`     |         |
| Change Form Offset         | `uint32`     |         |
| Global Data Table 3 Offset | `uint32`     |         |
| Global Date Table 1 Count  | `uint32`     |         |
| Global Date Table 2 Count  | `uint32`     |         |
| Global Date Table 3 Count  | `uint32`     |         |
| Change Form Count          | `uint32`     |         |
| Unused                     | `uint32[15]` |         |

### Global Data

| Field Name | Type            | Remarks                                      |
| ---------- | --------------- | -------------------------------------------- |
| type       | `uint32`        |                                              |
| lenght     | `uint32`        |                                              |
| data       | `uint8[lenght]` | Format of Data depends on type (noted below) |

| Number | Type                      |
| ------ | ------------------------- |
| 0      | [Misc Stats](#misc-stats) |
| 1      |                           |
| 2      |                           |
| 3      |                           |
| 4      |                           |
| 5      |                           |
| 6      |                           |
| 7      |                           |
| 8      |                           |
| 9      |                           |
| 10     |                           |
| 11     |                           |
| 100    |                           |
| 101    |                           |
| 102    |                           |
| 103    |                           |
| 105    |                           |
| 106    |                           |
| 109    |                           |
| 110    |                           |
| 111    |                           |
| 113    |                           |
| 114    |                           |
| 115    |                           |
| 116    |                           |
| 117    |                           |
| 1000   |                           |
| 1001   |                           |
| 1002   |                           |
| 1003   |                           |
| 1004   |                           |
| 1005   |                           |
| 1006   |                           |
| 1007   |                           |

#### Misc-Stats

| Field Name   | Type      | Remarks |
| ------------ | --------- | ------- |
| Name         | `wstring` |         |
| type/buffer? | `uint8`   |         |
| Number       | `uint32`  |         |
