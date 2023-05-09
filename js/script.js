// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"


/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {
  document.getElementById('answer').innerHTML = "Hello, World!"
  // input
  const lengthA = parseFloat(document.getElementById('side-a').value)
  const lengthB = parseFloat(document.getElementById('side-a').value)
  const lengthC = parseFloat(document.getElementById('side-a').value)

  // unsing the cosin law
  const angleA = Math.acos((lengthB**2 + lengthC**2 - lengthA**2) / (2 * lengthB * lengthC)) * (180/Math.PI)
const angleB = Math.acos((lengthC**2 + lengthA**2 - lengthB**2) / (2 * lengthC * lengthA)) * (180/Math.PI)
const angleC = Math.acos((lengthA**2 + lengthB**2 - lengthC**2) / (2 * lengthA * lengthB)) * (180/Math.PI)
  
  const sumOfAngles = Number((angleA).toFixed(2)) + Number((angleB).toFixed(2)) + Number((angleC).toFixed(2))

  document.getElementById('answer').innerHTML = angleA + "° " + angleB + "° " + angleC + "°" + sumOfAngles + "°"


  if (sumOfAngles == 180) {
    if (angleA != angleB || angleA != angleC || angleB != angleC) {
      if (angleA == 90 || angleB == 90 || angleC == 90) {
        document.getElementById('answer').innerHTML = "This is an right angle triangle"
      } else {
        document.getElementById('answer').innerHTML = "This is an scalene triangle"
      } if (angleA == angleB && angleA == angleC) {
        document.getElementById('answer').innerHTML = "This is an equilateral triangle"
      } else {
        if (angleA == 90 || angleB == 90 || angleC == 90) {
          document.getElementById('answer').innerHTML = "This is an right angle triangle"
        } else {
          document.getElementById('answer').innerHTML = "This is an isosceles triangle"
        }
      }
    } 
  } else { 
      document.getElementById('answer').innerHTML = "Please enter a valid triangle"
  }
}