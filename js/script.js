// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: Sep 2020
// This file contains the JS functions for index.html

"use strict"

if (navigator.serviceWorker) {
  navigator.serviceWorker.register("/ICS2O-Unit-6-03-Dominic-Madeira/sw.js", {
    scope: "/ICS2O-Unit-1-08-Dominic-Madeira/",
  })
}

const getImage = async (URLAddress) => {
  try {
    const result = await fetch(URLAddress)
    const jsonData = await result.json()
    console.log(jsonData)
    document.getElementById("api-image").innerHTML =
    '<img src="' + 
      jsonData.url + 
      '" alt="API image" class="center" ' +
      '>'
    if (jsonData.main.temp != "none") {
    
  } else {
    document.getElementById("image-artist").innerHTML = "<p>Artist: unknown</p>"
  }
  if (jsonData.weather.icon != "none") {
    document.getElementById("image-date").innerHTML = "<p>Date: unknown</p>"
  }
  } catch (err) {
    console.log(err)
  }
}

getImage("https://api.catboys.com/img")
