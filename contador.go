package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"math"
	"time"

	"github.com/wailsapp/wails"
)

type Contador struct {
	tempo   int
	runtime *wails.Runtime
}

func (c *Contador) WailsInit(runtime *wails.Runtime) error {
	c.runtime = runtime
	return nil
}

func (c *Contador) Calcular(peso float64, copo float64) {
	c.tempo = int(math.RoundToEven((float64(peso*35) / copo)))

	avisar := func(tempo int) {
		for tempo > 0 {
			err := beeep.Notify("Hidrate-se", "Hora de beber água!", "assets/information.png")
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Hour)
			tempo -= 1
		}
	}

	fmt.Println("Iniciando o contador!")

	go avisar(c.tempo)

}
