#!/usr/bin/env python
#-*- coding=utf-8 -*-

from fabric.api import *


env.hosts=[
    'root@10.139.98.32',
    'root@10.253.0.205',
    'root@10.253.101.103',
    'root@10.253.101.135',
]
env.passwords={
'root@10.139.98.32':'123',
'root@10.253.0.205':'123',
'root@10.253.101.103':'123',
'root@10.253.101.135':'123',
}
@task
def task1():
    run('df -h')

@task
def task2():
    run('uname -a')
@task
def dotask():
    execute(task1)
    execute(task2)


dotask()
