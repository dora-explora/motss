# removes the last 3 lines of every thread because im dumb

import os

posts = os.listdir("./archive/threads")
entries = []

for post in posts:
    path = f"./archive/threads/{post}"
    file = open(path, "r")
    lines = file.read().splitlines()
    file.close()
    for i in range(len(lines)):
        lines[i] += "\n"
    file = open(path, "w")
    file.writelines(lines[:-3])
    file.close()
