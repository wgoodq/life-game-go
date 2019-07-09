package domain

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const BoardSize = 7

type LifeGameService interface {
	RunGame() (int, string)
}

type LifeGame struct {
	mu       sync.Mutex
	cnt      int
	board    [BoardSize][BoardSize]int
	newBoard [BoardSize][BoardSize]int
}

func (lg *LifeGame) RunGame() (int, string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	if lg.cnt == 0 || lg.sumBoard() == 0 {

		lg.initLifeCells()
	} else {

		lg.evolution()
	}

	lg.cnt++
	return lg.writeResponse()
}

// 用随机数初始化棋格。
func (lg *LifeGame) initLifeCells() {
	rand.Seed(time.Now().UnixNano())

	for l := range lg.board {

		for c := range lg.board[l] {

			lg.board[l][c] = int(rand.Intn(2))
		}
	}

	log.Printf("init board: %v", lg.board)
}

func (lg *LifeGame) writeResponse() (int, string) {
	var cells []string

	for _, line := range lg.board {

		for _, cell := range line {

			if cell == 0 {

				cells = append(cells, "⬜")
			} else {

				cells = append(cells, "⬛")
			}
		}

		cells = append(cells, "<br>")
	}

	return lg.cnt, strings.Join(cells, "")
}

func (lg *LifeGame) evolution() {

	for l := range lg.board {

		for c := range lg.board[l] {

			sum := lg.sumNeighbor(l, c)
			if sum == 3 || sum == 4 {

				lg.newBoard[l][c] = 1
			} else {

				lg.newBoard[l][c] = 0
			}
		}
	}

	// 复制型赋值。
	lg.board = lg.newBoard
}

func (lg *LifeGame) sumNeighbor(l, c int) int {
	var idx = []int{-1, 0, 1}
	sum := 0

	for _, x := range idx {

		for _, y := range idx {

			if -1 < l+x && l+x < len(lg.board) && -1 < c+y && c+y < len(lg.board[l]) {

				sum += lg.board[l+x][c+y]
			}
		}
	}

	return sum
}

func (lg *LifeGame) sumBoard() int {
	sum := 0

	for _, line := range lg.board {

		for _, cell := range line {

			sum += cell
		}
	}

	return sum
}
