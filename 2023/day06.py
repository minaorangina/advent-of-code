import math


filepath = "input/day06.txt"
# filepath = "input/day06-test.txt"


def solve_quadratic(a, b, c):
    x1 = (b + math.sqrt(b * b - (4 * a * c))) / 2
    x2 = (b - math.sqrt(b * b - (4 * a * c))) / 2

    return x1, x2


def part1():
    with open(filepath, "r") as reader:
        input = reader.readlines()

    times = [int(float(x)) for x in input[0].split(":")[1].strip().split()]
    distances = [int(float(x)) for x in input[1].split(":")[1].strip().split()]


    sum = 0
    for i in range(len(distances)):
        d = distances[i]  # 9
        t = times[i]  # 7
        ranges = solve_quadratic(1, t, d)
        poss = math.floor(max(ranges)) - math.floor(min(ranges))
        if sum == 0:
            sum = poss
        else:
            sum *= poss

    print(sum)


def part2():
    with open(filepath, "r") as reader:
        input = reader.readlines()

    t = input[0].split(":")[1].strip().split()
    t = int(float("".join(t)))

    d = input[1].split(":")[1].strip().split()
    d = int(float("".join(d))) 


    sum = 0
    ranges = solve_quadratic(1, t, d)
    poss = math.floor(max(ranges)) - math.floor(min(ranges))
    print(ranges, poss)
    if sum == 0:
        sum = poss
    else:
        sum *= poss

    print(sum)


part1()
part2()
