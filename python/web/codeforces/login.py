from selenium.webdriver.firefox.webdriver import WebDriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


def login(driver: WebDriver, login_info):
    driver.get("https://codeforces.com/enter")

    handle_or_email = driver.find_element(By.ID, 'handleOrEmail')
    password = driver.find_element(By.ID, 'password')
    submit = driver.find_element(By.CLASS_NAME, 'submit')

    handle_or_email.send_keys(login_info[0])
    password.send_keys(login_info[1])
    submit.click()

    url_to_wait_for = 'https://codeforces.com/'
    WebDriverWait(driver, 20).until(EC.url_to_be(url_to_wait_for))
