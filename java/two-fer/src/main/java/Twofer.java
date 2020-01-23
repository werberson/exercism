class Twofer {

    private static final String DEFAULT_NAME = "you";

    String twofer(String name) {
        return "One for " + validOrDefault(name) + ", one for me.";
    }

    String validOrDefault(String name) {
        if (name == null || name.trim().isEmpty()) {
            return DEFAULT_NAME;
        }
        return name;
    }
}
