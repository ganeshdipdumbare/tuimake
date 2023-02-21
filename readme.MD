# gmake
gmake is a simple command line tool to use Makefile in interactive way

![](https://github.com/ganeshdipdumbare/gmake/blob/main/demo.gif)

## Description
- The gmake is a command line tool to use with Makefile. It gets the targets and their description and shows it on TUI.
  - First it gets all the targets and their description at the start from Makefile
  - The target and their description must be in the following format in the Makefile.
  ```sh
  target: ## description of the target
    command
  ``` 
- For showing it on beautiful TUI it uses [bubbletea](https://github.com/charmbracelet/bubbletea).
- You can interact and select the target to execute on TUI.

## Installation
- Requirements:-
    - Go 1.20
- To install the tool, just run the following command
    ```bash
    go install github.com/ganeshdipdumbare/gmake@v0.1.0
    ```
## Improvements
- Add unit tests 
