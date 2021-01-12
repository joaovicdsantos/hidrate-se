package main

import (
	"fmt"
	"math"
	"time"

	"github.com/gen2brain/beeep"

	"github.com/wailsapp/wails"
)

// Lint bugou

// Contador é o cara responsável por iniciar o time
type Contador struct {
	tempo   int
	runtime *wails.Runtime
}

// WailsInit - função sempre executada no início de uma conexão
func (c *Contador) WailsInit(runtime *wails.Runtime) error {
	c.runtime = runtime
	return nil
}

// CalcularTempo - onde o tempo é calculado e incia-se o time
func (c *Contador) CalcularTempo(peso float64, copo float64) string {
	quantidadeDeAgua := float64(peso * 35)
	c.tempo = int(math.RoundToEven((quantidadeDeAgua / copo)))

	avisar := func(tempo int) {
		for tempo > 0 {
			err := beeep.Notify("Hidrate-se", "Hora de beber água!", "")
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Hour)
			tempo--
		}
	}

	fmt.Println("Iniciando o contador!")

	go avisar(c.tempo)

	return fmt.Sprintf("Pelos dados que você nos passou você precisa beber %v ml de água. Sendo assim, iremos te notificar a cada 1 hora durante %v hora(s)", quantidadeDeAgua, c.tempo)
}
