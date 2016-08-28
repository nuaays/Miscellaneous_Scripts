import os, sys, time, subprocess, re
import shutil, zipfile, tempfile
import requests
import json
import pprint

def exectue_external_cmd( cmd,  msg_in='' ):
    proc = subprocess.Popen( cmd,  shell = True,  stdin = subprocess.PIPE,  stdout = subprocess.PIPE,  stderr = subprocess.PIPE)
    stdout_info,  stderr_info = proc.communicate( msg_in)
    return stdout_info





if __name__ == "__main__":
    hostIp = "localhost"
    if len(sys.argv) > 1:
       hostIp = sys.argv[1]
    flag = "="
    flag_number = 50
    images_kept_number = 3
    


    images_restful_api = "http://%s:2375/images/json" % hostIp
    images_delete_restful_api = "http://%s:2375/images/" % hostIp
    containers_restful_api = "http://%s:2375/containers/json?all=1" % hostIp
    containers_delete_restful_api = "http://%s:2375/containers/" % hostIp
    try:
       res = requests.get(images_restful_api, timeout=5)
    except Exception, e:
       print "[ERROR] Cannot reach to %s" % hostIp
       sys.exit() 
    

    containers_all = requests.get(containers_restful_api, timeout=5) #requests.get('http://localhost:2375/containers/json?all=1')    
    containers_running_json = json.loads(containers_all.text)
    """
    #container which is exited
    containers_exited_uuids = [ i["Id"] for i in containers_running_json if "Exited" in i["Status"]]
    #clean stopped containers
    print "Begin to clean stop containers %d ... ..." % len(containers_exited_uuids)
    for c in containers_exited_uuids:
        break
        deleteurl = containers_delete_restful_api + c
        print "[Container Exited]", deleteurl
        r = requests.delete(deleteurl)
        if r.status_code == 204:
           print "Success delete container %s" % c
        else:
           print "Failed to delete container %s " % c
           print r.text
    """
    #images used by container
    images_used_ids = list(set([ i["ImageID"] for i in containers_running_json]))
    images_used_tags= list(set([ ':'.join(i["Image"].split(":")[:-1]) for i in containers_running_json]))
    #print images_used_ids
    
    #images_all = json.loads( requests.get('http://localhost:2375/images/json').text)
    images_all = json.loads( requests.get(images_restful_api, timeout=5).text )
    images_repotags = []
    images_info = {}
    images_candeleted_ids=[]
    images_none = []
 
    for i in images_all:
        #print flag*flag_number
        #print i
        if i["RepoTags"] == "<none>:<none>":
            print "[NONE]", i["RepoTags"], i["Id"]
            images_none.append(i["Id"])
            continue
        for img in i["RepoTags"]:
            image_name = ":".join(img.split(":")[:-1])
            image_tag  = img.split(":")[-1]
            if image_name not in images_info:
               images_info[image_name] = []
            images_info[image_name].append({"tag":image_tag, "created":i["Created"], "Id":i["Id"] })

        images_repotags.extend(i["RepoTags"])
        
       
        if i["Id"] in images_used_ids:
           #print "Found!!!!!!" ,i["RepoTags"], i["Id"]
           pass
        else:
           #print "Can Deleted", i["RepoTags"], i["Id"]
           images_candeleted_ids.append(i["Id"])

    #
    print flag*flag_number, "\n Containers List on %s" % hostIp 
    pprint.pprint([ c["Id"] for c in containers_running_json ])
    #       
    print flag*flag_number, "\n Images Used: ", len(images_used_ids)
    pprint.pprint( images_used_tags)
    #
    print flag*flag_number, "\n Images List: ", len(images_repotags)
    pprint.pprint(sorted(images_repotags))
    #
    print flag*flag_number, "\n"



    #
    for img, tags in images_info.items():
        if img in images_used_tags:
           #img is in used, keep the latest 3 images
           tags_deleted = sorted(tags, key=lambda x:x["created"], reverse=True)[images_kept_number:]
        else:
           #img is not used current
           tags_deleted = sorted(tags, key=lambda x:x["created"], reverse=True)
        if len(tags_deleted):
           print flag*flag_number
           print "Begin to clean %s ...\n" % img 
        for tag in tags_deleted:
            if tag["Id"] in images_used_ids:
               continue
            repotag = "%s:%s" % (img, tag["tag"])
            #deleteurl = "http://localhost:2375/images/%s" % repotag
            deleteurl = images_delete_restful_api + repotag
            deleteurl = images_delete_restful_api + tag["Id"]
            print "[%s] curl -X DELETE %s ... " % (repotag, deleteurl)
            r = requests.delete(deleteurl)
            if r.status_code == 200:
               print "Success to delete %s" %  tag["Id"]
            else:
               print "Failed to delete %s" % deleteurl
               print r.text
         #clean <none> images
        if len(images_none):
           print "\nBegin to clean <none> ... ...\n"
           for uuid in images_none:
               print "[NONE]", uuid
               deleteurl = images_delete_restful_api + uuid
               r = requests.delete(deleteurl)
               if r.status_code == 200:
                  print "Success to delete %s " % uuid
               else:
                  print "Failed to delete %s " % uuid
                  print r.text
   
    #
    images_after_clean = json.loads( requests.get(images_restful_api, timeout=5).text )
    print "[%s]Before Clean Images Number:" % hostIp, len(images_all)
    print "[%s]After  Clean Images Number:" % hostIp, len(images_after_clean)
    print "\n\n\n", flag*flag_number, "\n\n\n"
   
 
               
     

    
    
