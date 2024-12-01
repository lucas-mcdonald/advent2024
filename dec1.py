# Advent of Code 2024, December 1
with open("inputs/dec1.txt") as f:
    data = f.read().splitlines()
    # Split by whitespace
    a = []
    b = []
    for line in data:
        split = line.split()
        a.append(int(split[0]))
        b.append(int(split[1]))

    a = sorted(a)
    b = sorted(b)

    total = 0
    similarity = 0
    for i in range(len(a)):
       # Part 1
       total += abs(b[i] - a[i])
       # Part 2
       similarity += a[i] * b.count(a[i])

    print(f'Part1: {total}')
    print(f'Part2: {similarity}')