#!/usr/bin/python
#coding:utf-8
#Author: Yang Sen


import jenkinsapi
from jenkinsapi.jenkins import Jenkins
import optparse
import os, sys, time
import xml.dom.minidom
job_config_template    = os.path.join(os.path.dirname(os.path.realpath(__file__)), "jon_config_template.xml")



# def get_plugin_details(server):
#     for plugin in server.get_plugins().values():
#         print "-"*10
#         print "Short Name:%s" %(plugin.shortName)
#         print "Long Name:%s" %(plugin.longName)
#         print "Version:%s" %(plugin.version)
#         print "URL:%s" %(plugin.url)
#         print "Active:%s" %(plugin.active)
#         print "Enabled:%s" %(plugin.enabled)


def create_folder():
    ###instsall plugins:　CloudBees Folders Plugin, https://wiki.jenkins-ci.org/display/JENKINS/CloudBees+Folders+Plugin
    #http://updates.jenkins-ci.org/latest/cloudbees-folder.hpi

    #
    pass



def create_job(jenkinshandle, jobname="mytest", jobxmlconfig="jon_config_template.xml"):
    xmlhandle=xml.dom.minidom.parse(job_config_template)
    commands = xmlhandle.documentElement.getElementsByTagName('command')
    commands[0].childNodes[0].nodeValue = "uname -a;ifconfig;whoami;"
    custom_configxml_content = xmlhandle.toxml('UTF-8')


    #print(custom_configxml_content)
    job = jenkinshandle.create_job(jobname=jobname, xml=custom_configxml_content)
    my_job = jenkinshandle[jobname]
    print(my_job)
    # jenkinshandle.delete_job(jobname)




if __name__ == '__main__':
    option_parser = optparse.OptionParser()
    option_parser.add_option('-v','--view',       dest="clearcaseview", default='', help="ClearCase Linux View")
    option_parser.add_option('-p','--project'  ,  dest="project"      , default='SoC'  , help="")
    option_parser.add_option('-f','--frametype',  dest="frameType"    , default='TDD'  , help="")
    option_parser.add_option('-c','--csls'      ,  dest="csls"        , default='senya', help="")
    options, args = option_parser.parse_args()
    
    #Jenkins(jenkins_url, username = 'foouser', password = 'foopassword')
    j = Jenkins('http://10.253.6.128:9999')
    print "============================="
    print j.version
    for k,v in j.get_jobs():
        print k,v
    for i in j.get_jobs():
        job_instance = j.get_job(i[0])
        #print dir(job_instance)
        print 'Job Name:%s' %(job_instance.name)
        print 'Job Description:%s' %(job_instance.get_description())
        print 'Is Job running:%s' %(job_instance.is_running())
        print 'Is Job enabled:%s' %(job_instance.is_enabled())

    print "============================="
    #print j.poll(tree='jobs[name,color,url]')
    print j.get_folders()

    create_job(jenkinshandle=j, jobname="mytest1111111", jobxmlconfig="jon_config_template.xml")

    # print j.create_folder("xxwfwgfg43g34455111111112211111")

    # k = Jenkins("http://10.253.6.128:9999/job/xxwfwgfg43g34455111111112211111/")
    # k.create_folder("xxxxx")


    # time.sleep(5)
    # print j.delete_folder("xxwfwgfg43g34455111111112211111")
    # print j.delete_folder("HAHAfolder1111111")
    # print j.delete_folder("HAHAfolder2222")
    # print j.delete_folder("folderdesc")
    # print j.delete_folder("HAHAfolder")

    #get_plugin_details(j)

    # try:
    #    j.create(options.clearcaseview, job_config_template)
    # except e:
    #    print e
    

#     #job cmdline
#     cmdline="""#!/usr/bin/env bash
# uptime
# python /home/senya/LteCloud/TaaS/L2_build_test.py --view=%s --project=%s  --frametype=%s
# """ % (options.clearcaseview, options.project, options.frameType)

#     #custom
#     xmlhandle=xml.dom.minidom.parse(job_config_template)
#     commands = xmlhandle.documentElement.getElementsByTagName('command')
#     commands[0].childNodes[0].nodeValue = cmdline
#     custom_configxml_content = xmlhandle.toxml('UTF-8')

#     try:
#        j.set_config_xml(options.clearcaseview, custom_configxml_content)

#        #trigger
#        j.build(options.clearcaseview)
#     except e:
#        print e
