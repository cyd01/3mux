package main

import "fmt"

func showHelp() {
	fmt.Println(helpText)
}

const helpText = `3mux
The terminal multiplexer inspired by i3

USAGE:
    3mux                  Interactive 3mux interface
    3mux ls               List session names (has alias '3mux ps')
    3mux attach <name>    Attach to a session
    3mux detach           Detach from the current session
    3mux new <name>       Create a new session
    3mux kill <name>      Kill a session

SHORTCUTS:
	Alt+N/Alt+Enter   Create new pane
	Alt+Shift+Q       Close pane
	Alt+Shift+F       Make pane fullscreen
	Alt+Shift+Arrow   Move pane
	Alt+Arrow         Move selection
	Alt+/             Toggle search

Key Bindings
	Alt+Enter / Alt+N  Create a new pane
	Alt+Shift+F        Make the selected pane fullscreen. Useful for copying text
	Alt+←/↓/↑/→        Select an adjacent pane
	Alt+h/j/k/l        Select an adjacent pane
	Alt+Shift+←/↓/↑/→  Move the selected pane
	Alt+Shift+h/j/k/l  Move the selected pane
	Alt+R              Enter resize mode. Resize selected pane with arrow keys or h/j/k/l. Exit using any other key(s)
	Alt+/              Enter search mode. Type query, navigate between results with arrow keys or n/N
	Scroll 	           Move through scrollback
	Shift              Many terminal emulators support selecting text while pressing this key

Supported tmux Bindings
	Ctrl+b "           Split horizontally
	Ctrl+b %           Split vertically
	Ctrl+b {           Move pane left
	Ctrl+b }           Move pane right

Supported screen Bindings
	Ctrl+a |           Split horizontally
	Ctrl+a S           Split vertically
	Ctrl+a Tab         Cycle forward through panes
`
