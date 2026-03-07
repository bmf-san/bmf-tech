---
title: Road to Vimmer
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
I will list the commands I learned to become a Vimmer.

# Motivation
- Although I got used to the shortcuts of GUI editors, I feel that my text editing speed has plateaued.
- Reducing dependency on GUI editors (which is more cost-effective: memorizing shortcuts for GUI editors or memorizing shortcuts for the mature editor VIM? vi is included as the standard editor in all OS).
- The use of shortcuts and arrow keys causes my hands to leave the home position, which decreases my typing speed.
- My desire to complete text editing quickly on the terminal has increased.

# Navigation
h Move left by logical line
j Move down by logical line
k Move up by logical line
l Move right by logical line

^ Move to the beginning
0 Move to the beginning, ignoring indentation
$ Move to the end
\+ Move to the beginning of the next line
\- Move to the beginning of the previous line

:3 Move to line 3

w Move to the beginning of the next word (including spaces)
b Move to the beginning of the previous word (including spaces)

e Move to the end of the next word (excluding spaces)
ge Move to the end of the previous word (excluding spaces)

% Jump to the matching closing tag for (, [, {

ctrl+f Move forward one screen
ctrl+b Move back one screen

ctrl+d Move forward half a screen
ctrl+u Move back half a screen

H Jump to the top of the screen
M Jump to the middle of the screen
L Jump to the bottom of the screen

{ Move to the previous empty line
} Move to the next empty line

ctrl+y Scroll up while keeping the cursor fixed
ctrl+e Scroll down while keeping the cursor fixed

zEnter or zt Make the line with the cursor the top of the screen
zz Make the line with the cursor the center of the screen
z- or zb Make the line with the cursor the bottom of the screen

fWord Jump to Word (any character)
tWord Jump to one character before Word (any character)

FWord Jump to Word (any character) in the opposite direction
tWord Jump to one character before Word (any character) in the opposite direction

;/, Move to the next/previous search result for any character

. Repeat
; Repeat (in the opposite direction)

v Visual mode
ctrl+v Visual block mode
V Visual line mode

# Editing
i Insert mode
I Insert mode from the beginning
a Insert mode from one character behind the cursor
A Insert mode from the end of the line
l Insert mode from the beginning of the line
o Insert mode from the next line
O Insert mode from the previous line
s Delete the character under the cursor and enter insert mode
S Delete the line under the cursor and enter insert mode

r Edit one character (normal mode after Enter)
R Edit multiple characters (normal mode after ESC)

dd Delete the cursor line
3+dd Delete 3 lines

d$ Delete to the end
d^ Delete to the beginning

dw Delete to the beginning of the next word (including spaces)
db Delete to the beginning of the previous word (including spaces)

de Delete to the end of the next word (excluding spaces)
dge Delete to the end of the previous word (excluding spaces)

diw Delete the word under the cursor and enter insert mode
daw Delete the word under the cursor and the following spaces, then enter insert mode

u Undo
ctrl+r Redo

p Paste the cut or copied line

yy Yank the cursor line
yi( Yank the contents surrounded by symbols

cw Change the word, insert mode
c$ Change to the end, insert mode
c^ Change to the beginning, insert mode
c0 Change to the beginning without indentation, insert mode
ci( Change the contents surrounded by symbols
cit Delete the contents of the tag and enter insert mode

\>> Decrease indentation by one
<< Increase indentation by one

ctrl+p Select word completion from all files opened in vim / Select the previous completion candidate
ctrl+n Select the next completion candidate
ctrl+y Confirm the selected completion candidate
ctrl+x ctrl+l Line completion (completes with lines matching those in opened files)
ctrl+x ctrl+f File path completion

:s/thee/the Replace the first found word in the line where the cursor is
:s/thee/the/g Replace all found words in the entire line where the cursor is
:1,100s/thee/the Replace the first found word in each line from 1 to 100
:1,100s/thee/the/g Replace all found words in lines from 1 to 100
:%s/thee/the/g Replace all found words in the entire file
:%s/thee/the/gc Replace all found words in the entire file while confirming

# Search
/ Search for phrases n moves forward, N moves backward

# Insert Normal Mode
ctrl+o Execute a normal mode command once and return to insert mode

# Other Notes
- Memorize basic commands (repetitive practice)
- Operator + number + motion (application of basics)

# Practice Drill Created
- [github - road-to-vimmer](https://github.com/bmf-san/road-to-vimmer)

# References
- vimtutor (the tutorial included in vim)