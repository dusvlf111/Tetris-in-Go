package main

import (
	"github.com/nsf/termbox-go" // 터미널 기반의 UI를 처리하기 위한 패키지
)

const (
	boardWidth  = 10 // 게임 보드의 너비
	boardHeight = 20 // 게임 보드의 높이
)

// Game 구조체는 게임 보드와 현재 활성화된 테트로미노를 포함합니다.
type Game struct {
	board   [][]int   // 게임 보드의 상태를 나타내는 2차원 배열
	current Tetromino // 현재 활성화된 테트로미노
}

// NewGame 함수는 새로운 게임 인스턴스를 생성하고 초기화합니다.
func NewGame() *Game {
	// 게임 보드를 초기화합니다. 빈 배열을 사용하여 초기 상태를 설정합니다.
	board := make([][]int, boardHeight)
	for i := range board {
		board[i] = make([]int, boardWidth)
	}

	// 새 게임 인스턴스를 생성하고 초기화합니다.
	game := &Game{
		board: board,
	}
	// 새로운 테트로미노를 생성하여 게임에 설정합니다.
	game.spawnTetromino()
	return game
}

// spawnTetromino 함수는 게임에 새로운 테트로미노를 생성하고 설정합니다.
func (g *Game) spawnTetromino() {
	// 무작위로 새로운 테트로미노를 생성하여 현재 테트로미노로 설정합니다.
	g.current = NewRandomTetromino()
}

// DrawBoard 함수는 게임 보드와 현재 테트로미노를 화면에 그립니다.
func (g *Game) DrawBoard() {
	// 게임 보드의 각 셀을 화면에 그립니다.
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if g.board[y][x] != 0 {
				// 보드 셀에 블록이 있으면 그립니다.
				termbox.SetCell(x, y, '█', termbox.ColorDefault, termbox.ColorDefault)
			} else {
				// 보드 셀에 블록이 없으면 빈 공간으로 표시합니다.
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}

	// 현재 테트로미노를 화면에 그립니다.
	for y := 0; y < len(g.current.shape); y++ {
		for x := 0; x < len(g.current.shape[y]); x++ {
			if g.current.shape[y][x] != 0 {
				// 테트로미노의 블록을 현재 위치에 그립니다.
				termbox.SetCell(g.current.pos.x+x, g.current.pos.y+y, '█', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}

	// 화면에 모든 변경 사항을 적용합니다.
	termbox.Flush()
}

// MoveTetromino 함수는 현재 테트로미노를 주어진 방향으로 이동시킵니다.
func (g *Game) MoveTetromino(dx, dy int) {
	// 테트로미노의 위치를 이동시킵니다.
	g.current.pos.x += dx
	g.current.pos.y += dy
}

// Update 함수는 게임 상태를 주기적으로 업데이트합니다.
func (g *Game) Update() {
	// 현재 테트로미노를 아래로 이동시킵니다.
	g.MoveTetromino(0, 1)
}
