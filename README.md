# drudge
A CLI that prints the drudge page headlines to stdout 
## Install
### Homebrew
```
brew tap mikegetz/drudge
brew install drudge
```
### Shell
#### curl
```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/mikegetz/drudge/main/tools/install.sh)"
```
#### wget
```
sh -c "$(wget -qO- https://raw.githubusercontent.com/mikegetz/drudge/main/tools/install.sh)"
```
## Usage

```
Usage: drudge [options]
Options:
  -v        Print the version
  -s        Print the output as text without ANSI Escaped links
  -h        Print this help menu
  -w [n]    Watch for updates every n seconds (default is 30 seconds)
  -r        Parse RSS feed (default)
  -d        Parse DOM
```

![image](https://github.com/user-attachments/assets/6f7185ba-aa63-42cc-b456-0ea4c1055add)
