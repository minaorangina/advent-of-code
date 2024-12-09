def multiply(x,y):
    return x * y
def add(x,y):
    return x+y
def concatenate(x,y):
    return int(str(x)+str(y))

with open("input/day07.txt", "r") as reader:
    lines = reader.readlines()
    lines = [l.split(" ") for l in lines]
    lines = [[l.strip() for l in y] for y in lines]
    lines = [[l.strip(":") for l in y] for y in lines]
    lines = [[int(l) for l in x] for x in lines]

    # ops_str = "*+|"
    ops = [multiply, add, concatenate]
    from itertools import product

    total = 0
    for l in lines:
        expected = l[0]
        seq = l[1:]
        end_idx = len(seq)-1
        runs = list(product(ops, repeat = end_idx))
        # runs_str = list(product(ops_str, repeat = end_idx))
        print
        match_found = False

        for r in runs:
            if match_found:
                break
            outcome = seq[0]
            for i, v in enumerate(seq[1:]):
                next = seq[i+1]
                outcome = r[i](outcome, next)
            if outcome == expected:
                match_found = True
                total += expected

    print("total", total)