const colorCode = {
    black: "0",
    brown: "1",
    red: "2",
    orange: "3",
    yellow: "4",
    green: "5",
    blue: "6",
    violet: "7",
    grey: "8",
    white: "9"
};

/**
 * Decode the resistor color names as two digit number.
 * @param colors represents the colors of the first 2 or more color bands of a resistor
 * @returns {number} the combined number value of the first 2 color bands
 */
export const decodedValue = (colors) => Number(colorCode[colors[0]] + colorCode[colors[1]]);
