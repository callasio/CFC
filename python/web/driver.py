from selenium import webdriver
from selenium.webdriver.firefox.options import Options

import internal
from codeforces import login


class ThreadSafeDriver:
    def __init__(self):
        self.driver = None
        self.contest_url = ''

    def enter(self, contest_url: str):
        self.driver = webdriver.Firefox(
            options=Options().add_argument('-headless')
        )
        self.login()
        self.set_contest(contest_url)

    def login(self):
        login.login(self.driver, internal.read_login_info())

    def set_contest(self, contest_url: str):
        self.contest_url = contest_url

    def quit_(self):
        self.driver.quit()
