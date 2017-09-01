package numutils

func SwapInt(a, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
}

func ExchanggeSymbol(a int) (ret int) {
	ret = ^a + 1
	return
}

func CheckOdd(a int) bool {
	if a%2 > 0 {
		return true
	}
	return false
}

func CalcAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}

/*
	calculate the number of 1 in a
*/
func CalcOneNum(a int) int {
	a = ((a & 0xAAAA) >> 1) + (a & 0x5555)
	a = ((a & 0xCCCC) >> 2) + (a & 0x3333)
	a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F)
	a = ((a & 0xFF00) >> 8) + (a & 0x00FF)
	return a
}

/*
	reverse the order according to the int's binary byte
*/
func ByteReverse(a int) int {
	a = ((a & 0xAAAA) >> 1) | ((a & 0x5555) << 1)
	a = ((a & 0xCCCC) >> 2) | ((a & 0x3333) << 2)
	a = ((a & 0xF0F0) >> 4) | ((a & 0x0F0F) << 4)
	a = ((a & 0xFF00) >> 8) | ((a & 0x00FF) << 8)
	return a
}