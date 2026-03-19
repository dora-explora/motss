import os
import requests
from bs4 import BeautifulSoup as bs, Tag

group = "soc.motss"
path = f"./archive/{group}"
if not os.path.isdir(path):
    os.mkdir(path)

def process_thread(id: str):
    response = requests.get(f"https://usenet.trashworldnews.com/?thread={id}")
    soup = bs(response.content, 'html.parser')
    title = soup.find_all("h1")[0].getText()[12:]
    print(f"archiving thread {id}: {title}")
    output = title + "\n"
    ps = soup.find_all("p")
    pres = soup.find_all("pre")
    for i in range(len(ps)):
        output += process_message(ps[i], pres[i])
        output += "\n\n-----------------------------------------------------------------\n\n"
    outpath = f"{path}/{id}.txt"
    if os.path.isfile(outpath):
        return
    file = open(outpath, "w")
    file.write(output)
    file.close()

def process_message(p: Tag, pre: Tag):
    b = p.find("b")
    if b is None:
        print("No \'b\' element found in post's title.")
        exit(1)
    sender = b.string
    sender = "(sender was not found)" if sender is None else sender
    date = p.getText().split("(")[-1][:-1] # what have i become
    message = pre.string
    message = "(message was not found)" if message is None else message
    return sender + "\n" + date + "\n" + message

def sanitize_html(string: str):
    string = string.replace("&amp;#39;", "\'")
    string = string.replace("&amp;lt;", "<")
    string = string.replace("&amp;gt;", ">")
    string = string.replace("&lt;", "<")
    string = string.replace("&gt;", ">")
    return string

response = requests.get(f"https://usenet.trashworldnews.com/?group={group}&sort=date")
soup = bs(response.content, "html.parser")
links = soup.find_all("a")
for link in links[2:]:
    href = link.get("href")
    if href is not None:
        id = str(href)[9:]
        if os.path.isfile(f"{path}/{id}.txt"):
            print(f"skipping thread {id}")
            continue
        process_thread(id)
