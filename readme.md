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

<!--
### With F4Save.fos:

hexyl F4Save.fos -s 983248 -n 40

0 : 1971351040
1 : 1971676672
2 : 251737344
3 : 262167040
4 : 266353408
5 : 1259124480
6 : 3072
7 : 3584
8 : 2048
9 : 6871296
-->

### Decode in bytecode

uses F4Save_long.fos (after all the stats)

```nushell
hexyl F4Save_long.fos --border none --color never | save -f bytecode.txt
hexyl F4Save_long.fos -s 985567 --border none --color never | save -f bytecode.txt
```

Starting at line 61583....

985567 is the last index that I know will bring us further
