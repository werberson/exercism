class Twofer {
    String twofer(String name) {
        if(name == null || name.trim().isEmpty()){
            name = "you";
        }
        return "One for " + name + ", one for me.";
    }
}
