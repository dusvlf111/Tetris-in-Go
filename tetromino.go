package main

import (
	"math/rand" // 난수 생성 패키지
	"time"      // 현재 시간을 가져오기 위한 패키지
)

// Point 구조체는 테트로미노의 위치를 나타냅니다.
type Point struct {
	x, y int // x와 y 좌표
}

// Tetromino 구조체는 테트로미노의 형태와 위치를 나타냅니다.
type Tetromino struct {
	shape [][]int // 테트로미노의 형태를 정의하는 2차원 배열
	pos   Point   // 테트로미노의 현재 위치
}

// tetrominos 변수는 가능한 모든 테트로미노 형태를 정의합니다.
var tetrominos = [][][]int{
	// I 모양
	{
		{1, 1, 1, 1},
	},
	// O 모양
	{
		{1, 1},
		{1, 1},
	},
	// T 모양
	{
		{0, 1, 0},
		{1, 1, 1},
	},
	// S 모양
	{
		{0, 1, 1},
		{1, 1, 0},
	},
	// Z 모양
	{
		{1, 1, 0},
		{0, 1, 1},
	},
	// J 모양
	{
		{1, 0, 0},
		{1, 1, 1},
	},
	// L 모양
	{
		{0, 0, 1},
		{1, 1, 1},
	},
}

// init 함수는 패키지가 초기화될 때 호출됩니다.
// 랜덤 시드를 현재 시간으로 설정하여 매번 다른 난수를 생성합니다.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRandomTetromino 함수는 무작위로 새로운 테트로미노를 생성하여 반환합니다.
func NewRandomTetromino() Tetromino {
	// tetrominos 배열에서 랜덤으로 하나의 테트로미노 형태를 선택합니다.
	return Tetromino{
		shape: tetrominos[rand.Intn(len(tetrominos))],
		// 테트로미노의 초기 위치를 보드의 중간 상단에 설정합니다.
		pos: Point{x: boardWidth/2 - 2, y: 0},
	}
}
