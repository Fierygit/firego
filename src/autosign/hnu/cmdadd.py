import login
import sys
import time


def main():
    # login.login(sys.argv[1], sys.argv[2])

    with open("/home/firefly/firego/src/autosign/hnu/user", 'r') as lines:
        for line in lines:
            try:
                [id, name] = parse(line)
                cookies = login.login(id, name)
                # get_info(cookies)
                login.add(cookies)
                time.sleep(1)
            except:
                print(id)
                print("warning warning -------------------------------------")


def parse(line):
    for i, val in enumerate(line):
        if val == "=":
            return line[:i], line[i+1:len(line)-1]


if __name__ == "__main__":
    main()
