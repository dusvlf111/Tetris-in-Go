package main

import (
	"time" // 시간 관련 기능을 사용하기 위한 패키지

	"github.com/nsf/termbox-go" // termbox 패키지를 임포트하여 터미널 기반의 UI를 처리합니다.
)

func main() {
	// termbox를 초기화합니다. 화면을 터미널에 출력하기 위해 필요합니다.
	err := termbox.Init()
	if err != nil {
		// 초기화 오류가 발생하면 패닉을 발생시킵니다.
		panic(err)
	}
	// 함수가 종료될 때 termbox를 닫아 리소스를 정리합니다.
	defer termbox.Close()

	// 게임을 초기화합니다.
	game := NewGame()

	// 게임 상태를 주기적으로 업데이트하기 위해 타이머를 설정합니다.
	ticker := time.NewTicker(time.Millisecond * 50) // 500밀리초마다 타이머 이벤트 발생
	defer ticker.Stop()                             // 함수가 종료되면 타이머를 정리합니다.

	// 메인 게임 루프
	for {
		select {
		// 타이머가 트리거되면 이 케이스가 실행됩니다.
		case <-ticker.C:
			// 게임 상태를 업데이트합니다. 예를 들어, 테트로미노를 아래로 이동시킵니다.
			game.Update()
			// 게임 보드를 화면에 그립니다.
			game.DrawBoard()
		default:
			// 사용자 입력 이벤트를 처리합니다.
			ev := termbox.PollEvent()        // termbox에서 이벤트를 폴링합니다.
			if ev.Type == termbox.EventKey { // 이벤트가 키 입력인지 확인합니다.
				switch ev.Key {
				// 왼쪽 화살표 키를 누르면 테트로미노를 왼쪽으로 이동시킵니다.
				case termbox.KeyArrowLeft:
					game.MoveTetromino(-1, 0)
				// 오른쪽 화살표 키를 누르면 테트로미노를 오른쪽으로 이동시킵니다.
				case termbox.KeyArrowRight:
					game.MoveTetromino(1, 0)
				// 아래쪽 화살표 키를 누르면 테트로미노를 아래로 빠르게 이동시킵니다.
				case termbox.KeyArrowDown:
					game.MoveTetromino(0, 1)
				// ESC 키를 누르면 게임을 종료합니다.
				case termbox.KeyEsc:
					return // 메인 루프를 종료하고 프로그램을 종료합니다.
				}
			}
		}
	}
}
