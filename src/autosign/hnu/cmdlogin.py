import login
import sys

def main():
    cookies = login.login(sys.argv[1], sys.argv[2])
    # login.add(cookies)

if __name__ == "__main__":
    main()
