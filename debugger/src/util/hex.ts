export const toHex = (value: number, padding: number) => {
    let hexValue = value.toString(16);

    while (hexValue.length < padding) {
        hexValue = '0' + hexValue
    }
    return '0x' + hexValue;
}