#!/usr/bin/python
# -*- coding: utf-8 -*-


url = 'https://httpbin.org/post'
from logged_requests import LoggedRequests
req = LoggedRequests()
resp = req.post(url, json={"name":"jsoh"})

res = req.get("https://httpbin.org")