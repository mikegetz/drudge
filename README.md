# drudge
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
A CLI that prints the drudge page headlines to stdout 

```
Usage: drudge [options]
Options:
  -v        Print the version
  -s        Print the output as a string
  -h        Print this help menu
  -w [n]    Watch for updates every n seconds (default is 30 seconds)
  ```

![image](https://github.com/user-attachments/assets/6f7185ba-aa63-42cc-b456-0ea4c1055add)
