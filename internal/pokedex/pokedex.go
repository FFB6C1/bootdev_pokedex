package pokedex

import (
	"fmt"
	"math/rand"
)

type Pokedex struct {
	Pokemon map[string]Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		Pokemon: map[string]Pokemon{},
	}
}

func (p *Pokedex) Check(newPokemon string) bool {
	_, ok := p.Pokemon[newPokemon]
	return ok
}

func (p *Pokedex) Catch(newPokemon Pokemon) bool {
	random := rand.Intn(450)
	if random >= newPokemon.BaseExperience || (newPokemon.Name == "blissey" && random >= newPokemon.BaseExperience-200) {
		p.Pokemon[newPokemon.Name] = newPokemon
		return true
	}
	return false
}

func (p *Pokedex) RemovePokemon(key string) error {
	if _, ok := p.Pokemon[key]; !ok {
		err := fmt.Errorf("%s does not exist in pokedex, cannot remove.", key)
		return err
	}
	delete(p.Pokemon, key)
	return nil
}

func (p Pokedex) Report(key string) {
	pokemon := p.Pokemon[key]
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("	-%s\n", types.Type.Name)
	}
}

func (p Pokedex) AllPokemon() {
	if len(p.Pokemon) == 0 {
		fmt.Println("No captured pokemon!")
		return
	}
	for _, pokemon := range p.Pokemon {
		fmt.Println("	- " + pokemon.Name)
	}
}
