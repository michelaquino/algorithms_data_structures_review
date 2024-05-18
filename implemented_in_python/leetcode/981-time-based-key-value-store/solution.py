import collections


class TimeMap:
    def __init__(self):
        self.data = {}

    def set(self, key, value, timestamp):
        if key not in self.data:
            self.data[key] = []

        self.data[key].append((timestamp, value))

    def get(self, key, timestamp):
        result = ""
        values = self.data.get(key, [])
        left, right = 0, len(values) - 1

        while left <= right:
            middle = left + ((right - left) // 2)
            t, v = values[middle][0], values[middle][1]

            if t <= timestamp:
                result = v
                left = middle + 1
            elif t > timestamp:
                right = middle - 1

        return result


t = TimeMap()
# t.set("foo", "bar", 0)
# t.set("foo", "bar1", 4)
# t.set("foo", "bar2", 7)
# t.set("foo", "bar3", 3)

# print(t.data)


# ["TimeMap", "set", "get", "get", "set", "get", "get"]
# [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
t.set("foo", "bar", 1)
print("get -> bar: ", t.get("foo", 1))
print("get -> bar: ", t.get("foo", 333))
t.set("foo", "bar2", 4)
print("get -> bar2: ", t.get("foo", 4))
print("get -> bar2: ", t.get("foo", 5))


# ["TimeMap","set","set","get","get","get","get","get"]
# [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
# [null,null,null,"","high","high","low","low"]
# t.set("love", "high", 10)
# t.set("love", "low", 20)
# print("get -> '': ", t.get("foo", 1))
# print("get -> 'high': ", t.get("love", 5))
# print("get -> 'high': ", t.get("love", 10))
# print("get -> 'low': ", t.get("love", 15))
# print("get -> 'low': ", t.get("love", 20))
# print("get -> 'low': ", t.get("love", 25))
