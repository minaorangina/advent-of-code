
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



print(day02_1())