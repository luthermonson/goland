# goland cli
Simple command line interface to open goland from a terminal window for a specific directory. Currently works on Windows and MacOSX.

## Project Opening
Tool works best when setting the Project Opening setting, default is `Confirm window to open project in`, if you 
set to `Open project in new window` behavior will be similar to VSCode's `code` cli. If you choose `Open project in the 
same window` it will switch your current open window, or start one if it's not running.

## Usage
```
goland ~/go/my-project
goland .
goland /Users/user/go/my-project
``` 
