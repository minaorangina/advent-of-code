filepath = "input/day05-1.txt"

def part1():
    with open(filepath, "r") as reader:
        data = reader.readlines()


    def process(ranges, seed):
        # get ranges
        v = seed
        for r in ranges:
            ss = r[1]
            lr = r[2]
            ds = r[0]

            if ss <= v <= ss + lr:
                # you belong here
                v = v + ds - ss
                return v

        # if we get here, there was no match
        return v


    seeds = [int(float(x)) for x in data[0].split()[1:]]

    instructions = []

    instruct = []
    for row in data[3:]:
        row = row.strip()

        if row == "":
            continue
        if row[0].isnumeric():
            # collect
            instruct.append([int(float(x)) for x in row.split()])

        else:
            instructions.append(instruct)
            instruct = []

    instructions.append(instruct)

    output = []

    outs = []
    for s in seeds:
        for i in range(len(instructions)):
            s = process(instructions[i], s)
        outs.append(s)

    print(min(outs))