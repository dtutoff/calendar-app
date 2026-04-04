package main

import "testing"

// func TestAdd(t *testing.T) {
// 	result := Add(1, 2)
// 	if result != 3 {
// 		t.Errorf("Expected 3, got %d", result)
// 	}
// }

func TestCheckPositive(t *testing.T) {
	num := 5
	err := CheckPositive(num)
	if err != nil {
		t.Errorf("Не ожидалось ошибки для положительного числа, получен %v", err)
	}

	num = -3
	err = CheckPositive(num)
	if err == nil {
		t.Errorf("Ожидал ошибку для отрицательного числа, не получил ни одной")
	}

	num = 0
	err = CheckPositive(0)
	if err == nil {
		t.Errorf("Ожидал ошибку, равную нулю, но не получил её")
	}
}
func BenchmarkAdd(b *testing.B) { // передаем testing.B
	for i := 0; i < b.N; i++ { // b.N — достаточное количество итераций (выбирается компилятором)
		_ = Add(123, 456) // вызов тестируемой функции
	}
}
