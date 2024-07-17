# LocalSearch

## Introduction

A lightweight, CLI based application for finding matching strings in speccified directory

## Develop Note

### Bugs to be fixed

- Highlight marks misalignment in multiple cases, e.g. below:
- Incorrect recognition of Chinese character length causes highlight mark misalignment
	- utf8 3 -> 2
	- utf8mb4 4 -> 2
	- ...
- ~~Heuristic algorithm unable to search some sort of text file, e.g. nohup.out~~
- ~~error occurs when walking on symlink~~

### Features to be added

- ~~Help page with command `help` or flag `-h`~~
- ~~Search current directory if directory `-d` is not specified~~
- Use `config.yaml` replace constant define
- ~~Ignore file or directory with flag `-i`~~
- Hide excessive results
	- Initially displays `and ... more matched results`
	- switch to `cat` styled after pressing arrow down key
	- Allow using arrow keys to switch pages, i.e. `cat` styled
- Allows printing of all results at once with flag `-a` (save to file with `>`)
- Allows ignore with patterns like `*/` `**/`
- When contextLines is `-1`, print entire file

### Performance to be improved

- Use multiple goroutine when directory have way too many chilren

### Documentation to be written

- `README.md`: more info regarding intro, usage, etc.
- `DEVLOG.md`: move develop related info to a new file.
- `help`: `-h` help page to be completed.
