# Hear Pete Speak!

Hear Pete Speak is a catalogue of Mayor Pete's interviews. [Check out the site](https://hearpetespeak.com)

## Contributing

### Content

First and foremost, Hear Pete Speak is valuable because of its content! If you're interested in contributing and a little familiar with GitHub, read on. (We'll be thinking of a system that works well for users without GitHub accounts)

Currently, all interview data is stored in `yaml` files in the `/data` directory. The file naming scheme is `yearmonthday-interview-name.yaml`.

Each interview file has two required fields: 

* `title`: a name for the overall interview, often the name of the show / podcast / interviewer
* `date`: the date the interview took place. The format is `{year}-{month}-{day}`.

Here's an example:

```yaml
title: The Van Jones Show
date: 2019-03-30
```

Then we can include each of the individual topics discussed, and include direct links to those moments in the interview. An individual `topic` has three components: 

* `categories`: this is the top-level category . e.g. `Healthcare`, `Democracy`. The category here _must_ be one listed in [the category definitions file](https://github.com/pH14/hearpetespeak/blob/master/data/categories.yaml). If the category you're looking for isn't listed there, feel free to add it or file an issue and we can discuss!
* `direct_link`: this is a link to the content with a timestamp. YouTube, Facebook and Overcast.fm (for podcasts) allow you to share links that jump directly to certain moments in the video or audio, which is strongly preferable.
* `tags`: these are the descriptors for the link so people can find the discussions they're most interested in. We think it's best if the language on the tags comes from Pete's own language in the interview.

Let's say we created a file named  `20190330-van-jones.yaml` that contained:

```yaml
title: The Van Jones Show
date: 2019-03-30

topics:
  -
    categories: "Democracy"
    direct_link: "https://youtu.be/PFZgbDNaEko?t=1033"
    tags:
        - Electoral College
        - structural remedies
        - Constitution
  -
    categories: "Values"
    direct_link: "https://youtu.be/PFZgbDNaEko?t=1480"
    tags:
        - faith
        - humility
        - community
        - relationships
        - trust
```

Here's what would be shown on the site:

![example of van jones interview topic under Democracy header](https://github.com/pH14/hearpetespeak/blob/master/contributing/vanjones_democracy_example.png)

![example of van jones interview topic under Value header](https://github.com/pH14/hearpetespeak/blob/master/contributing/vanjones_values_example.png)

As a method of coordination, you can file an Issue first saying that you're working on a particular interview.

### Code

HPS is a simple static site built with [Hugo](https://gohugo.io/). It's hosted on Amazon S3 and uses Cloudflare CDN. The site's intent is to work well on spotty Internet connections, browsers old and new, devices large and small, some with displays and others not (screenreaders).

If you're interested in contributing, check out the open [code-related Issues](https://github.com/pH14/hearpetespeak/issues?q=is%3Aopen+is%3Aissue+label%3Acode) or file your own.

### Feature Requests / Suggestions / Bugs

Anything bugging you about the site? [Create an Issue](https://github.com/pH14/hearpetespeak/issues/new)

### Design Goals

We want to make Pete's content as universally accessible as possible. Timestamped links, accurate tags, fast page loads, and simple page structure are all part of this.
