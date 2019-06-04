package vending

import "testing"

func TestInsertCoinWith_T_F_TW_O_MachineShouldHave_18_Bath(t *testing.T) {
	v := NewVendingMachine()
	expectedResult := 18

	v.insertCoins("T")
	v.insertCoins("F")
	v.insertCoins("TW")
	actualResult := v.insertCoins("O")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestInsertCoinWith_T_MachineShouldHave_10_Bath(t *testing.T) {

	v := NewVendingMachine()
	expectedResult := 10

	actualResult := v.insertCoins("T")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestInsertCoinWith_F_MachineShouldHave_5_Bath(t *testing.T) {
	v := NewVendingMachine()
	expectedResult := 5

	actualResult := v.insertCoins("F")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestInsertCoinWith_T_F_MachineShouldHave_15_Bath(t *testing.T) {
	v := NewVendingMachine()
	expectedResult := 15

	v.insertCoins("T")
	actualResult := v.insertCoins("F")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}