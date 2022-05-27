# Find top ten IPs from web server log file. Here’s what the log looks like:

# $ cat logfile.log
"""
1222124699	GET /cats/meow	2.3.4.5
1222159481	GET /cats/meow	2.3.4.5
1222384903	GET /bird/chirp	1.3.6.5
1222154356	GET /cats/meow	2.3.4.5
1223493000	GET /cow/moo 	2.3.7.5
1223849499	GET /cats/meow	1.3.4.5
1223983884	GET /cow/moo 	5.6.4.5
1224029939	GET /cow/moo 	8.3.6.5
1224198747	GET /cats/meow	8.3.6.5
1224524688	GET /cats/meow	2.3.7.5
122487298	GET /bird/peep	8.3.6.5
1225128619	GET /cow/moo	1.3.6.5
1225354356	GET /cats/meow	8.3.6.5
1225711424	GET /dog/bark	1.3.4.5
1225852885	GET /cow/moo	1.3.6.5
1226254356	GET /bird/peep	1.3.4.5
1226620691	GET /bird/chirp	8.3.6.5
"""

# file structure

# <line>\n...

# main.py


def top_ten_ips(fpath: str) -> dict[str, int]:
    # read in the file
    try:
        with open(fpath, "r") as f:
            contents = f.read()
    except OSError:
        pass

    # {"1": "2.3..."}
    entries = contents.split("\n")
    entry_maps = []

    # get the amount of requests
    for entry in entries:
        print(entry)
        timestamp, rest = entry.split("\t")
        method, path, ip = rest.split(" ")
        entry_map = {
            "timestamp": timestamp,
            "method": method,
            "path": path,
            "ip": ip
        }
        entry_maps.append(entry_map)

    # sort the entries
    entry_maps = dict(sorted(entry_maps.items(), key=lambda x: x[1]))
    return entry_maps.items()[:10]


if __name__ == "__main__":
    print(top_ten_ips("logfile.log"))
