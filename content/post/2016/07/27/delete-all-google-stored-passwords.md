+++
author = ""
comments = true
date = "2016-07-27T22:02:00+00:00"
description = "How to delete all the passwords stored in Google passwords https://passwords.google.com"
draft = false
featured = false
image = "/forestryio/images/Screenshot from 2016-07-27 21:47:45-1.png"
share = true
slug = "delete-google-passwords"
tags = ["google", "security", "password"]
title = "Delete all google stored passwords"

+++

I was trying to delete all the passwords stored in google at {{< exLink "Google passwords" "https://passwords.google.com" >}}, way too many to do manually!

Couldn’t find a solution online so here is my easy one:

*   Go to {{< exLink "Google passwords" "https://passwords.google.com" >}}, sign in to your account.
*   Open the development tools in your browser, usually with F12\. I’ve used Chrome.
*   Paste the following code in the console.

{{< highlight javascript >}}

// Gets all the elements to be clicked to delete the passwords.

var a = document.getElementsByClassName(“dga”)

// Click all the elements.

for(i=0; i < a.length; i++){a[i].click();}

{{< /highlight >}}
