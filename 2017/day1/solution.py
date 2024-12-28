import inputs

title = "Inverse Captcha"


def part1():
    data = inputs.string(1)
    return captcha(data, 1)


def part2():
    data = inputs.string(1)
    return captcha(data, len(data)//2)


def captcha(data: str, offset: int):
    return sum([
        int(l) for (l, r)
        in zip(data, data[offset:] + data[:offset])
        if l == r
    ])
