# removes the text file for any archived threads without valid messages

import os

def process_post(path: str):
    file = open(path, "r")
    lines = file.read().splitlines()
    if lines[3] == "(message was not found)": # this can be replaced with any condition
        print(f"{path[:-4]} is invalid, removing")
        os.remove(path)
    file.close()

walk = os.walk("./archive")
socposts = []
netposts = []
entries = []
for i, t in enumerate(walk):
    if i == 1:
       socposts = t[2]
    elif i == 2:
        netposts = t[2]

for post in socposts:
    path = f"./archive/soc.motss/{post}"
    process_post(path)

for post in netposts:
    path = f"./archive/net.motss/{post}"
    process_post(path)
