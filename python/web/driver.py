from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.firefox.options import Options


class ThreadSafeDriver:
    def __init__(self):
        self.driver = webdriver.Firefox(
            options=Options()
        )

    def login(self):
        self.driver.get("https://codeforces.com/enter")

        handle_or_email = self.driver.find_element(By.ID, 'handleOrEmail')
        password = self.driver.find_element(By.ID, 'password')
        submit = self.driver.find_element(By.CLASS_NAME, 'submit')

        handle_or_email.send_keys(
            # hidden
        )
        password.send_keys(
            # hidden
        )
        submit.click()

    def quit_(self):
        self.driver.quit()