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
        line = reader.readline() #two1nine

        all_nums = []

        c_idx = 0
        while line != '':
            digits = []
            num = ''
            while c_idx < len(line):
                c = line[c_idx]
                try:
                    int(float(c))
                    digits.append(c)
                except:
                    # print("looking in %s for %s @ %d"%(line[c_idx:],c, 0))
                    found = find_word(line[c_idx:], 0, known_numbers)
                    if found != '':
                        print("we found one!!!", found)
                        # map string num to actual num
                        digits.append(known_numbers[found])
                        print("collected thus far", digits)
                c_idx += 1

            if len(digits) != 0:
                print("num",num, digits)
                num += digits[0]
                num += digits[-1]
                all_nums.append(num)
                print(all_nums)
                line = reader.readline()

        print(all_nums)


def find_word(line: str, i: int, words: Dict[str, str]):
    if len(words) == 0:
         return ''
    # print(">>>looking for "+line[i], i, words, line)
    if i >= len(line):
        # we've overshot it
        return ''

    fewer_words = {}
    # smallify dict
    for word in words.keys():
         if i >= len(word):
            print("returning with", word)
            return word
         if line[i] == word[i]:
              fewer_words[word] = words[word]

    # print('we keep looking')
    return find_word(line, i+1, fewer_words)


part2()