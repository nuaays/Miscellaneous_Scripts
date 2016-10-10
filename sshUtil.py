#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author:
# Date:


#
#
import os, sys
sys.path.append( os.path.dirname( os.path.abspath(__file__) ) )
#sys.path.insert(0, "Lib")

import Crypto, paramiko
import time, datetime, subprocess, re
import pyping
#from paramikoe import SSHClientInteraction



def exectue_external_cmd( cmd,  msg_in='' ):
    proc = subprocess.Popen( cmd,  shell = True,  stdin = subprocess.PIPE,  stdout = subprocess.PIPE,  stderr = subprocess.PIPE)
    stdout_info,  stderr_info = proc.communicate( msg_in)
    return stdout_info #,  stderr_info

class sshConnect(object):
    def __init__(self, usrName, passWd, host, port=22, rootName='root', rootPassWord='****',logfile = "paramiko_ssh.log"):
        self.Flag = '=';
        self.numofFlag= 102;
        self.logfile  = logfile #self.logfile = os.path.join(os.getcwd(),"paramiko.log")
        self.username = usrName
        self.password = passWd
        self.rootname = rootName
        self.rootpswd = rootPassWord
        self.host_ip  = host
        self.host_port= port
        self.host_ping_online_pattern = 'Lost = 0 \(0% loss\)'
        self.ping_times = 3
        #
        self.send_cmd_timeout = 2
        #
        self.channel = None
        self.sftp = None
        self.rootshell = None
        #inter
        self.prompt = '$' #'fots@fotsies-ubuntu-testlab:~\$ '
        self.root_prompt = '#'

        # Launch ssh connecting to host
        # self.connect()

    def check_online(self):
        print "[INFO] Checking Host %s is online or not ... ..." % self.host_ip
        try:
            subprocess.check_call('ping %s -c %d -w %d' % (self.host_ip, self.ping_times, self.ping_timeout) , stdin = subprocess.PIPE, stdout= subprocess.PIPE, stderr=subprocess.STDOUT, shell = True)
        except Exception, e:
            print "[INFO] %s is Offline" % self.host_ip
            print e
            return False
        else:
            print "[INFO] %s is Online" % self.host_ip
            return True

    def connect(self):
        if not self.check_online():
            raise Exception("SSH Error As to Target is Offline!!!")
        self.channel = paramiko.SSHClient()
        self.channel.load_system_host_keys()
        self.channel.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        try:
            self.channel.connect( self.host_ip, self.host_port, username = self.username, password= self.password, pkey=None, timeout=None, allow_agent=False, look_for_keys=False)
        except Exception, e:
            print("Failed to SSH HOST: %s: %s\n" %(e.errno, e.strerror))
            #self.close()
            return False
        else:
            self.sftp = self.channel.open_sftp()     #self.sftp = paramiko.SFTPClient.from_transport(self.channel)
        return True

        #http://code.ohloh.net/file?fid=bXGaiYrEo7Fqnskvxx2H1Tex1hI&cid=p_NS_z9lX2s&s=&fp=507736&projSelected=true#L0
        # Create a client interaction class which will interact with the host
        #interact = SSHClientInteraction(self.channel, timeout=10, display=True)
        #interact.expect(self.prompt)


    def is_active(self):
        is_active = False
        if self.channel != None:
            try:
                trans = self.channel.get_transport()
                if trans.is_active():
                    is_active = True
            except:
                is_active = False
        return is_active

    def su_root(self):
        '''
        http://www.jb51.net/article/44070.htm
        '''
        if self.rootshell == None:
            self.rootshell = self.channel.invoke_shell()
            time.sleep(0.5)
        if self.username == 'root':
           print "Already root account!"
           return True
        # su to root account
        #self.rootshell.send( 'su %s\n' % self.rootname )
        self.rootshell.send( 'sudo -i\n' )
        # wait to input password
        buff = ''
        while not buff.endswith(': '):
            resp = self.rootshell.recv(65535)
            buff +=resp
        self.rootshell.send( '%s\n' % self.password )
        # wait login into root account
        while not buff.endswith('%s ' % self.root_prompt):
            resp = self.rootshell.recv(65535)
            buff +=resp
        if self.rootshell.send_ready():
            print "[INFO] Su to root account successfully!"
            return True
        else:
            print "[ERROR] Su to root account Failed!!!"
            return False

    def reboot(self, rebootcmd='reboot -n -f'):
        if self.rootshell == None:
            self.su_root()
        print "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-="
        print "Rebooting Host: %s , Command is: %s " % ( self.host_ip,rebootcmd )
        print "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-="
        self.rootshell.send( rebootcmd )
        self.rootshell.send('\n')
        start_time = datetime.datetime.now()
        print "... ... ... ... ... ... ... ..."
        time.sleep(30)
        self.close()

        

        res = False
        for i in range(10):
            res = self.check_online()
            print "... ... ... ... ... ... ... ..."
            time.sleep(10)
            if res:
                res = True
                now_time = datetime.datetime.now()
                print "[INFO] Reboot Used: %d s" % (now_time - start_time).seconds
                return res
            else:
                continue
        return res

    def send_cmd(self, cmd_list, timeout=2 ):
        ssh_stdin, ssh_stdout, ssh_stderr = self.channel.exec_command(cmd_list)
        ssh_stdin.write('\n')
        ssh_stdin.flush()
        output = ssh_stdout.read()
        return output

    def sudoExectueCommand(self, cmdList, verbose=False):
        if len(cmdList) == 0:
            return
        if self.rootshell == None:
            self.su_root()

        for cmd in cmdList:
            while not self.rootshell.send_ready():
                time.sleep(1)
            print "[Commands] %s" % cmd
            self.rootshell.send( cmd )
            self.rootshell.send('\r')
            buff = ''
            while not buff.endswith('%s ' % self.root_prompt):
                resp = self.rootshell.recv(9999)
                buff +=resp
            print "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-="
            print buff
            print "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-="
        return buff

    def exectueCommand(self, cmdList, verbose=False):
        if len(cmdList) == 0:
            return
        #startTimeStamp = time.time()

        fh = open(self.logfile, 'a+')
        print self.Flag*self.numofFlag
        for cmd in cmdList:
            if verbose:
                #print self.Flag*self.numofFlag
                print "[Commands] %s" % cmd
                #for subcmd in cmd.split(";"):
                #    print subcmd
                #print self.Flag*self.numofFlag
            fh.write(self.Flag*self.numofFlag + "\n")
            fh.write("\nExectuing Command: %s\n" % cmd)

            stdin,stdout,stderr = self.channel.exec_command(cmd)
            out = stdout.readlines()
            if verbose:
                for i in out:
                    print i.rstrip()
                    fh.write(i)
                print self.Flag*self.numofFlag

        fh.close()
        #print "Used Time:", time.time() - startTimeStamp, "s"

    def upload(self, localFile, remoteFile, verbose=False ):
        #startTimeStamp = time.time()
        #print "[Upload] Uploading LocalFile=",localFile," To RemoteFile=",remoteFile
        print "[Upload] Uploading LocalFile=%-90s To RemoteFile=%-80s " % (localFile,remoteFile)
        if verbose:
            self.sftp.put(localFile, remoteFile, self._callback)
        else:
            self.sftp.put(localFile, remoteFile)
        #print "Used Time:", time.time() - startTimeStamp, "s"

    #sftp.get(remotepath, localpath)
    def download(self, remoteFile, localFile, verbose=False):
        print "[DownLoad] Downloading RemoteFile=",remoteFile," To LocalFile=",localFile
        #startTimeStamp = time.time()
        if verbose:
            self.sftp.get(remoteFile, localFile, self._callback)
        else:
            self.sftp.get(remoteFile, localFile)
        #print "Used Time:", time.time() - startTimeStamp, "s"

    def _callback(self,a,b):
        # http://eleveni386.7axu.com/index.php/2012/09/10/python-%E8%BF%9B%E5%BA%A6%E6%9D%A1/
        #sys.stdout.write('Data Transmission %.2f M [%3.2f%%]\r' %(a/1024./1024.,a*100./int(b)))
        #if (int(a/1024./1024.) % 2 == 0 and int(a/1024./1024. *100) % 5 == 0 ):
        if (int(a*100./b) % 10 == 0 and int(a*1000/b) % 10 < 2):
            sys.stdout.write('|%-40s| %.2f M [%3.2f%%]\r' %('>'*int(a*40./b), a/1024./1024.,a*100./int(b)))
            sys.stdout.flush()

    def close(self):
        #if self.channel.is_active():
        self.sftp.close
        self.channel.close()

    def __enter__(self):
        return self

    def __exit__(self, type, value, tb):
        self.sftp.close
        self.channel.close()
        self.rootshell.close()

