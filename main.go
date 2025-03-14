package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const STUDY_MINUTES int = 50 * 60
const BREAK_MINUTES int = 10 * 60

const screenW = int32(900)
const screenH = int32(300)

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
	rl.InitWindow(screenW, screenH, "Raylib Example")

	defer rl.CloseWindow()

	fontSize := 140
	font := rl.LoadFontEx("assets/JetBrainsMono-VariableFont_wght.ttf", int32(fontSize), nil)

	studyTime := STUDY_MINUTES
	breakTime := BREAK_MINUTES

	studyTimerHms := newHms(studyTime)
	breakTimerHms := newHms(breakTime)

	studyText := fmt.Sprintf("%.2d:%.2d:%.2d", studyTimerHms.hrs, studyTimerHms.mins, studyTimerHms.secs)
	breakText := fmt.Sprintf("%.2d:%.2d:%.2d", breakTimerHms.hrs, breakTimerHms.mins, breakTimerHms.secs)

	var currentScreen GameState
	currentScreen = waitSTUDY

	rl.SetExitKey(rl.KeyNull)

	frames := 0

	fontPosition := rl.NewVector2(100, 100)

	studyTextSize := rl.MeasureTextEx(font, studyText, float32(fontSize), 0)
	breakTextSize := rl.MeasureTextEx(font, studyText, float32(fontSize), 0)

	studyTextPosition := rl.NewVector2(float32(screenW)/2 - (studyTextSize.X/2), 90)
	studyShadowPosition1 := studyTextPosition
	studyShadowPosition1.X += 2
	studyShadowPosition1.Y += 5

	studyShadowPosition2 := studyShadowPosition1
	studyShadowPosition2.X += 2
	studyShadowPosition2.Y += 5

	breakTextPosition := rl.NewVector2(float32(screenW)/2 - (breakTextSize.X/2), 90)
	breakShadowPosition1 := breakTextPosition
	breakShadowPosition1.X += 2
	breakShadowPosition1.Y += 5

	breakShadowPosition2 := breakShadowPosition1
	breakShadowPosition2.X += 2
	breakShadowPosition2.Y += 5

	rl.SetTextureFilter(font.Texture, rl.TextureFilterNearest)

	bgColour := rl.GetColor(0x282828ff)
	textColour := rl.GetColor(0xfbf1c7ff)
	textShadowColour1 := rl.GetColor(0xa89984ff)
	textShadowColour2 := rl.GetColor(0x655c54ff)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		switch currentScreen {
		case waitSTUDY:
			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyY) {
				currentScreen = STUDY
			}
		case STUDY:
			frames++
			if frames == 60 {
				studyTime--
				studyTimerHms = newHms(studyTime)
				studyText = fmt.Sprintf("%.2d:%.2d:%.2d", studyTimerHms.hrs, studyTimerHms.mins, studyTimerHms.secs)
				frames = 0
			}
			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyY) {
				currentScreen = waitBREAK
				studyTime = STUDY_MINUTES
				studyTimerHms = newHms(studyTime)
				studyText = fmt.Sprintf("%.2d:%.2d:%.2d", studyTimerHms.hrs, studyTimerHms.mins, studyTimerHms.secs)
			}
		case waitBREAK:
			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyY) {
				currentScreen = BREAK
			}
		case BREAK:
			frames++
			if frames == 60 {
				breakTime--
				breakTimerHms = newHms(breakTime)
				frames = 0
				breakText = fmt.Sprintf("%.2d:%.2d:%.2d", breakTimerHms.hrs, breakTimerHms.mins, breakTimerHms.secs)
			}
			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyY) {
				currentScreen = waitSTUDY 
				breakTime = BREAK_MINUTES 
				breakTimerHms = newHms(breakTime)
				breakText = fmt.Sprintf("%.2d:%.2d:%.2d", breakTimerHms.hrs, breakTimerHms.mins, breakTimerHms.secs)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(bgColour)
		// rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))

		switch currentScreen {
		case waitSTUDY:
			rl.DrawTextEx(font, "Study time!", fontPosition, float32(fontSize), 0, textColour)
		case STUDY:
			rl.DrawTextEx(font, studyText, studyShadowPosition2, float32(fontSize), 0, textShadowColour2)
			rl.DrawTextEx(font, studyText, studyShadowPosition1, float32(fontSize), 0, textShadowColour1)
			rl.DrawTextEx(font, studyText, studyTextPosition, float32(fontSize), 0, textColour)
		case waitBREAK:
			rl.DrawTextEx(font, "Break time!", fontPosition, float32(fontSize), 0, textColour)
		case BREAK:
			rl.DrawTextEx(font, breakText, breakShadowPosition2, float32(fontSize), 0, textShadowColour2)
			rl.DrawTextEx(font, breakText, breakShadowPosition1, float32(fontSize), 0, textShadowColour1)
			rl.DrawTextEx(font, breakText, breakTextPosition, float32(fontSize), 0, textColour)
		}	

		rl.EndDrawing()
	}

	rl.UnloadFont(font)
	rl.UnloadDroppedFiles()
	rl.CloseWindow()
}
