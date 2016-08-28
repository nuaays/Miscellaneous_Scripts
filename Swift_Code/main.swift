//
//  main.swift
//  HelloWorld
//
//  Created by senya on 16/2/29.
//  Copyright © 2016年 senya. All rights reserved.
//

import Foundation
import AppKit


print("Hello, YangSen!")
NSLog("Hello NSLog, Calling %@", #function)
//NSLog("backtrace: %@", NSThread.callStackSymbols())

var greeting: String = "Hello World"
print(greeting)

let maxNumberOfLoginAttempts = 10
var currentLoginAttempt = 0

var x = 1, y = 2, z = 3
let myInt : Int = 42


print( UInt8.max )

let three = 3
let pointNumber = 0.14159
let pi = Double(three) + pointNumber
print(pi)
let integerPi = Int(pi)
print(integerPi)


let isPassWordCorrect = false
let areYouChinese = true
if isPassWordCorrect {
    print("欢迎！")
} else {
    print("对不起！")
}

var varableString = "Hello,"
varableString += " Swift"

let yanSign: Character = "¥"

var emptyString = ""
var anotherEmptyString = String()
emptyString += varableString
anotherEmptyString.appendContentsOf(varableString)

print(emptyString)
print(anotherEmptyString)
    

if emptyString.isEmpty {
    print("Nothing to see here!")
}



//let age = -3
//assert(age >= 0, "年龄不能小于0")

var shoppingList :[String] = ["Eggs", "Apple", "Milk", "123"]
//var shoppingList = [ "Eggs", "Apple", "Milk", 123, true]
for item in shoppingList {
    print("shoppingList中包含:\(item)")
}

shoppingList.insert("Mapple", atIndex:0)
shoppingList.append("Flour")
shoppingList += ["Chocolate", "Cheese"]
shoppingList.removeLast()
print(shoppingList.count)
print(shoppingList)
print(shoppingList[4...6])
print("=======================")

for item in shoppingList {
    print(item)
}
for (index, value) in EnumerateSequence(shoppingList){
    print("Item \(index+1) : \(value)")
}











