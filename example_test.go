package btcutil_test

import (
	"fmt"
	"math"

	"github.com/jadeblaquiere/cttutil"
)

func ExampleAmount() {

	a := btcutil.Amount(0)
	fmt.Println("Zero Mystiko:", a)

	a = btcutil.Amount(1e8)
	fmt.Println("100,000,000 Mystikos:", a)

	a = btcutil.Amount(1e5)
	fmt.Println("100,000 Mystikos:", a)
	// Output:
	// Zero Mystiko: 0 CTT
	// 100,000,000 Mystikos: 1 CTT
	// 100,000 Mystikos: 0.001 CTT
}

func ExampleNewAmount() {
	amountOne, err := btcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := btcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := btcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := btcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 CTT
	// 0.01234567 CTT
	// 0 CTT
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := btcutil.Amount(44433322211100)

	fmt.Println("Mystiko to kCTT:", amount.Format(btcutil.AmountKiloCTT))
	fmt.Println("Mystiko to CTT:", amount)
	fmt.Println("Mystiko to MilliCTT:", amount.Format(btcutil.AmountMilliCTT))
	fmt.Println("Mystiko to MicroCTT:", amount.Format(btcutil.AmountMicroCTT))
	fmt.Println("Mystiko to Mystiko:", amount.Format(btcutil.AmountMystiko))

	// Output:
	// Mystiko to kCTT: 444.333222111 kCTT
	// Mystiko to CTT: 444333.222111 CTT
	// Mystiko to MilliCTT: 444333222.111 mCTT
	// Mystiko to MicroCTT: 444333222111 μCTT
	// Mystiko to Mystiko: 44433322211100 Mystiko
}
