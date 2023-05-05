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
  // input
  const lengthA = parseInt(document.getElementById('side-a').value)
  const lengthB = parseInt(document.getElementById('side-a').value)
  const lengthC = parseInt(document.getElementById('side-a').value)

  // unsing the cosin law
  const angleA = Number (Math.acos((lengthB**2 + lengthC**2 - lengthA**2) / (2 * lengthB * lengthC)) * (180/Math.PI)).toFixed(2)
  const angleB = Number (Math.acos((lengthC**2 + lengthA**2 - lengthB**2) / (2 * lengthC * lengthA)) * (180/Math.PI)).toFixed(2)
  const angleC = Number (Math.acos((lengthA**2 + lengthB**2 - lengthC**2) / (2 * lengthA * lengthB)) * (180/Math.PI)).toFixed(2)


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