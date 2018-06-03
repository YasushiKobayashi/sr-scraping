# -*- coding: utf-8 -*-

import logging.config

from selenium import webdriver
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import Select, WebDriverWait

logging.config.fileConfig("logging.conf")
logger = logging.getLogger()


def app_log(message):
    logging.info(message)


def app_error_log(message):
    logging.error(message)



def start_chrome():
    d = DesiredCapabilities.CHROME
    d['loggingPrefs'] = {'browser': 'ALL'}
    options = Options()
    driver = webdriver.Chrome(options=options)
    driver.set_window_size(1366, 768)
    driver.implicitly_wait(15)
    return driver
