package main

import (
    "bufio"
	"log"
	"os"
	"fmt"
    "regexp"
	"github.com/ernestosuarez/itertools"
    "sync"
)


// Функция генерирующая Product - все возможные сочетания
func NextIndex(ix []int, lens int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens {
			return
		}
		ix[j] = 0
	}
}

// на списки такого размера поделим результирующий список
const batchSize = 1000


func process(data [][]int, m map[string]bool) {
    // Мютекс нужен для работы с общим мапом.
    // делим данные на батчи и скармливаем из функции процесс
	mu := new(sync.Mutex)
	for start, end := 0, 0; start <= len(data)-1; start = end {
		end = start + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[start:end]
        fmt.Println(end)
		processBatch(batch, m, mu)
	}
	fmt.Println("done processing all data")
}

func renes2string(runes []int) string {
    // Не совсем руны в строку. Массив интов переводим в строку
	str := ""
	for _, value := range runes {
		str += string(value)
	}
	return str
}


func shortArr(short_vowels []int, m map[string]bool, mu *sync.Mutex, wg *sync.WaitGroup) {
    // Функция мапит короткий набор гласных на итоговое слово
	// Например short_vowels :=[]int{111, 101, 97} пусть будет aie фуенкция составит все возможные сочетания с мними
    // и отматчит только те что подходят под чередование
	defer wg.Done()
	pattern := "^[ aeiouy]?h[ aeiouy]?p[ aeiouy]?p[ aeiouy]?n[ aeiouy]?$"
    re, _ := regexp.Compile(pattern)
	word := []int{104, 112, 112, 110} // Итоговое слово hppn - захардкожено
	word = append(word, short_vowels...)
	for v := range itertools.PermutationsInt(word, len(word)) {
		st := re.FindAllString(renes2string(v), -1)
        if st != nil {
			mapper := renes2string(v)	
            mu.Lock()
			if !m[mapper] {
				m[mapper] = true
			}
			mu.Unlock()
		}
	}
}

func fullArr(vowels []int, m map[string]bool, mu *sync.Mutex, wg *sync.WaitGroup) {
    // Функция чередует элементы массив vowels и word через одного. Работает только для vowels фиксированного размера
	// join 2 arrays of len 4 and 5
	// vowels := []int{97, 32, 111, 105, 101}
    defer wg.Done()
	word := []int{104, 112, 112, 110} // hppn захардкожено
    var res []int
	for idx, _ := range vowels {
        if idx < len(word) {
			res = append(res,  vowels[idx])
            res = append(res, word[idx])
		} else {
			res = append(res, vowels[idx])
		}        
	}
	mapper := renes2string(res)
	mu.Lock()
	if !m[mapper] {
		m[mapper] = true
	}
	mu.Unlock()
}


func processBatch(batch [][]int, m map[string]bool, mu *sync.Mutex) {
	wg := new(sync.WaitGroup)
	for _, sequence := range batch {
        wg.Add(1)
		if len(sequence) == 5 {
			go fullArr(sequence, m, mu, wg)
		} else {
			go shortArr(sequence, m, mu, wg)
		}
	}
    wg.Wait()
}


func main() {
	resultMap := make(map[string]bool)  // результирующая карта
	c := "aeioui " // буквы на замену - все гласные
    fmt.Println([]rune("hppn"))
	var d []int
    for _, value := range []rune(c) {       // переводим строку в слайс интов
        d = append(d, int(value))
    }
	res := make([][]int, 0) // в этот массив слайсов собираем все сочетания букв на замену
    for k:=1; k<6;k++ {     // словов hppn можно заменить максимум *h*p*p*n* те. 5 звездочками - заменами, поэтому меньше 6
    	lens := len(d)
	    for ix := make([]int, k); ix[0] < lens; NextIndex(ix, lens) {
            r := make([]int, k)
		    for i, j := range ix {
			    r[i] = d[j]
    		}
	    	res = append(res, r)
    	}
	}
	fmt.Println(len(res)) // длина списка списков всех сочетаний замен
	process(res, resultMap)
	fmt.Println(len(resultMap))

	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)

	for k := range resultMap {         // Loop
		fmt.Println(k)
		datawriter.WriteString(k + "\n")
	}
	datawriter.Flush()
	file.Close()

}