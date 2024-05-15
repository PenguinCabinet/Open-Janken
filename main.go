package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/mattn/go-gtk/gtk"
)

type Game_Status_t int
type Game_Result_t int
type Hand_t int

const (
	Game_wait Game_Status_t = iota
	Game_running
	Game_lock
)
const (
	Game_Result_draw Game_Result_t = iota
	Game_Result_lose
	Game_Result_win
)

var Result_msg = []string{"DRAW", "YOU LOSE", "YOU WIN"}

const (
	Hand_gu Hand_t = iota
	Hand_chyoki
	Hand_pa
)

var (
	Hand_strs = []string{"グー", "チョキ", "パー"}
)

type Hand_with_Label_t struct {
	Label *gtk.Label
	Hand  Hand_t
}

func (self *Hand_with_Label_t) Update(x Hand_t) {
	self.Hand = x
	if self.Label != nil {
		self.Label.SetText(Hand_strs[self.Hand])
	}
}

var Game_Status Game_Status_t = Game_wait
var Game_Result Game_Result_t = Game_Result_draw
var Player Hand_with_Label_t
var Enemy Hand_with_Label_t
var Enemy_AI AI_t = Random_AI{}
var msg_label *gtk.Label

func Judge_Game(player Hand_t, enemy Hand_t) Game_Result_t {
	if player == Hand_gu && enemy == Hand_chyoki {
		return Game_Result_win
	}
	if player == Hand_chyoki && enemy == Hand_pa {
		return Game_Result_win
	}
	if player == Hand_pa && enemy == Hand_gu {
		return Game_Result_win
	}

	if player == Hand_gu && enemy == Hand_pa {
		return Game_Result_lose
	}
	if player == Hand_chyoki && enemy == Hand_gu {
		return Game_Result_lose
	}
	if player == Hand_pa && enemy == Hand_chyoki {
		return Game_Result_lose
	}

	return Game_Result_draw

}

func Game_Start() {
	if Game_Status == Game_wait {
		Game_Status = Game_running
		go func() {

			msg_label.SetText("最初はグー")
			Player.Update(Hand_gu)
			Enemy.Update(Hand_gu)
			time.Sleep(time.Millisecond * 2000)

			msg_label.SetText("じゃんけん")
			time.Sleep(time.Millisecond * 2000)

			msg_label.SetText("ぽん！")
			Game_Status = Game_lock
			Enemy.Update(Enemy_AI.Get(Player.Hand))
			result := Judge_Game(Player.Hand, Enemy.Hand)
			time.Sleep(time.Millisecond * 3000)

			msg_label.SetText(Result_msg[result])

			Game_Status = Game_wait
		}()
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Open Janken GTK")
	window.SetSizeRequest(800, 600)
	window.Connect("destroy", gtk.MainQuit)

	vbox1 := gtk.NewVBox(true, 1)

	Enemy_img := gtk.NewLabel("")
	Enemy.Label = Enemy_img
	Enemy.Update(Hand_gu)

	My_img := gtk.NewLabel("")
	Player.Label = My_img
	Player.Update(Hand_gu)

	hbox_buttons := gtk.NewHBox(true, 1)

	var Hand_buttons []*gtk.Button
	for hand, v := range Hand_strs {
		button := gtk.NewButton()
		button.Add(gtk.NewLabel(v))
		temp := hand
		button.Clicked(func() {
			if Game_Status == Game_lock {
				return
			}
			Player.Update(Hand_t(temp))
		})
		Hand_buttons = append(Hand_buttons, button)
	}
	for _, v := range Hand_buttons {
		hbox_buttons.Add(v)
	}

	msg_label = gtk.NewLabel("")

	hbox_Game_buttons := gtk.NewHBox(true, 1)
	Game_Start_button := gtk.NewButton()
	Game_Start_button.Add(gtk.NewLabel("スタート"))
	Game_Start_button.Clicked(Game_Start)
	hbox_Game_buttons.Add(Game_Start_button)

	vbox1.Add(Enemy_img)
	vbox1.Add(My_img)
	vbox1.Add(msg_label)
	vbox1.Add(hbox_buttons)
	vbox1.Add(hbox_Game_buttons)

	window.Add(vbox1)

	window.ShowAll()
	gtk.Main()

}
