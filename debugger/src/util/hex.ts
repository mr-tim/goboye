export const toHex = (value: number, padding: number, prefix = true) => {
    let hexValue = value.toString(16);

    while (hexValue.length < padding) {
        hexValue = '0' + hexValue
    }
    if (prefix) {
        hexValue = '0x' + hexValue;
    }
    return hexValue;
}