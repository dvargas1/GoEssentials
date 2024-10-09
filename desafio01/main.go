package main

import ("fmt"
		"time")

// Dado um array inteiro nums, retorne true se qualquer valor aparecer 
// pelo menos duas vezes no array e retorne false se cada elemento for distinto.

func contemDuplicado(nums []int) bool {
	n := len(nums)

	if n < 2 {
		return false
	}

	seen := make([]*int, n)

	for i := 0; i < n; i++ {
		num:=nums[i]

		for j := 0; j < n; j++ {
			if seen[j] != nil && *seen[j] == num{
				return true;
			}
		}

		for j:= 0; j < n; j++ {
			if seen[j] == nil {
				seen[j] = new(int)
				*seen[j] = num
				break
			}
		}
	}
	return false
}

func contemDuplicadoGPT(nums []int) bool {
	seen := make(map[int]struct{}) // Usamos um mapa para armazenar os números já vistos

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true // Se já existe, encontramos um duplicado
		}
		// Caso contrário, adiciona o número ao mapa
		seen[num] = struct{}{}
	}

	return false
}


func main() {

	nums := []int{1, 2, 3, 4, 5, 1}
	numsWithoutDup := []int{1, 2, 3, 4, 5}

	numsLarge := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		numsLarge[i] = i % 100 
	}

	numsLargeWithoutDup := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		numsLargeWithoutDup[i] = i
	}

	start := time.Now()
	contemDuplicado(nums)
	duration := time.Since(start) 
	fmt.Println(" -- Teste array pequena com duplicata ")
	fmt.Printf("Tempo de execução com array pequena Daniel : %v\n", duration)

	start = time.Now()
	contemDuplicadoGPT(nums)
	duration = time.Since(start) 
	fmt.Printf("Tempo de execução com array pequena GPT : %v\n", duration)


	start = time.Now()
	contemDuplicado(numsLarge)
	duration = time.Since(start) 
	fmt.Println("\n -- Teste array grande com duplicata ")
	fmt.Printf("Tempo de execução com array grande Daniel: %v\n", duration)

	start = time.Now()
	contemDuplicadoGPT(numsLarge)
	duration = time.Since(start) 
	fmt.Printf("Tempo de execução com array grande GPT: %v\n", duration)

	start = time.Now()
	contemDuplicado(numsWithoutDup)
	duration = time.Since(start) 
	fmt.Println("\n -- Teste array pequena sem duplicata ")
	fmt.Printf("Tempo de execução Daniel: %v\n", duration)

	start = time.Now()
	contemDuplicadoGPT(numsWithoutDup)
	duration = time.Since(start) 
	fmt.Printf("Tempo de execução GPT: %v\n", duration)


	start = time.Now()
	contemDuplicado(numsLarge)
	duration = time.Since(start) 
	fmt.Println("\n -- Teste array grande sem duplicata ")
	fmt.Printf("Tempo de execução Daniel: %v\n", duration)

	start = time.Now()
	contemDuplicadoGPT(numsLarge)
	duration = time.Since(start) 
	fmt.Printf("Tempo de execução GPT: %v\n", duration)
}

