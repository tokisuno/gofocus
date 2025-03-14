package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const STUDY_MINUTES int = 50 * 60
const BREAK_MINUTES int = 10 * 60

const screenW = int32(800)
const screenH = int32(600)

type GameState int

const (
	waitSTUDY GameState = iota
	STUDY
	waitBREAK	
	BREAK
)

type hms struct {
	hrs  int
	mins int
	secs int
}

func newHms(seconds int) *hms {
	var time hms

	time.hrs = seconds / 3600
	time.mins = (seconds / 60) % 60
	time.secs = seconds % 60

	return &time
}

func main() {
	rl.InitWindow(800, 600, "Raylib Example")
	defer rl.CloseWindow()

	studyTime := STUDY_MINUTES
	breakTime := BREAK_MINUTES
	studyTimerHms := newHms(studyTime)
	breakTimerHms := newHms(breakTime)

	var currentScreen GameState
	currentScreen = waitSTUDY

	frames := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		switch currentScreen {
		case waitSTUDY:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreen = STUDY
			}
		case STUDY:
			frames++
			if frames > 60 {
				studyTime--
				studyTimerHms = newHms(studyTime)
				frames = 0
			}
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreen = waitBREAK
				studyTime = STUDY_MINUTES
				studyTimerHms = newHms(studyTime)
			}
		case waitBREAK:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreen = BREAK
			}
		case BREAK:
			frames++
			if frames > 60 {
				breakTime--
				breakTimerHms = newHms(breakTime)
				frames = 0
			}
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreen = waitSTUDY 
				breakTime = BREAK_MINUTES 
				breakTimerHms = newHms(breakTime)
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		// rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))

		switch currentScreen {
		case waitSTUDY:
			rl.DrawText("Study time!", 100, 100, 20, rl.White)
		case STUDY:
			rl.DrawText(
				fmt.Sprintf("%.2d:%.2d:%.2d", studyTimerHms.hrs, studyTimerHms.mins, studyTimerHms.secs),
				100, 100, 20, rl.White)

		case waitBREAK:
			rl.DrawText("Break time!", 100, 100, 20, rl.White)
		case BREAK:
			rl.DrawText(
				fmt.Sprintf("%.2d:%.2d:%.2d", breakTimerHms.hrs, breakTimerHms.mins, breakTimerHms.secs),
				100, 100, 20, rl.White)
		}	

		rl.EndDrawing()
	}
}
