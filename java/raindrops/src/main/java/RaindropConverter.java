class RaindropConverter {

    /**
     * Convert a number to a string, the contents of which depend on the number's factors.
     *
     * If the number has 3 as a factor, output 'Pling'.
     * If the number has 5 as a factor, output 'Plang'.
     * If the number has 7 as a factor, output 'Plong'.
     * If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
     *
     * @param number Number to be converted
     * @return String number converted to sound
     */
    String convert(int number) {
        String result = "";
        if (isMultiple(number, 3)) {
            result = "Pling";
        }
        if (isMultiple(number, 5)) {
            result += "Plang";
        }
        if (isMultiple(number, 7)) {
            result += "Plong";
        }
        return result.isEmpty() ? Integer.toString(number) : result;
    }

    /**
     * If the number is multiple of factor
     *
     * @param number number to check
     * @param factor factor
     * @return true if the number is multiple of factor, false if is not.
     */
    private boolean isMultiple(int number, int factor) {
        return number % factor == 0;
    }

}
