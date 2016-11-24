+++
author = ""
comments = true
date = "2016-11-24T23:05:23+00:00"
description = ""
draft = true
featured = false
image = "/forestryio/images/treasure-chest-1558932.jpg"
share = true
slug = "umbraco-simple-caching"
tags = ["umbraco", "cache", "programming"]
title = "Umbraco Simple Caching"

+++
I once looked into a pretty straightforward, very recent, website that took more than 5 seconds for the {{< exLink "TTFB" "https://en.wikipedia.org/wiki/Time_To_First_Byte" >}} from Azure. I created a default controller for Umbraco and set the ASP caching on: TTFB dropped to 25 m!

ince then