# class Solution:

#     def encode(self, strs: List[str]) -> str:
#         encoded = ""

#         for s in strs:
#             lenStr = len(s)
#             encoded += lenStr + "#" + s

#         return encoded

#     def decode(self, s: str) -> List[str]:
#         return []


def encode(strs):
    encoded = ""

    for s in strs:
        lenStr = len(s)
        encoded += str(lenStr) + "#" + s

    return encoded


def decode(encodedStr):
    result = []

    quantityStr = ""
    index = 0

    while index < len(encodedStr):
        if encodedStr[index] == "#":
            quantity = int(quantityStr)

            word = encodedStr[index + 1 : index + quantity + 1]
            result.append(word)

            quantityStr = ""
            index += quantity + 1
            continue

        quantityStr += encodedStr[index]
        index += 1

    return result


# strs = ["abc#d#8-0!as", "code", "love", "you"]
# encoded = encode(strs)
# print(encoded)
# decoded = decode(encoded)
# print(decoded)

strs = [""]
encoded = encode(strs)
print(encoded)
decoded = decode(encoded)
print(decoded)
