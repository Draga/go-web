+++
author = ""
comments = true
date = "2016-05-21T20:51:00+00:00"
description = "How to write content, deploy automatically and host your website"
draft = false
featured = true
image = "/images/hugo.jpg"
share = true
slug = "hugo-write-deploy-host"
tags = ["hugo", "git", "github", "deployment", "nginx", "pagespeed"]
title = "Hugo: write, deploy, host"
[menu]

+++
In this port I'll show you how I write my posts, deploy them to the server, process them with {{< exLink Hugo "https://gohugo.io" >}} and host them on my {{< exLink DigitalOcean "https://m.do.co/c/875cd23a5c97" >}} server.

I wanted this process to be as easy and quick as possible, with as few extra softwares and services as possible. There are a lot of solutions out there like Wercker but I didn't want more things in my way.

# Writing
## Local

I write the posts from IntelliJ IDEA, which is my IDE of choice. Screw Sublime, Atom and all the rest! It's quite heavier but is worth it. You get spell checking, quite good live previews with the Multimarkdown plugin (shame the embedded HTML doesn't render), git integration, and in general a powerful IDE.

Just to be clear I like Sublime and Atom, but every time I use them there is something wrong. Right now the terminal in Atom doesn't work for me. There is an open issue for this but in the meantime you get a blinking unmovable cursor staring at you.

I feel that, since Hugo is quite new, the tools haven't caught up with its popularity yet. Intellij is flexible enough to allow you to have shortcuts ({{< exLink "live templates" "https://www.jetbrains.com/help/idea/2016.1/live-templates.html" >}}) for your shortcodes. E.g.:  
{{< highlight xml >}}
  <template name=".shortcode" value="{{&lt; $SHORTCODE$ &gt;}}" description="A Hugo shortcode" toReformat="false" toShortenFQNames="true">
    <variable name="SHORTCODE" expression="" defaultValue="" alwaysStopAt="true" />
    <context>
      <option name="OTHER" value="true" />
    </context>
  </template>
{{< /highlight >}}
This allows me to type `.s`, press enter (if the live template `.shortcode` is selected), type the name of the shortcode and its parameters and press enter to keep typing without having to type too many symbols or having to move your cursor.

