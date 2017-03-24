##adb命令

```
adb connect xxx:5037
adb start-server
adb kill-server
adb remount
```

* adb devices

* adb root
* adb -s XXX shell
```
adb -s xx shell 'xxx'
adb -s 724e5658 shell logcat -b radio
adb -s 724e5658 shell cat /proc/meminfo
adb -s 724e5658 shell cat /proc/cpuinfo
adb -s 724e5658 shell cat /system/build.prop
	ro.product.model=MI-ONE Plus
	ro.product.brand=Xiaomi
	ro.product.name=mione_plus
	ro.product.device=mione_plus
	ro.product.board=MIONE
	ro.product.cpu.abi=armeabi-v7a
	ro.product.cpu.abi2=armeabi
	ro.product.manufacturer=Xiaomi
	ro.board.platform=msm8660
	ro.build.product=mione_plus
	ro.build.description=mione_plus-user 4.0.4 IMM76D ICS24.0 release-keys
	ro.build.fingerprint=Xiaomi/mione_plus/mione_plus:4.0.4/IMM76D/ICS24.0:user/release-keys
adb -s 724e5658 shell ls /system/lib
``

* 获取系统版本
```
adb shell getprop ro.build.version.release
4.0.4
```

* 获取系统api版本
```
adb shell getprop ro.build.version.sdk
15
```

* kill -3 $PID
```
/data/anr/trace.txt
```

* adb logcat
```
senya:~ root# adb logcat -g
/dev/log/main: ring buffer is 1024Kb (538Kb consumed), max entry is 4096b, max payload is 4076b
/dev/log/system: ring buffer is 256Kb (255Kb consumed), max entry is 4096b, max payload is 4076b

adb logcat -v time -v process -v brief 
adb logcat WifiHW:D *:S
adb logcat | grep Wifi
```

* adb pull/push

```
adb pull /data/app/xxxx.apk .
adb push xxxx /data/app/xxx.apk
adb install -r xxx.apk
```


* adb get-product 

* adb get-serialno

* adb bugreport

* adb screencap

```
adb shell screencap -p /data/screen.png
adb pull /sdcard/screen.png
adb shell rm /sdcard/screen.png

export LC_COLLATE='C'
export LC_CTYPE='C'
adb shell screencap -p | sed 's/\r$//' > screen.png
```

* battery and power

```
root@android:/ # dumpsys battery
Current Battery Service state:
  AC powered: false
  USB powered: true
  status: 2
  health: 2
  present: true
  level: 89
  scale: 100
  voltage:4163
  temperature: 240
  technology: Li-poly
```

* top
```
senya:~ root# adb shell top -n 1 -d 0.5 -s cpu
User 9%, System 45%, IOW 0%, IRQ 0%
User 1 + Nice 0 + Sys 5 + Idle 5 + IOW 0 + IRQ 0 + SIRQ 0 = 11

  PID PR CPU% S  #THR     VSS     RSS PCY UID      Name
19141  1  54% R     1   1044K    432K  fg root     top
    2  1   0% S     1      0K      0K  fg root     kthreadd
    3  0   0% S     1      0K      0K  fg root     ksoftirqd/0
    5  0   0% D     1      0K      0K  fg root     kworker/u:0



senya:~ root# adb shell top -m 10 -s cpu
User 2%, System 2%, IOW 0%, IRQ 0%
User 16 + Nice 0 + Sys 16 + Idle 579 + IOW 2 + IRQ 0 + SIRQ 0 = 613

  PID PR CPU% S  #THR     VSS     RSS PCY UID      Name
  389  0   1% S    85 561192K  61948K  fg system   system_server
19049  0   1% R     1   1044K    432K  fg root     top
  133  1   1% S    10  53804K  11408K  fg system   /system/bin/surfaceflinger
  653  0   0% S    17 488840K  62452K  fg app_28   com.miui.home
  506  0   0% S    19 493708K  64076K  fg system   com.android.systemui
18985  0   0% S     1      0K      0K  fg root     kworker/u:1
   73  0   0% S     1      0K      0K  fg root     irq/379-lis3dh
18908  0   0% S     1      0K      0K  fg root     kworker/u:3
16974  1   0% S     1      0K      0K  fg root     kworker/1:1
  542  1   0% S     1      0K      0K  fg root     dhd_dpc
