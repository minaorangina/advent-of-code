import numpy as np

# filepath = "input/day09-test.txt"
filepath = "input/day09.txt"

with open(filepath, "r") as reader:
    file = reader.readlines()
    for i, row in enumerate(file):
        file[i] = [int(float(x)) for x in file[i].split()]


def do(nums):
    n = 0
    if len(set(nums)) == 1 and nums[0] == 0:
        return nums[-1] 
    
    diff = np.diff(nums)
    n = diff[-1]

    return do(diff) + nums[-1]


def part1(file):
    sum = 0
    for row in file:
        sum += do(row)

    print(sum)


part1(file)
