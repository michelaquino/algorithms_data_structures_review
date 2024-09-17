# https://leetcode.com/problems/meeting-scheduler
#
# O(A + B) time, where A is length of slots1 and B is length of slots2
# O(1) memory
def findTimeSlot(slot1, slot2, duration):
    s1Pointer, s2Pointer = 0, 0

    while s1Pointer < len(slot1) and s2Pointer < len(slot2):
        if slot1[s1Pointer][1] < slot2[s2Pointer][0]:
            s1Pointer += 1
            continue

        if slot2[s2Pointer][1] < slot1[s1Pointer][0]:
            s2Pointer += 1
            continue

        overlapSlot = [
            max(slot1[s1Pointer][0], slot2[s2Pointer][0]),
            min(slot1[s1Pointer][1], slot2[s2Pointer][1]),
        ]

        if overlapSlot[1] - overlapSlot[0] >= duration:
            return [overlapSlot[0], overlapSlot[0] + duration]
        else:
            if slot1[1] < slot2[1]:
                s1Pointer += 1
            else:
                s2Pointer += 1

    return []


slots1 = [[10, 50], [60, 120], [140, 210]]
slots2 = [[0, 15], [60, 70]]
duration = 8
print(findTimeSlot(slots1, slots2, duration))

slots1 = [[10, 50], [60, 120], [140, 210]]
slots2 = [[0, 15], [60, 70]]
duration = 12
print(findTimeSlot(slots1, slots2, duration))
