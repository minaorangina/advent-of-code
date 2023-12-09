import numpy as np

# filepath = "input/day09-test.txt"
filepath = "input/day09.txt"

with open(filepath, "r") as reader:
    file = reader.readlines()
    for i, row in enumerate(file):
        file[i] = [int(float(x)) for x in file[i].split()]


def do(nums, idx):
    if len(set(nums)) == 1 and nums[0] == 0:
        return nums[idx]

    diff = np.diff(nums)

    rtn = do(diff, idx)
    return rtn + nums[idx]


def do2(nums, idx):
    if len(set(nums)) == 1 and nums[0] == 0:
        return nums[idx]

    diff = np.diff(nums)

    rtn = do2(diff, idx)
    return nums[idx] - rtn


def part1(file):
    sum = 0
    for row in file:
        sum += do(row, -1)

    print(sum)


def part2(file):
    sum = 0
    for row in file:
        sum += do2(row, 0)

    print(sum)


part1(file)
part2(file)
