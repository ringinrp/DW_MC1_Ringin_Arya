package main

import (
	"fmt"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {

	if name == "Sasuke" {
		return Profile{
			Name:   name,
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	}
	if name == "Goku" {
		return Profile{
			Name:   name,
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	}
	if name == "Naruto" {
		return Profile{
			Name:   name,
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	}

	return Profile{}

}

func PowerUp(profile Profile, multiplier int) Profile {
	return Profile{

		//nilai = nilai+(nilai*multiplier)

		Name:   profile.Name,
		Health: profile.Health + (profile.Health * multiplier),
		Power:  profile.Power + (profile.Power * multiplier),
		Exp:    profile.Exp + (profile.Exp * multiplier),
	}
}

func main() {

	fmt.Println("==GOKU==")
	var Goku = MakeProfile("Goku")
	fmt.Println("Name :", Goku.Name)
	fmt.Println("Health :", Goku.Health)
	fmt.Println("Power :", Goku.Power)
	fmt.Println("Exp :", Goku.Exp)
	fmt.Println("===HEROES POWER UP===")
	var GokuPowerUp = PowerUp(Goku, 2)
	fmt.Println("Name :", GokuPowerUp.Name)
	fmt.Println("Health :", GokuPowerUp.Health)
	fmt.Println("Power :", GokuPowerUp.Power)
	fmt.Println("Exp :", GokuPowerUp.Exp)

	fmt.Println("==SASUKE==")
	var Sasuke = MakeProfile("Sasuke")
	fmt.Println("Name :", Sasuke.Name)
	fmt.Println("Health :", Sasuke.Health)
	fmt.Println("Power :", Sasuke.Power)
	fmt.Println("Exp :", Sasuke.Exp)
	fmt.Println("===HEROES POWER UP===")
	var SasukePowerUp = PowerUp(Sasuke, 3)
	fmt.Println("Name :", SasukePowerUp.Name)
	fmt.Println("Health :", SasukePowerUp.Health)
	fmt.Println("Power :", SasukePowerUp.Power)
	fmt.Println("Exp :", SasukePowerUp.Exp)

	fmt.Println("==NARUTO==")
	var Naruto = MakeProfile("Naruto")
	fmt.Println("Name :", Naruto.Name)
	fmt.Println("Health :", Naruto.Health)
	fmt.Println("Power :", Naruto.Power)
	fmt.Println("Exp :", Naruto.Exp)
	fmt.Println("===HEROES POWER UP===")
	var NarutoPowerUp = PowerUp(Sasuke, 3)
	fmt.Println("Name :", NarutoPowerUp.Name)
	fmt.Println("Health :", NarutoPowerUp.Health)
	fmt.Println("Power :", NarutoPowerUp.Power)
	fmt.Println("Exp :", NarutoPowerUp.Exp)
}
