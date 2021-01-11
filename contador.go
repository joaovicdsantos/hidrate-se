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

func (c *Contador) Calcular(peso float64, copo float64) string {
	quantidadeDeAgua := float64(peso * 35)
	c.tempo = int(math.RoundToEven((quantidadeDeAgua / copo)))

	avisar := func(tempo int) {
		for tempo > 0 {
			err := beeep.Notify("Hidrate-se", "Hora de beber água!", "garrafinha.png")
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Hour)
			tempo -= 1
		}
	}

	fmt.Println("Iniciando o contador!")

	go avisar(c.tempo)

	return fmt.Sprintf("Pelos dados que você nos passou você precisa beber %v ml de água. Sendo assim, iremos te notificar a cada 1 hora durante %v hora(s)", quantidadeDeAgua, c.tempo)
}
