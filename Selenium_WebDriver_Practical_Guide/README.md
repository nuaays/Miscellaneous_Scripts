


## Selenium Webdriver Practical Guide
* http://techbus.safaribooksonline.com/book/web-development/9781782168850


## For Testing iOS and Android Apps
*  Available Tools
```
In order to automate the testing of your applications on mobile devices, there are many software tools available. The following are some of the tools that are built based on Selenium WebDriver:

1. AndroidDriver: This driver is a direct implementation of WebDriver, which is similar to FirefoxDriver, IEDriver, and so on. It acts as the client library with which your test script interacts. Its server side is the AndroidWebDriver that is installed on the device, or the emulator and executes all the test script commands that gets forwarded from AndroidDriver.

2. iPhoneDriver: This driver works very similar to AndroidDriver, but only on iOS platforms. In order to use it, you need to set up a server on the simulator or on the device. iPhoneDriver, however, is no longer supported and is deprecated.

3. iOSDriver: As the name says, this driver is used for automating native, hybrid, and m.site applications on iOS platforms. It uses native UI Automation libraries to automate on the iOS platform. For the test scripts, all this is transparent because it can still continue to use the WebDriver API in its favorite client language bindings. The test script communicates with the iOS Driver using the JSON wire protocol. However, if you want to execute your test scripts against the Android platform, you cannot use this driver.

4. Selendroid: This driver is similar to iOSDriver and can execute your native, hybrid, and m.site application test scripts on the Android platform. It uses the native UI Automator library provided by Google. The test scripts communicate with the Selendroid driver over the JSON wire protocol while using its favorite client language bindings.

5.Appium: This is another tool that can let you execute your test scripts against Android and iOS platforms without your having to change the underlying driver. Appium can also work with Firefox OS platforms. In the rest of the chapter, we will see how we can work with Appium.
```
