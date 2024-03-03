class StringBuilder:
    def __init__(self):
        self.strings = []

    def append(self, word):
        self.strings.append(word)

    def to_string(self):
        return " ".join(self.strings)


sb = StringBuilder()
sb.append("aaa")
sb.append("b")
sb.append("cc")
sb.append("dddd")

print(sb.to_string())
