package vime
import "fmt"
import "math/rand"
import "time"
import "strings"
import "strconv"
import "github.com/nsf/termbox-go"

type Vime struct {
    points int
    player_x int
    player_y int
    field [][]string
    text []string
    text_default []string
    result string
    last string
    lost bool
    auto bool
    launch_count int
    death string

    Field_limit int
    Win_condition int
    Text []string

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

    // Set these to -1 to disable them.
    // A 0 value gives them default values during initialization.
    Obstruction_prob int
    Objective_prob int
    Danger_prob int
    Platform_prob int
    Penalty_prob int
    Launcher_r_prob int
    Launcher_l_prob int
    Launcher_u_prob int
    Launcher_d_prob int

    Key_r string
    Key_l string
    Key_u string
    Key_d string
    Key_R string
    Key_L string
    Key_U string
    Key_D string
    Key_ping string
    Key_quit string
}
func (this *Vime) Initialize() {
    if this.Field_limit == 0 { this.Field_limit = 31 }
    if this.player_x == 0 { this.player_x = this.Field_limit / 2 }
    if this.player_y == 0 { this.player_y = this.Field_limit / 2 }
    if this.Win_condition == 0 { this.Win_condition = 20 }

    if this.Empty == "" { this.Empty = " " }
    if this.Player == "" { this.Player = "⍟" }
    if this.Danger == "" { this.Danger = "⚠" }
    if this.Penalty == "" { this.Penalty = "-" }
    if this.Platform == "" { this.Platform = "⛀" }
    if this.Objective == "" { this.Objective = "+" }
    if this.Launcher_r == "" { this.Launcher_r = "→" }
    if this.Launcher_l == "" { this.Launcher_l = "←" }
    if this.Launcher_u == "" { this.Launcher_u = "↑" }
    if this.Launcher_d == "" { this.Launcher_d = "↓" }
    if this.Player_alt == "" { this.Player_alt = "✪" }
    if this.Obstruction == "" { this.Obstruction = "𝌆" }

    if this.Danger_prob == 0 { this.Danger_prob = 10 }
    if this.Penalty_prob == 0 { this.Penalty_prob = 20 }
    if this.Platform_prob == 0 { this.Platform_prob = 5 }
    if this.Objective_prob == 0 { this.Objective_prob = 4 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Launcher_l_prob == 0 { this.Launcher_l_prob = 2 }
    if this.Obstruction_prob == 0 { this.Obstruction_prob = 15 }

    if this.Key_r == "" { this.Key_r = "ldfu" }
    if this.Key_l == "" { this.Key_l = "habo" }
    if this.Key_u == "" { this.Key_u = "kwp." }
    if this.Key_d == "" { this.Key_d = "jsne" }
    if this.Key_R == "" { this.Key_R = "LDFU" }
    if this.Key_L == "" { this.Key_L = "HABO" }
    if this.Key_U == "" { this.Key_U = "KWP>" }
    if this.Key_D == "" { this.Key_D = "JSNE" }
    if this.Key_ping == "" { this.Key_ping = "zZ" }
    if this.Key_quit == "" { this.Key_quit = "qQ" }

    this.lost = false
    this.auto = false
    this.last = this.Empty

    this.field = make([][]string,this.Field_limit)
    for i := 0; i < this.Field_limit; i++ { this.field[i] = make([]string,this.Field_limit) }
    this.populate()

    this.Text = make([]string,this.Field_limit)
    this.text = make([]string,this.Field_limit)
    this.text_default = make([]string,this.Field_limit)
    this.text_default[0] = "Objective: collect %Win_condition%"
    this.text_default[1] = "%Objective% is your objective"
    this.text_default[2] = "%Obstruction% obstructs you"
    this.text_default[3] = "%Penalty% is counterproductive"
    this.text_default[4] = "%Danger% will end you"
    this.text_default[5] = "%Platform% will allow you once"
    this.text_default[6] = "%Launcher_r%%Launcher_l%%Launcher_u%%Launcher_d% will move you"
    this.text_default[7] = "Quit with q (by default)"
    this.text_default[8] = "Ping yourself with z (by default)"
    this.text_default[9] = "Execute action with \"Enter\""
    this.text_default[10] = "Points: %Points%"
    this.text_default[11] = "Last Step: %result%"

}
func (this *Vime) populate() {
    rand.Seed(time.Now().UTC().UnixNano())
    var objectives_generated int = 0
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
                if this.field[i][j] == this.Objective { objectives_generated -= 1 }
                this.field[i][j] = this.Obstruction
            }
        }
    }
    this.field[this.player_x][this.player_y] = this.Player
    if objectives_generated < this.Win_condition + 2 { this.populate() }
}
func (this *Vime) flush() {
    fmt.Println(strings.Repeat("\n", 100))
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
func (this *Vime) execute(instruction string) {
    if strings.Contains(this.Key_r, instruction) { this.right(1) }
    if strings.Contains(this.Key_l, instruction) { this.left(1) }
    if strings.Contains(this.Key_u, instruction) { this.up(1) }
    if strings.Contains(this.Key_d, instruction) { this.down(1) }
    if strings.Contains(this.Key_R, instruction) { this.right(2) }
    if strings.Contains(this.Key_L, instruction) { this.left(2) }
    if strings.Contains(this.Key_U, instruction) { this.up(2) }
    if strings.Contains(this.Key_D, instruction) { this.down(2) }
    if strings.Contains(this.Key_ping, instruction) { this.ping(3) }
    if strings.Contains(this.Key_quit, instruction) { this.lost = true }
    this.launch_count = 0
}
func (this *Vime) status() {
    this.flush()
    output := ""
    for i := 0; i < this.Field_limit; i++ {
        if this.Text[i] != "" {
            this.text[i] = this.Text[i]
        } else {
            this.text[i] = this.text_default[i]
        }
        this.text[i] = strings.Replace(this.text[i], "%Points%", strconv.FormatInt(int64(this.points),10), -1)
        this.text[i] = strings.Replace(this.text[i], "%player_x%", strconv.FormatInt(int64(this.player_x),10), -1)
        this.text[i] = strings.Replace(this.text[i], "%player_y%", strconv.FormatInt(int64(this.player_y),10), -1)
        this.text[i] = strings.Replace(this.text[i], "%Win_condition%", strconv.FormatInt(int64(this.Win_condition),10), -1)
        this.text[i] = strings.Replace(this.text[i], "%Field_limit%", strconv.FormatInt(int64(this.Field_limit),10), -1)
        this.text[i] = strings.Replace(this.text[i], "%result%", this.result, -1)
        this.text[i] = strings.Replace(this.text[i], "%Objective%", this.Objective, -1)
        this.text[i] = strings.Replace(this.text[i], "%Obstruction%", this.Obstruction, -1)
        this.text[i] = strings.Replace(this.text[i], "%Penalty%", this.Penalty, -1)
        this.text[i] = strings.Replace(this.text[i], "%Danger%", this.Danger, -1)
        this.text[i] = strings.Replace(this.text[i], "%Platform%", this.Platform, -1)
        this.text[i] = strings.Replace(this.text[i], "%Launcher_r%", this.Launcher_r, -1)
        this.text[i] = strings.Replace(this.text[i], "%Launcher_l%", this.Launcher_l, -1)
        this.text[i] = strings.Replace(this.text[i], "%Launcher_u%", this.Launcher_u, -1)
        this.text[i] = strings.Replace(this.text[i], "%Launcher_d%", this.Launcher_d, -1)
        output += fmt.Sprintf("%v %v\n", this.field[i], this.text[i])
    }
    fmt.Println(output)
}
func (this *Vime) Run() {
    termbox.Init()
    defer termbox.Close()

    // Play game
    this.ping(5)
    for {
        if this.points >= this.Win_condition { break }
        if this.lost { break }
        this.status()

        event := termbox.PollEvent()
        this.execute(fmt.Sprintf("%c", event.Ch))
    }

    // Handle end of game
    this.flush()
    if this.points >= this.Win_condition {
        fmt.Println("You Win")
    } else {
        switch this.death {
        case "danger":
            fmt.Println("You were ended.")

        case "obstruction":
            fmt.Println("You were launched up against a wall until you lost conciousness.")

        case "launch":
            fmt.Println("As you endlessly bounce between the launchers, you slowly resign yourself to your strange fate.")
            fmt.Println("You are absolutely sure that there are ways to die that are more stupid and trivial than this, but you cannot seem to think of any.")
            fmt.Println("Oh well, plenty of time for that now.")
            fmt.Println("Game Over")
        }
    }
    termbox.PollEvent()  // Pause until next keypress
}
