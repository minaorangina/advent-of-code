from functools import reduce

def part1():
    with open("input/day02-1.txt", "r") as reader:
        line = reader.readline()

        rules = {"red": 12, "green": 13, "blue": 14}

        line_num = 1
        sum = 0

        while line != "":
            sets = line.split(":")[1].split(";")
            valid = True

            for s in sets:
                draws = [x.strip() for x in s.split(",")]

                for d in draws:
                    bits = d.split(" ")
                    num_cubes = int(float(bits[0]))
                    colour = bits[1]

                    if rules[colour] < num_cubes:
                        valid = False

            if valid:
                sum += line_num

            line_num += 1

            # print(sets)

            line = reader.readline()

        print(sum)


def part2():
    # filepath = "input/day02-test.txt"
    filepath = "input/day02-1.txt"
    with open(filepath, "r") as reader:
        line = reader.readline()

        sum = 0

        while line != "":
            max = {"red": 0, "green": 0, "blue": 0}

            sets = line.split(":")[1].split(";")

            for s in sets:
                draws = [x.strip() for x in s.split(",")]

                for d in draws:
                    bits = d.split(" ")

                    num_cubes = int(float(bits[0]))
                    colour = bits[1]

                    if max[colour] < num_cubes:
                        max[colour] = num_cubes

            power = reduce(lambda x, y: x * y, max.values())

            sum += power

            line = reader.readline()

        print(sum)


part1()
part2()