参数含义：
PID  : progress identification，应用程序ID
S    : 进程的状态，其中S表示休眠，R表示正在运行，Z表示僵死状态，N表示该进程优先值是负数
#THR : 程序当前所用的线程数
VSS  : Virtual Set Size虚拟耗用内存（包含共享库占用的内存）
RSS  : Resident Set Size实际使用物理内存（包含共享库占用的内存）
PCY  : 前台(fg)和后台(bg)进程
UID  : User　Identification，用户身份ID
Name : 应用程序名称
```

* adb dumpsys

```
adb shell dumpsys meminfo <package_name>

senya:~ root# adb shell ps |grep intel
app_40    17533 134   466016 42888 ffffffff 400a84e0 S com.intel.senyang.myapplication
参数含义：
dalvik : dalvik使用的内存
native : native堆上的内存，指C\C++堆的内存（android 3.0以后bitmap就是放在这儿）
other  : 除了dalvik和native的内存，包含C\C++非堆内存······
Pss    : 该内存指将共享内存按比例分配到使用了共享内存的进程
allocated : 已使用的内存
free      : 空闲的内存
private dirty : 非共享，又不能被换页出去的内存（比如linux系统中为了提高分配内存速度而缓冲的小对象，即使你的进程已经退出，该内存也不会被释放）
share dirty   : 共享，但有不能被换页出去的内存
senya:~ root# adb shell dumpsys meminfo 17533
Applications Memory Usage (kB):
Uptime: 14224624 Realtime: 316325823

** MEMINFO in pid 17533 [com.intel.senyang.myapplication] **
                         Shared  Private     Heap     Heap     Heap
                   Pss    Dirty    Dirty     Size    Alloc     Free
                ------   ------   ------   ------   ------   ------
       Native     2037     1140     1996     5996     5407      100
       Dalvik     1731    13996     1276    12511    10960     1551
       Cursor        0        0        0
       Ashmem        2        4        0
    Other dev      292       32      288
     .so mmap      917     2024      284
    .jar mmap        0        0        0
    .apk mmap       97        0        0
    .ttf mmap        1        0        0
    .dex mmap     1140        0       28
   Other mmap      318       20       36
      Unknown     1795      324     1784
        TOTAL     8330    17540     5692    18507    16367     1651

 Objects
               Views:       14         ViewRootImpl:        1
         AppContexts:        3           Activities:        1
              Assets:        3        AssetManagers:        3
       Local Binders:        5        Proxy Binders:       13
    Death Recipients:        1
     OpenSSL Sockets:        0

 SQL
                heap:        0          MEMORY_USED:        0
  PAGECACHE_OVERFLOW:        0          MALLOC_SIZE:        0
```


* adb shell procrank

```
可以看到，在Linux下表示内存的耗用情况有四种不同的表现形式：
 VSS - Virtual Set Size 虚拟耗用内存（包含共享库占用的内存）
 RSS - Resident Set Size 实际使用物理内存（包含共享库占用的内存）
 PSS - Proportional Set Size 实际使用的物理内存（比例分配共享库占用的内存）
 USS - Unique Set Size 进程独自占用的物理内存（不包含共享库占用的内存）
senya:~ root# adb shell procrank
  PID      Vss      Rss      Pss      Uss  cmdline
 1260   78280K   78164K   36854K   31816K  com.qihoo.appstore
  578   71168K   71020K   33973K   30744K  com.sohu.inputmethod.sogou
  948   65344K   65256K   27901K   23736K  com.qihoo.daemon
  506   68428K   64148K   26349K   23688K  com.android.systemui
  653   67900K   62900K   26340K   24272K  com.miui.home
  389   62020K   61980K   25117K   22284K  system_server
  754   51796K   51656K   17432K   15536K  android.process.acore
18254   48324K   48156K   14109K   11548K  sogou.mobile.explorer.hotwords
17003   47188K   47108K   11333K    9548K  com.miui.gallery
VSS：VSS表示一个进程可访问的全部内存地址空间的大小。这个大小包括了进程已经申请但尚未使用的内存空间。在实际中很少用这种方式来表示进程占用内存的情况，用它来表示单个进程的内存使用情况是不准确的。
RSS：表示一个进程在RAM中实际使用的空间地址大小，包括了全部共享库占用的内存，这种表示进程占用内存的情况也是不准确的。
PSS：表示一个进程在RAM中实际使用的空间地址大小，它按比例包含了共享库占用的内存。假如有3个进程使用同一个共享库，那么每个进程的PSS就包括了1/3大小的共享库内存。这种方式表示进程的内存使用情况较准确，但当只有一个进程使用共享库时，其情况和RSS一模一样。
USS：表示一个进程本身占用的内存空间大小，不包含其它任何成分，这是表示进程内存大小的最好方式！
可以看到:VSS>=RSS>=PSS>=USS
```
