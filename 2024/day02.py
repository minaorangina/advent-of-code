
def day02_1():
    total_safe_reports = 0

    with open("input/day02.txt", "r") as reader:
        lines = reader.readlines()

    for line in lines:
        report = list(map(int, line.split()))
        print("new report", report)

        increasing, decreasing = 0, 0
        skipped = False

        for i, v in enumerate(report):
            print("new val")

            if i == len(report) - 1:
                break
            
            diff = report[i+1] - v
            # print(report[i+1], "minus", v, diff)

            # absolute diff can't exceed 3
            if abs(diff) > 3 or abs(diff) == 0:
                print("skipping big diff of ", diff)
                skipped = True
                break

            if diff >= 0:
                increasing += 1
            else:
                decreasing += 1

        if skipped:
            continue

        print("inc", increasing, "dec", decreasing)
        if increasing > 0 and decreasing > 0:
            continue
        
        total_safe_reports += 1
    
    return total_safe_reports


def get_possibilities(line):
    all = [line]


    for i, v in enumerate(line):
        # if i != len(line)-1:
            all.append(line[0:i] + line[i+1:])

    return all

def count_safe(reports):
    total_safe_reports = 0
    problems = 0
    safe_found = False

    for report in reports:
        skipped = False
        increasing, decreasing = 0, 0
        if safe_found:
            break

        print("\n", report)
        for i, v in enumerate(report):
            if i == len(report) - 1:
                break

            diff = report[i+1] - v
            #print("diff", diff)

            # absolute diff can't exceed 3
            if abs(diff) > 3 or abs(diff) == 0:
                print("    problem: diff too big (or too small)")
                problems += 1
                skipped = True
                break

            if diff >= 0:
                increasing += 1
            else:
                decreasing += 1

        # print("total problems for row: ", problems)
        if skipped:
            continue

        if increasing > 0 and decreasing > 0:
            print("    problem: inc and dec")
            problems += 1
            continue

        safe_found = True
        total_safe_reports += 1
        print("safe found for ", report)


    return total_safe_reports

def day02_2():
    total_safe_reports = 0

    with open("input/day02.txt", "r") as reader:
        lines = reader.readlines()

    for line in lines:
        report = list(map(int, line.split()))

        reports = get_possibilities(report)
        # print(reports)

        safe = count_safe(reports)
        print("safe total for possibilities", safe)
        total_safe_reports += safe

    return total_safe_reports



print(day02_2())
