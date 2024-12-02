
def day01_1(lines):
    left = []
    right = []
    total = 0

    for l in lines:
        split_line = l.split()
        left.append(int(split_line[0]))
        right.append(int(split_line[1]))

    left.sort()
    right.sort()

    for i, _ in enumerate(left):
        total += abs(left[i] - right[i])

    return total

def day01_2(lines):
    left = []
    right_count = {}
    total = 0

    for l in lines:
        split_line = l.split()

        l, r = int(split_line[0]), int(split_line[1])
        left.append(l)

        if r in right_count:
            right_count[r] = right_count[r]+1
        else:
            right_count[r] = 1

    for v in left:
        total += (v * right_count.get(v, 0))

    return total

with open("input/day01.txt", "r") as reader:
    lines = reader.readlines()
    print(day01(lines))
    print(day02(lines))
