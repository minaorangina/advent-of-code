import re

with open("input/day11.txt", "r") as reader:
    input = reader.readlines()
    input[0] = input[0].strip()
    input[0] = input[0].split(" ")
    input = input[0]

    with open("input/day11_out.txt", "w") as the_file:
        the_file.write("")

    src = input.copy()
    for blink in range (0, 25):
        output = []
        for i, v in enumerate(src):
            idx = len(output) -1

            if v == "0":
                output.append("1")
            
            elif v == "1":
                output.append(str(int(v) * 2024))
            
            elif len(v) % 2 == 0:
                left = v[0:int(len(v)/2)]
                right = v[int(len(v)/2):]

                pattern = r"^0+$"
                if re.match(pattern, left):
                    left = "0"
                if re.match(pattern, right):
                    right = "0"

                if len(left) > 1:
                    left = left.lstrip("0")
                if len(right) > 1:
                    right = right.lstrip("0")

                # TODO - deal with zeros
                output.append(left)
                output.append(right)
            
            else:
                output.append(str(int(v)*2024))
            
        got = " ".join(str(x) for x in output)
        # print(f"blink {blink+1}: {got}")
        src = output.copy()
    print(f"number of stones: {len(src)}")