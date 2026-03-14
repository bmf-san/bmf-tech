---
title: Road to Vimmer
description: 'Master vim navigation, editing, and text manipulation commands for efficient terminal-based text editing without GUI shortcuts.'
slug: road-to-vimmer
date: 2018-04-22T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - vim
  - Editor
translation_key: road-to-vimmer
---

# Overview
A list of commands to remember on the journey to becoming a Vimmer.

# Motivation
- Although I am accustomed to GUI editor shortcuts, I feel my text editing speed has plateaued.
- Moving away from GUI editor dependency (considering whether learning shortcut keys for GUI editors or the well-established VIM is more cost-effective. vi is included as the standard editor on all OS).
- Using shortcuts and arrow keys causes hands to leave the home position, which decreases typing speed.
- A growing desire to complete text editing quickly on the terminal.

# Navigation
h Move left on the logical line
j Move down on the logical line
k Move up on the logical line
l Move right on the logical line

^ Move to the start
0 Move to the start ignoring indentation
$ Move to the end
\+ Move to the start of the next line
\- Move to the start of the previous line

:3 Move to line 3

w Move to the start of the next word (including spaces)
b Move to the start of the previous word (including spaces)

e Move to the end of the next word (excluding spaces)
ge Move to the end of the previous word (excluding spaces)

% Jump to the matching closing tag for (, [, {, etc.

ctrl+f Move forward one screen
ctrl+b Move back one screen

ctrl+d Move forward half a screen
ctrl+u Move back half a screen

H Move the cursor to the top of the screen
M Move the cursor to the middle of the screen
L Move the cursor to the bottom of the screen

{ Move to the previous blank line
} Move to the next blank line

ctrl+y Scroll up while keeping the cursor fixed
ctrl+e Scroll down while keeping the cursor fixed

zEnter or zt Move the line with the cursor to the top of the screen
zz Move the line with the cursor to the center of the screen
z- or zb Move the line with the cursor to the bottom of the screen

fWord Jump to Word (any character)
tWord Jump to one character before Word (any character)

FWord Jump to Word (any character) in reverse direction
tWord Jump to one character before Word (any character) in reverse direction

;/, Move forward/backward through search results for any character

. Repeat
; Repeat (reverse direction)

v visual mode
ctrl+v visual block mode
V visual line mode

# Editing
i insert mode
I insert mode from the start
a insert mode from one character after the cursor
A insert mode from the end of the line
l insert mode from the start of the line
o insert mode from the next line
O insert mode from the previous line
s delete one character under the cursor and enter insert mode
S delete the line under the cursor and enter insert mode

r Edit one character (return to normal mode after Enter)
R Edit multiple characters (return to normal mode after ESC)

dd Delete the line under the cursor
3+dd Delete 3 lines

d$ Delete to the end
d^ Delete to the start

dw Delete to the start of the next word (including spaces)
db Delete to the start of the previous word (including spaces)

de Delete to the end of the next word (excluding spaces)
dge Delete to the end of the previous word (excluding spaces)

diw Delete the word under the cursor and enter insert mode
daw Delete the word under the cursor and the following spaces, then enter insert mode

u Undo
ctrl+r Redo

p Paste the cut or copied line

yy Yank the line under the cursor
yi( Yank the contents enclosed by symbols

cw Change the word, enter insert mode
c$ Change to the end, enter insert mode
c^ Change to the start, enter insert mode
c0 Change to the start ignoring indentation, enter insert mode
ci( Change the contents enclosed by symbols
cit Delete the contents of a tag and enter insert mode

\>> Indent one level
<< Unindent one level

ctrl+p Autocomplete words that appear in all files opened in vim/select the previous completion candidate
ctrl+n Select the next completion candidate
ctrl+y Confirm the selected completion candidate
ctrl+x ctrl+l Line completion (complete with lines matching those in the opened file)
ctrl+x ctrl+f File path completion

:s/thee/the Replace the first occurrence of a word on the line with the cursor
:s/thee/the/g Replace all occurrences of a word on the line with the cursor
:1,100s/thee/the Replace the first occurrence of a word on each line from line 1 to 100
:1,100s/thee/the/g Replace all occurrences of a word from line 1 to 100
:%s/thee/the/g Replace all occurrences of a word in the entire file
:%s/thee/the/gc Replace all occurrences of a word in the entire file with confirmation

# Search
/ Search for a phrase, move forward with n, backward with N

# Insert Normal Mode
ctrl+o Execute a normal mode command once and return to insert mode

# Other Notes
- Remember basic commands (repeated practice)
- Operator + number + motion (basic application)

# Created a Practice Drill
- [github - road-to-vimmer](https://github.com/bmf-san/road-to-vimmer)

# References
- vimtutor (tutorial included in vim)
