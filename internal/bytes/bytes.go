package bytes

import "math/big"

func PadBytes(length int, data []byte) []byte {
	dataLength := len(data);
	if dataLength >= length {
		return data
	}

	ret := make([]byte, length);
	copy(ret[length - dataLength:], data)

	return ret
}

func ReverseBytes(data []byte) []byte {
	length := len(data)
	newData := make([]byte, length);
	for i := 0; i < length; i++ {
		newData[i] = data[length - i - 1];
	}

	return newData
}

func BytesToBigIntLittleEndian(data []byte) *big.Int {
	return big.NewInt(0).SetBytes(ReverseBytes(data))
}

func BigEndianBytesFromLittleEndianBigInt(padding int, bigInt *big.Int) []byte {
	return ReverseBytes(PadBytes(padding, bigInt.Bytes()))
}
