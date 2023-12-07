from functools import cmp_to_key, reduce

filepath = "input/day07.txt"
# filepath = "input/day07-test.txt"


def compare(h1, h2):
    r1 = get_score(h1)
    r2 = get_score(h2)

    if r1 < r2:
        return -1
    if r1 > r2:
        return 1
    # the case where the hands have the same score
    for i in range(5):
        if h1[i] != h2[i]:
            return cards.index(h1[i]) - cards.index(h2[i])


def get_score(hand):
    for i, fn in enumerate(funcs):
        if fn(hand):
            return i+1


def separate(hand):
    vals = {}
    for v in hand:
        if v not in vals:
            vals[v] = 1
        else:
            vals[v] += 1

    return vals


def is_five_of_a_kind(hand):
    first = hand[0]

    for v in hand[1:]:
        if v != first:
            return False

    return True


def is_four_of_a_kind(hand):
    vals = separate(hand)
    return len(vals.keys()) == 2 and 4 in vals.values()


def is_full_house(hand):
    vals = separate(hand)
    return len(vals.keys()) == 2 and 2 in vals.values()


def is_three_of_a_kind(hand):
    vals = separate(hand)
    return len(vals.keys()) == 3 and 1 in vals.values()


def is_two_pair(hand):
    vals = separate(hand)
    if len(vals.keys()) == 3:
        return sorted(vals.values()) == [1, 2, 2]
    return False


def is_one_pair(hand):
    vals = separate(hand)
    return len(vals.keys()) == 4


def is_high_card(hand):
    vals = separate(hand)
    return len(vals.keys()) == 5


funcs = [
    is_high_card,
    is_one_pair,
    is_two_pair,
    is_three_of_a_kind,
    is_full_house,
    is_four_of_a_kind,
    is_five_of_a_kind,
]

cards = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"]


with open(filepath, "r") as reader:
    file = reader.readlines()
    hands = {}
    for line in file:
        l = line.split()
        hands[l[0]] = l[1]

    sorted_hands = sorted(list(hands.keys()), key=cmp_to_key(compare))

    print(hands.keys(), sorted_hands)

    sum = 0
    for i, h in enumerate(sorted_hands):
        bid = int(float(hands[h]))
        sum += bid * (i+1)

    print(sum)
