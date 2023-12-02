import functools
from typing import List, Dict

filepath = "input/day01-1.txt"


# filepath = 'input/day01-1-test.txt'
def part1():
    with open(filepath, "r") as reader:
        line = reader.readline()

        results = []

        while line != "":  # '' == EOF
            num = ""

            for c in line:
                try:
                    int(float(c))
                    num = num + c
                    break
                except:
                    pass

            line2 = line[::-1]
            for c in line2:
                try:
                    int(float(c))
                    num = num + c
                    break
                except:
                    pass

            results.append(num)
            line = reader.readline()

        results2 = functools.reduce(lambda x, y: int(float(x) + float(y)), results)
        print(results2)


filepath = 'input/day01-1.txt'
# filepath = "input/day01-2-test.txt"


def part2():
    known_numbers = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    with open(filepath, "r") as reader:
        line = reader.readline()

        all_nums = []

        while line != "":
            c_idx = 0
            digits = []
            num = ""

            while c_idx < len(line):
                c = line[c_idx]
                try:
                    int(float(c))
                    digits.append(c)

                except:
                    found = find_word(line[c_idx:], 0, known_numbers)

                    if found != "":
                        # map string num to actual num
                        digits.append(known_numbers[found])
                        
                c_idx += 1

            if len(digits) != 0:
                num += digits[0]
                num += digits[-1]
                all_nums.append(num)

            line = reader.readline()

        # print(all_nums)

        answer = functools.reduce(lambda x,y: int(float(x)) + int(float(y)), all_nums)
        print(answer)


def find_word(line: str, i: int, words: Dict[str, str]):
    if len(words) == 0:
        return ""

    if i >= len(line):
        # we've overshot it
        return ""

    # whittling down
    fewer_words = {}

    for word in words.keys():
        if i >= len(word):
            return word

        if line[i] == word[i]:
            fewer_words[word] = words[word]

    return find_word(line, i + 1, fewer_words)


part1()
part2()
