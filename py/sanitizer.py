# removes the text file for any archived threads without valid messages

import os

posts = os.listdir("./archive/threads")
entries = []

for post in posts:
    path = f"./archive/threads/{post}"
    file = open(path, "r")
    lines = file.read().splitlines()
    if lines[3] == "(message was not found)": # this can be replaced with any condition
        print(f"{path[:-4]} is invalid, removing")
        os.remove(path)
    file.close()
