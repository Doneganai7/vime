package main
import "github.com/Doneganai7/vime"

func main() {
    var game vime.Vime
    game.Player = "0"
    game.Player_alt = " "
    game.Obstruction = "#"
    game.Danger = "!"
    game.Platform = "@"
    game.Launcher_r = ">"
    game.Launcher_l = "<"
    game.Launcher_u = "^"
    game.Launcher_d = "v"
    game.Initialize()
    game.Run()
}
