from flask import Flask
import os
from driver import ThreadSafeDriver

app = Flask(__name__)
driver = ThreadSafeDriver()


@app.route('/e/<contest_number>', methods=['GET'])
def enter(contest_number):
    print("From server - enter:", contest_number)
    contest_url = f"https://codeforces.com/contest/{contest_number}"
    driver.driver.get(contest_url)
    return f"contest {contest_number}"


@app.route('/q', methods=['GET'])
def quit_():
    driver.quit_()
    # noinspection PyUnresolvedReferences
    os._exit(0)


if __name__ == '__main__':
    app.run(port=5000)
