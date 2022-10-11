package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const screenWidth = 900
	const screenHeight = 700

	const p_width = 25
	const p_height = 90

	rl.InitWindow(screenWidth, screenHeight, "Hello, Raylib!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ballPosition := rl.Vector2{screenWidth / 2, screenHeight / 2}
	ballSpeed := rl.Vector2{5, 3}
	ballRadius := 12

	aScore := 0
	bScore := 0

	p_one := rl.Rectangle{30, float32(rl.GetScreenHeight()) / 2, p_width, p_height}
	p_two := rl.Rectangle{float32(rl.GetScreenWidth() - 60), float32(rl.GetScreenHeight()) / 2, p_width, p_height}

	pause := false

	for !rl.WindowShouldClose() {

		//ball := rl.Rectangle {float32(ballPosition.X), float32(ballPosition.Y), 20, 20}

		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		if !pause {
			ballPosition.X += ballSpeed.X
			ballPosition.Y += ballSpeed.Y

			// when collided with player bar
			if rl.CheckCollisionCircleRec(ballPosition, float32(ballRadius), p_one) || rl.CheckCollisionCircleRec(ballPosition, float32(ballRadius), p_two) {
				ballSpeed.X *= -1
			}

			// when collided with window border
			if ballPosition.X >= float32(rl.GetScreenWidth())-float32(ballRadius) {
				aScore++
				ballPosition = rl.Vector2{screenWidth / 2, screenHeight / 2}
			}

			if ballPosition.X <= float32(ballRadius) {
				bScore++
				ballPosition = rl.Vector2{screenWidth / 2, screenHeight / 2}

			}

			if ballPosition.Y >= float32(rl.GetScreenHeight())-float32(ballRadius) || ballPosition.Y <= float32(ballRadius) {
				ballSpeed.Y *= -1
			}

			if rl.IsKeyDown(rl.KeyDown) {
				p_two.Y += 6
			}
			if rl.IsKeyDown(rl.KeyUp) {
				p_two.Y -= 6
			}
			if rl.IsKeyDown(rl.KeyC) {
				p_one.Y += 6
			}
			if rl.IsKeyDown(rl.KeyV) {
				p_one.Y -= 6
			}

			if p_one.Y <= float32(rl.GetScreenHeight()) {
				p_one = rl.Rectangle{p_one.X, p_one.Y, p_width, p_height}
			}

		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawCircle(int32(ballPosition.X), int32(ballPosition.Y), float32(ballRadius), rl.DarkBlue)
		//rl.DrawRectangle(int32(ball.X), int32(ball.Y), int32(ball.Width), int32(ball.Height), rl.DarkBlue)

		if pause {
			rl.DrawText("Paused", 100, 100, 100, rl.White)
		}

		rl.DrawRectangle(int32(p_one.X), int32(p_one.Y), 20, 90, rl.Purple)
		rl.DrawRectangle(int32(p_two.X), int32(p_two.Y), 20, 90, rl.Purple)

		text := "%d : %d"
		output := fmt.Sprintf(text, aScore, bScore)

		rl.DrawText(output, int32(rl.GetScreenWidth())/2, 50, 20, rl.White)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
