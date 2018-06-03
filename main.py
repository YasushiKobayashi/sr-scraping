#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import threading
import time
import traceback
from datetime import datetime

from selenium import webdriver
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import Select, WebDriverWait

from utils import *

URL = 'https://www.showroom-live.com/'


account_id = os.getenv('account_id', '')
password = os.getenv('password', '')

# multi thred is not run
targets = ['48_IMAMURA_MITSUKI']

class Driver():
    def __init__(self):
        super().__init__()
        self.driver = None

    def main(self):
        threads = []
        for v in targets:
            t = threading.Thread(target=self.run, args=(v,))
            threads.append(t)
            t.start()
            # self.run(v)


    def run(self, target):
        try:
            self.driver = start_chrome()
            self.driver.get(URL+target)
            is_delivery = self.is_delivery()
            if is_delivery == False:
                raise Exception(target + 'delivery is ended.')
            self.run_js('showLoginDialog();')
            self.login()
            self.comment()

            app_log(targets + "count done")
            self.driver.close()
        except Exception as e:
            self.driver.close()
            print('\033[91m')
            print(target)
            print(traceback.format_exc())
            app_error_log(target)
            app_error_log(traceback.format_exc())
            print('\033[0m')


    def login(self):
        # self.send_key("account_id", account_id)
        # self.send_key("password", password)
        self.send_key_js("account_id", account_id)
        self.send_key_js("password", password)
        self.click_by_id("js-login-submit")

    def comment(self):
        i = 0
        while True:
            time.sleep(5)
            i = i+1
            print(i)
            self.send_key_by_id('js-chat-input-comment', i)
            self.driver.find_element_by_xpath("//*[@id='js-room-comment']//*[@type='submit']").click()
            if i == 50:
                break

    def is_delivery(self):
        try:
            self.driver.find_element_by_xpath("//*[@class='html5-video-container']")
            return False
        except:
            return True

    def click_by_id(self, name):
        element = self.find_element_by_id(name)
        element.click()


    def send_key(self, name, val):
        element = self.find_element_by_name(name)
        element.send_keys(val)


    def send_key_by_id(self, name, val):
        element = self.find_element_by_id(name)
        element.send_keys(val)


    def send_key_js(self, name, val):
        code = '''
$("[name='{name}']").val('{val}');
'''.format(name=name, val=val).strip()
        self.run_js(code)


    def find_element_by_id(self, name):
        return self.driver.find_element_by_id(name)


    def find_element_by_name(self, name):
        print(name)
        element = WebDriverWait(self.driver, 20).until(
            EC.presence_of_element_located(
                (By.XPATH, "//*[@name='" + name + "']"))
        )
        return element


    def run_js(self, code):
        self.driver.execute_script(code)


if __name__ == "__main__":
    driver = Driver()
    driver.main()
