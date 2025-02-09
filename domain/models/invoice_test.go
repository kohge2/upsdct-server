package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcCommission(t *testing.T) {
	commissionRate := 0.04
	t.Run("正常系", func(t *testing.T) {
		invoice := Invoice{
			PaidAmount:     10000,
			CommissionRate: &commissionRate,
		}

		commission, err := invoice.CalcCommission()

		assert.Equal(t, 400, commission)
		assert.NoError(t, err)
	})

	t.Run("正常系_手数料率をかけて小数になる場合", func(t *testing.T) {
		invoice := Invoice{
			PaidAmount:     1001,
			CommissionRate: &commissionRate,
		}

		commission, err := invoice.CalcCommission()

		assert.Equal(t, 40, commission)
		assert.NoError(t, err)
	})
}

func TestCalcBilledAmount(t *testing.T) {
	taxRate := 0.1
	t.Run("正常系", func(t *testing.T) {
		commision := 400

		invoice := Invoice{
			PaidAmount: 10000,
			Commission: &commision,
		}

		billedAmount, err := invoice.CalcBilledAmount(taxRate)

		assert.Equal(t, 10440, billedAmount)
		assert.NoError(t, err)
	})

	t.Run("正常系_消費税が小数", func(t *testing.T) {
		commision := 41

		invoice := Invoice{
			PaidAmount: 1000,
			Commission: &commision,
		}

		billedAmount, err := invoice.CalcBilledAmount(taxRate)

		assert.Equal(t, 1045, billedAmount)
		assert.NoError(t, err)
	})

	t.Run("異常系_手数料がnil", func(t *testing.T) {
		invoice := Invoice{
			PaidAmount: 1000,
		}
		_, err := invoice.CalcBilledAmount(0.08)
		assert.Error(t, err)
	})
}
