package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const STUDY_MINUTES int = 50 * 60
const BREAK_MINUTES int = 10 * 60
const SESSION_AMNT int = 4

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
	rl.InitWindow(screenW, screenH, "GoFocus!")
	rl.InitAudioDevice()

	defer rl.CloseWindow()

	fontSize := 100
	font := rl.LoadFontEx("assets/JetBrainsMono-VariableFont_wght.ttf", int32(fontSize), nil)

	studyTime := STUDY_MINUTES
	breakTime := BREAK_MINUTES

	studyTimerHms := newHms(studyTime)
	breakTimerHms := newHms(breakTime)

	session_begin := rl.LoadSound("assets/menu.ogg")
	session_end := rl.LoadSound("assets/ping.ogg")
	bonk := rl.LoadSound("assets/bonk.ogg")
	rl.SetSoundVolume(session_begin, 0.6)
	rl.SetSoundVolume(session_end, 0.6)
	rl.SetSoundVolume(bonk, 0.6)

	sound_played := false

	studyText := fmt.Sprintf("%.2d:%.2d:%.2d", studyTimerHms.hrs, studyTimerHms.mins, studyTimerHms.secs)
	breakText := fmt.Sprintf("%.2d:%.2d:%.2d", breakTimerHms.hrs, breakTimerHms.mins, breakTimerHms.secs)

	sessionCounter := 0
	totalSessions := SESSION_AMNT

	sessionFontSize := 60
	sessionFont := rl.LoadFontEx("assets/JetBrainsMono-VariableFont_wght.ttf", int32(sessionFontSize), nil)
	sessionText := fmt.Sprintf("Session %d/%d!", sessionCounter, totalSessions)
	sessionTextSize := rl.MeasureTextEx(sessionFont, sessionText, float32(sessionFontSize), 0)
	sessionTextPosition := rl.NewVector2(float32(screenW)/2-(sessionTextSize.X/2), (float32(screenH)-sessionTextSize.Y)-30)

	var currentScreen GameState
	currentScreen = waitSTUDY

	rl.SetExitKey(rl.KeyNull)

	frames := 0

	studySplash := "Study time!"
	studySplashSize := rl.MeasureTextEx(font, studySplash, float32(fontSize), 0)
	studySplashPosition := rl.NewVector2(float32(screenW)/2-(studySplashSize.X/2), 70)

	studySplashPosition1 := studySplashPosition
	studySplashPosition1.X += 2
	studySplashPosition1.Y += 5

	studySplashPosition2 := studySplashPosition1
	studySplashPosition2.X += 2
	studySplashPosition2.Y += 5

	breakSplash := "Time to rest :3"
	breakSplashSize := rl.MeasureTextEx(font, breakSplash, float32(fontSize), 0)
	breakSplashPosition := rl.NewVector2(float32(screenW)/2-(breakSplashSize.X/2), 70)

	breakSplashPosition1 := breakSplashPosition
	breakSplashPosition1.X += 2
	breakSplashPosition1.Y += 5

	breakSplashPosition2 := breakSplashPosition1
	breakSplashPosition2.X += 2
	breakSplashPosition2.Y += 5

	studyTextSize := rl.MeasureTextEx(font, studyText, float32(fontSize), 0)
	breakTextSize := rl.MeasureTextEx(font, studyText, float32(fontSize), 0)

	studyTextPosition := rl.NewVector2(float32(screenW)/2-(studyTextSize.X/2), 70)
	studyShadowPosition1 := studyTextPosition
	studyShadowPosition1.X += 2
	studyShadowPosition1.Y += 5

	studyShadowPosition2 := studyShadowPosition1
	studyShadowPosition2.X += 2
	studyShadowPosition2.Y += 5

	breakTextPosition := rl.NewVector2(float32(screenW)/2-(breakTextSize.X/2), 70)
	breakShadowPosition1 := breakTextPosition
	breakShadowPosition1.X += 2
	breakShadowPosition1.Y += 5

	breakShadowPosition2 := breakShadowPosition1
	breakShadowPosition2.X += 2
	breakShadowPosition2.Y += 5

	rl.SetTextureFilter(font.Texture, rl.TextureFilterNearest)

	bgColour := rl.GetColor(0x1d2021ff)
	textColour := rl.GetColor(0xebdbb2ff)
	textShadowColour1 := rl.GetColor(0x504945ff)
	textShadowColour2 := rl.GetColor(0x282828ff)
	sessionTextColour := rl.GetColor(0xffff0045)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyA) {
			sessionCounter++
			sessionText = fmt.Sprintf("Session %d/%d!", sessionCounter, totalSessions)
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyX) {
			sessionCounter--
			sessionText = fmt.Sprintf("Session %d/%d!", sessionCounter, totalSessions)
		}
		switch currentScreen {
		case waitSTUDY:
			if sound_played == false {
				rl.PlaySound(bonk)
				sound_played = true
			}

			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyY) {
				currentScreen = STUDY
				rl.PlaySound(session_begin)
				sound_played = false
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
				rl.PlaySound(session_end)
			}
		case waitBREAK:
			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyY) {
				currentScreen = BREAK
				rl.PlaySound(session_begin)
				sound_played = false
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

		switch currentScreen {
		case waitSTUDY:
			rl.DrawTextEx(font, studySplash, studySplashPosition2, float32(fontSize), 0, textShadowColour2)
			rl.DrawTextEx(font, studySplash, studySplashPosition1, float32(fontSize), 0, textShadowColour1)
			rl.DrawTextEx(font, studySplash, studySplashPosition, float32(fontSize), 0, textColour)
			rl.DrawTextEx(sessionFont, sessionText, sessionTextPosition, float32(sessionFontSize), 0, sessionTextColour)
		case STUDY:
			rl.DrawTextEx(font, studyText, studyShadowPosition2, float32(fontSize), 0, textShadowColour2)
			rl.DrawTextEx(font, studyText, studyShadowPosition1, float32(fontSize), 0, textShadowColour1)
			rl.DrawTextEx(font, studyText, studyTextPosition, float32(fontSize), 0, textColour)
			rl.DrawTextEx(sessionFont, sessionText, sessionTextPosition, float32(sessionFontSize), 0, sessionTextColour)
		case waitBREAK:
			rl.DrawTextEx(font, breakSplash, breakSplashPosition2, float32(fontSize), 0, textShadowColour2)
			rl.DrawTextEx(font, breakSplash, breakSplashPosition1, float32(fontSize), 0, textShadowColour1)
			rl.DrawTextEx(font, breakSplash, breakSplashPosition, float32(fontSize), 0, textColour)
			rl.DrawTextEx(sessionFont, sessionText, sessionTextPosition, float32(sessionFontSize), 0, sessionTextColour)
		case BREAK:
			rl.DrawTextEx(font, breakText, breakShadowPosition2, float32(fontSize), 0, textShadowColour2)
			rl.DrawTextEx(font, breakText, breakShadowPosition1, float32(fontSize), 0, textShadowColour1)
			rl.DrawTextEx(font, breakText, breakTextPosition, float32(fontSize), 0, textColour)
			rl.DrawTextEx(sessionFont, sessionText, sessionTextPosition, float32(sessionFontSize), 0, sessionTextColour)
		}

		rl.EndDrawing()
	}

	rl.UnloadFont(font)
	rl.UnloadDroppedFiles()
	rl.CloseWindow()
}
