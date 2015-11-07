package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/nsf/termbox-go"
)

// Colors
const backgroundColor = termbox.ColorBlue
const boardColor = termbox.ColorBlack
const instructionsColor = termbox.ColorYellow

var pieceColors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
}

// Layout
const stretchFactor = 2

const widthRatio = 1
const heightRatio = 1.6

const defaultMarginWidth = 2
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight

const boardWidth = 10 * widthRatio * stretchFactor
const boardHeight = boardWidth * heightRatio * stretchFactor

const cellWidth = 2 * widthRatio * stretchFactor
const boardEndX = boardStartX + boardWidth*cellWidth
const boardEndY = boardStartY + boardHeight
const instructionsStartX = boardEndX + defaultMarginWidth
const instructionsStartY = boardStartY

// Text in the UI
const title = "TETRIS WRITTEN IN GO"

var instructions = []string{
	"Goal: Fill in 5 lines!",
	"",
	"left   Left",
	"right  Right",
	"up     Rotate",
	"down   Down",
	"space  Fall",
	"s      Start",
	"p      Pause",
	"esc,q  Exit",
	"",
	"Level: %v",
	"Lines: %v",
	"",
	"GAME OVER!",
}

// This takes care of rendering everything.
func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	tbprint(titleStartX, titleStartY, instructionsColor, backgroundColor, title)
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			cellValue := g.board[y][x]
			absCellValue := int(math.Abs(float64(cellValue)))
			cellColor := pieceColors[absCellValue]
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(
					boardStartX+cellWidth*x+i,
					boardStartY+y,
					' ',       // rune
					cellColor, // fg
					cellColor /*bg*/)
			}
		}
	}
	for y, instruction := range instructions {
		if strings.HasPrefix(instruction, "Level:") {
			instruction = fmt.Sprintf(instruction, g.level)
		} else if strings.HasPrefix(instruction, "Lines:") {
			instruction = fmt.Sprintf(instruction, g.numLines)
		} else if strings.HasPrefix(instruction, "GAME OVER") && g.state != gameOver {
			instruction = ""
		}
		tbprint(instructionsStartX, instructionsStartY+y, instructionsColor, backgroundColor, instruction)
	}
	termbox.Flush()
}

// Function tbprint draws a string.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
