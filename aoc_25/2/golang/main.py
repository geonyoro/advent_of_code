with open("input") as wfile:
    parts = wfile.read().split(",")
    total = 0
    for part in parts:
        if not part:
            continue
        start, end = part.split("-")
        print(repr(part), repr(start), repr(end))
        for i in range(int(start), int(end) + 1):
            s = str(i)
            length = len(s)
            if length % 2 != 0:
                continue
            if s[length//2:] * 2 == s:
                total += i
                # print(i)
    print(total)