def test(HostIp):
    CurrentPath=os.path.dirname(os.path.realpath(__file__))
    Commands = ['pwd','who|grep nodeb','free|grep Mem']
    t = sshConnect('root','***',HostIp,22)

    #print t.is_active()
    #print t.create_shell()

    #print t.send_cmd('free|grep Mem;users;whoami')
    #t.reboot()
    #self.expect(["SDCxM-SDCAM-1-root"])
    t.exectueCommand(["whoami"],verbose=True)
    #t.su_root()
    t.exectueCommand(["docker ps"],verbose=True)
    #t.sudoExectueCommand(["whoami;pwd","uptime","ifconfig","uname -a;users","which python;python -V"],verbose=True)
    #t.sudoExectueCommand(["ls /store","ls -l /etc/ssh/sshd_config"],verbose=True)
    t.upload(os.path.join(CurrentPath,"check.sh"), "/tmp/check.sh", verbose=True)

    t.exectueCommand(["bash /tmp/check.sh"],verbose=True)
    #if t.reboot():
    #    print "[INFO]Reboot Successfully!"
    #    t.exectueCommand(["whoami"],verbose=True)
    #    t.exectueCommand(["uptime"],verbose=True)

    #t.exectueCommand(Commands,verbose=True)
    #t.download("/localdisk/Store/Temp/moc.jpg",os.path.join(currentPath,"moc.jpg")
    t.close()


def test_prd():
    t = sshConnect('admin', '******', '10.139.**.*', 22)
    t.exectueCommand(["whoami"],verbose=True)
    #t.exectueCommand(["docker ps"],verbose=True)
    t.sudoExectueCommand(["docker ps","whoami", "bash /tmp/check.sh"],verbose=True)
    t.close()

def test_test():
    t = sshConnect('root', '****', '10.253.*.***', 22)
    if not t.check_online():
       return False
    t.connect()
    t.exectueCommand(["whoami"],verbose=True)
    #t.exectueCommand(["docker ps"],verbose=True)
    t.sudoExectueCommand(["docker ps","whoami"],verbose=True)
    t.close()




if __name__ == '__main__':
    HostIp="10.253.*.***"
    if len(sys.argv) > 1:
       HostIp=sys.argv[1]

    #r=pyping.ping('10.139.*.***', timeout=10, count=3, packet_size=55, own_id=None, quiet_output=False, udp=False)
    #print dir(r)
    #print r.__hash__
    test_test()
    print "Over!"
