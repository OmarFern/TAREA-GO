package tp0

import (
	//"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	a := 10
	b := 5
	Swap(&a, &b)
	require.EqualValues(t, 5, a)
	require.EqualValues(t, 10, b)

	Swap(&a, &b)
	require.EqualValues(t, 10, a)
	require.EqualValues(t, 5, b)

	Swap(&a, &a)
	require.EqualValues(t, 10, a)

	a = 3
	b = 3
	Swap(&a, &b)
	require.EqualValues(t, 3, a)
	require.EqualValues(t, 3, b)
}



func TestComparar(t *testing.T) {
	{
		a := []int{10}
		b := []int{10}
		require.Equal(t, 0, Comparar(a, b))
	}
	{
		a := []int{10, 20}
		b := []int{10, 20}
		require.Equal(t, 0, Comparar(a, b))
	}
	{
		a := []int{}
		b := []int{}
		require.Equal(t, 0, Comparar(a, b))
	}
	{
		a := []int{1, 2, 3}
		b := []int{0, 2, 3}
		require.Equal(t, 1, Comparar(a, b))
		require.Equal(t, -1, Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{1, 2}
		require.Equal(t, 1, Comparar(a, b))
		require.Equal(t, -1, Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{1, 2, 2, 4}
		require.Equal(t, 1, Comparar(a, b))
		require.Equal(t, -1, Comparar(b, a))
	}
}



func TestSuma(t *testing.T) {
	var (
		vacio = []int{}
		unico = []int{8}
		vec1  = []int{3, 5, 4, 2, 1}
		vec2  = []int{4, 8, 15, 16, 23, 42}
		vec3  = []int{-38, -46, -65, -78}
		vec4  = []int{10, 9, -15, 0, 7, -12, 1}
	)

	require.Equal(t, 0, Suma(vacio))
	require.Equal(t, 8, Suma(unico))
	require.Equal(t, 15, Suma(vec1))
	require.Equal(t, 108, Suma(vec2))
	require.Equal(t, -227, Suma(vec3))
	require.Equal(t, 0, Suma(vec4))
}
