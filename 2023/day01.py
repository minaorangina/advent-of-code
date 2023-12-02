import functools
from typing import List, Dict

filepath = 'input/day01-1.txt'
# filepath = 'input/day01-1-test.txt'
def part1():
    with open(filepath, 'r') as reader:
        line = reader.readline()

        results = []

        while line != '': # '' == EOF
            print(line, end='')
            num = ''
            
            for c in line:
                try:
                    int(float(c))
                    num = num+ c
                    break
                except:
                    pass

            line2 = line[::-1]
            for c in line2:
                try:
                    int(float(c))
                    num=num+c
                    break
                except:
                    pass

            results.append(num)
            line = reader.readline()

        print(results) 
        results2 = functools.reduce(lambda x, y: int(float(x) + float(y)), results)
        print(results2) 

# filepath = 'input/day01-1.txt'
filepath = 'input/day01-2-test.txt'
def part2():

    known_numbers = {
        'one': '1', 'two': '2', 'three': '3', 'four': '4', 'five': '5', 'six': '6', 'seven': '7', 'eight': '8', 'nine': '9',
    }

    with open(filepath, 'r') as reader:
        line = reader.readline()

        num = ''

        for c in line:
            try:
                int(float(c))
                num = num + int(float(c))
            except:
                found = find_word(line, 0, known_numbers)
                if len(found) == 1:
                    # map string num to actual num
                    for k, v in found.items():
                        num = num + v
                        break
                    print(num)
                    break


def find_word(line: str, i: int, words: Dict[str, str]):
    print(i, words, line)
    if len(words) < 2:
         return words
    if i >= len(line):
        # we've overshot it
        return {}

    fewer_words = {}
    # smallify dict
    for word in words.keys():
         if line[i] == word[i]:
              fewer_words[word] = words[word]

    return find_word(line, i+1, fewer_words)


part2()