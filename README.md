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
func printHelp() {
	fmt.Println("Usage: drudge [options]")
	fmt.Println("Options:")
	fmt.Println("  -v        Print the version")
	fmt.Println("  -s        Print the output as a without ANSI Escaped links")
	fmt.Println("  -h        Print this help menu")
	fmt.Println("  -w [n]    Watch for updates every n seconds (default is 30 seconds)")
	fmt.Println("  -r        Parse RSS feed (default)")
	fmt.Println("  -d        Parse DOM")
}
```

![image](https://github.com/user-attachments/assets/6f7185ba-aa63-42cc-b456-0ea4c1055add)
