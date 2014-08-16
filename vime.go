package vime
import "fmt"
import "math/rand"
import "time"

type Vime struct {
    points int
    player_x int
    player_y int
    field [][]string
    instruction string
    result string
    last string
    lost bool
    auto bool
    launch_count int
    death string

    Win_condition int
    Field_limit int

    Obstruction string
    Objective string
    Danger string
    Platform string
    Penalty string
    Empty string
    Launcher_r string
    Launcher_l string
    Launcher_u string
    Launcher_d string
    Player string
    Player_alt string

    Obstruction_prob int
    Objective_prob int
    Danger_prob int
    Platform_prob int
    Penalty_prob int
    Launcher_r_prob int
    Launcher_l_prob int
    Launcher_u_prob int
    Launcher_d_prob int
}
func (this *Vime) Initialize() {
    this.lost = false
    this.auto = false
    rand.Seed(time.Now().UTC().UnixNano())

    if this.Field_limit == 0 { this.Field_limit = 31 }
    if this.player_x == 0 { this.player_x = this.Field_limit / 2 }
    if this.player_y == 0 { this.player_y = this.Field_limit / 2 }
    if this.Win_condition == 0 { this.Win_condition = 20 }

    if this.Empty == "" { this.Empty = " " }
    if this.Player == "" { this.Player = "⍟" }
    if this.Player_alt == "" { this.Player_alt = "✪" }
    if this.Obstruction == "" { this.Obstruction = "𝌆" }
    if this.Objective == "" { this.Objective = "+" }
    if this.Danger == "" { this.Danger = "⚠" }
    if this.Penalty == "" { this.Penalty = "-" }
    if this.Platform == "" { this.Platform = "⛀" }
    if this.Launcher_r == "" { this.Launcher_r = "→" }
    if this.Launcher_l == "" { this.Launcher_l = "←" }
    if this.Launcher_u == "" { this.Launcher_u = "↑" }
    if this.Launcher_d == "" { this.Launcher_d = "↓" }

    if this.Danger_prob == 0 { this.Danger_prob = 10 }
    if this.Penalty_prob == 0 { this.Penalty_prob = 20 }
    if this.Obstruction_prob == 0 { this.Obstruction_prob = 15 }
    if this.Platform_prob == 0 { this.Platform_prob = 5 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Objective_prob == 0 { this.Objective_prob = 4 }

    this.last = this.Empty
    this.field = make([][]string,this.Field_limit)
    for i := 0; i < this.Field_limit; i++ {
        this.field[i] = make([]string,this.Field_limit)
    }
    var objectives_generated int
    objectives_generated = 0
    for i := 0; i < this.Field_limit; i++ {
        for j := 0; j < this.Field_limit; j++ {
            this.field[i][j] = this.Empty
            if rand.Intn(100-0) < this.Danger_prob { this.field[i][j] = this.Danger }
            if rand.Intn(100-0) < this.Penalty_prob { this.field[i][j] = this.Penalty }
            if rand.Intn(100-0) < this.Obstruction_prob { this.field[i][j] = this.Obstruction }
            if rand.Intn(100-0) < this.Platform_prob { this.field[i][j] = this.Platform }
            if rand.Intn(100-0) < this.Launcher_r_prob { this.field[i][j] = this.Launcher_r }
            if rand.Intn(100-0) < this.Launcher_l_prob { this.field[i][j] = this.Launcher_l }
            if rand.Intn(100-0) < this.Launcher_u_prob { this.field[i][j] = this.Launcher_u }
            if rand.Intn(100-0) < this.Launcher_d_prob { this.field[i][j] = this.Launcher_d }
            if rand.Intn(100-0) < this.Objective_prob { this.field[i][j] = this.Objective; objectives_generated += 1 }
            if i < 2 || i >= this.Field_limit - 2 || j < 2 || j >= this.Field_limit - 2 {
                this.field[i][j] = this.Obstruction
            }
        }
    }
    this.field[this.player_x][this.player_y] = this.Player
    if objectives_generated < this.Win_condition + 2 { this.Initialize() }
}
func (this *Vime) flush() {
    for i := 0; i < 100; i++ { fmt.Println("") }
}
func (this *Vime) ping(input int) {
    for i := 0; i < input; i++ {
        this.flush()
            this.field[this.player_y][this.player_x] = this.Player_alt
            this.status()
            time.Sleep(200000000)
            this.field[this.player_y][this.player_x] = this.Player
            this.status()
            time.Sleep(200000000)
    }
}
func (this *Vime) step_on() {
    switch this.result {
        case this.Objective: this.points += 1
        case this.Penalty: this.points -= 1
        case this.Danger: this.lost = true; this.death = "danger"
        case this.Launcher_r, this.Launcher_l, this.Launcher_u, this.Launcher_d:
            this.auto = true
            this.launch_count += 1
    }
    this.field[this.player_y][this.player_x] = this.last
}
func (this *Vime) step_off() {
    this.last = this.field[this.player_y][this.player_x]
    switch this.last {
        case this.Penalty, this.Objective: this.last = this.Empty
        case this.Platform: this.last = this.Obstruction
    }
    this.field[this.player_y][this.player_x] = this.Player
}
func (this *Vime) right(distance int) {
    this.result = this.field[this.player_y][this.player_x+distance]
    if this.result != this.Obstruction {
        this.step_on()
        this.player_x += distance
        this.step_off()
    }
    this.automove()
}
func (this *Vime) down(distance int) {
    this.result = this.field[this.player_y+distance][this.player_x]
    if this.result != this.Obstruction {
        this.step_on()
        this.player_y += distance
        this.step_off()
    }
    this.automove()
}
func (this *Vime) left(distance int) {
    this.result = this.field[this.player_y][this.player_x-distance]
    if this.result != this.Obstruction {
        this.step_on()
        this.player_x -= distance
        this.step_off()
    }
    this.automove()
}
func (this *Vime) up(distance int) {
    this.result = this.field[this.player_y-distance][this.player_x]
    if this.result != this.Obstruction {
        this.step_on()
        this.player_y -= distance
        this.step_off()
    }
    this.automove()
}
func (this *Vime) automove() {
    if this.launch_count > 1000 {
        this.auto = false
        this.lost = true
        this.death = "launch"
    }
    if this.result == this.Obstruction && this.auto == true {
        this.auto = false
        this.lost = true
        this.death = "obstruction"
    }
    if this.auto == true {
        switch this.result {
            case this.Launcher_r:
                if this.last != this.Launcher_l {
                    this.right(1)
                } else {
                    this.lost = true
                    this.death = "launch"
                }
            case this.Launcher_l:
                if this.last != this.Launcher_r {
                    this.left(1)
                } else {
                    this.lost = true
                    this.death = "launch"
                }
            case this.Launcher_u:
                if this.last != this.Launcher_d {
                    this.up(1)
                } else {
                    this.lost = true
                    this.death = "launch"
                }
            case this.Launcher_d:
                if this.last != this.Launcher_u {
                    this.down(1)
                } else {
                    this.lost = true
                    this.death = "launch"
                }
        }
    this.auto = false
    }
}
func (this *Vime) execute() {
    var final_letter string = string(this.instruction[len(this.instruction)-1])
    switch final_letter {
    case "h": this.left(1)
    case "l": this.right(1)
    case "k": this.up(1)
    case "j": this.down(1)
    case "H": this.left(2); this.instruction = "h"
    case "L": this.right(2); this.instruction = "l"
    case "K": this.up(2); this.instruction = "k"
    case "J": this.down(2); this.instruction = "j"
    case "q": this.lost = true
    case "p": this.ping(3)
    }
    this.launch_count = 0
}
func (this *Vime) status() {
        this.flush()
        for i := 0; i < this.Field_limit; i++ {
            switch i {
                case 0: fmt.Println(this.field[i], " ", "Objective: collect", this.Win_condition)
                case 1: fmt.Println(this.field[i], " ", this.Objective, "is your objective")
                case 2: fmt.Println(this.field[i], " ", this.Obstruction, "blocks the way")
                case 3: fmt.Println(this.field[i], " ", this.Penalty, "will harm you")
                case 4: fmt.Println(this.field[i], " ", this.Danger, "will end you")
                case 5: fmt.Println(this.field[i], " ", this.Platform, "will allow you once")
                case 6: fmt.Println(this.field[i], " ", this.Launcher_r, this.Launcher_l, this.Launcher_d, this.Launcher_u, "will move you")
                case 7: fmt.Println(this.field[i], " ", "Move with h j k l")
                case 8: fmt.Println(this.field[i], " ", "Hop with H J K L")
                case 9: fmt.Println(this.field[i], " ", "Quit with q")
                case 10: fmt.Println(this.field[i], " ", "Ping yourself with with p")
                case 11: fmt.Println(this.field[i], " ", "Execute action with \"Enter\"")
                case 12: fmt.Println(this.field[i], " ", "Points:", this.points)
                default: fmt.Println(this.field[i])
            }
        }
}
func (this *Vime) Run() {
    this.ping(5)
    for {
        if this.points >= this.Win_condition { break }
        if this.lost { break }
        this.status()
        fmt.Scanf("%s",&this.instruction)
        this.execute()
    }
        this.flush()
    if this.points >= this.Win_condition {
        fmt.Println("You Win")
    } else {
        switch this.death {
            case "danger":
                fmt.Println("You were ended.")
                fmt.Scanf("%s",&this.instruction)
            case "obstruction":
                fmt.Println("You were launched up against a wall until you lost conciousness.")
                fmt.Scanf("%s",&this.instruction)
            case "launch":
                fmt.Println("As you endlessly bounce between the launchers, you slowly resign yourself to your strange fate.")
                fmt.Println("You are absolutely sure that there are ways to die that are more stupid and trivial than this, but you cannot seem to think of any.")
                fmt.Println("Oh well, plenty of time for that now.")
                fmt.Scanf("%s",&this.instruction)
            default:
                fmt.Println("Game Over")
                fmt.Scanf("%s",&this.instruction)
        }
    }
}
