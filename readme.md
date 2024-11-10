# f4stw

fallout 4 save to website

generates websites with data from the fallout 4 savegames it finds

## goals

- only use stdlib
- fast
- small
- modular style [ref](https://en.wikipedia.org/wiki/Modular_programming), [ref](https://best-practice-and-impact.github.io/qa-of-code-guidance/modular_code.html)

## infos

- read fallout4 savefile
  - read bytecode (my max 17.5 Mib) / together 8.2 GiB
  - turn into one webpage per file
  - recursively read

## save

ref: [1](https://gist.github.com/SirTony/5832ad8a2b8fd4acb636), [2](docs/save.md),
[3](https://fallout.wiki/wiki/FOS_file_format), [4](https://en.uesp.net/wiki/Skyrim_Mod:Save_File_Format)

### help

ref: [1](https://lucasklassmann.com/blog/2018-07-21-handling-binary-files-in-go/)

hexyl F4Save.fos -s 0xF00F0 -n 130
