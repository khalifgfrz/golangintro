package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printSegitiga(n int) {
	if n < 0 {
		fmt.Println("Nilai n harus positif")
	}

	for i := n; i >= 1; i-- {
		for j := 0 ; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func genPass(password string, level string) string {
	// fmt.Println("Password before ensuring minimum length:", password)
	if len(password) < 6 {
		fmt.Println("Password harus lebih dari 6 karakter")
		return password
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var extraChars string
	switch level {
	case "low":
		extraChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "med":
		extraChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case "strong":
		extraChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	default:
		return ""
	}

	passwordRunes := []rune(password)
	for i := range passwordRunes {
		j := r.Intn(i + 1)
		passwordRunes[i], passwordRunes[j] = passwordRunes[j], passwordRunes[i]
	}
	// shuffled := string(passwordRunes)
	// fmt.Println("Password after first shuffle:", shuffled)

	for len(passwordRunes) < 10 {
		passwordRunes = append(passwordRunes, rune(extraChars[r.Intn(len(extraChars))]))
	}
	// shuffled = string(passwordRunes)
	// fmt.Println("Password after adding extra characters:", shuffled)

	for i := range passwordRunes {
		j := r.Intn(i + 1)
		passwordRunes[i], passwordRunes[j] = passwordRunes[j], passwordRunes[i]
	}
	finalPassword := string(passwordRunes)
	// fmt.Println("Final password:", finalPassword)

	return finalPassword
}


func findMovies(films []int, flights int) (int, int, bool) {
    for i := range films {
        for j := i + 1; j < len(films); j++ {
            if films[i] + films[j] == flights {
                return films[i], films[j], true
            }
        }
    }
    return 0, 0, false
}

func main() {
	printSegitiga(5)

	fmt.Println(genPass("abcdef", "strong"))

	data := []int{1, 5, 3, 3, 8, 9}
    n := 6
    movie1, movie2, found := findMovies(data, n)
    if found {
        fmt.Printf("Dua durasi film yang cocok adalah %d dan %d", movie1, movie2)
    } else {
        fmt.Println("Tidak ada dua film yang durasinya menambah menjadi", n)
    }
}
