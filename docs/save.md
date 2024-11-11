<h1 align="center">Fallout 4 Save File Format</h1>

initial file taken from [here](https://gist.github.com/SirTony/5832ad8a2b8fd4acb636)

The binary format for Fallout 4 PC save files.
This document was created by reverse-engineering files from version 1.10.984.0 of the game.

_**Note**: This document is incomplete!_

## Table of Contents

- _[Table of Contents](#table-of-contents)_
- _[Types](#types)_
- _[Format](#format)_
  - _[Header](#header)_
  - _[Statistics](#statistics)_

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

| Field Name          | Type                           | Remarks                                                                                              |
| ------------------- | ------------------------------ | ---------------------------------------------------------------------------------------------------- |
| Magic ID            | `char[12]`                     | Always `FO4_SAVEGAME`                                                                                |
| Header Size         | `uint32`                       | The total size (in bytes) of the header                                                              |
| Header              | `header`                       | See: [Header](#header)                                                                               |
| Snapshot            | `uint8[Width * Height * 4]`    | An array containing raw pixel data for the thumbnail. The array is stored as 32-bits-per-pixel ARGB. |
| Format Version      | `uint8`                        | The save file format version (?). Current value is 61                                                |
| Game Version        | `wstring`                      | The game's patch version when the save was created in dot-notation (ex `1.2.37.0`)                   |
| Plugin Info Size    | `uint32`                       | The total size (in bytes) of the plugin information                                                  |
| Plugins Count       | `uint8`                        | The number of plugins used by this save                                                              |
| Plugins             | `wstring[Plugins Count]`       | Each string is a file name for a `.esm` or `.esp` file in the `Data` directory.                      |
| Light Plugins Count | `uint8`                        | The number of light plugins used by this save                                                        |
| Light Plugins       | `wstring[Plugins Count]`       | Each string is a file name for a `.esm` or `.esp` file in the `Data` directory.                      |
| unknown             | `uint8[105]`                   | Note: I have absolutely no Idea what this is.                                                        |
| Statistic Size      | `uint32`                       |                                                                                                      |
| Statistic Count     | `uint32`                       |                                                                                                      |
| Statistic           | `statistics[Statistics Count]` |                                                                                                      |
| unknown             |                                | Note: Further Research required!                                                                     |

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

### Statistics

| Field Name   | Type      | Remarks |
| ------------ | --------- | ------- |
| Name         | `wstring` |         |
| type/buffer? | `uint8`   |         |
| Number       | `uint32`  |         |
