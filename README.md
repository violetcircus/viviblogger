# viviblogger
i'm writing this so i don't have to rely on something someone else made to convert markdown files i write into blog posts for my website. it's probably going to be jank and exclusively tailored to my use-case. hooray!

it works now! make sure to set up your config file before use in `~/.config/vvblogger/`. keep in mind this program is very idiosyncratic and largely exists to facilitate one thing I want to do - there's far more flexible/practical/generalised options out there, but I wanted to do something from scratch with no external libraries for fun and the sense of pride in using tools largely created by and for myself.

## config layout:
every item in the config is stored as "field=value", with no quotes around any value. extraneous spaces are trimmed before the value is parsed.
- SiteDir: directory containing your website files.
- PostsDir: subdirectory containing your website's posts. will be added onto the end of "SiteDir", so it assumes the format "posts/" as shown in the default value.
- ImageDir: where your website images are stored. same principle as the posts directory.
- SourceImageDir: where the site should look for images referenced in your markdown files to be copied to the site image folder.
- TemplateFile: the html template file path. this one should be a full path - i thought you should be able to put this wherever you want.
- DateTimeFormat: this is a string Go will use to format the time and date for tracking uploaded/created date in your blog posts. Go's reference time is the second of january 2006 at 03:04:05PM, so rearrange that into whatever format the dates in your notes use. It's a little odd, but it's not too complicated.

## usage
after filling out your config file simply run viviblogger FILE.md and it'll process it and stick it in your posts directory.

## installation instructions:
```
git clone https://github.com/violetcircus/viviscrobbler
cd viviscrobbler
go install
```

if you run into any issues, please let me know, but again: this program is like, mostly for me. I don't intend to support every possible use-case.
