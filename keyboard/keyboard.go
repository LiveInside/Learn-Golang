//Package keyboard Работа с данными вводимыми пользователем
package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadS создаёт буфер, записывает туда данные введёные пользователем с клавиатуры
// Убирает ненужные знаки, возвращает строку
func ReadS(line string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(line)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}
