package main

import (
	"fmt"
	"strconv"
)

func FibonacciIterative(n int) int {
	// TODO: імплементуйте ітераційно.
	// Підказка: тримайте два останні значення й оновлюйте їх у циклі.
	// Вхід вважаємо: n >= 0.
	// При отриманні негативного n повертаємо його без змін
	if n <= 0 {
		return n
	}

	previous, current := 0, 1

	for i := 2; i <= n; i++ {
		next := current + previous
		previous = current
		current = next
	}

	return current
}

func FibonacciRecursive(n int) int {
	// TODO: імплементуйте рекурсивно.
	// База: n==0 -> 0; n==1 -> 1.
	// Рекурсія: F(n-1)+F(n-2)
	// При отриманні негативного n повертаємо його без змін

	if n <= 1 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	// TODO: імплементуйте перевірку на просте число.
	// Підказка: n<=1 -> false; 2 -> true; парні >2 -> false;
	// Далі перевіряйте дільники до sqrt(n).

	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsBinaryPalindrome(n int) bool {
	// TODO: імплементуйте перевірку числа на те що це паліндром.
	// Підказка: перетворіть n у строку (strconv ефективніший за fmt)
	// потім перевірте паліндромність.
	if n < 0 {
		return false
	}

	bin := strconv.FormatInt(int64(n), 2)

	left, right := 0, len(bin)-1

	for left < right {
		if bin[left] != bin[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func ValidParentheses(s string) bool {
	// TODO: імплементуйте перевірку дужок.
	// Правила:
	// 1. Допустимі дужки (, [, {, ), ], }
	// 2. У кожної відкритої дужки є відповідна закриваюча дужка того ж типу
	// 3. Закриваючі дужки стоять у правильному порядку
	// "[{}]" - правильно
	// "[{]}" - не правильно
	// 4. Кожна закриваюча дужка має відповідну відкриваючу дужку
	// Підказка: використовуйте стек (можна зробити через масив рун []rune)
	stack := make([]rune, 0, len(s))

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	opens := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}

	for _, ch := range s {
		if opens[ch] {
			stack = append(stack, ch)
			continue
		}

		if open, ok := pairs[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}

			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	return len(stack) == 0
}

func Increment(num string) int {
	// TODO: Імплементуйте функцію яка на вхід отримує строку яка складається лише з символів 0 та 1
	// Тобто строка містить певне число у бінарному вигляді
	// Потрібно повернути число на один більше
	// Додайте валідацію вхідної строки, якщо вона містить недопустимі символи, повертайте 0
	for _, ch := range num {
		if ch != '0' && ch != '1' {
			return 0
		}
	}

	if len(num) == 0 {
		return 0
	}

	value, err := strconv.ParseInt(num, 2, 64)

	if err != nil {
		return 0
	}

	return int(value + 1)
}

func main() {
	fmt.Println(FibonacciIterative(0), "expected:", 0)     // сценарій 1: база F(0)
	fmt.Println(FibonacciIterative(1), "expected:", 1)     // сценарій 2: база F(1)
	fmt.Println(FibonacciIterative(5), "expected:", 5)     // сценарій 3: середнє число
	fmt.Println(FibonacciIterative(10), "expected:", 55)   // сценарій 4: більше число
	fmt.Println(FibonacciIterative(-3), "expected:", -3)   // сценарій 5: негативні значення
	fmt.Println(FibonacciIterative(20), "expected:", 6765) // сценарій 6: велике число

	fmt.Println(FibonacciRecursive(0), "expected:", 0)    // сценарій 1: база F(0)
	fmt.Println(FibonacciRecursive(1), "expected:", 1)    // сценарій 2: база F(1)
	fmt.Println(FibonacciRecursive(5), "expected:", 5)    // сценарій 3: середнє число
	fmt.Println(FibonacciRecursive(8), "expected:", 21)   // сценарій 4: інше валідне число
	fmt.Println(FibonacciRecursive(-4), "expected:", -4)  // сценарій 5: негативне число
	fmt.Println(FibonacciRecursive(12), "expected:", 144) // сценарій 6: більше число

	fmt.Println(IsPrime(-5), "expected:", false) // сценарій 1: негативне число
	fmt.Println(IsPrime(0), "expected:", false)  // сценарій 2: 0 — не просте
	fmt.Println(IsPrime(1), "expected:", false)  // сценарій 3: 1 — не просте
	fmt.Println(IsPrime(2), "expected:", true)   // сценарій 4: найменше просте число
	fmt.Println(IsPrime(17), "expected:", true)  // сценарій 5: просте число
	fmt.Println(IsPrime(20), "expected:", false) // сценарій 6: складене число

	fmt.Println(IsBinaryPalindrome(-1), "expected:", false) // сценарій 1: негативне число
	fmt.Println(IsBinaryPalindrome(0), "expected:", true)   // сценарій 2: 0 у бінарі — "0", паліндром
	fmt.Println(IsBinaryPalindrome(1), "expected:", true)   // сценарій 3: 1 у бінарі — "1", паліндром
	fmt.Println(IsBinaryPalindrome(3), "expected:", true)   // сценарій 4: 3 → "11", паліндром
	fmt.Println(IsBinaryPalindrome(9), "expected:", true)   // сценарій 5: 9 → "1001", паліндром
	fmt.Println(IsBinaryPalindrome(6), "expected:", false)  // сценарій 6: 6 → "110", не паліндром

	fmt.Println(ValidParentheses(""), "expected:", true)       // сценарій 1: порожній рядок — валідний
	fmt.Println(ValidParentheses("()"), "expected:", true)     // сценарій 2: одна пара дужок
	fmt.Println(ValidParentheses("()[]{}"), "expected:", true) // сценарій 3: кілька правильних дужок
	fmt.Println(ValidParentheses("[{}]"), "expected:", true)   // сценарій 4: вкладені дужки
	fmt.Println(ValidParentheses("[{]}"), "expected:", false)  // сценарій 5: неправильне вкладення
	fmt.Println(ValidParentheses("(]"), "expected:", false)    // сценарій 6: неправильна відповідність

	fmt.Println(Increment("0"), "expected:", 1)     // сценарій 1: найменше бінарне число
	fmt.Println(Increment("1"), "expected:", 2)     // сценарій 2: одиниця → 2
	fmt.Println(Increment("10"), "expected:", 3)    // сценарій 3: 2 у бінарі → 3
	fmt.Println(Increment("111"), "expected:", 8)   // сценарій 4: 7 → 8
	fmt.Println(Increment("1010"), "expected:", 11) // сценарій 5: 10 → 11
	fmt.Println(Increment("10201"), "expected:", 0) // сценарій 6: невалідна строка, містить символ 2
}
