# creates a directory.txt that contains the id, title, sender, and date of every post in both groups

import os

class Entry:
    def __init__(self, id: int, title: str, sender: str, date: str, day: int):
        self.id: int = id
        self.title: str = title
        self.sender: str = sender
        self.date: str = date
        self.day: int = day

    def __str__(self):
        return f"{self.title} - posted on {self.date} by {self.sender} - [{self.id}]"

def process_post(file):
    id = int(post[:-4])
    title = file.readline()[:-1]
    sender = file.readline()[:-1]
    date = file.readline()[:-1]
    day = (int(date[0:2])) * 31 + (int(date[3:5])) + (int(date[6:8]) * 366) # shhh you see nothing
    entry = Entry(id, title, sender, date, day)
    entries.append(entry)

file = open("./archive/directory.txt", "w")
file.write("")
file.close()
file = open("./archive/directory.txt", "a")

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
    postfile = open(f"./archive/soc.motss/{post}", "r")
    process_post(postfile)
    postfile.close()

for post in netposts:
    postfile = open(f"./archive/net.motss/{post}", "r")
    process_post(postfile)
    postfile.close()


entries.sort(key=lambda x: x.day)
for entry in entries:
    lines: list[str] = []
    lines.append(str(entry.id) + "\n")
    lines.append(str(entry.title) + "\n")
    lines.append(str(entry.sender) + "\n")
    lines.append(str(entry.date) + "\n\n")
    file.writelines(lines)
