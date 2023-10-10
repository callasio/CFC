import requests

if __name__ == '__main__':
    try:
        requests.get(f"http://127.0.0.1:5000/q")
    except requests.exceptions.ConnectionError:
        print("Quit successfully")