I keep switching between IntelliJ and the localhost page on my browser, where my hugo server is showing how the end result will be (just launch "hugo server" from the root folder). If you have the page open you don't even need to refresh it thanks to Hugo's live reload feature (and don't even have to save the file since IntelliJ does it when it loses focus). This is very quick to do, which is particularly nice when you want to make final tweaks to make sure your post will display as you want.

## Online
If you are using Github and want to write from your browser, or you are not at home, you can use {{< exLink "prose.io" "https://prose.io" >}}. It's a software developed for Jekyll, but works just fine with Hugo. You get some nice shortcuts to format your text with markdown in case you are not too familiar with it. It also allows you to upload images which are committed to your repo immediately. I don't find it quite as powerful as IntelliJ but it's a nice option.

You can have a look at {{< exLink "my prose configuration" "https://github.com/Draga/go-web/blob/master/_prose.yml" >}} to get you started. It's important to know that prose.io can only generate content with yaml metadata. You will have to configure Hugo to do the same by adding `MetaDataFormat: yaml`.

I've looked at many more free online content editor:
<ul>
    <li>{{< exLink sofish "http://sofish.github.io/pen/" >}}</li>
    <li>{{< exLink webhook "http://www.webhook.com/" >}}</li>
    <li>{{< exLink contentful "https://www.contentful.com" >}}</li>
    <li>{{< exLink prismic "https://prismic.io/" >}}</li>
    <li>{{< exLink stackedit "https://stackedit.io" >}}</li>
    <li>{{< exLink dillinger "http://dillinger.io/" >}}</li>
</ul>
I found them actually extremely good...just not very compatible with Hugo.

Unfortunately I couldn't find any online resource to write and preview Hugo content, so editing content locally and preview it with Hugo server is still the best option for me. I reckon you could host a test environment to preview your draft, deploying with a different branch, by adding `-buildDrafts` to your `hugo` or `hugo server` commands.

Hopefully some tool that works perfectly with Hugo will come out soon. Spf13, the creator of Hugo, {{< exLink "is thinking about creating one" "https://discuss.gohugo.io/t/web-based-editor/155/41" >}}.

# Deployment
I can deploy my website in 2 ways:

- Push you branch to the repo on the server.
- Push the master branch to the main repo. In my case this is on Github but can work with pretty much anything else.

These both trigger a small script which I've adapted from {{< exLink "this tutorial on Digitalocean" "https://www.digitalocean.com/community/tutorials/how-to-deploy-a-hugo-site-to-production-with-git-hooks-on-ubuntu-14-04" >}}.

{{< highlight bash >}}
#!/bin/bash

# Better be specific rather than using $HOME because it wouldn't work with all users.
GIT_REPO=/home/draga/go-web
WORKING_DIRECTORY=$GIT_REPO/public
PUBLIC_WWW=/home/draga/public_html
BACKUP_WWW=/home/draga/backup_html

set -e

# Clean working dir.
rm -rf $WORKING_DIRECTORY
# Backup current website. Use -z if moving across network to save bandwidth.
rsync -aq $PUBLIC_WWW/ $BACKUP_WWW
# Restore backup if an error occurs from now on.
trap "echo 'A problem occurred.  Reverting to backup.'; rsync -aq --del $BACKUP_WWW/ $PUBLIC_WWW; rm -rf $WORKING_DIRECTORY" EXIT

# fetch repo to make sure it's up to date if working with remote webhook.
cd $GIT_REPO
git fetch
git checkout origin/master
# Build the website. GO HUGO!
/usr/bin/hugo -s $GIT_REPO -d $WORKING_DIRECTORY
# Mirror the output of Hugo to the directory where the website is hosted.
# Use -c to copy only files that have changed to avoid messing with webhost caches.
# Use -z if moving across network to save bandwidth.
rsync -aqc --delete $WORKING_DIRECTORY/ $PUBLIC_WWW
# Ask Google webmaster tools to parse your sitemap.
curl https://google.com/webmasters/sitemaps/ping?sitemap=https://stefano.chiodino.uk/sitemap.xml
curl https://www.bing.com/ping?sitemap=https%3A%2F%2Fstefano.chiodino.uk/sitemap.xml
# Disable trap.
trap - EXIT
{{< /highlight >}}

I used the method described in the tutorial above for a while, but it bothered me to push twice since I had to type the password for my ssh key (I've disabled the password access to the server and can now only be accessed with the key). Also what if I wasn't home or wanted to use prose.io?

Solution: git(hub) webhooks!

Thanks to <a href="https://github.com/adnanh/webhook" target="_blank">github.com/adnanh/webhook</a>, little software written in go, you can have you server listen on a port and execute a script when invoked.

I currently have it connecting in https, with a secret key, and executing the script when the master branch is pushed.

Now I can even use Github itself to edit my posts. Pretty neat!

# Hosting
Now this was the ~~long~~ fun part. I wanted it all so it took a lot of learning, hence the fun.

I heard of nginx a lot and some research showed that it seems to be the go-to solution. Here are some, more performance related, details of how it performs for me: <a href="{{< relref "server-load-test.md" >}}">server load test</a>.

I've then enabled the pagespeed module which is very complex and heavy, but definitively worth it!
The website now scores <a href="https://developers.google.com/speed/pagespeed/insights/?url=https%3A%2F%2Fstefano.chiodino.uk%2F" target="_blank">100/100 in Google PageSpeed</a>, both in mobile and desktop!

I score 99/100 in <a href="https://gtmetrix.com/reports/stefano.chiodino.uk/FsGkxV4j" target="_blank">GTmetrix</a> just because there is a bug in the pagespeed module ({{< exLink 1064 "https://github.com/pagespeed/ngx_pagespeed/issues/1064" >}}) and the very header is added twice, and for some reason GTmetrix seems to think that is not there...? Or maybe is because pagespeed seems to add it on subsequent reloads. Another small penalty is for the G analytics code being cached for just 2 hours, it's out of my control and don't think is worth self hosting it.

YSlow seems to be mostly concerned of the lack of CDN.

You can find all the scripts that I've used to set up my server on {{< exLink "github" "https://github.com/Draga/serverScripts" >}}. They are ordered in the same way they are in the filesystem, so you should get the idea. There is a script to install nginx with the brotli and pagespeed module, directives to use webhook as a service, the entire nginx configuration, etc.
