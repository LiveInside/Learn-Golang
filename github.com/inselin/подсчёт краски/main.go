// подсчёт краски на стену
package main

import (
	"fmt"
	"log"
	"strconv"
	"keyboard"
)

func main() {
		defer fmt.Println("Прощайте.")
		for flagChoice := false; flagChoice == false;{
			choice()
			flagC := keyboard.ReadS("Выйти из программы?(Да || Нет): ")
			if flagC == "Да" || flagC == "1" || flagC == "да"{
				flagChoice = true
			}
		}
}

func choice() {
	input := keyboard.ReadS("Выберите программу или справочник(цифра - 0):")
	choice, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	for flag := false; flag == false;{
		if choice == 1 {
			err := paintNeeded()
			if err != nil {
				fmt.Println(err)
			}
			flag = true
		} else if choice == 2 {
			fmt.Println("СЛАВА РОССИИ!")
			flag = true
		} else if choice == 0 {
			fmt.Println("1-Подсчёт краски для стен. \n2-Секретное сообщение. \n\"-1\"-Выйти из программы")
			flag = true
		} else if choice == -1 {
			break
		} else {
			panic("Извините, пока такой программы нет.")
		}
	}
}

func paintNeeded() error{
	input := keyboard.ReadS("Width:")
	width, err := strconv.ParseFloat(input, 32)
	if err != nil {
		log.Fatal(err)
	}
	if width < 0 {
		return fmt.Errorf("a width of %0.2f, is invalid", width)
	}

	input = keyboard.ReadS("Height:")
	height, err := strconv.ParseFloat(input, 32)
	if err != nil {
		log.Fatal(err)
	}
	if height < 0 {
		return fmt.Errorf("a height of %0.2f, is invalid", height)
	}

	fmt.Println("Width:", width, "Height:", height)
	area := width * height
 	fmt.Printf("%.2f liters nedeed\n", area/10.0)
	return nil
}
