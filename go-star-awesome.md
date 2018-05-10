# Awesome Go

[![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://camo.githubusercontent.com/984828c0b020357921853f59eaaa65aaee755542/68747470733a2f2f73332e65752d63656e7472616c2d312e616d617a6f6e6177732e636f6d2f6e6774756e612f6a6f696e2d75732d6f6e2d736c61636b2e706e67)](http://gophers.slack.com/messages/awesome)

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Files](#files)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Game Development](#game-development)
    - [Generation and Generics](#generation-and-generics)
    - [Geographic](#geographic)
    - [Go Compilers](#go-compilers)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [IoT](#iot-internet-of-things)
    - [Logging](#logging)
    - [Machine Learning](#machine-learning)
    - [Messaging](#messaging)
    - [Miscellaneous](#miscellaneous)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
    - [Text Processing](#text-processing)
    - [Third-party APIs](#third-party-apis)
    - [Utilities](#utilities)
    - [Validation](#validation)
    - [Version Control](#version-control)
    - [Video](#video)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares)
            - [Actual middlewares](#actual-middlewares)
            - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
        - [Routers](#routers)
    - [Windows](#windows)
    - [XML](#xml)

- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
	- [Go Generate Tools](#go-generate-tools)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)

- [Server Applications](#server-applications)

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [EasyMIDI](https://github.com/algoGuy/EasyMIDI):star:9 - EasyMidi is a simple and reliable library for working with standard midi file (SMF).  *2018-03-22T05:16:53Z*
* [flac](https://github.com/eaburns/flac):star:72 - Native Go FLAC decoder.  *2017-10-03T20:06:20Z*
* [flac](https://github.com/mewkiz/flac):star:60 - Native Go FLAC decoder.  *2018-04-08T15:47:05Z*
* [gaad](https://github.com/Comcast/gaad):star:39 - Native Go AAC bitstream parser.  *2018-02-20T21:07:56Z*
* [go-sox](https://github.com/krig/go-sox):star:66 - libsox bindings for go.  *2018-03-29T15:43:18Z*
* [go_mediainfo](https://github.com/zhulik/go_mediainfo):star:16 - libmediainfo bindings for go.  *2015-12-24T20:44:59Z*
* [gosamplerate](https://github.com/dh1tw/gosamplerate):star:3 - libsamplerate bindings for go.  *2017-02-24T00:19:51Z*
* [id3v2](https://github.com/bogem/id3v2):star:60 - Fast and stable ID3 parsing and writing library for Go.  *2018-04-07T21:16:34Z*
* [malgo](https://github.com/gen2brain/malgo):star:22 - Mini audio library.  *2017-11-14T04:07:30Z*
* [minimp3](https://github.com/tosone/minimp3):star:8 - Lightweight MP3 decoder library.  *2018-03-26T09:00:42Z*
* [mix](https://github.com/go-mix/mix):star:75 - Sequence-based Go-native audio mixer for music apps.  *2017-06-24T18:22:10Z*
* [mp3](https://github.com/tcolgate/mp3):star:66 - Native Go MP3 decoder.  *2017-04-26T19:24:21Z*
* [music-theory](https://github.com/go-music-theory/music-theory):star:191 - Music theory models in Go.  *2016-05-20T05:08:32Z*
* [PortAudio](https://github.com/gordonklaus/portaudio):star:188 - Go bindings for the PortAudio audio I/O library.  *2017-08-06T14:41:49Z*
* [portmidi](https://github.com/rakyll/portmidi):star:163 - Go bindings for PortMidi.  *2017-07-16T03:23:45Z*
* [taglib](https://github.com/wtolson/go-taglib):star:56 - Go bindings for taglib.  *2015-11-14T19:47:54Z*
* [vorbis](https://github.com/mccoyst/vorbis):star:20 - "Native" Go Vorbis decoder (uses CGO, but has no dependencies).  *2017-10-14T15:28:25Z*
* [waveform](https://github.com/mdlayher/waveform):star:207 - Go package capable of generating waveform images from audio streams.  *2016-07-07T15:25:25Z*

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [authboss](https://github.com/volatiletech/authboss):star:0 - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.  *2017-11-13T16:33:40Z*
* [casbin](https://github.com/hsluoyz/casbin):star:0 - Authorization library that supports access control models like ACL, RBAC, ABAC.  *2018-04-22T05:00:05Z*
* [cookiestxt](https://github.com/mengzhuo/cookiestxt):star:0 - provides parser of cookies.txt file format.  *2017-10-09T12:14:37Z*
* [Go-AWS-Auth](https://github.com/smartystreets/go-aws-auth):star:197 - AWS (Amazon Web Services) request signing library.  *2017-05-04T20:50:21Z*
* [go-jose](https://github.com/square/go-jose):star:849 - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.  *2018-04-11T19:38:52Z*
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server):star:826 - Standalone, specification-compliant,  OAuth2 server written in Golang.  *2018-02-23T03:52:02Z*
* [gologin](https://github.com/dghubble/gologin):star:824 - chainable handlers for login with OAuth1 and OAuth2 authentication providers.  *2018-01-26T01:50:18Z*
* [gorbac](https://github.com/mikespook/gorbac):star:693 - provides a lightweight role-based access control (RBAC) implementation in Golang.  *2017-10-29T23:37:38Z*
* [goth](https://github.com/markbates/goth):star:0 - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.  *2018-04-12T22:15:10Z*
* [httpauth](https://github.com/goji/httpauth):star:137 - HTTP Authentication middleware.  *2016-06-01T13:53:02Z*
* [jwt](https://github.com/robbert229/jwt):star:40 - Clean and easy to use implementation of JSON Web Tokens (JWT).  *2018-03-07T03:24:34Z*
* [jwt](https://github.com/pascaldekloe/jwt):star:1 - Lightweight JSON Web Token (JWT) library.  *2018-04-10T22:13:33Z*
* [jwt-auth](https://github.com/adam-hanna/jwt-auth):star:110 - JWT middleware for Golang http servers with many configuration options.  *2018-02-06T06:40:14Z*
* [jwt-go](https://github.com/dgrijalva/jwt-go):star:0 - Golang implementation of JSON Web Tokens (JWT).  *2018-03-08T23:13:08Z*
* [loginsrv](https://github.com/tarent/loginsrv):star:576 - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.  *2018-03-14T08:22:35Z*
* [oauth2](https://github.com/golang/oauth2):star:0 - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.  *2018-04-07T16:16:47Z*
* [osin](https://github.com/RangelReale/osin):star:0 - Golang OAuth2 server library.  *2018-02-22T13:56:16Z*
* [paseto](https://github.com/o1egl/paseto):star:102 - Golang implementation of Platform-Agnostic Security Tokens (PASETO)  *2018-03-11T21:25:29Z*
* [permissions2](https://github.com/xyproto/permissions2):star:260 - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.  *2018-04-17T08:10:19Z*
* [securecookie](https://github.com/chmike/securecookie):star:19 - Efficient secure cookie encoding/decoding.  *2017-10-16T08:25:04Z*
* [session](https://github.com/icza/session):star:60 - Go session management for web servers (including support for Google App Engine - GAE).  *2018-03-21T16:35:34Z*
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go):star:4 - Go session management using the SessionGate Redis module.  *2018-03-02T08:40:08Z*
* [sessions](https://github.com/adam-hanna/sessions):star:29 - Dead simple, highly performant, highly customizable sessions service for go http servers.  *2017-11-28T18:30:58Z*
* [yubigo](https://github.com/GeertJohan/yubigo):star:84 - Yubikey client package that provides a simple API to integrate the Yubico Yubikey into a go application.  *2014-05-21T14:15:43Z*

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [argparse](https://github.com/akamensky/argparse):star:25 - Command line argument parser inspired by Python's argparse module.  *2018-04-05T04:43:06Z*
* [argv](https://github.com/cosiner/argv):star:7 - Go library to split command line string as arguments array using the bash syntax.  *2017-02-25T14:54:30Z*
* [cli](https://github.com/mkideal/cli):star:361 - Feature-rich and easy to use command-line package based on golang struct tags.  *2017-03-24T15:56:50Z*
* [cli](https://github.com/teris-io/cli):star:27 - Simple and complete API for building command line interfaces in Go.  *2018-02-22T13:27:41Z*
* [cli-init](https://github.com/tcnksm/gcli):star:811 - The easy way to start building Golang command line applications.  *2017-01-29T03:38:39Z*
* [climax](http://github.com/tucnak/climax):star:131 - Alternative CLI with "human face", in spirit of Go command.  *2016-01-10T10:13:00Z*
* [cobra](https://github.com/spf13/cobra):star:0 - Commander for modern Go CLI interactions.  *2018-04-12T12:08:29Z*
* [commandeer](https://github.com/jaffee/commandeer):star:56 - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   *2018-02-04T21:59:45Z*
* [complete](https://github.com/posener/complete):star:437 - Write bash completions in Go + Go command bash completion.  *2018-04-13T09:18:13Z*
* [docopt.go](https://github.com/docopt/docopt.go):star:924 - Command-line arguments parser that will make you smile.  *2018-01-11T23:17:33Z*
* [drive](https://github.com/odeke-em/drive):star:0 - Google Drive client for the commandline.  *2017-12-19T11:03:11Z*
* [env](https://github.com/codingconcepts/env):star:18 - Tag-based environment configuration for structs.  *2017-10-05T08:04:04Z*
* [flag](https://github.com/cosiner/flag):star:76 - Simple but powerful command line option parsing library for Go supporting subcommand.  *2017-07-16T09:01:35Z*
* [go-arg](https://github.com/alexflint/go-arg):star:472 - Struct-based argument parsing in Go.  *2018-04-20T14:57:08Z*
* [go-flags](https://github.com/jessevdk/go-flags):star:0 - go command line option parser.  *2018-03-31T12:42:32Z*
* [gocmd](https://github.com/devfacet/gocmd):star:14 - Go library for building command line applications.  *2018-04-01T21:33:31Z*
* [kingpin](https://github.com/alecthomas/kingpin):star:0 - Command line and flag parser supporting sub commands.  *2018-03-12T05:05:58Z*
* [liner](https://github.com/peterh/liner):star:438 - Go readline-like library for command-line interfaces.  *2018-03-10T18:12:06Z*
* [mitchellh/cli](https://github.com/mitchellh/cli):star:797 - Go library for implementing command-line interfaces.  *2018-04-14T17:04:47Z*
* [mow.cli](https://github.com/jawher/mow.cli):star:508 - Go library for building CLI applications with sophisticated flag and argument parsing and validation.  *2018-04-07T11:02:20Z*
* [pflag](https://github.com/spf13/pflag):star:366 - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.  *2018-04-12T12:09:13Z*
* [readline](https://github.com/chzyer/readline):star:0 - Pure golang implementation that provides most features in GNU-Readline under MIT license.  *2017-12-08T01:17:16Z*
* [sflags](https://github.com/octago/sflags):star:52 - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.  *2018-02-04T00:40:24Z*
* [strumt](https://github.com/antham/strumt):star:20 - Library to create prompt chain.  *2018-02-17T11:11:54Z*
* [ukautz/clif](https://github.com/ukautz/clif):star:84 - Small command line interface framework.  *2016-01-09T22:08:11Z*
* [urfave/cli](https://github.com/urfave/cli):star:0 - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).  *2018-02-26T03:02:53Z*
* [wlog](https://github.com/dixonwille/wlog):star:23 - Simple logging interface that supports cross-platform color and concurrency.  *2017-07-20T23:52:29Z*
* [wmenu](https://github.com/dixonwille/wmenu):star:44 - Easy to use menu structure for cli applications that prompts users to make choices.  *2018-01-03T10:33:19Z*

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [aurora](https://github.com/logrusorgru/aurora):star:290 - ANSI terminal colors that supports fmt.Printf/Sprintf.  *2018-04-19T16:45:47Z*
* [cfmt](https://github.com/mingrammer/cfmt):star:38 - Contextual fmt inspired by bootstrap color classes.  *2018-03-19T01:09:33Z*
* [chalk](https://github.com/ttacon/chalk):star:247 - Intuitive package for prettifying terminal/console output.  *2016-06-26T20:24:18Z*
* [color](https://github.com/fatih/color):star:0 - Versatile package for colored terminal output.  *2018-02-13T13:34:03Z*
* [colourize](https://github.com/TreyBastian/colourize):star:12 - Go library for ANSI colour text in terminals.  *2016-05-10T09:49:57Z*
* [go-ataman](https://github.com/workanator/go-ataman):star:5 - Go library for rendering ANSI colored text templates in terminals.  *2017-09-25T06:07:02Z*
* [go-colorable](https://github.com/mattn/go-colorable):star:267 - Colorable writer for windows.  *2018-03-10T13:32:14Z*
* [go-colortext](https://github.com/daviddengcn/go-colortext):star:168 - Go library for color output in terminals.  *2018-04-09T17:49:41Z*
* [go-isatty](https://github.com/mattn/go-isatty):star:232 - isatty for golang.  *2017-11-07T05:05:31Z*
* [go-prompt](https://github.com/c-bata/go-prompt):star:0 - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).  *2018-04-19T20:20:40Z*
* [gocui](https://github.com/jroimartin/gocui):star:0 - Minimalist Go library aimed at creating Console User Interfaces.  *2018-04-11T15:50:10Z*
* [gommon/color](https://github.com/labstack/gommon/tree/master/color):star:221 - Style terminal text.  *2016-09-25T18:04:09Z*
* [mpb](https://github.com/vbauerster/mpb):star:212 - Multi progress bar for terminal applications.  *2018-04-21T08:01:06Z*
* [progressbar](https://github.com/schollz/progressbar):star:265 - Basic thread-safe progress bar that works in every OS.  *2018-01-29T13:08:54Z*
* [termbox-go](https://github.com/nsf/termbox-go):star:0 - Termbox is a library for creating cross-platform text-based interfaces.  *2018-04-18T04:23:32Z*
* [termtables](https://github.com/apcera/termtables):star:988 - Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output.  *2018-01-23T19:08:53Z*
* [termui](https://github.com/gizak/termui):star:0 - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).  *2017-12-03T06:49:38Z*
* [tui-go](https://github.com/marcusolsson/tui-go):star:0 - Go UI library for building rich terminal applications.  *2018-03-16T21:29:05Z*
* [uilive](https://github.com/gosuri/uilive):star:643 - Library for updating terminal output in realtime.  *2017-03-23T04:15:06Z*
* [uiprogress](https://github.com/gosuri/uiprogress):star:0 - Flexible library to render progress bars in terminal applications.  *2017-02-24T06:39:37Z*
* [uitable](https://github.com/gosuri/uitable):star:425 - Library to improve readability in terminal apps using tabular data.  *2016-04-04T20:39:58Z*

## Configuration

*Libraries for configuration parsing.*

* [config](https://github.com/olebedev/config):star:150 - JSON or YAML configuration wrapper with environment variables and flags parsing.  *2017-09-14T11:20:07Z*
* [configure](https://github.com/paked/configure):star:34 - Provides configuration through multiple sources, including JSON, flags and environment variables.  *2017-06-26T17:05:06Z*
* [confita](https://github.com/heetch/confita):star:36 - Load configuration in cascade from multiple backends into a struct.  *2018-04-19T16:40:39Z*
* [env](https://github.com/caarlos0/env):star:455 - Parse environment variables to Go structs (with defaults).  *2018-03-16T12:55:29Z*
* [envcfg](https://github.com/tomazk/envcfg):star:87 - Un-marshaling environment variables to Go structs.  *2017-06-19T15:53:18Z*
* [envconf](https://github.com/ian-kent/envconf):star:6 - Configuration from environment.  *2014-10-26T12:11:21Z*
* [envconfig](https://github.com/vrischmann/envconfig):star:115 - Read your configuration from environment variables.  *2017-03-06T15:51:51Z*
* [envh](https://github.com/antham/envh):star:85 - Helpers to manage environment variables.  *2017-12-29T00:17:43Z*
* [gcfg](https://github.com/go-gcfg/gcfg):star:80 - read INI-style configuration files into Go structs; supports user-defined types and subsections.  *2018-03-28T16:41:26Z*
* [go-up](https://github.com/ufoscout/go-up):star:10 - A simple configuration library with recursive placeholders resolution and no magic.  *2018-03-16T19:28:39Z*
* [goConfig](https://github.com/crgimenes/goConfig):star:65 - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.  *2018-02-27T12:51:44Z*
* [godotenv](https://github.com/joho/godotenv):star:0 - Go port of Ruby's dotenv library (Loads environment variables from `.env`).  *2018-04-05T05:36:34Z*
* [gofigure](https://github.com/ian-kent/gofigure):star:52 - Go application configuration made easy.  *2017-05-02T19:22:41Z*
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf):star:21 - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.  *2017-06-20T09:24:12Z*
* [hjson](https://github.com/hjson/hjson-go):star:118 - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.  *2018-01-24T22:37:38Z*
* [ingo](https://github.com/schachmat/ingo):star:16 - Flags persisted in an ini-like config file.  *2017-04-03T01:15:06Z*
* [ini](https://github.com/go-ini/ini):star:894 - Go package to read and write INI files.  *2018-04-20T15:00:25Z*
* [joshbetz/config](https://github.com/joshbetz/config):star:192 - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.  *2017-05-06T16:09:22Z*
* [mini](https://github.com/sasbury/mini):star:16 - Golang package for parsing ini-style configuration files.  *2016-12-24T19:37:50Z*
* [store](https://github.com/tucnak/store):star:222 - Lightweight configuration manager for Go.  *2017-09-05T11:38:34Z*
* [viper](https://github.com/spf13/viper):star:0 - Go configuration with fangs.  *2018-03-22T07:00:28Z*
* [xdg](https://github.com/OpenPeeDeeP/xdg) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone):star:0 - Drone is a Continuous Integration platform built on Docker, written in Go.  *2018-04-10T16:07:28Z*
* [gomason](https://github.com/nikogura/gomason):star:1 - Test, Build, Sign, and Publish your go binaries from a clean workspace.  *2018-04-03T23:32:19Z*
* [goveralls](https://github.com/mattn/goveralls):star:472 - Go integration for Coveralls.io continuous code coverage tracking system.  *2018-03-19T02:19:29Z*
* [overalls](https://github.com/go-playground/overalls):star:88 - Multi-Package go project coverprofile for tools like goveralls.  *2018-02-01T14:43:45Z*
* [roveralls](https://github.com/LawrenceWoodman/roveralls):star:7 - Recursive coverage testing tool.  *2017-11-19T19:38:43Z*

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [c6](https://github.com/c9s/c6):star:401 - High performance SASS compatible-implementation compiler written in Go.  *2017-08-06T12:20:50Z*
* [gcss](https://github.com/yosssi/gcss):star:402 - Pure Go CSS Preprocessor.  *2014-10-12T14:02:05Z*
* [go-libsass](https://github.com/wellington/go-libsass):star:77 - Go wrapper to the 100% Sass compatible libsass project.  *2018-02-17T23:48:10Z*

## Data Structures

*Generic datastructures and algorithms in Go.*

* [binpacker](https://github.com/zhuangsirui/binpacker):star:80 - Binary packer and unpacker helps user build custom binary stream.  *2017-10-16T02:34:05Z*
* [bit](https://github.com/yourbasic/bit):star:17 - Golang set data structure with bonus bit-twiddling functions.  *2018-03-13T07:44:24Z*
* [bitset](https://github.com/willf/bitset):star:354 - Go package implementing bitsets.  *2017-09-05T00:26:39Z*
* [bloom](https://github.com/zhenjl/bloom):star:111 - Bloom filters implemented in Go.  *2015-10-26T23:31:58Z*
* [bloom](https://github.com/yourbasic/bloom):star:4 - Golang Bloom filter implementation.  *2017-06-02T16:39:13Z*
* [boomfilters](https://github.com/tylertreat/BoomFilters):star:931 - Probabilistic data structures for processing continuous, unbounded streams.  *2018-03-23T14:25:36Z*
* [concurrent-writer](https://github.com/free/concurrent-writer):star:5 - Highly concurrent drop-in replacement for `bufio.Writer`.  *2017-11-17T21:28:31Z*
* [conjungo](https://github.com/InVisionApp/conjungo):star:31 - A small, powerful and flexible merge library.  *2018-01-03T17:05:07Z*
* [count-min-log](https://github.com/seiflotfy/count-min-log):star:37 - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).  *2017-02-12T13:09:20Z*
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter):star:300 - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.  *2018-01-10T12:30:52Z*
* [encoding](https://github.com/zhenjl/encoding):star:81 - Integer Compression Libraries for Go.  *2017-12-23T22:15:21Z*
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree):star:32 - Go implementation of Adaptive Radix Tree.  *2017-03-24T18:00:00Z*
* [go-datastructures](https://github.com/Workiva/go-datastructures):star:0 - Collection of useful, performant, and thread-safe data structures.  *2018-04-12T13:39:22Z*
* [go-ef](https://github.com/amallia/go-ef):star:5 - A Go implementation of the Elias-Fano encoding.  *2017-09-25T20:03:57Z*
* [go-geoindex](https://github.com/hailocab/go-geoindex):star:274 - In-memory geo index.  *2016-01-27T13:48:10Z*
* [go-rquad](https://github.com/aurelien-rainone/go-rquad):star:77 - Region quadtrees with efficient point location and neighbour finding.  *2017-06-18T12:41:46Z*
* [gods](https://github.com/emirpasic/gods):star:0 - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.  *2017-12-26T20:51:25Z*
* [golang-set](https://github.com/deckarep/golang-set):star:718 - Thread-Safe and Non-Thread-Safe high-performance sets for Go.  *2018-03-28T15:30:22Z*
* [goset](https://github.com/zoumo/goset):star:6 - A useful Set collection implementation for Go.  *2017-08-25T10:23:07Z*
* [goskiplist](https://github.com/ryszard/goskiplist):star:150 - Skip list implementation in Go.  *2015-03-12T22:13:10Z*
* [gota](https://github.com/kniren/gota):star:495 - Implementation of dataframes, series, and data wrangling methods for Go.  *2017-10-24T06:28:51Z*
* [hilbert](https://github.com/google/hilbert):star:128 - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.  *2017-11-30T06:07:56Z*
* [hyperloglog](https://github.com/axiomhq/hyperloglog):star:581 - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.  *2018-03-17T13:17:30Z*
* [levenshtein](https://github.com/agext/levenshtein):star:12 - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.  *2017-02-17T06:30:20Z*
* [levenshtein](https://github.com/agnivade/levenshtein):star:15 - Implementation to calculate levenshtein distance in Go.  *2018-03-03T09:57:33Z*
* [mafsa](https://github.com/smartystreets/mafsa):star:265 - MA-FSA implementation with Minimal Perfect Hashing.  *2016-03-22T21:56:34Z*
* [merkletree](https://github.com/cbergoon/merkletree):star:47 - Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.  *2018-02-28T02:27:59Z*
* [roaring](https://github.com/RoaringBitmap/roaring):star:388 - Go package implementing compressed bitsets.  *2018-04-21T13:36:03Z*
* [skiplist](https://github.com/gansidui/skiplist):star:48 - Skiplist implementation in Go.  *2014-11-21T05:13:32Z*
* [trie](https://github.com/derekparker/trie):star:257 - Trie implementation in Go.  *2018-02-12T16:26:25Z*
* [ttlcache](https://github.com/diegobernardes/ttlcache):star:53 - In-memory LRU string-interface{} map with expiration for golang.  *2018-03-13T12:29:36Z*
* [willf/bloom](https://github.com/willf/bloom):star:430 - Go package implementing Bloom filters.  *2017-05-05T22:16:40Z*

## Database

*Databases implemented in Go.*

* [badger](https://github.com/dgraph-io/badger):star:0 - Fast key-value store in Go.  *2018-04-06T05:54:07Z*
* [BigCache](https://github.com/allegro/bigcache):star:0 - Efficient key/value cache for gigabytes of data.  *2018-03-29T15:28:23Z*
* [bolt](https://github.com/boltdb/bolt):star:0 - Low-level key/value database for Go.  *2018-03-02T18:00:52Z*
* [buntdb](https://github.com/tidwall/buntdb):star:0 - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.  *2018-03-16T00:57:49Z*
* [cache2go](https://github.com/muesli/cache2go):star:482 - In-memory key:value cache which supports automatic invalidation based on timeouts.  *2018-04-17T22:28:18Z*
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache):star:21 - BigCache with clustering support and individual item expiration.  *2018-01-22T22:02:53Z*
* [cockroach](https://github.com/cockroachdb/cockroach):star:0 - Scalable, Geo-Replicated, Transactional Datastore.  *2018-04-21T22:04:10Z*
* [couchcache](https://github.com/codingsince1985/couchcache):star:31 - RESTful caching micro-service backed by Couchbase server.  *2017-10-25T01:28:42Z*
* [dgraph](https://github.com/dgraph-io/dgraph):star:0 - Scalable, Distributed, Low Latency, High Throughput Graph Database.  *2018-04-20T06:41:20Z*
* [diskv](https://github.com/peterbourgon/diskv):star:579 - Home-grown disk-backed key-value store.  *2018-03-12T05:41:25Z*
* [eliasdb](https://github.com/krotik/eliasdb):star:486 - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.  *2017-09-09T08:08:44Z*
* [forestdb](https://github.com/couchbase/goforestdb):star:28 - Go bindings for ForestDB.  *2016-12-15T00:13:43Z*
* [GCache](https://github.com/bluele/gcache):star:411 - Cache library with support for expirable Cache, LFU, LRU and ARC.  *2017-10-10T15:56:17Z*
* [go-cache](https://github.com/pmylund/go-cache):star:0 - In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.  *2017-07-22T04:01:10Z*
* [goleveldb](https://github.com/syndtr/goleveldb):star:0 - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.  *2018-04-17T20:23:10Z*
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) - Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.
* [groupcache](https://github.com/golang/groupcache):star:0 - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.  *2018-02-03T14:35:32Z*
* [influxdb](https://github.com/influxdb/influxdb):star:0 - Scalable datastore for metrics, events, and real-time analytics.  *2018-04-20T14:14:40Z*
* [jaeger](https://github.com/jaegertracing/jaeger):star:0 - A distributed tracing system.  *2018-04-21T19:31:09Z*
* [ledisdb](https://github.com/siddontang/ledisdb):star:0 - Ledisdb is a high performance NoSQL like Redis based on LevelDB.  *2018-04-17T14:13:38Z*
* [levigo](https://github.com/jmhodges/levigo):star:328 - Levigo is a Go wrapper for LevelDB.  *2016-11-15T19:34:49Z*
* [moss](https://github.com/couchbase/moss):star:494 - Moss is a simple LSM key-value storage engine written in 100% Go.  *2018-02-28T20:48:26Z*
* [piladb](https://github.com/fern4lvarez/piladb):star:155 - Lightweight RESTful database engine based on stack data structures.  *2018-03-29T15:59:25Z*
* [prometheus](https://github.com/prometheus/prometheus):star:0 - Monitoring system and time series database.  *2018-04-20T16:35:42Z*
* [rqlite](https://github.com/rqlite/rqlite):star:0 - The lightweight, distributed, relational database built on SQLite.  *2018-04-21T22:43:20Z*
* [Scribble](https://github.com/nanobox-io/golang-scribble):star:123 - Tiny flat file JSON store.  *2017-10-26T14:29:21Z*
* [tempdb](https://github.com/rafaeljesus/tempdb):star:10 - Key-value store for temporary items.  *2018-02-14T19:03:10Z*
* [tidb](https://github.com/pingcap/tidb):star:0 - TiDB is a distributed SQL database. Inspired by the design of Google F1.  *2018-04-21T13:58:24Z*
* [tiedot](https://github.com/HouzuoGuo/tiedot):star:0 - Your NoSQL database powered by Golang.  *2018-01-04T21:35:38Z*
* [Vasto](https://github.com/chrislusf/vasto):star:65 - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.  *2018-04-14T19:46:27Z*

*Database schema migration.*

* [darwin](https://github.com/GuiaBolso/darwin):star:54 - Database schema evolution library for Go.  *2017-02-10T19:16:49Z*
* [go-fixtures](https://github.com/RichardKnop/go-fixtures):star:7 - Django style fixtures for Golang's excellent built-in database/sql library.  *2017-12-19T14:59:45Z*
* [gondolier](https://github.com/emvicom/gondolier):star:15 - Gondolier is a library to auto migrate database schemas using structs.  *2018-01-15T13:58:30Z*
* [goose](https://github.com/steinbacher/goose):star:69 - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.  *2016-07-25T13:16:29Z*
* [gormigrate](https://github.com/go-gormigrate/gormigrate):star:138 - Database schema migration helper for Gorm ORM.  *2018-03-24T22:04:29Z*
* [migrate](https://github.com/mattes/migrate):star:0 - Database migrations. CLI and Golang library.  *2017-12-08T21:48:26Z*
* [pravasan](https://github.com/pravasan/pravasan):star:20 - Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.  *2015-03-20T07:06:02Z*
* [soda](https://github.com/markbates/pop/tree/master/soda):star:563 - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.  *2018-02-25T22:33:14Z*
* [sql-migrate](https://github.com/rubenv/sql-migrate):star:926 - Database migration tool. Allows embedding migrations into the application using go-bindata.  *2018-02-17T20:35:53Z*

*Database tools.*

* [chproxy](https://github.com/Vertamedia/chproxy):star:107 - HTTP proxy for ClickHouse database.  *2018-04-09T08:14:58Z*
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk):star:58 - Collects small insterts and sends big requests to ClickHouse servers.  *2018-03-22T23:02:30Z*
* [go-mysql](https://github.com/siddontang/go-mysql):star:940 - Go toolset to handle MySQL protocol and replication.  *2018-04-07T13:08:29Z*
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch):star:0 - Sync your MySQL data into Elasticsearch automatically.  *2018-04-04T14:22:02Z*
* [kingshard](https://github.com/flike/kingshard):star:0 - kingshard is a high performance proxy for MySQL powered by Golang.  *2018-04-12T08:38:03Z*
* [myreplication](https://github.com/2tvenom/myreplication):star:95 - MySql binary log replication listener. Supports statement and row based replication.  *2015-02-04T21:04:42Z*
* [orchestrator](https://github.com/github/orchestrator):star:0 - MySQL replication topology manager & visualizer.  *2018-04-16T12:42:53Z*
* [pgweb](https://github.com/sosedoff/pgweb):star:0 - Web-based PostgreSQL database browser.  *2018-03-17T01:15:06Z*
* [prep](https://github.com/hexdigest/prep):star:19 - Use prepared SQL statements without changing your code.  *2017-12-19T17:35:41Z*
* [pREST](https://github.com/nuveo/prest):star:0 - Serve a RESTful API from any PostgreSQL database.  *2018-04-02T15:29:48Z*
* [rwdb](https://github.com/andizzle/rwdb):star:8 - rwdb provides read replica capability for multiple database servers setup.  *2017-11-08T09:10:06Z*
* [vitess](https://github.com/youtube/vitess):star:0 - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.  *2018-04-21T03:51:51Z*

*SQL query builder, libraries for building and using SQL.*

* [dat](https://github.com/mgutz/dat):star:515 - Go Postgres Data Access Toolkit.  *2017-03-28T13:55:41Z*
* [Dotsql](https://github.com/gchaincl/dotsql):star:351 - Go library that helps you keep sql files in one place and use them with ease.  *2015-12-21T00:04:13Z*
* [gendry](https://github.com/didi/gendry):star:210 - Non-invasive SQL builder and powerful data binder.  *2018-04-09T01:32:02Z*
* [godbal](https://github.com/xujiajun/godbal):star:2 - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.  *2018-04-15T03:10:27Z*
* [goqu](https://github.com/doug-martin/goqu):star:407 - Idiomatic SQL builder and query library.  *2017-12-27T16:44:32Z*
* [igor](https://github.com/galeone/igor):star:64 - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.  *2017-04-02T19:50:40Z*
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx):star:230 - Powerful data retrieval methods as well as DB-agnostic query building capabilities.  *2018-03-01T02:55:00Z*
* [scaneo](https://github.com/variadico/scaneo):star:129 - Generate Go code to convert database rows into arbitrary structs.  *2015-07-16T09:39:32Z*
* [sqrl](https://github.com/elgris/sqrl):star:108 - SQL query builder, fork of Squirrel with improved performance.  *2018-03-06T09:54:26Z*
* [Squirrel](https://github.com/Masterminds/squirrel):star:0 - Go library that helps you build SQL queries.  *2018-03-19T15:22:03Z*
* [xo](https://github.com/knq/xo):star:0 - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.  *2017-08-23T07:01:48Z*

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [avatica](https://github.com/Boostport/avatica):star:40 - Apache Phoenix/Avatica SQL driver for database/sql.  *2018-03-04T22:21:27Z*
    * [bgc](https://github.com/viant/bgc):star:7 - Datastore Connectivity for BigQuery for go.  *2018-04-20T01:44:39Z*
    * [firebirdsql](https://github.com/nakagami/firebirdsql):star:77 - Firebird RDBMS SQL driver for Go.  *2018-02-27T01:55:50Z*
    * [go-adodb](https://github.com/mattn/go-adodb):star:70 - Microsoft ActiveX Object DataBase driver for go that uses database/sql.  *2017-07-03T23:43:26Z*
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb):star:723 - Microsoft MSSQL driver for Go.  *2018-04-16T13:29:13Z*
    * [go-oci8](https://github.com/mattn/go-oci8):star:287 - Oracle driver for go that uses database/sql.  *2018-04-05T04:00:22Z*
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql):star:0 - MySQL driver for Go.  *2018-04-13T18:15:57Z*
    * [go-sqlite3](https://github.com/mattn/go-sqlite3):star:0 - SQLite3 driver for go that uses database/sql.  *2018-04-19T07:32:57Z*
    * [gofreetds](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
    * [goracle](https://github.com/go-goracle/goracle):star:81 - Oracle driver for Go, using the ODPI-C driver  *2018-04-07T15:22:34Z*
    * [pgx](https://github.com/jackc/pgx):star:0 - PostgreSQL driver supporting features beyond those exposed by database/sql.  *2018-04-14T15:02:08Z*
    * [pq](https://github.com/lib/pq):star:0 - Pure Go Postgres driver for database/sql.  *2018-03-27T07:18:24Z*

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go):star:256 - Aerospike client in Go language.  *2018-04-13T06:54:53Z*
    * [arangolite](https://github.com/solher/arangolite):star:59 - Lightweight golang driver for ArangoDB.  *2017-09-10T09:39:06Z*
    * [asc](https://github.com/viant/asc):star:3 - Datastore Connectivity for Aerospike for go.  *2018-03-13T13:05:06Z*
    * [cayley](https://github.com/google/cayley):star:0 - Graph database with support for multiple backends.  *2018-04-17T13:58:24Z*
    * [dsc](https://github.com/viant/dsc):star:5 - Datastore connectivity for SQL, NoSQL, structured files.  *2018-04-17T04:31:38Z*
    * [dynago](https://github.com/underarmour/dynago):star:53 - Dynago is a principle of least surprise client for DynamoDB.  *2017-03-28T01:08:44Z*
    * [go-couchbase](https://github.com/couchbase/go-couchbase):star:273 - Couchbase client in Go.  *2018-04-16T17:41:43Z*
    * [go-couchdb](https://github.com/fjl/go-couchdb):star:45 - Yet another CouchDB HTTP API wrapper for Go.  *2014-07-04T15:13:33Z*
    * [gocb](https://github.com/couchbase/gocb):star:243 - Official Couchbase Go SDK.  *2018-04-06T23:59:23Z*
    * [gocql](http://gocql.github.io) - Go language driver for Apache Cassandra.
    * [gomemcache](https://github.com/bradfitz/gomemcache/):star:894 - memcache client library for the Go programming language.  *2017-02-08T21:30:04Z*
    * [gorethink](https://github.com/dancannon/gorethink):star:0 - Go language driver for RethinkDB.  *2017-12-14T15:00:17Z*
    * [goriak](https://github.com/zegl/goriak):star:20 - Go language driver for Riak KV.  *2017-11-12T15:44:38Z*
    * [mgo](https://github.com/globalsign/mgo):star:608 - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms  *2018-04-03T08:58:42Z*
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver):star:313 - Official MongoDB driver for the Go language.  *2018-04-20T21:49:46Z*
    * [neo4j](https://github.com/cihangir/neo4j):star:21 - Neo4j Rest API Bindings for Golang.  *2015-04-02T17:38:47Z*
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO):star:66 - Neo4j REST Client in golang.  *2012-11-24T03:12:52Z*
    * [neoism](https://github.com/jmcvetta/neoism):star:329 - Neo4j client for Golang.  *2017-03-06T11:01:38Z*
    * [redigo](https://github.com/gomodule/redigo):star:0 - Redigo is a Go client for the Redis database.  *2018-04-01T19:18:55Z*
    * [redis](https://github.com/go-redis/redis):star:0 - Redis client for Golang.  *2018-04-17T06:18:16Z*
    * [redis](https://github.com/hoisie/redis):star:563 - Simple, powerful Redis client for Go.  *2016-07-30T15:44:56Z*
    * [redis](https://github.com/bsm/redeo):star:165 - Redis-protocol compatible TCP servers/services.  *2018-04-12T14:20:38Z*
    * [xredis](https://github.com/shomali11/xredis):star:7 - Typesafe, customizable, clean & easy to use Redis client.  *2017-11-11T03:10:15Z*

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve):star:0 - Modern text indexing library for go.  *2018-04-21T20:17:50Z*
    * [elastic](https://github.com/olivere/elastic):star:0 - Elasticsearch client for Go.  *2018-04-21T09:44:53Z*
    * [elasticsql](https://github.com/cch123/elasticsql):star:160 - Convert sql to elasticsearch dsl in Go.  *2017-12-26T08:53:44Z*
    * [elastigo](https://github.com/mattbaird/elastigo):star:904 - Elasticsearch client library.  *2017-01-23T22:00:20Z*
    * [goes](https://github.com/OwnLocal/goes):star:20 - Library to interact with Elasticsearch.  *2017-03-02T23:59:55Z*
    * [riot](https://github.com/go-ego/riot):star:0 - Go Open Source, Distributed, Simple and efficient Search Engine  *2018-04-21T13:02:19Z*
    * [skizze](https://github.com/seiflotfy/skizze):star:50 - probabilistic data-structures service and storage.  *2016-05-03T16:18:20Z*

## Date and Time

*Libraries for working with dates and times.*

* [carbon](https://github.com/uniplaces/carbon):star:189 - Simple Time extension with a lot of util methods, ported from PHP Carbon library.  *2018-03-05T12:16:29Z*
* [date](https://github.com/rickb777/date):star:12 - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.  *2018-02-19T22:30:10Z*
* [dateparse](https://github.com/araddon/dateparse):star:595 - Parse date's without knowing format in advance.  *2018-04-19T00:36:22Z*
* [durafmt](https://github.com/hako/durafmt):star:177 - Time duration formatting library for Go.  *2018-02-26T20:55:55Z*
* [feiertage](https://github.com/wlbr/feiertage):star:11 - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...  *2017-05-15T15:39:04Z*
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar):star:32 - The implementation of the Persian (Solar Hijri) Calendar in Go (golang).  *2016-08-05T20:26:46Z*
* [go-sunrise](https://github.com/nathan-osman/go-sunrise):star:4 - Calculate the sunrise and sunset times for a given location.  *2017-11-21T20:49:56Z*
* [goweek](https://github.com/grsmv/goweek):star:13 - Library for working with week entity in golang.  *2017-01-03T20:24:25Z*
* [now](https://github.com/jinzhu/now):star:0 - Now is a time toolkit for golang.  *2018-03-14T13:20:04Z*
* [NullTime](https://github.com/kirillDanshin/nulltime):star:6 - Nullable `time.Time`.  *2017-03-22T04:30:28Z*
* [strftime](https://github.com/awoodbeck/strftime):star:1 - C99-compatible strftime formatter.  *2018-02-21T15:59:08Z*
* [timespan](https://github.com/SaidinWoT/timespan):star:49 - For interacting with intervals of time, defined as a start time and a duration.  *2016-04-03T21:07:42Z*
* [timeutil](https://github.com/leekchan/timeutil):star:143 - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.  *2015-08-02T14:26:58Z*
* [tuesday](https://github.com/osteele/tuesday):star:4 - Ruby-compatible Strftime function.  *2017-08-30T13:24:14Z*

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [celeriac](https://github.com/svcavallar/celeriac.v1):star:34 - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.  *2018-02-21T05:52:47Z*
* [consistent](https://github.com/buraksezer/consistent):star:84 - Consistent hashing with bounded loads.  *2018-04-19T13:43:00Z*
* [digota](https://github.com/digota/digota):star:200 - grpc ecommerce microservice.  *2017-10-24T11:32:44Z*
* [drmaa](https://github.com/dgruber/drmaa):star:19 - Job submission library for cluster schedulers based on the DRMAA standard.  *2017-02-27T18:55:26Z*
* [emitter-io](https://github.com/emitter-io/emitter):star:0 - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.  *2018-04-15T07:18:33Z*
* [flowgraph](https://github.com/vectaport/flowgraph):star:35 - MPI-style ready-send coordination layer.  *2017-12-05T22:26:53Z*
* [gleam](https://github.com/chrislusf/gleam):star:0 - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.  *2018-04-19T07:11:57Z*
* [glow](https://github.com/chrislusf/glow):star:0 - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.  *2017-07-11T06:56:44Z*
* [go-health](https://github.com/InVisionApp/go-health):star:172 - Library for enabling asynchronous dependency health checks in your service.  *2018-04-11T15:24:25Z*
* [go-jump](https://github.com/dgryski/go-jump):star:180 - Port of Google's "Jump" Consistent Hash function.  *2017-04-09T06:50:14Z*
* [go-kit](https://github.com/go-kit/kit):star:0 - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.  *2018-04-02T01:41:32Z*
* [gorpc](https://github.com/valyala/gorpc):star:444 - Simple, fast and scalable RPC library for high load.  *2016-05-19T17:16:14Z*
* [grpc-go](https://github.com/grpc/grpc-go):star:0 - The Go language implementation of gRPC. HTTP/2 based RPC.  *2018-04-20T00:17:21Z*
* [heimdall](https://github.com/gojektech/heimdall):star:449 - An enchanced http client with retry and hystrix capabilities.  *2018-04-16T11:23:53Z*
* [hprose](https://github.com/hprose/hprose-golang):star:738 - Very newbility RPC Library, support 25+ languages now.  *2018-04-19T11:32:46Z*
* [jsonrpc](https://github.com/osamingo/jsonrpc):star:59 - The jsonrpc package helps implement of JSON-RPC 2.0.  *2018-03-27T02:20:43Z*
* [jsonrpc](https://github.com/ybbus/jsonrpc):star:40 - JSON-RPC 2.0 HTTP client implementation.  *2018-04-11T22:23:09Z*
* [KrakenD](https://github.com/devopsfaith/krakend):star:411 - Ultra performant API Gateway framework with middlewares.  *2018-04-07T11:41:12Z*
* [micro](https://github.com/micro/micro):star:0 - Pluggable microservice toolkit and distributed systems platform.  *2018-04-20T13:42:58Z*
* [NATS](https://github.com/nats-io/gnatsd):star:0 - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.  *2018-04-11T23:10:07Z*
* [raft](https://github.com/hashicorp/raft):star:0 - Golang implementation of the Raft consensus protocol, by HashiCorp.  *2018-02-12T22:15:04Z*
* [raft](https://github.com/coreos/etcd/tree/master/raft):star:0 - Go implementation of the Raft consensus protocol, by CoreOS.  *2018-01-12T07:41:59Z*
* [ringpop-go](https://github.com/uber/ringpop-go):star:410 - Scalable, fault-tolerant application-layer sharding for Go applications.  *2017-11-21T18:50:45Z*
* [rpcx](https://github.com/smallnest/rpcx):star:0 - Distributed pluggable RPC service framework like alibaba Dubbo.  *2018-04-19T04:15:54Z*
* [sleuth](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).
* [tendermint](https://github.com/tendermint/tendermint):star:0 - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.  *2018-04-17T08:59:36Z*
* [torrent](https://github.com/anacrolix/torrent):star:0 - BitTorrent client package.  *2018-04-14T11:44:50Z*
    * [dht](https://godoc.org/github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation.
    * [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix):star:318 - Video streaming torrent client.  *2016-07-18T08:38:44Z*

## Email

*Libraries and tools that implement email creation and sending.*

* [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.
* [douceur](https://github.com/aymerick/douceur):star:118 - CSS inliner for your HTML emails.  *2018-03-22T21:03:07Z*
* [email](https://github.com/jordan-wright/email):star:872 - A robust and flexible email library for Go.  *2018-01-15T03:29:44Z*
* [go-dkim](https://github.com/toorop/go-dkim):star:37 - DKIM library, to sign & verify email.  *2017-12-03T08:00:27Z*
* [go-imap](https://github.com/emersion/go-imap):star:408 - IMAP library for clients and servers.  *2018-03-19T18:19:01Z*
* [go-message](https://github.com/emersion/go-message):star:45 - Streaming library for the Internet Message Format and mail messages.  *2018-03-21T17:30:51Z*
* [Gomail](https://github.com/go-gomail/gomail/):star:0 - Gomail is a very simple and powerful package to send emails.  *2016-04-11T21:29:32Z*
* [Hectane](https://github.com/hectane/hectane):star:141 - Lightweight SMTP client providing an HTTP API.  *2018-04-14T20:06:30Z*
* [hermes](https://github.com/matcornic/hermes):star:0 - Golang package that generates clean, responsive HTML e-mails.  *2018-01-29T21:27:27Z*
* [MailHog](https://github.com/mailhog/MailHog):star:0 - Email and SMTP testing with web and API interface.  *2017-04-16T23:42:21Z*
* [SendGrid](https://github.com/sendgrid/sendgrid-go):star:343 - SendGrid's Go library for sending email.  *2018-02-07T00:05:56Z*
* [smtp](https://github.com/mailhog/smtp):star:39 - SMTP server protocol state machine.  *2016-11-19T23:01:07Z*

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [agora](https://github.com/PuerkitoBio/agora):star:293 - Dynamically typed, embeddable programming language in Go.  *2015-01-25T17:12:44Z*
* [anko](https://github.com/mattn/anko):star:599 - Scriptable interpreter written in Go.  *2018-04-14T11:54:08Z*
* [binder](https://github.com/alexeyco/binder):star:0 - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).  *2018-04-16T13:52:41Z*
* [gisp](https://github.com/jcla1/gisp):star:392 - Simple LISP in Go.  *2014-06-29T08:28:47Z*
* [go-duktape](https://github.com/olebedev/go-duktape):star:568 - Duktape JavaScript engine bindings for Go.  *2018-03-02T12:15:09Z*
* [go-lua](https://github.com/Shopify/go-lua):star:0 - Port of the Lua 5.2 VM to pure Go.  *2017-12-04T20:47:06Z*
* [go-php](https://github.com/deuill/go-php):star:453 - PHP bindings for Go.  *2018-01-10T21:28:50Z*
* [go-python](https://github.com/sbinet/go-python):star:642 - naive go bindings to the CPython C-API.  *2018-03-26T13:41:58Z*
* [golua](https://github.com/aarzilli/golua):star:391 - Go bindings for Lua C API.  *2018-03-02T08:44:41Z*
* [gopher-lua](https://github.com/yuin/gopher-lua):star:0 - Lua 5.1 VM and compiler written in Go.  *2018-04-16T13:52:41Z*
* [ngaro](https://github.com/db47h/ngaro):star:10 - Embeddable Ngaro VM implementation enabling scripting in Retro.  *2016-08-25T13:35:45Z*
* [otto](https://github.com/robertkrimen/otto):star:0 - JavaScript interpreter written in Go.  *2018-03-05T04:20:45Z*
* [purl](https://github.com/ian-kent/purl):star:20 - Perl 5.18.2 embedded in Go.  *2014-12-07T17:45:28Z*

## Files

*Libraries for  handling files and file systems.*

* [afero](https://github.com/spf13/afero):star:0 - FileSystem Abstraction System for Go.  *2018-04-01T12:27:15Z*
* [go-csv-tag](https://github.com/artonge/go-csv-tag):star:20 - Load csv file using tag.  *2017-10-06T12:58:27Z*
* [go-gtfs](https://github.com/artonge/go-gtfs):star:5 - Load gtfs files in go.  *2017-11-29T20:48:28Z*
* [notify](https://github.com/rjeczalik/notify):star:335 - File system event notification library with simple API, similar to os/signal.  *2018-03-12T21:30:58Z*
* [skywalker](https://github.com/dixonwille/skywalker):star:26 - Package to allow one to concurrently go through a filesystem with ease.  *2017-08-04T20:24:56Z*
* [tarfs](https://github.com/posener/tarfs) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.

## Financial

*Packages for accounting and finance.*

* [accounting](https://github.com/leekchan/accounting):star:365 - money and currency formatting for golang.  *2018-04-10T04:38:18Z*
* [decimal](https://github.com/shopspring/decimal):star:898 - Arbitrary-precision fixed-point decimal numbers.  *2018-03-19T17:08:23Z*
* [go-finance](https://github.com/FlashBoys/go-finance):star:548 - Comprehensive financial markets data in Go.  *2018-03-09T02:50:38Z*
* [go-finance](https://github.com/alpeb/go-finance):star:10 - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.  *2018-02-05T10:34:48Z*
* [go-money](https://github.com/rhymond/go-money):star:414 - Implementation of Fowler's Money pattern.  *2018-04-04T10:04:18Z*
* [ofxgo](https://github.com/aclindsa/ofxgo):star:40 - Query OFX servers and/or parse the responses (with example command-line client).  *2018-04-15T11:21:52Z*
* [techan](https://github.com/sdcoffey/techan):star:21 - Technical analysis library with advanced market analysis and trading strategies.  *2018-04-16T16:02:54Z*
* [transaction](https://github.com/claygod/transaction):star:15 - Embedded transactional database of accounts, running in multithreaded mode.  *2018-03-06T05:09:36Z*
* [vat](https://github.com/dannyvankooten/vat):star:41 - VAT number validation & EU VAT rates.  *2017-07-14T08:07:27Z*

## Forms

*Libraries for working with forms.*

* [bind](https://github.com/robfig/bind):star:18 - Bind form data to any Go values.  *2014-08-16T17:03:50Z*
* [binding](https://github.com/mholt/binding):star:666 - Binds form and JSON data from net/http Request to struct.  *2017-09-17T04:34:19Z*
* [conform](https://github.com/leebenson/conform):star:123 - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.  *2017-11-29T16:31:47Z*
* [form](https://github.com/go-playground/form):star:257 - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.  *2018-04-10T05:13:05Z*
* [formam](https://github.com/monoculum/formam):star:96 - decode form's values into a struct.  *2017-08-14T19:04:38Z*
* [forms](https://github.com/albrow/forms):star:79 - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.  *2017-02-15T23:14:05Z*
* [gorilla/csrf](https://github.com/gorilla/csrf):star:272 - CSRF protection for Go web applications & services.  *2017-09-12T15:53:03Z*
* [nosurf](https://github.com/justinas/nosurf):star:850 - CSRF protection middleware for Go.  *2017-10-23T06:46:27Z*

## Game Development

*Awesome game development libraries.*

* [Azul3D](https://github.com/azul3d/engine):star:354 - 3D game engine written in Go.  *2017-07-01T20:39:10Z*
* [Ebiten](https://github.com/hajimehoshi/ebiten):star:842 - dead simple 2D game library in Go.  *2018-04-21T16:40:21Z*
* [engo](https://github.com/EngoEngine/engo):star:762 - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.  *2018-04-11T03:40:19Z*
* [GarageEngine](https://github.com/vova616/GarageEngine):star:289 - 2d game engine written in Go working on OpenGL.  *2013-09-03T11:21:00Z*
* [glop](https://github.com/runningwild/glop):star:77 - Glop (Game Library Of Power) is a fairly simple cross-platform game library.  *2015-09-24T02:43:44Z*
* [go-astar](https://github.com/beefsack/go-astar):star:242 - Go implementation of the A\* path finding algorithm.  *2017-10-24T23:10:11Z*
* [go-collada](https://github.com/GlenKelley/go-collada):star:11 - Go package for working with the Collada file format.  *2013-09-26T21:17:46Z*
* [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [go3d](https://github.com/ungerik/go3d):star:136 - Performance oriented 2D/3D math package for Go.  *2015-04-05T21:01:24Z*
* [gonet](https://github.com/xtaci/gonet):star:922 - Game server skeleton implemented with golang.  *2017-05-12T07:31:41Z*
* [goworld](https://github.com/xiaonanln/goworld):star:577 - Scalable game server engine, featuring space-entity framework and hot-swapping  *2018-04-08T16:33:02Z*
* [Leaf](https://github.com/name5566/leaf):star:0 - Lightweight game server framework.  *2017-10-10T10:00:04Z*
* [nano](https://github.com/lonnng/nano):star:429 - Lightweight, facility, high performance golang based game server framework  *2018-03-11T02:35:25Z*
* [Oak](https://github.com/oakmound/oak):star:466 - Pure Go game engine.  *2018-02-07T19:51:16Z*
* [Pixel](https://github.com/faiface/pixel):star:0 - Hand-crafted 2D game library in Go.  *2018-03-29T14:18:46Z*
* [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [termloop](https://github.com/JoelOtter/termloop):star:866 - Terminal-based game engine for Go, built on top of Termbox.  *2017-11-05T18:05:45Z*

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [efaceconv](https://github.com/t0pep0/efaceconv):star:30 - Code generation tool for high performance conversion from interface{} to immutable type without allocations.  *2017-10-12T07:16:31Z*
* [gen](https://github.com/clipperhouse/gen):star:874 - Code generation tool for ‘generics’-like functionality.  *2016-01-08T02:14:47Z*
* [go-enum](https://github.com/abice/go-enum):star:38 - Code generation for enums from code comments.  *2018-04-18T16:07:27Z*
* [go-linq](https://github.com/ahmetalpbalkan/go-linq):star:0 - .NET LINQ-like query methods for Go.  *2017-02-26T20:32:25Z*
* [goderive](https://github.com/awalterschulze/goderive):star:554 - Derives functions from input types.  *2017-12-09T15:02:55Z*
* [interfaces](https://github.com/rjeczalik/interfaces):star:138 - Command line tool for generating interface definitions.  *2017-06-14T12:01:17Z*
* [jennifer](https://github.com/dave/jennifer):star:677 - Generate arbitrary Go code without templates.  *2018-03-20T06:18:48Z*
* [pkgreflect](https://github.com/ungerik/pkgreflect):star:61 - Go preprocessor for package scoped reflection.  *2017-09-05T12:27:26Z*

## Geographic

*Geographic tools and servers*

* [geocache](https://github.com/melihmucuk/geocache):star:82 - In-memory cache that is suitable for geolocation based applications.  *2016-06-21T16:53:17Z*
* [osm](https://github.com/paulmach/osm):star:6 - Library for reading, writing and working with OpenStreetMap data and APIs.  *2018-03-29T04:42:29Z*
* [pbf](https://github.com/maguro/pbf):star:6 - OpenStreetMap PBF golang encoder/decoder.  *2018-01-02T16:04:06Z*
* [S2 geometry](https://github.com/golang/geo):star:612 - S2 geometry library in Go.  *2018-03-11T23:59:00Z*
* [Tile38](https://github.com/tidwall/tile38):star:0 - Geolocation DB with spatial index and realtime geofencing.  *2018-04-20T18:15:07Z*

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs):star:0 - Compiler from Go to JavaScript.  *2018-04-08T17:42:01Z*
* [llgo](https://github.com/go-llvm/llgo):star:883 - LLVM-based compiler for Go.  *2015-01-05T01:15:37Z*
* [tardisgo](https://github.com/tardisgo/tardisgo):star:364 - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.  *2016-11-19T18:08:38Z*

## Goroutines

*Tools for managing and working with Goroutines.*

* [cyclicbarrier](https://github.com/marusama/cyclicbarrier):star:9 - CyclicBarrier for golang.  *2018-01-31T13:52:41Z*
* [go-floc](https://github.com/workanator/go-floc):star:136 - Orchestrate goroutines with ease.  *2018-01-16T07:13:36Z*
* [go-flow](https://github.com/kamildrazkiewicz/go-flow):star:73 - Control goroutines execution order.  *2017-09-19T07:20:06Z*
* [GoSlaves](https://github.com/themester/GoSlaves):star:18 - Simple and Asynchronous Goroutine pool library.  *2018-04-17T01:35:58Z*
* [goworker](https://github.com/benmanns/goworker):star:0 - goworker is a Go-based background worker.  *2017-03-29T11:16:05Z*
* [grpool](https://github.com/ivpusic/grpool):star:330 - Lightweight Goroutine pool.  *2017-08-04T09:21:34Z*
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn):star:13 - Run functions in parallel.  *2018-01-01T20:34:46Z*
* [pool](https://github.com/go-playground/pool):star:364 - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.  *2016-08-23T19:35:13Z*
* [semaphore](https://github.com/kamilsk/semaphore):star:27 - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.  *2018-03-06T20:59:58Z*
* [semaphore](https://github.com/marusama/semaphore):star:27 - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).  *2018-01-18T11:13:39Z*
* [tunny](https://github.com/Jeffail/tunny):star:754 - Goroutine pool for golang.  *2018-03-04T20:46:16Z*
* [worker-pool](https://github.com/vardius/worker-pool):star:10 - goworker is a Go simple async worker pool.  *2018-03-13T08:27:41Z*
* [workerpool](https://github.com/gammazero/workerpool):star:14 - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.  *2018-01-03T20:36:09Z*

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [app](https://github.com/murlokswarm/app):star:0 - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.  *2018-04-13T23:39:58Z*
* [go-astilectron](https://github.com/asticode/go-astilectron):star:0 - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).  *2018-04-21T16:20:57Z*
* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter):star:954 - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.  *2018-04-20T19:31:31Z*
* [gotk3](https://github.com/gotk3/gotk3):star:411 - Go bindings for GTK3.  *2018-04-09T20:08:03Z*
* [gowd](https://github.com/dtylman/gowd):star:104 - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.  *2017-10-16T06:35:07Z*
* [qt](https://github.com/therecipe/qt):star:0 - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).  *2018-04-18T13:32:40Z*
* [ui](https://github.com/andlabs/ui):star:0 - Platform-native GUI library for Go. Cross platform.  *2018-03-28T00:19:27Z*
* [walk](https://github.com/lxn/walk):star:0 - Windows application library kit for Go.  *2018-04-19T09:42:52Z*
* [webview](https://github.com/zserge/webview):star:0 - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).  *2018-03-25T11:24:30Z*

*Interaction*

* [gosx-notifier](https://github.com/deckarep/gosx-notifier):star:422 - OSX Desktop Notifications library for Go.  *2018-02-01T03:58:17Z*
* [robotgo](https://github.com/go-vgo/robotgo):star:0 - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.  *2018-04-20T15:31:22Z*
* [systray](https://github.com/getlantern/systray):star:382 - Cross platform Go library to place an icon and menu in the notification area.  *2018-03-20T02:04:50Z*
* [trayhost](https://github.com/shurcooL/trayhost):star:121 - Cross-platform Go library to place an icon in the host operating system's taskbar.  *2017-02-19T23:43:24Z*


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [bild](https://github.com/anthonynsimon/bild):star:0 - Collection of image processing algorithms in pure Go.  *2018-04-06T13:42:23Z*
* [bimg](https://github.com/h2non/bimg):star:508 - Small package for fast and efficient image processing using libvips.  *2018-03-01T16:19:43Z*
* [geopattern](https://github.com/pravj/geopattern):star:944 - Create beautiful generative image patterns from a string.  *2017-05-09T18:42:50Z*
* [gg](https://github.com/fogleman/gg):star:0 - 2D rendering in pure Go.  *2018-03-16T14:14:17Z*
* [gift](https://github.com/disintegration/gift):star:0 - Package of image processing filters.  *2017-10-26T20:55:04Z*
* [go-cairo](https://github.com/ungerik/go-cairo):star:69 - Go binding for the cairo graphics library.  *2018-04-16T13:30:00Z*
* [go-gd](https://github.com/bolknote/go-gd):star:42 - Go binding for GD library.  *2018-03-31T19:25:05Z*
* [go-nude](https://github.com/koyachi/go-nude):star:247 - Nudity detection with Go.  *2015-04-10T13:47:30Z*
* [go-opencv](https://github.com/lazywei/go-opencv):star:914 - Go bindings for OpenCV.  *2017-08-17T08:25:32Z*
* [go-webcolors](https://github.com/jyotiska/go-webcolors):star:22 - Port of webcolors library from Python to Go.  *2015-08-21T04:56:56Z*
* [gocv](https://github.com/hybridgroup/gocv):star:865 - Go package for computer vision using OpenCV 3.3+.  *2018-03-26T12:59:29Z*
* [goimagehash](https://github.com/corona10/goimagehash):star:81 - Go Perceptual image hashing package.  *2018-04-01T07:55:03Z*
* [govatar](https://github.com/o1egl/govatar):star:234 - Library and CMD tool for generating funny avatars.  *2018-04-13T00:07:57Z*
* [imagick](https://github.com/gographics/imagick):star:718 - Go binding to ImageMagick's MagickWand C API.  *2018-04-02T11:05:04Z*
* [imaginary](https://github.com/h2non/imaginary):star:0 - Fast and simple HTTP microservice for image resizing.  *2018-04-15T10:27:50Z*
* [imaging](https://github.com/disintegration/imaging):star:0 - Simple Go image processing package.  *2018-04-11T22:55:52Z*
* [img](https://github.com/hawx/img):star:116 - Selection of image manipulation tools.  *2015-05-01T15:11:19Z*
* [ln](https://github.com/fogleman/ln):star:0 - 3D line art rendering in Go.  *2017-02-23T13:55:21Z*
* [mpo](https://github.com/donatj/mpo):star:3 - Decoder and conversion tool for MPO 3D Photos.  *2017-10-13T19:21:16Z*
* [picfit](https://github.com/thoas/picfit):star:828 - An image resizing server written in Go.  *2018-02-27T22:58:02Z*
* [pt](https://github.com/fogleman/pt):star:0 - Path tracing engine written in Go.  *2017-06-19T01:24:16Z*
* [resize](https://github.com/nfnt/resize):star:0 - Image resizing for Go with common interpolation methods.  *2018-02-21T19:10:11Z*
* [rez](https://github.com/bamiaux/rez):star:144 - Image resizing in pure Go and SIMD.  *2017-07-31T18:41:18Z*
* [smartcrop](https://github.com/muesli/smartcrop):star:0 - Finds good crops for arbitrary images and crop sizes.  *2018-02-28T07:50:44Z*
* [svgo](https://github.com/ajstarks/svgo):star:0 - Go Language Library for SVG generation.  *2018-02-26T02:51:33Z*
* [tga](https://github.com/ftrvxmtrx/tga):star:15 - Package tga is a TARGA image format decoder/encoder.  *2015-05-24T08:11:24Z*

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [connectordb](https://github.com/connectordb/connectordb):star:97 - Open-Source Platform for Quantified Self & IoT.  *2018-04-10T06:51:05Z*
* [devices](https://github.com/goiot/devices):star:201 - Suite of libraries for IoT devices, experimental for x/exp/io.  *2016-07-08T21:40:26Z*
* [eywa](https://github.com/xcodersun/eywa):star:16 - Project Eywa is essentially a connection manager that keeps track of connected devices.  *2017-04-12T07:37:10Z*
* [flogo](https://github.com/tibcosoftware/flogo):star:561 - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.  *2018-04-22T00:06:59Z*
* [gatt](https://github.com/paypal/gatt):star:654 - Gatt is a Go package for building Bluetooth Low Energy peripherals.  *2015-10-11T22:09:35Z*
* [gobot](https://github.com/hybridgroup/gobot/):star:0 - Gobot is a framework for robotics, physical computing, and the Internet of Things.  *2018-04-20T14:58:56Z*
* [iot](https://github.com/vaelen/iot/):star:14 - IoT is a simple framework for implementing a Google IoT Core device.  *2018-04-06T07:22:28Z*
* [mainflux](https://github.com/Mainflux/mainflux):star:444 - Industrial IoT Messaging and Device Management Server.  *2018-04-19T11:40:22Z*
* [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.
* [sensorbee](https://github.com/sensorbee/sensorbee):star:137 - Lightweight stream processing engine for IoT.  *2017-12-31T00:14:28Z*

## Logging

*Libraries for generating and working with log files.*

* [distillog](https://github.com/amoghe/distillog):star:5 - distilled levelled logging (think of it as stdlib + log levels).  *2017-11-28T07:58:41Z*
* [glg](https://github.com/kpango/glg):star:10 - glg is simple and fast leveled logging library for Go.  *2017-09-15T10:40:56Z*
* [glog](https://github.com/golang/glog):star:0 - Leveled execution logs for Go.  *2016-01-25T20:49:56Z*
* [go-cronowriter](https://github.com/utahta/go-cronowriter):star:8 - Simple writer that rotate log files automatically based on current date and time, like cronolog.  *2018-02-21T05:11:55Z*
* [go-log](https://github.com/siddontang/go-log):star:15 - Log lib supports level and multi handlers.  *2014-08-18T08:44:07Z*
* [go-log](https://github.com/ian-kent/go-log):star:26 - Log4j implementation in Go.  *2016-01-13T21:12:17Z*
* [go-logger](https://github.com/apsdehal/go-logger):star:176 - Simple logger of Go Programs, with level handlers.  *2018-02-21T16:48:00Z*
* [gologger](https://github.com/sadlil/gologger):star:34 - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.  *2018-01-31T03:17:57Z*
* [gomol](https://github.com/aphistic/gomol):star:9 - Multiple-output, structured logging for Go with extensible logging outputs.  *2018-03-06T16:35:34Z*
* [gone/log](https://github.com/One-com/gone/tree/master/log):star:21 - Fast, extendable, full-featured, std-lib source compatible log library.  *2017-09-06T13:52:59Z*
* [journald](https://github.com/ssgreg/journald):star:4 - Go implementation of systemd Journal's native API for logging.  *2017-12-05T05:50:22Z*
* [log](https://github.com/apex/log):star:531 - Structured logging package for Go.  *2018-01-18T05:06:19Z*
* [log](https://github.com/go-playground/log):star:245 - Simple, configurable and scalable Structured Logging for Go.  *2018-04-05T16:14:17Z*
* [log](https://github.com/teris-io/log):star:17 - Structured log interface for Go cleanly separates logging facade from its implementation.  *2017-12-04T18:53:44Z*
* [log-voyage](https://github.com/firstrow/logvoyage):star:75 - Full-featured logging saas written in golang.  *2017-05-24T19:48:17Z*
* [log15](https://github.com/inconshreveable/log15):star:725 - Simple, powerful logging for Go.  *2016-11-12T20:41:34Z*
* [logdump](https://github.com/ewwwwwqm/logdump):star:4 - Package for multi-level logging.  *2018-04-02T00:28:11Z*
* [logex](https://github.com/chzyer/logex):star:33 - Golang log lib, supports tracking and level, wrap by standard log lib.  *2017-03-29T06:48:20Z*
* [logger](https://github.com/azer/logger):star:114 - Minimalistic logging library for Go.  *2018-01-30T04:12:15Z*
* [logo](https://github.com/mbndr/logo):star:2 - Golang logger to different configurable writers.  *2017-10-19T19:33:22Z*
* [logrus](https://github.com/Sirupsen/logrus):star:0 - Structured logger for Go.  *2018-03-29T22:59:52Z*
* [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
* [logutils](https://github.com/hashicorp/logutils):star:212 - Utilities for slightly better logging in Go (Golang) extending the standard logger.  *2015-06-09T07:04:31Z*
* [logxi](https://github.com/mgutz/logxi):star:324 - 12-factor app logger that is fast and makes you happy.  *2016-10-27T14:08:23Z*
* [lumberjack](https://github.com/natefinch/lumberjack):star:797 - Simple rolling logger, implements io.WriteCloser.  *2017-09-11T14:04:57Z*
* [mlog](https://github.com/jbrodriguez/mlog):star:13 - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.  *2016-05-01T15:51:40Z*
* [ozzo-log](https://github.com/go-ozzo/ozzo-log):star:86 - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).  *2016-07-03T17:57:02Z*
* [seelog](https://github.com/cihub/seelog):star:0 - Logging functionality with flexible dispatching, filtering, and formatting.  *2017-01-30T13:45:32Z*
* [spew](https://github.com/davecgh/go-spew):star:0 - Implements a deep pretty printer for Go data structures to aid in debugging.  *2018-02-21T22:46:20Z*
* [stdlog](https://github.com/alexcesaro/log):star:44 - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.  *2015-09-15T19:59:37Z*
* [tail](https://github.com/hpcloud/tail):star:0 - Go package striving to emulate the features of the BSD tail program.  *2017-08-14T16:06:53Z*
* [xlog](https://github.com/xfxdev/xlog):star:2 - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.  *2017-01-10T16:36:36Z*
* [xlog](https://github.com/rs/xlog):star:121 - Structured logger for `net/context` aware HTTP handlers with flexible dispatching.  *2017-12-27T18:52:59Z*
* [zap](https://github.com/uber-go/zap):star:0 - Fast, structured, leveled logging in Go.  *2018-04-13T23:02:07Z*
* [zerolog](https://github.com/rs/zerolog):star:0 - Zero-allocation JSON logger.  *2018-04-19T20:12:29Z*

## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian):star:545 - Naive Bayesian Classification for Golang.  *2016-12-10T17:52:30Z*
* [CloudForest](https://github.com/ryanbressler/CloudForest):star:586 - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.  *2016-12-01T19:44:07Z*
* [fonet](https://github.com/Fontinalis/fonet):star:12 - A Deep Neural Network library written in Go.  *2018-04-10T13:24:15Z*
* [gago](https://github.com/MaxHalford/gago):star:458 - Multi-population, flexible, parallel genetic algorithm.  *2018-04-17T07:15:31Z*
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster):star:10 - Go implementation of the k-modes and k-prototypes clustering algorithms.  *2017-11-09T16:40:40Z*
* [go-deep](https://github.com/patrikeh/go-deep):star:167 - A feature-rich neural network library in Go.  *2018-02-21T15:28:38Z*
* [go-fann](https://github.com/white-pony/go-fann):star:94 - Go bindings for Fast Artificial Neural Networks(FANN) library.  *2015-02-03T21:53:31Z*
* [go-galib](https://github.com/thoj/go-galib):star:156 - Genetic Algorithms library written in Go / golang.  *2015-12-28T16:27:45Z*
* [go-pr](https://github.com/daviddengcn/go-pr):star:50 - Pattern recognition package in Go lang.  *2013-06-08T10:16:25Z*
* [gobrain](https://github.com/goml/gobrain):star:243 - Neural Networks written in go.  *2017-04-26T00:10:04Z*
* [godist](https://github.com/e-dard/godist):star:18 - Various probability distributions, and associated methods.  *2015-05-11T10:38:44Z*
* [goga](https://github.com/tomcraven/goga):star:64 - Genetic algorithm library for Go.  *2016-04-17T17:35:15Z*
* [GoLearn](https://github.com/sjwhitworth/golearn):star:0 - General Machine Learning library for Go.  *2018-03-24T00:32:46Z*
* [golinear](https://github.com/danieldk/golinear):star:33 - liblinear bindings for Go.  *2017-04-18T08:41:42Z*
* [goml](https://github.com/cdipaolo/goml):star:852 - On-line Machine Learning in Go.  *2016-10-30T20:48:43Z*
* [goRecommend](https://github.com/timkaye11/goRecommend):star:110 - Recommendation Algorithms library written in Go.  *2014-07-29T04:49:34Z*
* [gorgonia](https://github.com/chewxy/gorgonia):star:0 - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.  *2018-03-27T11:27:56Z*
* [goscore](https://github.com/asafschers/goscore):star:12 - Go Scoring API for PMML.  *2018-04-04T18:20:55Z*
* [gosseract](https://github.com/otiai10/gosseract):star:489 - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.  *2018-03-21T17:42:05Z*
* [libsvm](https://github.com/datastream/libsvm):star:55 - libsvm golang version derived work based on LIBSVM 3.14.  *2016-05-09T03:47:10Z*
* [mlgo](https://github.com/NullHypothesis/mlgo):star:3 - This project aims to provide minimalistic machine learning algorithms in Go.  *2015-12-07T16:08:14Z*
* [neat](https://github.com/jinyeom/neat):star:39 - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).  *2017-07-03T11:55:47Z*
* [neural-go](https://github.com/schuyler/neural-go):star:56 - Multilayer perceptron network implemented in Go, with training via backpropagation.  *2013-10-18T15:30:23Z*
* [probab](https://github.com/ThePaw/probab):star:8 - Probability distribution functions. Bayesian inference. Written in pure Go.  *2013-08-02T11:19:58Z*
* [regommend](https://github.com/muesli/regommend):star:204 - Recommendation & collaborative filtering engine.  *2018-02-21T03:32:59Z*
* [shield](https://github.com/eaigner/shield):star:114 - Bayesian text classifier with flexible tokenizers and storage backends for Go.  *2013-04-15T14:22:53Z*
* [tfgo](https://github.com/galeone/tfgo):star:757 - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.  *2018-04-14T08:02:37Z*
* [Varis](https://github.com/Xamber/Varis):star:9 - Golang Neural Network.  *2017-12-13T10:54:48Z*

## Messaging

*Libraries that implement messaging systems.*

* [Benthos](https://github.com/Jeffail/benthos):star:203 - A message streaming bridge between a range of protocols.  *2018-04-21T11:46:07Z*
* [Centrifugo](https://github.com/centrifugal/centrifugo):star:0 - Real-time messaging (Websockets or SockJS) server in Go.  *2018-04-21T10:04:19Z*
* [dbus](https://github.com/godbus/dbus):star:210 - Native Go bindings for D-Bus.  *2018-04-06T11:14:27Z*
* [drone-line](https://github.com/appleboy/drone-line) - Sending [Line](https://business.line.me/en/services/bot) notifications using a binary, docker or Drone CI.
* [emitter](https://github.com/olebedev/emitter):star:197 - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.  *2017-12-21T09:59:05Z*
* [event](https://github.com/agoalofalife/event):star:11 - Implementation of the pattern observer.  *2018-02-19T12:11:26Z*
* [EventBus](https://github.com/asaskevich/EventBus):star:357 - The lightweight event bus with async compatibility.  *2018-03-15T14:05:47Z*
* [gaurun-client](https://github.com/osamingo/gaurun-client):star:6 - Gaurun Client written in Go.  *2017-07-03T14:02:03Z*
* [Glue](https://github.com/desertbit/glue):star:277 - Robust Go and Javascript Socket Library (Alternative to Socket.io).  *2017-10-18T14:27:42Z*
* [go-notify](https://github.com/TheCreeper/go-notify):star:32 - Native implementation of the freedesktop notification spec.  *2016-02-03T00:10:56Z*
* [go-nsq](https://github.com/nsqio/go-nsq):star:0 - the official Go package for NSQ.  *2017-10-30T00:05:02Z*
* [go-socket.io](https://github.com/googollee/go-socket.io):star:0 - socket.io library for golang, a realtime application framework.  *2018-04-20T11:43:31Z*
* [go-vitotrol](https://github.com/maxatome/go-vitotrol):star:7 - Client library to Viessmann Vitotrol web service.  *2017-12-15T21:43:25Z*
* [Gollum](https://github.com/trivago/gollum):star:604 - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.  *2018-03-30T19:11:42Z*
* [golongpoll](https://github.com/jcuga/golongpoll):star:367 - HTTP longpoll server library that makes web pub-sub simple.  *2016-08-21T02:51:52Z*
* [goose](https://github.com/ian-kent/goose):star:32 - Server Sent Events in Go.  *2014-12-21T09:00:59Z*
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster):star:0 - gopush-cluster is a go push server cluster.  *2017-05-25T04:25:54Z*
* [gorush](https://github.com/appleboy/gorush):star:97 - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).  *2017-02-14T17:04:21Z*
* [guble](https://github.com/smancke/guble):star:120 - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.  *2017-10-31T19:15:40Z*
* [machinery](https://github.com/RichardKnop/machinery):star:0 - Asynchronous task queue/job queue based on distributed message passing.  *2018-04-22T03:04:23Z*
* [mangos](https://github.com/go-mangos/mangos):star:0 - Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.  *2018-04-11T17:31:04Z*
* [melody](https://github.com/olahol/melody):star:907 - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.  *2018-02-27T13:42:53Z*
* [messagebus](https://github.com/vardius/message-bus):star:14 - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.  *2018-03-13T08:28:18Z*
* [NATS Go Client](https://github.com/nats-io/nats):star:0 - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.  *2018-03-31T19:16:09Z*
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus):star:28 - A tiny wrapper around NSQ topic and channel.  *2018-02-15T22:29:58Z*
* [oplog](https://github.com/dailymotion/oplog):star:87 - Generic oplog/replication system for REST APIs.  *2015-11-07T00:50:07Z*
* [pubsub](https://github.com/tuxychandru/pubsub):star:188 - Simple pubsub package for go.  *2018-02-25T01:06:11Z*
* [rabbus](https://github.com/rafaeljesus/rabbus):star:37 - A tiny wrapper over amqp exchanges and queues.  *2018-04-20T20:44:16Z*
* [rabtap](https://github.com/jandelgado/rabtap):star:15 - RabbitMQ swiss army knife cli app.  *2018-04-08T20:05:06Z*
* [RapidMQ](https://github.com/sybrexsys/RapidMQ):star:37 - RapidMQ is a lightweight and reliable library for managing of the local messages queue.  *2017-12-07T08:34:02Z*
* [sarama](https://github.com/Shopify/sarama):star:0 - Go library for Apache Kafka.  *2018-04-18T15:09:32Z*
* [Uniqush-Push](https://github.com/uniqush/uniqush-push):star:962 - Redis backed unified push service for server-side notifications to mobile devices.  *2018-04-01T21:41:21Z*
* [zmq4](https://github.com/pebbe/zmq4):star:15 - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).  *2018-04-10T11:05:40Z*

## Miscellaneous

*These libraries were placed here because none of the other categories seemed to fit.*

* [alice](https://github.com/magic003/alice):star:14 - Additive dependency injection container for Golang.  *2017-04-26T06:23:16Z*
* [anagent](https://github.com/mudler/anagent):star:2 - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.  *2018-02-01T21:42:46Z*
* [antch](https://github.com/antchfx/antch):star:67 - A fast, powerful and extensible web crawling & scraping framework.  *2017-12-25T03:45:34Z*
* [archiver](https://github.com/mholt/archiver):star:852 - Library and command for making and extracting .zip and .tar.gz archives.  *2018-04-17T22:02:35Z*
* [autoflags](https://github.com/artyom/autoflags):star:19 - Go package to automatically define command line flags from struct fields.  *2017-03-01T06:33:24Z*
* [avgRating](https://github.com/kirillDanshin/avgRating):star:5 - Calculate average score and rating based on Wilson Score Equation.  *2017-08-05T19:37:34Z*
* [banner](https://github.com/dimiro1/banner):star:160 - Add beautiful banners into your Go applications.  *2016-11-08T15:12:23Z*
* [base64Captcha](https://github.com/mojocn/base64Captcha):star:318 - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.  *2018-04-08T03:54:05Z*
* [battery](https://github.com/distatus/battery):star:87 - Cross-platform, normalized battery information library.  *2017-05-21T01:04:19Z*
* [bitio](https://github.com/icza/bitio):star:47 - Highly optimized bit-level Reader and Writer for Go.  *2018-02-21T12:02:00Z*
* [browscap_go](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](http://browscap.org/).
* [captcha](https://github.com/steambap/captcha):star:17 - Package captcha provides an easy to use, unopinionated API for captcha generation.  *2018-03-26T03:08:21Z*
* [conv](https://github.com/cstockton/go-conv):star:317 - Package conv provides fast and intuitive conversions across Go types.  *2017-05-24T00:24:50Z*
* [datacounter](https://github.com/miolini/datacounter):star:21 - Go counters for readers/writer/http.ResponseWriter.  *2017-11-04T15:29:33Z*
* [errors](https://github.com/pkg/errors):star:0 - Package that provides simple error handling primitives.  *2018-03-11T21:45:15Z*
* [go-chat-bot](https://github.com/go-chat-bot/bot):star:322 - IRC, Slack & Telegram bot written in Go.  *2018-03-18T12:59:23Z*
* [go-commons-pool](https://github.com/jolestar/go-commons-pool):star:475 - Generic object pool for Golang.  *2017-08-09T06:54:32Z*
* [go-multierror](https://github.com/hashicorp/go-multierror):star:443 - Go (golang) package for representing a list of errors as a single error.  *2017-12-04T18:29:08Z*
* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.
* [go-resiliency](https://github.com/eapache/go-resiliency):star:608 - Resiliency patterns for golang.  *2018-03-26T13:24:23Z*
* [go-sarah](https://github.com/oklahomer/go-sarah):star:92 - Framework to build bot for desired chat services including LINE, Slack, Gitter and more.  *2018-03-08T15:34:17Z*
* [go-unarr](https://github.com/gen2brain/go-unarr):star:28 - Decompression library for RAR, TAR, ZIP and 7z archives.  *2017-11-13T23:52:08Z*
* [go.uuid](https://github.com/satori/go.uuid):star:0 - Implementation of Universally Unique Identifier (UUID). Supported both creation and parsing of UUIDs.  *2018-01-03T17:44:51Z*
* [gofakeit](https://github.com/brianvoe/gofakeit):star:107 - Random data generator written in go.  *2018-04-17T19:36:59Z*
* [goid](https://github.com/jakehl/goid):star:12 - Generate and Parse RFC4122 compliant V4 UUIDs.  *2017-06-03T15:56:40Z*
* [gopsutil](https://github.com/shirou/gopsutil):star:0 - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).  *2018-04-17T02:11:51Z*
* [gosms](https://github.com/haxpax/gosms):star:0 - Your own local SMS gateway in Go that can be used to send SMS.  *2017-07-07T11:45:17Z*
* [gountries](https://github.com/pariz/gountries):star:159 - Package that exposes country and subdivision data.  *2017-10-19T11:17:38Z*
* [hanu](https://github.com/sbstjn/hanu):star:48 - Framework for writing Slack bots.  *2017-06-03T21:35:53Z*
* [health](https://github.com/dimiro1/health):star:279 - Easy to use, extensible health check library.  *2017-01-02T17:58:17Z*
* [healthcheck](https://github.com/etherlabsio/healthcheck):star:36 - An opinionated and concurrent health-check HTTP handler for RESTful services.  *2018-03-09T05:13:18Z*
* [hostutils](https://github.com/Wing924/hostutils):star:3 - A golang library for packing and unpacking FQDNs list.  *2017-09-26T08:43:42Z*
* [indigo](https://github.com/osamingo/indigo):star:34 - Distributed unique ID generator of using Sonyflake and encoded by Base58.  *2018-04-02T08:51:37Z*
* [jobs](https://github.com/albrow/jobs):star:422 - Persistent and flexible background jobs library.  *2017-12-14T01:02:15Z*
* [lk](https://github.com/hyperboloide/lk):star:43 - A simple licensing library for golang.  *2018-02-19T15:52:03Z*
* [margelet](https://github.com/zhulik/margelet):star:46 - Framework for building Telegram bots.  *2016-09-18T11:47:01Z*
* [pdfgen](https://github.com/hyperboloide/pdfgen):star:7 - HTTP service to generate PDF from Json requests.  *2018-02-19T15:49:35Z*
* [persian](https://github.com/mavihq/persian):star:8 - Some utilities for Persian language in go.  *2017-12-03T07:19:40Z*
* [secdl](https://github.com/xor-gate/secdl):star:6 - Lighttpd ModSecDownload algorithm ported to go to secure download urls.  *2017-08-12T22:44:23Z*
* [shellwords](https://github.com/Wing924/shellwords):star:1 - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.  *2017-10-02T01:58:48Z*
* [shortid](https://github.com/teris-io/shortid):star:247 - Distributed generation of super short, unique, non-sequential, URL friendly IDs.  *2017-10-29T13:18:06Z*
* [slacker](https://github.com/shomali11/slacker):star:177 - Easy to use framework to create Slack bots.  *2018-04-07T02:43:28Z*
* [stats](https://github.com/go-playground/stats):star:97 - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...  *2016-09-07T12:50:21Z*
* [turtle](https://github.com/hackebrot/turtle):star:39 - Emojis for Go.  *2017-10-05T16:45:02Z*
* [uuid](https://github.com/agext/uuid):star:4 - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.  *2017-06-21T18:45:11Z*
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler):star:2 - Generate boilerplate http input and ouput handling.  *2016-04-11T16:09:01Z*
* [werr](https://github.com/txgruppi/werr):star:5 - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.  *2016-03-10T03:36:54Z*
* [wuid](https://github.com/edwingeng/wuid):star:148 - An extremely fast unique number generator, 10-135 times faster than UUID.  *2018-04-12T09:09:27Z*
* [xkg](https://github.com/go-xkg/xkg):star:33 - X Keyboard Grabber.  *2015-01-08T04:00:34Z*
* [xstrings](https://github.com/huandu/xstrings):star:463 - Collection of useful string functions ported from other languages.  *2018-02-14T04:29:31Z*

## Natural Language Processing

*Libraries for working with human languages.*

* [dpar](https://github.com/danieldk/dpar/):star:23 - Transition-based statistical dependency parser.  *2018-02-22T07:25:03Z*
* [getlang](https://github.com/rylans/getlang):star:10 - Fast natural language detection package.  *2018-03-26T17:12:30Z*
* [go-eco](https://github.com/ThePaw/go-eco):star:3 - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.  *2014-05-27T13:11:00Z*
* [go-i18n](https://github.com/nicksnyder/go-i18n/):star:594 - Package and an accompanying tool to work with localized text.  *2018-04-19T05:09:40Z*
* [go-mystem](https://github.com/dveselov/mystem):star:15 - CGo bindings to Yandex.Mystem - russian morphology analyzer.  *2016-10-05T05:53:17Z*
* [go-nlp](https://github.com/nuance/go-nlp):star:74 - Utilities for working with discrete probability distributions and other tools useful for doing NLP work.  *2011-11-15T17:49:50Z*
* [go-stem](https://github.com/agonopol/go-stem):star:43 - Implementation of the porter stemming algorithm.  *2015-06-30T11:33:28Z*
* [go-unidecode](https://github.com/mozillazg/go-unidecode):star:23 - ASCII transliterations of Unicode text.  *2016-11-07T13:59:01Z*
* [go2vec](https://github.com/danieldk/go2vec):star:27 - Reader and utility functions for word2vec embeddings.  *2017-05-17T12:05:21Z*
* [gojieba](https://github.com/yanyiwu/gojieba):star:0 - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.  *2017-08-28T13:40:18Z*
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer):star:13 - Go bindings for the snowball libstemmer library including porter 2.  *2014-06-17T16:04:55Z*
* [gounidecode](https://github.com/fiam/gounidecode):star:58 - Unicode transliterator (also known as unidecode) for Go.  *2015-06-29T11:25:15Z*
* [gse](https://github.com/go-ego/gse):star:199 - Go efficient text segmentation; support english, chinese, japanese and other.  *2018-04-19T16:20:32Z*
* [icu](https://github.com/goodsign/icu):star:16 - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.  *2012-12-27T17:27:29Z*
* [libtextcat](https://github.com/goodsign/libtextcat):star:9 - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.  *2012-12-27T17:23:17Z*
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
* [nlp](https://github.com/Shixzie/nlp):star:329 - Extract values from strings and fill your structs with nlp.  *2017-09-18T14:32:03Z*
* [nlp](https://github.com/james-bowman/nlp):star:122 - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).  *2018-04-21T18:37:18Z*
* [paicehusk](https://github.com/rookii/paicehusk):star:22 - Golang implementation of the Paice/Husk Stemming Algorithm.  *2013-06-19T05:40:30Z*
* [petrovich](https://github.com/striker2000/petrovich):star:11 - Petrovich is the library which inflects Russian names to given grammatical case.  *2017-10-12T23:03:20Z*
* [porter](https://github.com/a2800276/porter):star:6 - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.  *2013-10-03T10:46:03Z*
* [porter2](https://github.com/zhenjl/porter2):star:30 - Really fast Porter 2 stemmer.  *2015-08-29T21:01:52Z*
* [prose](https://github.com/jdkato/prose):star:559 - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more.  *2017-10-30T20:05:14Z*
* [RAKE.go](https://github.com/Obaied/RAKE.go):star:30 - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).  *2017-08-07T08:00:34Z*
* [segment](https://github.com/blevesearch/segment):star:37 - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)  *2016-09-15T18:50:41Z*
* [sentences](https://github.com/neurosnap/sentences):star:230 - Sentence tokenizer:  converts text into a list of sentences.  *2017-09-02T11:52:27Z*
* [shamoji](https://github.com/osamingo/shamoji):star:6 - The shamoji is word filtering package written in Go.  *2017-11-08T05:45:16Z*
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).
* [stemmer](https://github.com/dchest/stemmer):star:37 - Stemmer packages for Go programming language. Includes English and German stemmers.  *2016-12-07T10:24:02Z*
* [textcat](https://github.com/pebbe/textcat):star:55 - Go package for n-gram based text categorization, with support for utf-8 and raw text.  *2016-11-09T11:59:50Z*
* [whatlanggo](https://github.com/abadojack/whatlanggo):star:273 - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).  *2018-02-10T19:25:21Z*
* [when](https://github.com/olebedev/when):star:852 - Natural EN and RU language date/time parser with pluggable rules.  *2017-10-23T12:05:03Z*

## Networking

*Libraries for working with various layers of the network.*

* [arp](https://github.com/mdlayher/arp):star:129 - Package arp implements the ARP protocol, as described in RFC 826.  *2017-12-14T19:58:10Z*
* [buffstreams](https://github.com/stabbycutyou/buffstreams):star:194 - Streaming protocolbuffer data over TCP made easy.  *2016-04-16T13:46:48Z*
* [canopus](https://github.com/zubairhamed/canopus):star:98 - CoAP Client/Server implementation (RFC 7252).  *2018-02-07T12:39:28Z*
* [cidranger](https://github.com/yl2chen/cidranger):star:260 - Fast IP to CIDR lookup for Go.  *2018-02-14T08:18:26Z*
* [dhcp6](https://github.com/mdlayher/dhcp6):star:47 - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.  *2017-12-12T14:39:16Z*
* [dns](https://github.com/miekg/dns):star:0 - Go library for working with DNS.  *2018-04-06T15:09:55Z*
* [ether](https://github.com/songgao/ether):star:47 - Cross-platform Go package for sending and receiving ethernet frames.  *2016-04-05T02:55:33Z*
* [ethernet](https://github.com/mdlayher/ethernet):star:138 - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.  *2017-07-07T21:33:43Z*
* [fasthttp](https://github.com/valyala/fasthttp):star:0 - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.  *2017-12-07T12:09:41Z*
* [fortio](https://github.com/istio/fortio):star:272 - Load testing library and command line tool and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them.  *2018-04-16T23:55:51Z*
* [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).
* [go-getter](https://github.com/hashicorp/go-getter):star:486 - Go library for downloading files or directories from various sources using a URL.  *2018-04-10T17:49:45Z*
* [go-stun](https://github.com/ccding/go-stun):star:220 - Go implementation of the STUN client (RFC 3489 and RFC 5389).  *2017-12-06T15:03:02Z*
* [gobgp](https://github.com/osrg/gobgp):star:0 - BGP implemented in the Go Programming Language.  *2018-04-13T02:04:27Z*
* [golibwireshark](https://github.com/sunwxg/golibwireshark):star:11 - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.  *2017-10-24T11:56:00Z*
* [gopacket](https://github.com/google/gopacket):star:0 - Go library for packet processing with libpcap bindings.  *2018-04-10T00:34:31Z*
* [gopcap](https://github.com/akrennmair/gopcap):star:305 - Go wrapper for libpcap.  *2015-07-28T16:05:02Z*
* [goshark](https://github.com/sunwxg/goshark):star:6 - Package goshark use tshark to decode IP packet and create data struct to analyse packet.  *2017-10-24T11:36:12Z*
* [gosnmp](https://github.com/soniah/gosnmp):star:303 - Native Go library for performing SNMP actions.  *2018-04-20T13:00:27Z*
* [gotcp](https://github.com/gansidui/gotcp):star:322 - Go package for quickly writing tcp applications.  *2017-04-18T07:26:13Z*
* [grab](https://github.com/cavaliercoder/grab):star:301 - Go package for managing file downloads.  *2018-04-03T00:54:14Z*
* [graval](https://github.com/koofr/graval):star:17 - Experimental FTP server framework.  *2017-10-12T13:09:05Z*
* [jazigo](https://github.com/udhos/jazigo):star:77 - Jazigo is a tool written in Go for retrieving configuration for multiple network devices.  *2018-01-24T01:37:10Z*
* [kcp-go](https://github.com/xtaci/kcp-go):star:0 - KCP - Fast and Reliable ARQ Protocol.  *2018-02-03T13:32:37Z*
* [kcptun](https://github.com/xtaci/kcptun):star:0 - Extremely simple & fast udp tunnel based on KCP protocol.  *2018-03-16T05:44:44Z*
* [lhttp](https://github.com/fanux/lhttp):star:395 - Powerful websocket framework, build your IM server more easily.  *2018-04-08T08:06:08Z*
* [linkio](https://github.com/ian-kent/linkio):star:32 - Network link speed simulation for Reader/Writer interfaces.  *2017-08-07T20:57:55Z*
* [llb](https://github.com/kirillDanshin/llb):star:5 - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.  *2016-04-04T04:13:06Z*
* [mdns](https://github.com/hashicorp/mdns):star:410 - Simple mDNS (Multicast DNS) client/server library in Golang.  *2017-02-21T17:29:40Z*
* [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [portproxy](https://github.com/aybabtme/portproxy):star:34 - Simple TCP proxy which adds CORS support to API's which don't support it.  *2014-12-13T03:04:53Z*
* [publicip](https://github.com/polera/publicip):star:13 - Package publicip returns your public facing IPv4 address (internet egress).  *2016-12-29T04:30:29Z*
* [raw](https://github.com/mdlayher/raw):star:198 - Package raw enables reading and writing data at the device driver level for a network interface.  *2018-03-06T15:27:56Z*
* [sftp](https://github.com/pkg/sftp):star:493 - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.  *2018-04-19T20:08:40Z*
* [ssh](https://github.com/gliderlabs/ssh):star:703 - Higher-level API for building SSH servers (wraps crypto/ssh).  *2018-04-17T01:12:34Z*
* [sslb](https://github.com/eduardonunesp/sslb):star:98 - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.  *2018-03-09T23:00:10Z*
* [stun](https://github.com/go-rtc/stun):star:85 - Go implementation of RFC 5389 STUN protocol.  *2018-04-20T04:38:17Z*
* [tcp_server](https://github.com/firstrow/tcp_server):star:181 - Go library for building tcp servers faster.  *2017-06-07T16:33:39Z*
* [utp](https://github.com/anacrolix/utp):star:136 - Go uTP micro transport protocol implementation.  *2018-02-19T06:06:59Z*
* [water](https://github.com/songgao/water):star:396 - Simple TUN/TAP library.  *2018-04-20T06:47:39Z*
* [winrm](https://github.com/masterzen/winrm):star:147 - Go WinRM client to remotely execute commands on Windows machines.  *2018-02-24T16:03:50Z*
* [xtcp](https://github.com/xfxdev/xtcp):star:37 - TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.  *2017-03-06T16:58:23Z*
* [YANNFF](https://github.com/intel-go/yanff):star:412 - Framework for rapid development of performant network functions for cloud and bare-metal.  *2018-04-16T16:03:30Z*

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl):star:464 - Go bindings for OpenGL (generated via glow).  *2018-04-07T15:57:06Z*
* [glfw](https://github.com/go-gl/glfw):star:468 - Go bindings for GLFW 3.  *2017-08-14T18:07:46Z*
* [goxjs/gl](https://github.com/goxjs/gl):star:115 - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).  *2017-11-28T03:40:04Z*
* [goxjs/glfw](https://github.com/goxjs/glfw):star:50 - Go cross-platform glfw library for creating an OpenGL context and receiving events.  *2017-10-18T04:47:55Z*
* [mathgl](https://github.com/go-gl/mathgl):star:222 - Pure Go math package specialized for 3D math, with inspiration from GLM.  *2018-03-19T21:07:51Z*

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm):star:0 - Powerful orm framework for go. Support: pq/mysql/sqlite3.  *2017-11-30T12:26:34Z*
* [go-pg](https://github.com/go-pg/pg):star:0 - PostgreSQL ORM with focus on PostgreSQL specific features and performance.  *2018-04-08T07:39:39Z*
* [go-queryset](https://github.com/jirfag/go-queryset):star:289 - 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.  *2018-03-09T18:44:49Z*
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder):star:34 - A flexible and powerful SQL string builder library plus a zero-config ORM.  *2018-03-15T12:23:40Z*
* [go-store](https://github.com/gosuri/go-store):star:88 - Simple and fast Redis backed key-value store library for Go.  *2017-02-23T15:11:41Z*
* [gomodel](https://github.com/cosiner/gomodel):star:59 - Lightweight, fast, orm-like library helps interactive with database.  *2017-07-02T10:17:20Z*
* [GORM](https://github.com/jinzhu/gorm):star:0 - The fantastic ORM library for Golang, aims to be developer friendly.  *2018-04-16T14:20:02Z*
* [gorp](https://github.com/go-gorp/gorp):star:0 - Go Relational Persistence, ORM-ish library for Go.  *2018-04-10T15:54:28Z*
* [grimoire](https://github.com/Fs02/grimoire):star:9 - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).  *2018-04-19T15:21:32Z*
* [lore](https://github.com/abrahambotros/lore):star:3 - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.  *2017-10-21T18:26:29Z*
* [Marlow](https://github.com/dadleyy/marlow):star:27 - Generated ORM from project structs for compile time safety assurances.  *2018-03-17T16:04:11Z*
* [pop/soda](https://github.com/markbates/pop):star:563 - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.  *2018-04-11T19:05:04Z*
* [QBS](https://github.com/coocood/qbs):star:494 - Stands for Query By Struct. A Go ORM.  *2017-04-18T01:16:07Z*
* [reform](https://github.com/go-reform/reform):star:638 - Better ORM for Go, based on non-empty interfaces and code generation.  *2017-12-07T11:26:00Z*
* [SQLBoiler](https://github.com/volatiletech/sqlboiler):star:0 - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.  *2018-03-05T16:08:25Z*
* [upper.io/db](https://github.com/upper/db):star:0 - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.  *2018-02-11T22:07:03Z*
* [Xorm](https://github.com/go-xorm/xorm):star:0 - Simple and powerful ORM for Go.  *2018-04-19T00:25:31Z*
* [Zoom](https://github.com/albrow/zoom):star:197 - Blazing-fast datastore and querying engine built on Redis.  *2018-03-14T21:23:33Z*

## Package Management

*Libraries for package and dependency management.*

* [dep](https://github.com/golang/dep):star:0 - Go dependency tool.  *2018-04-19T14:13:32Z*
* [gigo](https://github.com/LyricalSecurity/gigo):star:190 - PIP-like dependency tool for golang, with support for private repositories and hashes.  *2015-08-06T13:08:18Z*
* [glide](https://github.com/Masterminds/glide):star:0 - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.  *2017-12-14T01:27:03Z*
* [godep](https://github.com/tools/godep):star:0 - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.  *2018-01-26T22:05:26Z*
* [gom](https://github.com/mattn/gom):star:0 - Go Manager - bundle for go.  *2016-12-10T12:33:26Z*
* [goop](https://github.com/nitrous-io/goop):star:774 - Simple dependency manager for Go (golang), inspired by Bundler.  *2014-10-02T03:34:31Z*
* [gop](https://github.com/lunny/gop):star:41 - Build and manage your Go applications out of GOPATH  *2018-04-10T09:03:22Z*
* [gopm](https://github.com/gpmgo/gopm):star:0 - Go Package Manager.  *2017-07-27T21:13:37Z*
* [govendor](https://github.com/kardianos/govendor):star:0 - Go Package Manager. Go vendor tool that works with the standard vendor file.  *2018-02-09T21:39:04Z*
* [gpm](https://github.com/pote/gpm):star:0 - Barebones dependency manager for Go.  *2017-09-07T17:13:04Z*
* [gvt](https://github.com/FiloSottile/gvt):star:733 - `gvt` is a simple vendoring tool made for Go native vendoring (aka GO15VENDOREXPERIMENT), based on gb-vendor.  *2016-11-18T22:29:26Z*
* [johnny-deps](https://github.com/VividCortex/johnny-deps):star:214 - Minimal dependency version using Git.  *2014-09-16T21:26:09Z*
* [nut](https://github.com/jingweno/nut):star:246 - Vendor Go dependencies.  *2015-06-25T18:20:16Z*
* [VenGO](https://github.com/DamnWidget/VenGO):star:109 - create and manage exportable isolated go virtual environments.  *2016-04-26T18:57:59Z*

## Query Language

* [graphql](https://github.com/tmc/graphql):star:45 - graphql parser + utilities.  *2017-06-02T05:21:03Z*
* [graphql](https://github.com/sevki/graphql):star:34 - GraphQL implementation in go.  *2016-01-05T14:03:01Z*
* [graphql](https://github.com/neelance/graphql-go):star:0 - GraphQL server with a focus on ease of use.  *2018-04-11T22:19:16Z*
* [graphql-go](https://github.com/graphql-go/graphql):star:0 - Implementation of GraphQL for Go.  *2018-04-16T15:35:35Z*
* [jsonql](https://github.com/elgs/jsonql):star:145 - JSON query expression library in Golang.  *2017-12-29T17:22:55Z*

## Resource Embedding

* [esc](https://github.com/mjibson/esc):star:294 - Embeds files into Go programs and provides http.FileSystem interfaces to them.  *2017-09-07T07:13:03Z*
* [fileb0x](https://github.com/UnnoTed/fileb0x):star:227 - Simple tool to embed files in go with focus on "customization" and ease to use.  *2018-04-18T01:28:12Z*
* [go-embed](https://github.com/pyros2097/go-embed):star:11 - Generates go code to embed resource files into your library or executable.  *2016-04-12T06:18:40Z*
* [go-resources](https://github.com/omeid/go-resources):star:144 - Unfancy resources embedding with Go.  *2017-10-25T03:09:15Z*
* [go.rice](https://github.com/GeertJohan/go.rice):star:0 - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.  *2017-04-20T13:57:05Z*
* [packr](https://github.com/gobuffalo/packr):star:813 - The simple and easy way to embed static files into Go binaries.  *2018-04-13T17:35:21Z*
* [statics](https://github.com/go-playground/statics):star:47 - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.  *2016-10-05T01:25:49Z*
* [statik](https://github.com/rakyll/statik):star:0 - Embeds static files into a Go executable.  *2018-03-30T20:43:14Z*
* [templify](https://github.com/wlbr/templify):star:11 - Embed external template files into Go code to create single file binaries.  *2017-09-01T22:22:01Z*
* [vfsgen](https://github.com/shurcooL/vfsgen):star:254 - Generates a vfsdata.go file that statically implements the given virtual filesystem.  *2018-04-14T15:34:24Z*

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [blas](https://github.com/ziutek/blas):star:118 - Implementation of BLAS (Basic Linear Algebra Subprograms).  *2013-11-12T12:36:31Z*
* [chart](https://github.com/vdobler/chart):star:477 - Simple Chart Plotting library for Go. Supports many graphs types.  *2015-09-14T06:53:44Z*
* [evaler](https://github.com/soniah/evaler):star:30 - Simple floating point arithmetic expression evaluator.  *2015-08-09T23:40:13Z*
* [ewma](https://github.com/VividCortex/ewma):star:218 - Exponentially-weighted moving averages.  *2017-08-04T03:51:56Z*
* [geom](https://github.com/skelterjohn/geom):star:35 - 2D geometry for golang.  *2018-01-03T14:24:17Z*
* [go-dsp](https://github.com/mjibson/go-dsp):star:516 - Digital Signal Processing for Go.  *2017-01-04T18:39:34Z*
* [go-fn](https://github.com/ematvey/go-fn):star:7 - Mathematical functions written in Go language, that are not covered by math pkg.  *2013-04-03T06:55:44Z*
* [go-gt](https://github.com/ThePaw/go-gt):star:3 - Graph theory algorithms written in "Go" language.  *2013-04-03T07:06:47Z*
* [go.matrix](https://github.com/skelterjohn/go.matrix):star:301 - linear algebra for go (has been stalled).  *2013-05-17T14:41:13Z*
* [gocomplex](https://github.com/varver/gocomplex):star:3 - Complex number library for the Go programming language.  *2009-11-19T07:52:14Z*
* [goent](https://github.com/kzahedi/goent):star:4 - GO Implementation of Entropy Measures  *2018-01-31T22:18:37Z*
* [gofrac](https://github.com/anschelsc/gofrac):star:6 - (goinstallable) fractions library for go with support for basic arithmetic.  *2015-08-17T01:11:11Z*
* [gohistogram](https://github.com/VividCortex/gohistogram):star:106 - Approximate histograms for data streams.  *2016-08-22T18:46:21Z*
* [gonum/mat64](https://github.com/gonum/matrix):star:463 - The general purpose package for matrix computation. Package mat64 provides basic linear algebra operations for float64 matrices.  *2018-01-24T22:46:53Z*
* [gonum/plot](https://github.com/gonum/plot):star:810 - gonum/plot provides an API for building and drawing plots in Go.  *2018-04-10T07:39:29Z*
* [goraph](https://github.com/gyuho/goraph):star:468 - Pure Go graph theory library(data structure, algorith visualization).  *2017-10-01T06:05:14Z*
* [gosl](https://github.com/cpmech/gosl):star:0 - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.  *2018-04-22T00:55:13Z*
* [gostat](https://github.com/ematvey/gostat):star:25 - Statistics library for the go language.  *2018-01-19T12:39:06Z*
* [graph](https://github.com/yourbasic/graph):star:134 - Library of basic graph algorithms.  *2017-09-21T19:29:28Z*
* [ode](https://github.com/ChristopherRabotin/ode):star:5 - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.  *2017-01-18T01:29:17Z*
* [orb](https://github.com/paulmach/orb):star:16 - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.  *2018-03-27T04:58:56Z*
* [pagerank](https://github.com/alixaxel/pagerank):star:34 - Weighted PageRank algorithm implemented in Go.  *2016-03-06T11:07:29Z*
* [PiHex](https://github.com/claygod/PiHex):star:3 - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.  *2017-10-02T05:18:46Z*
* [sparse](https://github.com/james-bowman/sparse):star:29 - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.  *2018-04-21T18:47:36Z*
* [stats](https://github.com/montanaflynn/stats):star:877 - Statistics package with common functions missing from the Golang standard library.  *2017-12-01T20:20:39Z*
* [streamtools](https://github.com/nytlabs/streamtools):star:0 - general purpose, graphical tool for dealing with streams of data.  *2015-04-02T01:29:13Z*
* [TextRank](https://github.com/DavidBelicza/TextRank):star:7 - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.  *2018-03-11T14:51:34Z*
* [vectormath](https://github.com/spate/vectormath):star:58 - Vectormath for Go, an adaptation of the scalar C functions from Sony's Vector Math library, as found in the Bullet-2.79 source code (currently inactive).  *2013-03-23T13:35:40Z*

## Security

*Libraries that are used to help make your application more secure.*

* [acmetool](https://github.com/hlandau/acme):star:0 - ACME (Let's Encrypt) client tool with automatic renewal.  *2018-02-12T23:18:36Z*
* [argon2pw](https://github.com/raja/argon2pw):star:58 - Argon2 password hash generation with constant-time password comparison.  *2018-03-15T21:42:34Z*
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.  
* [BadActor](https://github.com/jaredfolkins/badactor):star:219 - In-memory, application-driven jailer built in the spirit of fail2ban.  *2017-06-05T17:14:07Z*
* [Cameradar](https://github.com/Ullaakut/cameradar):star:0 - Tool and library to remotely hack RTSP streams from surveillance cameras.  *2018-03-13T09:24:25Z*
* [go-yara](https://github.com/hillu/go-yara):star:66 - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".  *2016-06-08T01:42:37Z*
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword):star:15 - A probably paranoid package for securely hashing and encrypting passwords.  *2017-10-20T17:38:03Z*
* [lego](https://github.com/xenolf/lego):star:0 - Pure Go ACME client library and CLI tool (for use with Let's Encrypt).  *2018-04-15T13:49:13Z*
* [memguard](https://github.com/awnumar/memguard):star:755 - A pure Go library for handling sensitive values in memory.  *2018-04-02T22:05:12Z*
* [nacl](https://github.com/kevinburke/nacl):star:349 - Go implementation of the NaCL set of API's.  *2017-11-04T02:06:57Z*
* [passlib](https://github.com/hlandau/passlib):star:180 - Futureproof password hashing library.  *2016-09-22T11:41:03Z*
* [secure](https://github.com/unrolled/secure):star:941 - HTTP middleware for Go that facilitates some quick security wins.  *2018-04-16T20:52:22Z*
* [simple-scrypt](https://github.com/elithrar/simple-scrypt):star:128 - Scrypt package with a simple, obvious API and automatic cost calibration built-in.  *2018-04-02T16:31:15Z*
* [ssh-vault](https://github.com/ssh-vault/ssh-vault):star:142 - encrypt/decrypt using ssh keys.  *2017-11-06T17:48:00Z*

## Serialization

*Libraries and tools for binary serialization.*

* [asn1](https://github.com/PromonLogicalis/asn1):star:27 - Asn.1 BER and DER encoding library for golang.  *2016-03-07T19:22:09Z*
* [bambam](https://github.com/glycerine/bambam):star:58 - generator for Cap'n Proto schemas from go.  *2016-10-07T18:28:00Z*
* [colfer](https://github.com/pascaldekloe/colfer):star:359 - Code generation for the Colfer binary format.  *2018-03-08T13:30:16Z*
* [csvutil](https://github.com/jszwec/csvutil):star:201 - High Performance, idiomatic CSV record encoding and decoding to native Go structures.  *2018-04-11T14:02:52Z*
* [fwencoder](https://github.com/o1egl/fwencoder):star:3 - Fixed width file parser (encoding and decoding library) for Go.  *2018-01-14T18:23:43Z*
* [go-capnproto](https://github.com/glycerine/go-capnproto):star:258 - Cap'n Proto library and parser for go.  *2017-01-31T22:54:36Z*
* [go-codec](https://github.com/ugorji/go):star:925 - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.  *2018-04-07T10:30:00Z*
* [gogoprotobuf](https://github.com/gogo/protobuf):star:0 - Protocol Buffers for Go with Gadgets.  *2018-03-30T17:46:43Z*
* [goprotobuf](https://github.com/golang/protobuf):star:0 - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.  *2018-03-28T16:31:53Z*
* [jsoniter](https://github.com/json-iterator/go):star:0 - High-performance 100% compatible drop-in replacement of "encoding/json".  *2018-04-20T08:10:56Z*
* [mapstructure](https://github.com/mitchellh/mapstructure):star:0 - Go library for decoding generic map values into native Go structures.  *2018-02-20T23:01:11Z*
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder):star:84 - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.  *2017-01-27T12:03:00Z*
* [structomap](https://github.com/tuvistavie/structomap):star:69 - Library to easily and dynamically generate maps from static structures.  *2015-06-04T13:46:32Z*
* [structs](https://github.com/fatih/structs):star:0 - Library with support for converting structs to maps, struct keys/values to slices, and more.  *2018-01-23T06:50:59Z*

## Server Applications

* [algernon](https://github.com/xyproto/algernon):star:426 - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.  *2018-04-19T22:23:09Z*
* [Caddy](https://github.com/mholt/caddy):star:0 - Caddy is an alternative, HTTP/2 web server that's easy to configure and use.  *2018-04-20T00:11:50Z*
* [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
* [devd](https://github.com/cortesi/devd):star:0 - Local webserver for developers.  *2018-04-12T21:18:44Z*
* [etcd](https://github.com/coreos/etcd):star:0 - Highly-available key value store for shared configuration and service discovery.  *2018-04-20T21:40:22Z*
* [Fider](https://github.com/getfider/fider):star:275 - Fider is an open platform to collect and organize customer feedback.  *2018-04-20T06:46:53Z*
* [Flagr](https://github.com/checkr/flagr):star:333 - Flagr is an open-source feature flagging and A/B testing service.  *2018-04-12T16:38:20Z*
* [jackal](https://github.com/ortuman/jackal):star:189 - An XMPP server written in Go.  *2018-04-21T10:14:11Z*
* [minio](https://github.com/minio/minio):star:0 - Minio is a distributed object storage server.  *2018-04-22T03:51:53Z*
* [nsq](http://nsq.io/) - A realtime distributed messaging platform.
* [yakvs](https://github.com/sci4me/yakvs):star:23 - Small, networked, in-memory key-value store.  *2018-02-12T23:04:46Z*

## Template Engines

*Libraries and tools for templating and lexing.*

* [ace](https://github.com/yosssi/ace):star:713 - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.  *2016-07-28T07:45:28Z*
* [amber](https://github.com/eknkc/amber):star:785 - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.  *2017-10-10T12:03:22Z*
* [damsel](https://github.com/dskinner/damsel):star:19 - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.  *2016-04-07T02:57:10Z*
* [ego](https://github.com/benbjohnson/ego):star:370 - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.  *2018-02-27T15:59:08Z*
* [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).
* [gofpdf](https://github.com/jung-kurt/gofpdf):star:0 - PDF document generator with high level support for text, drawing and images.  *2018-04-20T19:28:40Z*
* [grender](https://github.com/dannyvankooten/grender):star:73 - small wrapper around html/template for file-based templates that support extending other template files.  *2017-10-24T07:41:32Z*
* [hero](https://github.com/shiyanhui/hero):star:953 - Hero is a handy, fast and powerful go template engine.  *2018-04-05T01:48:29Z*
* [jet](https://github.com/CloudyKit/jet):star:455 - Jet template engine.  *2017-10-30T20:14:56Z*
* [kasia.go](https://github.com/ziutek/kasia.go):star:71 - Templating system for HTML and other text documents - go implementation.  *2015-07-22T13:57:53Z*
* [liquid](https://github.com/osteele/liquid):star:48 - Go implementation of Shopify Liquid templates.  *2018-03-02T16:14:42Z*
* [mustache](https://github.com/hoisie/mustache):star:922 - Go implementation of the Mustache template language.  *2016-08-04T23:50:33Z*
* [pongo2](https://github.com/flosch/pongo2):star:0 - Django-like template-engine for Go.  *2018-04-12T21:51:50Z*
* [quicktemplate](https://github.com/valyala/quicktemplate):star:952 - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.  *2017-06-23T15:57:32Z*
* [raymond](https://github.com/aymerick/raymond):star:252 - Complete handlebars implementation in Go.  *2018-03-22T19:33:09Z*
* [Razor](https://github.com/sipin/gorazor):star:624 - Razor view engine for Golang.  *2018-04-17T08:37:44Z*
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).
* [velvet](https://github.com/gobuffalo/velvet):star:58 - Complete handlebars implementation in Go.  *2017-03-20T14:41:06Z*

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [assert](https://github.com/go-playground/assert):star:9 - Basic Assertion Library used along side native go testing, with building blocks for custom assertions.  *2016-02-05T15:04:13Z*
    * [badio](https://github.com/cavaliercoder/badio):star:6 - Extensions to Go's `testing/iotest` package.  *2016-02-13T15:00:51Z*
    * [baloo](https://github.com/h2non/baloo):star:487 - Expressive and versatile end-to-end HTTP API testing made easy.  *2017-11-29T15:19:36Z*
    * [bro](https://github.com/marioidival/bro):star:20 - Watch files in directory and run tests for them.  *2016-10-21T01:45:24Z*
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy):star:38 - Simple snapshot testing addon for your test framework.  *2018-01-25T14:05:27Z*
    * [dbcleaner](https://github.com/khaiql/dbcleaner):star:29 - Clean database for testing purpose, inspired by `database_cleaner` in Ruby.  *2018-01-31T17:22:32Z*
    * [dsunit](https://github.com/viant/dsunit):star:14 - Datastore testing for SQL, NoSQL, structured files.  *2018-04-11T14:26:10Z*
    * [endly](https://github.com/viant/endly):star:22 - Declarative end to end functional testing.  *2018-04-20T22:09:23Z*
    * [frisby](https://github.com/verdverm/frisby):star:225 - REST API testing framework.  *2017-06-04T21:13:11Z*
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet):star:177 - Tool for viewing test coverage in terminal.  *2018-03-17T11:06:43Z*
    * [go-mutesting](https://github.com/zimmski/go-mutesting):star:132 - Mutation testing for Go source code.  *2017-10-31T21:02:12Z*
    * [go-vcr](https://github.com/dnaeon/go-vcr):star:189 - Record and replay your HTTP interactions for fast, deterministic and accurate tests.  *2018-03-23T18:07:43Z*
    * [goblin](https://github.com/franela/goblin):star:482 - Mocha like testing framework fo Go.  *2018-04-07T13:27:55Z*
    * [gocheck](http://labix.org/gocheck) - More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/):star:0 - BDD-style framework with web UI and live reload.  *2018-02-22T19:45:00Z*
    * [gocrest](https://github.com/corbym/gocrest):star:6 - Composable hamcrest-like matchers for Go assertions.  *2018-01-31T21:50:58Z*
    * [godog](https://github.com/DATA-DOG/godog):star:365 - Cucumber or Behat like BDD framework for Go.  *2018-03-30T09:55:25Z*
    * [gofight](https://github.com/appleboy/gofight):star:157 - API Handler Testing for Golang Router framework.  *2018-02-18T14:34:45Z*
    * [gogiven](https://github.com/corbym/gogiven):star:6 - YATSPEC-like BDD testing framework for Go.  *2018-03-01T09:29:52Z*
    * [gomega](http://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
    * [GoSpec](https://github.com/orfjackal/gospec):star:109 - BDD-style testing framework for the Go programming language.  *2014-07-31T18:58:59Z*
    * [gospecify](https://github.com/stesla/gospecify):star:49 - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.  *2011-06-18T18:45:37Z*
    * [gosuite](https://github.com/pavlo/gosuite):star:5 - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.  *2016-10-18T16:53:15Z*
    * [Hamcrest](https://github.com/rdrdr/hamcrest):star:24 - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.  *2011-02-14T22:42:06Z*
    * [httpexpect](https://github.com/gavv/httpexpect):star:867 - Concise, declarative, and easy to use end-to-end HTTP and REST API testing.  *2017-08-20T08:05:27Z*
    * [restit](https://github.com/yookoala/restit):star:45 - Go micro framework to help writing RESTful API integration test.  *2017-01-17T14:46:08Z*
    * [testfixtures](https://github.com/go-testfixtures/testfixtures):star:164 - A helper for Rails' like test fixtures to test database applications.  *2017-12-21T17:23:04Z*
    * [Testify](https://github.com/stretchr/testify):star:0 - Sacred extension to the standard go testing package.  *2018-03-14T08:05:35Z*
    * [wstest](https://github.com/posener/wstest):star:37 - Websocket client for unit-testing a websocket http.Handler.  *2018-02-17T13:36:18Z*

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter):star:225 - Tool for generating self-contained mock objects.  *2018-03-06T14:55:59Z*
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock):star:799 - Mock SQL driver for testing database interactions.  *2018-03-04T15:30:57Z*
    * [go-txdb](https://github.com/DATA-DOG/go-txdb):star:48 - Single transaction based database driver mainly for testing purposes.  *2017-09-06T06:36:44Z*
    * [gock](https://github.com/h2non/gock):star:468 - Versatile HTTP mocking made easy.  *2018-02-28T22:52:24Z*
    * [gomock](https://github.com/golang/mock):star:0 - Mocking framework for the Go programming language.  *2018-04-06T15:19:24Z*
    * [govcr](https://github.com/seborama/govcr):star:30 - HTTP mock for Golang: record and replay HTTP interactions for offline testing.  *2018-03-23T13:54:23Z*
    * [minimock](https://github.com/gojuno/minimock):star:124 - Mock generator for Go interfaces.  *2018-02-20T15:50:40Z*
    * [mockhttp](https://github.com/tv42/mockhttp):star:20 - Mock object for Go http.ResponseWriter.  *2014-10-29T22:12:23Z*

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz):star:0 - Randomized testing system.  *2018-04-21T06:56:10Z*
    * [gofuzz](https://github.com/google/gofuzz):star:415 - Library for populating go objects with random values.  *2017-06-12T17:47:53Z*
    * [Tavor](https://github.com/zimmski/tavor):star:185 - Generic fuzzing and delta-debugging framework.  *2017-11-30T22:30:20Z*

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp):star:115 - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.  *2018-04-08T19:29:04Z*
    * [chromedp](https://github.com/knq/chromedp):star:0 - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.  *2018-03-26T21:30:59Z*
    * [ggr](https://github.com/aerokube/ggr):star:122 - a lightweight server that routes and proxies Selenium Wedriver requests to multiple Selenium hubs.  *2018-04-12T16:58:19Z*
    * [selenoid](https://github.com/aandryashin/selenoid):star:5 - alternative Selenium hub server that launches browsers within containers.  *2018-03-27T14:51:19Z*

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align):star:37 - A general purpose application that aligns text.  *2017-09-10T03:36:06Z*
    * [allot](https://github.com/sbstjn/allot):star:29 - Placeholder and wildcard text parsing for CLI tools and bots.  *2016-10-24T16:41:07Z*
    * [bbConvert](https://github.com/CalebQ42/bbConvert):star:2 - Converts bbCode to HTML that allows you to add support for custom bbCode tags.  *2016-09-14T13:04:27Z*
    * [blackfriday](https://github.com/russross/blackfriday):star:0 - Markdown processor in Go.  *2018-02-18T09:36:01Z*
    * [bluemonday](https://github.com/microcosm-cc/bluemonday):star:882 - HTML Sanitizer.  *2018-03-27T21:19:28Z*
    * [colly](https://github.com/asciimoo/colly):star:0 - Fast and Elegant Scraping Framework for Gophers  *2018-04-17T23:41:46Z*
    * [commonregex](https://github.com/mingrammer/commonregex):star:482 - A collection of common regular expressions for Go  *2018-03-20T05:24:19Z*
    * [doi](https://github.com/hscells/doi):star:0 - Document object identifier (doi) parser in Go.  *2017-08-21T05:50:49Z*
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go):star:19 - Editorconfig file parser and manipulator for Go.  *2016-08-25T01:23:04Z*
    * [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).
    * [encdec](https://github.com/mickep76/encdec):star:2 - Package provides a generic interface to encoders and decodersa.  *2018-04-10T19:45:42Z*
    * [genex](https://github.com/alixaxel/genex):star:42 - Count and expand Regular Expressions into all matching Strings.  *2016-03-13T15:07:42Z*
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth):star:8 - Fixed-width text formatting (encoder/decoder with reflection).  *2018-04-16T19:55:36Z*
    * [go-humanize](https://github.com/dustin/go-humanize):star:0 - Formatters for time, numbers, and memory size to human readable format.  *2018-04-21T18:26:05Z*
    * [go-nmea](https://github.com/adrianmo/go-nmea):star:43 - NMEA parser library for the Go language.  *2018-04-16T15:21:04Z*
    * [go-runewidth](https://github.com/mattn/go-runewidth):star:125 - Functions to get fixed width of the character or string.  *2018-04-08T05:53:51Z*
    * [go-slugify](https://github.com/mozillazg/go-slugify):star:16 - Make pretty slug with multiple languages support.  *2016-08-13T15:06:51Z*
    * [go-vcard](https://github.com/emersion/go-vcard):star:7 - Parse and format vCard.  *2017-10-02T08:59:19Z*
    * [gofeed](https://github.com/mmcdole/gofeed):star:813 - Parse RSS and Atom feeds in Go.  *2018-02-20T20:01:45Z*
    * [gographviz](https://github.com/awalterschulze/gographviz):star:172 - Parses the Graphviz DOT language.  *2018-01-31T14:37:11Z*
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes):star:221 - Format bytes to string.  *2018-02-23T23:01:46Z*
    * [gonameparts](https://github.com/polera/gonameparts):star:25 - Parses human names into individual name parts.  *2017-05-26T02:47:05Z*
    * [goq](https://github.com/andrewstuart/goq):star:83 - Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).  *2017-11-29T23:50:11Z*
    * [GoQuery](https://github.com/PuerkitoBio/goquery):star:0 - GoQuery brings a syntax and a set of features similar to jQuery to the Go language.  *2018-03-24T16:22:12Z*
    * [goregen](https://github.com/zach-klippenstein/goregen):star:28 - Library for generating random strings from regular expressions.  *2016-03-03T16:20:51Z*
    * [gotext](https://github.com/leonelquinteros/gotext):star:154 - GNU gettext utilities for Go.  *2018-04-08T21:25:47Z*
    * [guesslanguage](https://github.com/endeveit/guesslanguage):star:38 - Functions to determine the natural language of a unicode text.  *2014-12-16T12:10:03Z*
    * [inject](https://github.com/facebookgo/inject):star:884 - Package inject provides a reflect based injector.  *2016-10-06T17:47:21Z*
    * [mxj](https://github.com/clbanning/mxj):star:245 - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.  *2018-04-19T12:09:38Z*
    * [sh](https://github.com/mvdan/sh):star:0 - Shell parser and formatter.  *2018-04-17T15:44:51Z*
    * [slug](https://github.com/gosimple/slug):star:188 - URL-friendly slugify with multiple languages support.  *2017-08-21T11:09:01Z*
    * [Slugify](https://github.com/avelino/slugify):star:17 - Go slugify application that handles string.  *2015-04-13T03:05:04Z*
    * [syndfeed](https://github.com/zhengchun/syndfeed):star:1 - A syndication feed for Atom 1.0 and RSS 2.0.  *2018-03-13T02:31:27Z*
    * [toml](https://github.com/BurntSushi/toml):star:0 - TOML configuration format (encoder/decoder with reflection).  *2017-06-26T11:06:00Z*
* Utility
    * [gotabulate](https://github.com/bndr/gotabulate):star:168 - Easily pretty-print your tabular data with Go.  *2017-03-15T14:24:10Z*
    * [kace](https://github.com/codemodus/kace):star:6 - Common case conversions covering common initialisms.  *2017-09-08T01:45:41Z*
    * [parseargs-go](https://github.com/nproc/parseargs-go):star:4 - string argument parser that understands quotes and backslashes.  *2017-01-24T21:53:56Z*
    * [parth](https://github.com/codemodus/parth):star:17 - URL path segmentation parsing.  *2017-02-08T00:17:52Z*
    * [radix](https://github.com/yourbasic/radix):star:49 - fast string sorting algorithm.  *2018-03-08T12:29:24Z*
    * [xj2go](https://github.com/stackerzzq/xj2go):star:7 - Convert xml or json to go struct.  *2017-11-15T18:13:27Z*
    * [xurls](https://github.com/mvdan/xurls):star:341 - Extract urls from text.  *2018-03-31T13:59:46Z*

## Third-party APIs

*Libraries for accessing third party APIs.*

* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).
* [anaconda](https://github.com/ChimeraCoder/anaconda):star:868 - Go client library for the Twitter 1.1 API.  *2018-03-05T15:13:39Z*
* [aws-sdk-go](https://github.com/aws/aws-sdk-go):star:0 - The official AWS SDK for the Go programming language.  *2018-04-20T18:22:12Z*
* [brewerydb](https://github.com/naegelejd/brewerydb):star:13 - Go library for accessing the BreweryDB API.  *2015-06-18T19:34:13Z*
* [cachet](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/).
* [circleci](https://github.com/jszwedko/go-circleci):star:22 - Go client library for interacting with CircleCI's API.  *2018-03-10T20:44:18Z*
* [clarifai](https://github.com/samuelcouch/clarifai):star:55 - Go client library for interfacing with the Clarifai API.  *2017-08-28T17:25:47Z*
* [codeship-go](https://github.com/codeship/codeship-go):star:12 - Go client library for interacting with Codeship's API v2.  *2018-04-19T13:05:55Z*
* [discordgo](https://github.com/bwmarrin/discordgo):star:553 - Go bindings for the Discord Chat API.  *2018-04-04T15:55:14Z*
* [ethrpc](https://github.com/onrik/ethrpc):star:63 - Go bindings for Ethereum JSON RPC API.  *2018-03-29T12:56:52Z*
* [facebook](https://github.com/huandu/facebook):star:620 - Go Library that supports the Facebook Graph API.  *2018-02-06T08:01:08Z*
* [fcm](https://github.com/maddevsio/fcm):star:25 - Go library for Firebase Cloud Messaging.  *2017-05-17T15:36:33Z*
* [gads](https://github.com/emiddleton/gads):star:32 - Google Adwords Unofficial API.  *2015-09-08T14:45:51Z*
* [gami](https://github.com/bit4bit/gami):star:22 - Go library for Asterisk Manager Interface.  *2017-06-10T06:09:03Z*
* [gcm](https://github.com/Aorioli/gcm):star:30 - Go library for Google Cloud Messaging.  *2015-12-04T14:37:05Z*
* [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](http://open.mapquestapi.com/nominatim/), [OpenCage](http://geocoder.opencagedata.com/api.html), [HERE](https://developer.here.com/rest-apis/documentation/geocoder), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.
* [github](https://github.com/google/go-github):star:0 - Go library for accessing the GitHub REST API v3.  *2018-04-20T21:30:56Z*
* [githubql](https://github.com/shurcooL/githubql):star:223 - Go library for accessing the GitHub GraphQL API v4.  *2018-03-22T22:43:15Z*
* [go-chronos](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler
* [go-hacknews](https://github.com/PaulRosset/go-hacknews):star:4 - Tiny Go client for HackerNews API.  *2017-08-15T07:51:27Z*
* [go-imgur](https://github.com/koffeinsource/go-imgur):star:6 - Go client library for [imgur](https://imgur.com)  *2016-04-30T09:37:53Z*
* [go-jira](https://github.com/andygrunwald/go-jira):star:243 - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)  *2018-04-03T13:29:23Z*
* [go-marathon](https://github.com/gambol99/go-marathon):star:174 - Go library for interacting with Mesosphere's Marathon PAAS.  *2018-03-03T22:16:09Z*
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans):star:4 - Go client library for the SPTrans Olho Vivo API.  *2017-09-26T02:37:06Z*
* [go-telegraph](https://github.com/toby3d/go-telegraph):star:45 - Telegraph publishing platform API client.  *2018-03-14T10:55:43Z*
* [go-tgbot](https://github.com/olebedev/go-tgbot):star:71 - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.  *2018-02-28T19:28:07Z*
* [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.
* [go-twitch](https://github.com/knspriggs/go-twitch):star:12 - Go client for interacting with the Twitch v3 API.  *2017-08-20T21:01:01Z*
* [go-twitter](https://github.com/dghubble/go-twitter):star:433 - Go client library for the Twitter v1.1 APIs.  *2017-09-10T03:52:29Z*
* [go-unsplash](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API.
* [go-xkcd](https://github.com/nishanths/go-xkcd):star:31 - Go client for the xkcd API.  *2016-05-31T09:29:01Z*
* [goamz](https://github.com/mitchellh/goamz) - Popular fork of [goamz](https://launchpad.net/goamz) which adds some missing API calls to certain packages.
* [golyrics](https://github.com/mamal72/golyrics):star:20 - Golyrics is a Go library to fetch music lyrics data from the Wikia website.  *2016-11-24T02:36:43Z*
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz):star:25 - Go MusicBrainz WS2 client library.  *2017-04-18T00:01:17Z*
* [google](https://github.com/google/google-api-go-client):star:0 - Auto-generated Google APIs for Go.  *2018-04-20T00:09:06Z*
* [google-analytics](https://github.com/chonthu/go-google-analytics):star:8 - Simple wrapper for easy google analytics reporting.  *2015-05-30T17:41:40Z*
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang):star:0 - Google Cloud APIs Go Client Library.  *2018-04-20T15:39:56Z*
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).
* [gostorm](https://github.com/jsgilmore/gostorm):star:106 - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.  *2017-10-09T12:00:26Z*
* [govkbot](https://github.com/nikepan/govkbot) - Simple Go [VK](https://vk.com) bot library.
* [hipchat](https://github.com/andybons/hipchat):star:107 - This project implements a golang client library for the Hipchat API.  *2016-03-24T19:12:10Z*
* [hipchat (xmpp):star:108](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP.  *2017-05-12T18:52:32Z*
* [igdb](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/).
* [Medium](https://github.com/Medium/medium-sdk-go):star:99 - Golang SDK for Medium's OAuth2 API.  *2017-12-30T20:12:02Z*
* [megos](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.
* [micha](https://github.com/onrik/micha):star:5 - Go Library for Telegram bot api.  *2018-02-15T11:45:13Z*
* [minio-go](https://github.com/minio/minio-go):star:457 - Minio Go Library for Amazon S3 compatible cloud storage.  *2018-04-13T18:58:16Z*
* [mixpanel](https://github.com/dukex/mixpanel):star:24 - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.  *2018-04-12T21:50:57Z*
* [patreon-go](https://github.com/mxpv/patreon-go):star:11 - Go library for Patreon API.  *2017-10-31T00:10:22Z*
* [paypal](https://github.com/logpacker/paypalsdk):star:215 - Wrapper for PayPal payment API.  *2018-03-22T16:02:01Z*
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk):star:0 - The Playlyfe Rest API Go SDK.  *2016-03-06T10:21:05Z*
* [pushover](https://github.com/gregdel/pushover):star:38 - Go wrapper for the Pushover API.  *2018-02-08T23:07:17Z*
* [rrdaclient](https://github.com/Omie/rrdaclient):star:5 - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.  *2014-09-19T16:35:52Z*
* [shopify](https://github.com/rapito/go-shopify):star:18 - Go Library to make CRUD request to the Shopify API.  *2015-01-20T22:06:42Z*
* [slack](https://github.com/nlopes/slack):star:0 - Slack API in Go.  *2018-04-20T12:55:33Z*
* [smite](https://github.com/sergiotapia/smitego):star:8 - Go package to wraps access to the Smite game API.  *2014-07-18T15:50:52Z*
* [spotify](https://github.com/rapito/go-spotify):star:13 - Go Library to access Spotify WEB API.  *2017-10-31T11:26:44Z*
* [steam](https://github.com/sostronk/go-steam):star:8 - Go Library to interact with Steam game servers.  *2018-03-13T06:54:38Z*
* [stripe](https://github.com/stripe/stripe-go):star:678 - Go client for the Stripe API.  *2018-04-17T18:46:17Z*
* [tbot](https://github.com/yanzay/tbot):star:145 - Telegram bot server with API similar to net/http.  *2018-03-01T06:43:24Z*
* [telebot](https://github.com/tucnak/telebot):star:618 - Telegram bot framework written in Go.  *2018-04-18T18:57:23Z*
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api):star:0 - Simple and clean Telegram bot client.  *2018-03-28T13:10:29Z*
* [textbelt](https://github.com/dietsche/textbelt):star:11 - Go client for the textbelt.com txt messaging API.  *2015-09-04T14:12:39Z*
* [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) - Simple golang package to communicate with [themoviedb.org](https://themoviedb.org).
* [translate](https://github.com/poorny/translate):star:21 - Go online translation package.  *2016-02-28T15:12:56Z*
* [Trello](https://github.com/adlio/trello):star:45 - Go wrapper for the Trello API.  *2018-04-13T21:22:48Z*
* [tumblr](https://github.com/mattcunningham/gumblr):star:5 - Go wrapper for the Tumblr v2 API.  *2016-10-30T23:45:19Z*
* [webhooks](https://github.com/go-playground/webhooks):star:165 - Webhook receiver for GitHub and Bitbucket.  *2018-04-08T16:15:57Z*
* [zooz](https://github.com/gojuno/go-zooz):star:2 - Go client for the Zooz API.  *2017-08-28T09:17:06Z*

## Utilities

*General utilities and tools to make your life easier.*

* [abutil](https://github.com/bahlo/abutil):star:45 - Collection of often-used Golang helpers.  *2015-09-02T12:39:24Z*
* [apm](https://github.com/topfreegames/apm):star:106 - Process manager for Golang applications with an HTTP API.  *2016-11-24T20:58:45Z*
* [backscanner](https://github.com/icza/backscanner):star:3 - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.  *2018-02-26T08:25:41Z*
* [boilr](https://github.com/tmrts/boilr):star:681 - Blazingly fast CLI tool for creating projects from boilerplate templates.  *2017-07-19T13:20:21Z*
* [chyle](https://github.com/antham/chyle):star:83 - Changelog generator using a git repository with multiple configuration possibilities.  *2018-02-26T08:21:44Z*
* [circuit](https://github.com/cep21/circuit):star:32 - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.  *2018-04-19T17:44:17Z*
* [circuitbreaker](https://github.com/rubyist/circuitbreaker):star:611 - Circuit Breakers in Go.  *2017-03-30T19:36:47Z*
* [clockwerk](http://github.com/onatm/clockwerk):star:42 - Go package to schedule periodic jobs using a simple, fluent syntax.  *2017-04-15T18:08:32Z*
* [command](https://github.com/txgruppi/command):star:6 - Command pattern for Go with thread safe serial and parallel dispatcher.  *2016-04-20T17:06:52Z*
* [coop](https://github.com/rakyll/coop):star:0 - Cheat sheet for some of the common concurrent flows in Go.  *2015-10-06T15:47:02Z*
* [copy-pasta](https://github.com/jutkko/copy-pasta):star:22 - Universal multi-workstation clipboard that uses S3 like backend for the storage.  *2017-08-06T21:30:51Z*
* [ctop](https://github.com/bcicen/ctop) - [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.
* [Death](https://github.com/vrecan/death):star:92 - Managing go application shutdown with signals.  *2018-03-27T23:21:19Z*
* [Deepcopier](https://github.com/ulule/deepcopier):star:145 - Simple struct copying for Go.  *2017-11-07T15:55:58Z*
* [delve](https://github.com/derekparker/delve):star:0 - Go debugger.  *2018-04-19T05:27:34Z*
* [dlog](https://github.com/kirillDanshin/dlog):star:13 - Compile-time controlled logger to make your release smaller without removing debug calls.  *2017-07-28T00:08:07Z*
* [ergo](https://github.com/cristianoliveira/ergo):star:174 - The management of multiple local services running over different ports made easy.  *2018-04-21T18:47:23Z*
* [evaluator](https://github.com/nullne/evaluator):star:7 - Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.  *2018-01-29T03:16:49Z*
* [excelize](https://github.com/360EntSecGroup-Skylar/excelize):star:0 - Golang library for reading and writing Microsoft Excel™ (XLSX) files.  *2018-04-09T11:44:08Z*
* [fastlz](https://github.com/digitalcrab/fastlz) - Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang.
* [filetype](https://github.com/h2non/filetype):star:277 - Small package to infer the file type checking the magic numbers signature.  *2018-01-11T11:44:05Z*
* [filler](https://github.com/yaronsumel/filler):star:10 - small utility to fill structs using "fill" tag.  *2017-04-10T08:03:29Z*
* [fzf](https://github.com/junegunn/fzf):star:0 - Command-line fuzzy finder written in Go.  *2018-04-12T08:49:52Z*
* [generate](https://github.com/go-playground/generate):star:8 - runs go generate recursively on a specified path or environment variable and can filter by regex.  *2017-01-10T00:20:54Z*
* [gentleman](https://github.com/h2non/gentleman):star:521 - Full-featured plugin-driven HTTP client library.  *2017-11-26T11:36:53Z*
* [git-time-metric](https://github.com/git-time-metric/gtm):star:521 - Simple, seamless, lightweight time tracking for Git.  *2018-04-21T16:22:18Z*
* [GJSON](https://github.com/tidwall/gjson):star:0 - Get a JSON value with one line of code.  *2018-03-01T18:15:57Z*
* [go-astitodo](https://github.com/asticode/go-astitodo):star:33 - Parse TODOs in your GO code.  *2018-04-14T14:54:03Z*
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin):star:154 - go:generate tool for wrapping symbols exported by golang plugins (1.8 only).  *2017-08-27T10:34:11Z*
* [go-cron](https://github.com/rk/go-cron):star:158 - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.  *2013-04-19T21:34:54Z*
* [go-debug](https://github.com/tj/go-debug) - Conditional debug logging for Golang libraries & applications.
* [go-dry](https://github.com/ungerik/go-dry):star:389 - DRY (don't repeat yourself) package for Go.  *2018-04-11T13:39:23Z*
* [go-funk](https://github.com/thoas/go-funk):star:371 - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).  *2018-01-10T16:49:51Z*
* [go-httpheader](https://github.com/mozillazg/go-httpheader):star:11 - Go library for encoding structs into Header fields.  *2017-06-27T23:22:55Z*
* [go-rate](https://github.com/beefsack/go-rate):star:253 - Timed rate limiter for Go.  *2018-04-08T01:11:53Z*
* [go-respond](https://github.com/nicklaw5/go-respond):star:8 - Go package for handling common HTTP JSON responses.  *2018-04-07T06:12:51Z*
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator):star:67 - XML Sitemap generator written in Go.  *2017-06-08T02:26:36Z*
* [go-torch](https://github.com/uber/go-torch):star:0 - Stochastic flame graph profiler for Go programs.  *2017-07-18T14:45:31Z*
* [go-trigger](https://github.com/sadlil/go-trigger):star:133 - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.  *2017-03-28T16:18:25Z*
* [go-underscore](https://github.com/tobyhede/go-underscore):star:950 - Useful collection of helpfully functional Go collection utilities.  *2015-06-29T22:40:57Z*
* [goback](https://github.com/carlescere/goback):star:36 - Go simple exponential backoff package.  *2015-03-14T18:09:17Z*
* [godaemon](https://github.com/VividCortex/godaemon):star:353 - Utility to write daemons.  *2015-09-10T21:22:27Z*
* [godropbox](https://github.com/dropbox/godropbox):star:0 - Common libraries for writing Go services/applications from Dropbox.  *2018-02-27T22:17:28Z*
* [gohper](https://github.com/cosiner/gohper):star:234 - Various tools/modules help for development.  *2017-08-12T06:52:59Z*
* [gojq](https://github.com/elgs/gojq):star:109 - JSON query in Golang.  *2016-04-21T19:40:50Z*
* [gojson](https://github.com/ChimeraCoder/gojson):star:0 - Automatically generate Go (golang) struct definitions from example JSON.  *2018-01-30T05:56:29Z*
* [golarm](https://github.com/msempere/golarm):star:29 - Fire alarms with system events.  *2015-08-24T13:33:34Z*
* [golog](https://github.com/mlimaloureiro/golog):star:37 - Easy and lightweight CLI tool to time track your tasks.  *2016-07-03T12:24:49Z*
* [gopencils](https://github.com/bndr/gopencils):star:410 - Small and simple package to easily consume REST APIs.  *2016-11-13T11:41:52Z*
* [goplaceholder](https://github.com/michiwend/goplaceholder):star:17 - a small golang lib to generate placeholder images.  *2016-01-17T18:22:32Z*
* [goreleaser](https://github.com/goreleaser/goreleaser):star:0 - Deliver Go binaries as fast and easily as possible.  *2018-04-20T11:26:04Z*
* [goreporter](https://github.com/wgliang/goreporter):star:0 - Golang tool that does static analysis, unit testing, code review and generate code quality report.  *2018-04-16T02:47:05Z*
* [goreq](https://github.com/franela/goreq):star:665 - Minimal and simple request library for Go language.  *2017-12-04T16:33:38Z*
* [goreq](https://github.com/smallnest/goreq):star:63 - Enhanced simplified HTTP client based on gorequest.  *2018-02-24T03:12:47Z*
* [gorequest](https://github.com/parnurzeal/gorequest):star:0 - Simplified HTTP client with rich features for Go.  *2017-10-15T11:04:55Z*
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs):star:8 - SeaweedFS client library with almost full features.  *2017-08-04T04:30:48Z*
* [gotenv](https://github.com/subosito/gotenv):star:81 - Load environment variables from `.env` or any `io.Reader` in Go.  *2017-01-20T14:09:13Z*
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter):star:8 - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.  *2018-04-20T06:08:34Z*
* [gpath](https://github.com/tenntenn/gpath):star:21 - Library to simplify access struct fields with Go's expression in reflection.  *2017-06-04T08:31:36Z*
* [grequests](https://github.com/levigross/grequests):star:0 - Elegant and simple `net/http` wrapper that follows Python's requests library.  *2017-10-09T01:03:47Z*
* [gron](https://github.com/roylee0704/gron):star:505 - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.  *2016-06-21T04:24:32Z*
* [htcat](https://github.com/htcat/htcat):star:449 - Parallel and Pipelined HTTP GET Utility.  *2015-06-08T20:36:41Z*
* [httpcontrol](https://github.com/facebookgo/httpcontrol):star:474 - Package httpcontrol allows for HTTP transport level control around timeouts and retries.  *2015-07-08T23:40:01Z*
* [hub](https://github.com/github/hub):star:0 - wrap git commands with additional functionality to interact with github from the terminal.  *2018-04-06T09:36:01Z*
* [hystrix-go](https://github.com/afex/hystrix-go):star:0 - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.  *2018-04-06T01:24:32Z*
* [immortal](https://github.com/immortal/immortal):star:480 - *nix cross-platform (OS agnostic) supervisor.  *2018-01-10T09:43:58Z*
* [intrinsic](https://github.com/mengzhuo/intrinsic):star:20 - Use x86 SIMD without writing any assembly code.  *2017-06-22T15:45:04Z*
* [JobRunner](https://github.com/bamzi/jobrunner):star:444 - Smart and featureful cron job scheduler with job queuing and live monitoring built in.  *2016-10-19T14:30:21Z*
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors):star:3 - Go bindings based on the JSON API errors reference.  *2016-11-17T15:59:34Z*
* [jsonf](https://github.com/miolini/jsonf):star:47 - Console tool for highlighted formatting and struct query fetching JSON.  *2016-07-08T00:43:10Z*
* [jsongo](https://github.com/ricardolonga/jsongo):star:78 - Fluent API to make it easier to create Json objects.  *2016-12-15T11:09:33Z*
* [jsonhal](https://github.com/RichardKnop/jsonhal):star:6 - Simple Go package to make custom structs marshal into HAL compatible JSON responses.  *2017-12-19T15:03:13Z*
* [kazaam](https://github.com/Qntfy/kazaam):star:75 - API for arbitrary transformation of JSON documents.  *2018-02-12T21:54:51Z*
* [lrserver](https://github.com/jaschaephraim/lrserver):star:91 - LiveReload server for Go.  *2017-11-29T20:29:58Z*
* [mc](https://github.com/minio/mc):star:716 - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.  *2018-04-13T20:57:07Z*
* [mergo](https://github.com/imdario/mergo):star:466 - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.  *2018-04-04T21:29:33Z*
* [minify](https://github.com/tdewolff/minify):star:0 - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.  *2018-04-19T03:59:41Z*
* [minquery](https://github.com/icza/minquery):star:30 - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).  *2017-02-17T09:44:09Z*
* [mmake](https://github.com/tj/mmake):star:0 - Modern Make.  *2018-01-10T01:11:38Z*
* [moldova](https://github.com/StabbyCutyou/moldova):star:144 - Utility for generating random data based on an input template.  *2017-09-04T15:04:37Z*
* [mp](https://github.com/sanbornm/mp):star:22 - Simple cli email parser. It currently takes stdin and outputs JSON.  *2016-05-11T19:40:54Z*
* [mssqlx](https://github.com/linxGnu/mssqlx):star:42 - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.  *2018-03-21T15:23:43Z*
* [multitick](https://github.com/VividCortex/multitick):star:52 - Multiplexor for aligned tickers.  *2016-08-22T18:54:26Z*
* [myhttp](https://github.com/inancgumus/myhttp):star:25 - Simple API to make HTTP GET requests with timeout support.  *2017-10-12T12:28:27Z*
* [netbug](https://github.com/e-dard/netbug):star:61 - Easy remote profiling of your services.  *2015-10-29T17:28:37Z*
* [okrun](https://github.com/xta/okrun):star:11 - go run error steamroller.  *2014-10-06T01:15:26Z*
* [onecache](https://github.com/adelowo/onecache):star:72 - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).  *2018-03-13T03:53:21Z*
* [panicparse](https://github.com/maruel/panicparse):star:0 - Groups similar goroutines and colorizes stack dump.  *2018-03-18T23:01:39Z*
* [peco](https://github.com/peco/peco):star:0 - Simplistic interactive filtering tool.  *2018-03-30T20:27:26Z*
* [pester](https://github.com/sethgrid/pester):star:243 - Go HTTP client calls with retries, backoff, and concurrency.  *2018-02-27T22:34:04Z*
* [pm](https://github.com/VividCortex/pm):star:56 - Process (i.e. goroutine) manager with an HTTP API.  *2018-03-20T19:51:00Z*
* [profile](https://github.com/pkg/profile):star:739 - Simple profiling support package for Go.  *2017-05-09T09:25:25Z*
* [rclient](https://github.com/zpatrick/rclient):star:19 - Readable, flexible, simple-to-use client for REST APIs.  *2017-06-22T21:53:04Z*
* [realize](https://github.com/tockins/realize):star:0 - Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.  *2018-04-17T07:31:41Z*
* [repeat](https://github.com/ssgreg/repeat):star:40 - Go implementation of different backoff strategies useful for retrying operations and heartbeating.  *2017-12-02T06:51:29Z*
* [request](https://github.com/mozillazg/request):star:283 - Go HTTP Requests for Humans™.  *2017-10-23T23:58:38Z*
* [rerate](https://github.com/abo/rerate):star:9 - Redis-based rate counter and rate limiter for Go.  *2017-03-28T02:22:07Z*
* [rerun](https://github.com/ivpusic/rerun):star:140 - Recompiling and rerunning go apps when source changes.  *2017-03-31T08:08:01Z*
* [resty](https://github.com/go-resty/resty):star:833 - Simple HTTP and REST client for Go inspired by Ruby rest-client.  *2018-04-20T05:34:30Z*
* [retry](https://github.com/kamilsk/retry):star:36 - Functional mechanism based on context to perform actions repetitively until successful.  *2018-03-06T20:47:25Z*
* [retry](https://github.com/thedevsaddam/retry):star:18 - Simple and easy retry mechanism package for Go.  *2018-03-22T14:34:34Z*
* [retry-go](https://github.com/rafaeljesus/retry-go):star:18 - Retrying made simple and easy for golang.  *2017-12-14T20:46:23Z*
* [robustly](https://github.com/VividCortex/robustly):star:126 - Runs functions resiliently, catching and restarting panics.  *2018-03-23T18:27:11Z*
* [rq](https://github.com/ddo/rq):star:15 - A nicer interface for golang stdlib HTTP client.  *2018-03-04T18:52:34Z*
* [scheduler](https://github.com/carlescere/scheduler):star:206 - Cronjobs scheduling made easy.  *2017-01-09T14:14:37Z*
* [sling](https://github.com/dghubble/sling):star:676 - Go HTTP requests builder for API clients.  *2017-06-29T03:54:44Z*
* [spinner](https://github.com/briandowns/spinner):star:498 - Go package to easily provide a terminal spinner with options.  *2018-01-23T22:20:39Z*
* [sqlx](https://github.com/jmoiron/sqlx):star:0 - provides a set of extensions on top of the excellent built-in database/sql package.  *2018-04-06T16:44:12Z*
* [Storm](https://github.com/asdine/storm):star:925 - Simple and powerful toolkit for BoltDB.  *2018-03-11T15:17:39Z*
* [structs](https://github.com/PumpkinSeed/structs):star:5 - Implement simple functions to manipulate structs.  *2017-10-23T13:03:16Z*
* [Task](https://github.com/go-task/task):star:838 - simple "Make" alternative.  *2018-04-07T18:36:37Z*
* [toolbox](https://github.com/viant/toolbox):star:23 - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.  *2018-04-20T16:59:33Z*
* [ugo](https://github.com/alxrm/ugo):star:15 - ugo is slice toolbox with concise syntax for Go.  *2016-06-30T19:18:16Z*
* [UNIS](https://github.com/esemplastic/unis):star:65 - Common Architecture™ for String Utilities in Go.  *2017-05-09T16:17:24Z*
* [usql](https://github.com/knq/usql):star:0 - usql is a universal command-line interface for SQL databases.  *2018-04-09T17:18:07Z*
* [util](https://github.com/shomali11/util):star:58 - Collection of useful utility functions. (strings, concurrency, manipulations, ...).  *2017-11-19T23:27:47Z*
* [wuzz](https://github.com/asciimoo/wuzz):star:0 - Interactive cli tool for HTTP inspection.  *2017-10-25T01:18:05Z*
* [xferspdy](https://github.com/monmohan/xferspdy):star:53 - Xferspdy provides binary diff and patch library in golang.  *2016-08-20T16:49:52Z*
* [xlsx](https://github.com/tealeg/xlsx):star:0 - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.  *2018-04-19T19:51:53Z*

## Validation

*Libraries for validation.*

* [govalidator](https://github.com/asaskevich/govalidator):star:0 - Validators and sanitizers for strings, numerics, slices and structs.  *2018-03-19T08:16:51Z*
* [govalidator](https://github.com/thedevsaddam/govalidator):star:230 - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.  *2018-02-21T10:10:54Z*
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation):star:637 - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.  *2018-03-22T18:43:44Z*
* [validate](https://github.com/markbates/validate):star:37 - This package provides a framework for writing validations for Go applications.  *2017-11-03T20:50:26Z*
* [validator](https://github.com/go-playground/validator):star:0 - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.  *2018-04-10T05:32:07Z*

## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh):star:55 - Scriptable server and net/http middleware for GitHub Webhooks.  *2017-07-25T20:41:03Z*
* [git2go](https://github.com/libgit2/git2go):star:0 - Go bindings for libgit2.  *2018-03-26T10:58:53Z*
* [go-vcs](https://github.com/sourcegraph/go-vcs):star:60 - manipulate and inspect VCS repositories in Go.  *2018-04-16T07:04:11Z*
* [hgo](https://github.com/beyang/hgo):star:11 - Hgo is a collection of Go packages providing read-access to local Mercurial repositories.  *2015-08-25T03:56:31Z*

## Video

*Libraries for manipulating video.*

* [gmf](https://github.com/3d0c/gmf):star:353 - Go bindings for FFmpeg av\* libraries.  *2018-03-09T10:48:51Z*
* [go-astisub](https://github.com/asticode/go-astisub):star:100 - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).  *2018-03-11T08:22:42Z*
* [go-astits](https://github.com/asticode/go-astits):star:200 - Parse and demux MPEG Transport Streams (.ts) natively in GO.  *2018-02-25T13:42:27Z*
* [goav](https://github.com/giorgisio/goav):star:440 - Comphrensive Go bindings for FFmpeg.  *2015-12-01T22:15:57Z*
* [gst](https://github.com/ziutek/gst):star:141 - Go bindings for GStreamer.  *2017-11-24T16:14:05Z*
* [libgosubs](https://github.com/wargarblgarbl/libgosubs):star:6 - Subtitle format support for go. Supports .srt, .ttml, and .ass.  *2018-01-29T05:41:32Z*
* [v4l](https://github.com/korandiz/v4l):star:12 - Video capture library for Linux, written in Go.  *2017-10-28T19:04:15Z*

## Web Frameworks

*Full stack web frameworks.*

* [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.
* [Air](https://github.com/sheng/air):star:62 - Ideal RESTful web framework for Go.  *2018-02-22T06:09:02Z*
* [Banjo](https://github.com/nsheremet/banjo):star:1 - Very simple and fast web framework for Go.  *2018-01-31T16:42:14Z*
* [Beego](https://github.com/astaxie/beego):star:0 - beego is an open-source, high-performance web framework for the Go programming language.  *2017-12-18T11:18:59Z*
* [Buffalo](http://gobuffalo.io) - Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo):star:0 - High performance, minimalist Go web framework.  *2018-04-16T17:20:21Z*
* [Fireball](https://github.com/zpatrick/fireball):star:33 - More "natural" feeling web framework.  *2018-03-28T17:47:39Z*
* [Florest](https://github.com/jabong/florest-core):star:38 - High-performance workflow based REST API framework.  *2017-01-31T06:30:16Z*
* [Gem](https://github.com/go-gem/gem):star:153 - Simple and fast web framework, friendly to REST API.  *2017-03-19T15:25:20Z*
* [Gin](https://github.com/gin-gonic/gin):star:0 - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.  *2018-04-20T02:27:44Z*
* [Gizmo](https://github.com/NYTimes/gizmo):star:0 - Microservice toolkit used by the New York Times.  *2018-04-02T21:02:05Z*
* [go-json-rest](https://github.com/ant0ine/go-json-rest):star:0 - Quick and easy way to setup a RESTful JSON API.  *2017-09-13T04:12:08Z*
* [go-relax](https://github.com/codehack/go-relax):star:148 - Framework of pluggable components to build RESTful API's.  *2017-11-16T23:05:03Z*
* [go-rest](https://github.com/ungerik/go-rest):star:106 - Small and evil REST framework for Go.  *2017-01-20T13:25:26Z*
* [goa](https://github.com/raphael/goa):star:0 - Framework for developing microservices based on the design of Ruby's Praxis.  *2018-04-20T23:16:35Z*
* [Golf](https://github.com/dinever/golf):star:216 - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.  *2017-02-24T16:57:24Z*
* [Gondola](https://github.com/rainycape/gondola):star:311 - The web framework for writing faster sites, faster.  *2017-02-20T16:16:56Z*
* [gongular](https://github.com/mustafaakin/gongular):star:390 - Fast Go web framework with input mapping/validation and (DI) Dependency Injection.  *2017-11-16T07:46:31Z*
* [Macaron](https://github.com/go-macaron/macaron):star:0 - Macaron is a high productive and modular design web framework in Go.  *2018-03-06T06:20:08Z*
* [mango](https://github.com/paulbellamy/mango):star:319 - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.  *2017-10-17T08:18:43Z*
* [Microservice](https://github.com/claygod/microservice):star:41 - The framework for the creation of microservices, written in Golang.  *2017-12-27T13:38:53Z*
* [neo](https://github.com/ivpusic/neo):star:360 - Neo is minimal and fast Go Web Framework with extremely simple API.  *2017-08-14T23:54:31Z*
* [Resoursea](https://github.com/resoursea/api):star:29 - REST framework for quickly writing resource based services.  *2015-02-01T22:58:19Z*
* [REST Layer](http://rest-layer.io) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Revel](https://github.com/revel/revel):star:0 - High-productivity web framework for the Go language.  *2018-03-21T16:43:36Z*
* [rex](https://github.com/goanywhere/rex):star:17 - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.  *2017-12-22T03:25:40Z*
* [sawsij](https://github.com/jaybill/sawsij):star:1 - lightweight, open-source web framework for building high-performance, data-driven web applications.  *2014-04-28T20:21:46Z*
* [tango](https://github.com/lunny/tango):star:709 - Micro & pluggable web framework for Go.  *2018-04-12T14:57:37Z*
* [tigertonic](https://github.com/rcrowley/go-tigertonic):star:972 - Go framework for building JSON web services inspired by Dropwizard.  *2017-04-20T12:38:39Z*
* [traffic](https://github.com/pilu/traffic):star:515 - Sinatra inspired regexp/pattern mux and web framework for Go.  *2015-11-26T21:31:07Z*
* [utron](https://github.com/gernest/utron):star:0 - Lightweight MVC framework for Go(Golang).  *2017-10-30T21:48:45Z*
* [violetear](https://github.com/nbari/violetear):star:85 - Go HTTP router.  *2018-01-06T13:42:01Z*
* [WebGo](https://github.com/bnkamalesh/webgo):star:31 - A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).  *2018-04-15T11:00:10Z*
* [YARF](https://github.com/yarf-framework/yarf):star:42 - Fast micro-framework designed to build REST APIs and web services in a fast and simple way.  *2017-09-14T14:41:36Z*
* [Zerver](https://github.com/cosiner/zerver):star:140 - Zerver is an expressive, modular, feature completed RESTful framework.  *2016-05-08T12:20:55Z*

### Middlewares

#### Actual middlewares

* [client-timing](https://github.com/posener/client-timing):star:8 - An HTTP client for Server-Timing header.  *2018-02-26T19:20:44Z*
* [CORS](https://github.com/rs/cors):star:757 - Easily add CORS capabilities to your API.  *2018-04-17T22:52:04Z*
* [formjson](https://github.com/rs/formjson):star:28 - Transparently handle JSON input as a standard form POST.  *2015-12-17T09:34:14Z*
* [go-server-timing](https://github.com/mitchellh/go-server-timing):star:689 - Add/parse Server-Timing header.  *2018-02-26T01:59:00Z*
* [Limiter](https://github.com/ulule/limiter):star:320 - Dead simple rate limit middleware for Go.  *2018-04-03T13:11:38Z*
* [Tollbooth](https://github.com/didip/tollbooth):star:788 - Rate limit HTTP request handler.  *2018-04-15T19:51:42Z*
* [XFF](https://github.com/sebest/xff):star:66 - Handle `X-Forwarded-For` header and friends.  *2016-09-10T04:38:05Z*

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice):star:0 - Painless middleware chaining for Go.  *2017-10-23T06:44:55Z*
* [catena](https://github.com/codemodus/catena):star:6 - http.Handler wrapper catenation (same API as "chain").  *2017-02-17T22:07:31Z*
* [chain](https://github.com/codemodus/chain):star:62 - Handler wrapper chaining with scoped data (net/context-based "middleware").  *2017-02-17T18:32:57Z*
* [go-wrap](https://github.com/go-on/wrap):star:52 - Small middlewares package for net/http.  *2014-12-15T18:18:31Z*
* [gores](https://github.com/alioygur/gores):star:69 - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.  *2017-02-20T07:43:40Z*
* [interpose](https://github.com/carbocation/interpose):star:281 - Minimalist net/http middleware for golang.  *2016-12-06T21:52:53Z*
* [muxchain](https://github.com/stephens2424/muxchain):star:203 - Lightweight middleware for net/http.  *2014-10-26T18:04:30Z*
* [negroni](https://github.com/urfave/negroni):star:0 - Idiomatic HTTP middleware for Golang.  *2018-01-30T04:45:49Z*
* [render](https://github.com/unrolled/render):star:0 - Go package for easily rendering JSON, XML, and HTML template responses.  *2017-11-02T16:21:32Z*
* [renderer](https://github.com/thedevsaddam/renderer):star:83 - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.  *2017-12-28T17:50:33Z*
* [rye](https://github.com/InVisionApp/rye):star:72 - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.  *2017-12-07T23:57:57Z*
* [stats](https://github.com/thoas/stats):star:472 - Go middleware that stores various information about your web application.  *2017-09-26T10:15:42Z*
* [Volatile](https://github.com/volatile/core):star:81 - Minimalist middleware stack promoting flexibility, good practices and clean code.  *2016-04-10T22:33:57Z*

### Routers

* [alien](https://github.com/gernest/alien):star:88 - Lightweight and fast http router from outer space.  *2018-01-29T09:17:43Z*
* [Bone](https://github.com/go-zoo/bone):star:0 - Lightning Fast HTTP Multiplexer.  *2018-01-30T19:34:01Z*
* [Bxog](https://github.com/claygod/Bxog):star:85 - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.  *2018-03-01T08:16:16Z*
* [chi](https://github.com/go-chi/chi):star:0 - Small, fast and expressive HTTP router built on net/context.  *2018-02-02T19:41:35Z*
* [fasthttprouter](https://github.com/buaazp/fasthttprouter):star:467 - High performance router forked from `httprouter`. The first router fit for `fasthttp`.  *2018-02-11T07:41:05Z*
* [FastRouter](https://github.com/razonyang/fastrouter):star:14 - a fast, flexible HTTP router written in Go.  *2017-11-02T01:22:00Z*
* [gocraft/web](https://github.com/gocraft/web):star:0 - Mux and middleware package in Go.  *2017-09-25T13:59:45Z*
* [Goji](https://github.com/goji/goji):star:623 - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.  *2016-11-14T00:05:42Z*
* [GoRouter](https://github.com/vardius/gorouter):star:32 - GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.  *2018-03-13T08:29:47Z*
* [gowww/router](https://github.com/gowww/router):star:139 - Lightning fast HTTP router fully compatible with the net/http.Handler interface.  *2018-03-27T19:52:01Z*
* [httprouter](https://github.com/julienschmidt/httprouter):star:0 - High performance router. Use this and the standard http handlers to form a very high performance web framework.  *2018-04-11T15:45:01Z*
* [httptreemux](https://github.com/dimfeld/httptreemux):star:321 - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.  *2018-02-13T07:44:14Z*
* [lars](https://github.com/go-playground/lars):star:352 - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.  *2017-10-31T12:50:11Z*
* [medeina](https://github.com/imdario/medeina):star:17 - Medeina is a HTTP routing tree based on HttpRouter, inspired by Roda and Cuba.  *2015-01-17T21:50:32Z*
* [mux](https://github.com/gorilla/mux):star:0 - Powerful URL router and dispatcher for golang.  *2018-04-16T20:45:19Z*
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing):star:276 - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.  *2018-02-10T02:24:22Z*
* [pat](https://github.com/bmizerany/pat):star:0 - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra.  *2017-08-15T01:04:13Z*
* [pure](https://github.com/go-playground/pure):star:55 - Is a lightweight HTTP router that sticks to the std "net/http" implementation.  *2017-09-11T04:25:16Z*
* [Siesta](https://github.com/VividCortex/siesta):star:347 - Composable framework to write middleware and handlers.  *2018-01-25T21:09:12Z*
* [vestigo](https://github.com/husobee/vestigo):star:218 - Performant, stand-alone, HTTP compliant URL Router for go web applications.  *2017-08-28T23:43:07Z*
* [xmux](https://github.com/rs/xmux):star:78 - High performance muxer based on `httprouter` with `net/context` support.  *2017-06-09T18:54:01Z*
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter):star:106 - A simple and fast HTTP router for Go.  *2018-03-19T03:15:06Z*
* [zeus](https://github.com/daryl/zeus):star:109 - Very simple and fast HTTP router for Go.  *2016-09-24T02:23:05Z*

## Windows

* [d3d9](https://github.com/gonutz/d3d9):star:66 - Go bindings for Direct3D9.  *2018-04-11T21:48:00Z*
* [go-ole](https://github.com/go-ole/go-ole):star:377 - Win32 OLE implementation for golang.  *2018-02-13T00:28:36Z*

## XML

*Libraries and tools for manipulating XML.*

* [XML-Comp](https://github.com/xml-comp/xml-comp):star:13 - Simple command line XML comparer that generates diffs of folders, files and tags.  *2018-03-21T03:05:23Z*
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter):star:3 - Procedural XML generation API based on libxml2's xmlwriter module.  *2017-05-24T13:39:52Z*
* [xpath](https://github.com/antchfx/xpath):star:57 - XPath package for Go.  *2018-02-24T01:49:21Z*
* [xquery](https://github.com/antchfx/xquery):star:115 - XQuery lets you extract data from HTML/XML documents using XPath expression.  *2017-12-05T03:07:50Z*

# Tools

*Go software and plugins.*

## Code Analysis

* [apicompat](https://github.com/bradleyfalzon/apicompat):star:141 - Checks recent changes to a Go project for backwards incompatible changes.  *2017-02-05T09:56:49Z*
* [dupl](https://github.com/mibk/dupl):star:104 - Tool for code clone detection.  *2017-11-19T16:48:37Z*
* [errcheck](https://github.com/kisielk/errcheck):star:990 - Errcheck is a program for checking for unchecked errors in Go programs.  *2018-03-03T00:00:09Z*
* [gcvis](https://github.com/davecheney/gcvis):star:804 - Visualise Go program GC trace data in real time.  *2017-12-22T22:29:38Z*
* [Go Metalinter](https://github.com/alecthomas/gometalinter):star:0 - Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form.  *2018-04-09T06:16:34Z*
* [go-checkstyle](https://github.com/qiniu/checkstyle):star:74 - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style refered to some points in Go Code Review Comments.  *2018-02-28T07:13:17Z*
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch):star:199 - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.  *2018-02-27T19:12:45Z*
* [go-outdated](https://github.com/firstrow/go-outdated):star:38 - Console application that displays outdated packages.  *2016-03-25T11:45:39Z*
* [goast-viewer](https://github.com/yuroyoro/goast-viewer):star:258 - Web based Golang AST visualizer.  *2017-12-06T07:17:43Z*
* [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/) - GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [GoLint](https://github.com/golang/lint):star:0 - Golint is a linter for Go source code.  *2018-03-02T06:23:24Z*
* [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple):star:0 - gosimple is a linter for Go source code that specialises on simplifying code.  *2018-04-03T05:13:39Z*
* [gostatus](https://github.com/shurcooL/gostatus):star:217 - Command line tool, shows the status of repositories that contain Go packages.  *2017-09-19T05:04:18Z*
* [interfacer](https://github.com/mvdan/interfacer):star:696 - Linter that suggests interface types.  *2018-03-26T10:46:26Z*
* [lint](https://github.com/surullabs/lint):star:58 - Run linters as part of go test.  *2017-10-03T11:47:19Z*
* [php-parser](https://github.com/z7zmey/php-parser):star:328 - A Parser for PHP written in Go.  *2018-04-10T21:58:57Z*
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck):star:0 - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.  *2017-11-10T19:24:41Z*
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp):star:8 - tarp finds functions and methods without direct unit tests in Go source code.  *2018-03-10T06:53:52Z*
* [unconvert](https://github.com/mdempsky/unconvert):star:206 - Remove unnecessary type conversions from Go source.  *2016-08-03T23:01:54Z*
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused):star:0 - unused checks Go code for unused constants, variables, functions and types.  *2017-11-08T10:17:53Z*
* [validate](https://github.com/mccoyst/validate):star:62 - Automatically validates struct fields with tags.  *2016-03-28T22:03:14Z*

## Editor Plugins

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server):star:8 - A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.  *2018-02-28T15:17:15Z*
* [go-mode](https://github.com/dominikh/go-mode.el):star:757 - Go mode for GNU/Emacs.  *2018-03-27T15:56:03Z*
* [go-plus](https://github.com/joefitzgerald/go-plus):star:0 - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.  *2018-04-02T22:26:22Z*
* [Goclipse](https://github.com/GoClipse/goclipse):star:751 - Eclipse plugin for Go.  *2017-08-25T16:53:24Z*
* [gocode](https://github.com/nsf/gocode):star:0 - Autocompletion daemon for the Go programming language.  *2018-03-25T12:21:35Z*
* [GoSublime](https://github.com/DisposaBoy/GoSublime):star:0 - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.  *2018-04-19T09:51:16Z*
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension):star:5 - Go language support for the Theia IDE.  *2018-03-01T14:32:17Z*
* [velour](https://github.com/velour/velour):star:14 - IRC client for the acme editor.  *2016-03-03T15:58:39Z*
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go):star:77 - Vim plugin to highlight syntax errors on save.  *2016-06-28T22:00:10Z*
* [vim-go](https://github.com/fatih/vim-go):star:0 - Go development plugin for Vim.  *2018-04-11T00:29:40Z*
* [vscode-go](https://github.com/Microsoft/vscode-go):star:0 - Extension for Visual Studio Code (VS Code) which provides support for the Go language.  *2018-04-12T05:39:07Z*
* [Watch](https://github.com/eaburns/Watch):star:143 - Runs a command in an acme win on file changes.  *2018-03-25T14:15:48Z*

## Go Generate Tools

* [generic](https://github.com/usk81/generic):star:14 - flexible data type for Go.  *2018-01-11T02:17:42Z*
* [genny](https://github.com/cheekybits/genny):star:606 - Elegant generics for Go.  *2017-03-28T20:00:08Z*
* [gonerics](http://github.com/bouk/gonerics):star:96 - Idiomatic Generics in Go.  *2014-09-29T15:04:46Z*
* [gotests](https://github.com/cweill/gotests):star:0 - Generate Go tests from your source code.  *2018-03-24T02:14:56Z*
* [re2dfa](https://github.com/opennota/re2dfa):star:139 - Transform regular expressions into finite state machines and output Go source code.  *2018-03-24T02:25:03Z*

## Go Tools

* [colorgo](https://github.com/songgao/colorgo):star:89 - Wrapper around `go` command for colorized `go build` output.  *2016-10-28T04:37:18Z*
* [depth](https://github.com/KyleBanks/depth):star:220 - Visualize dependency trees of any package by analyzing imports.  *2018-02-14T14:08:59Z*
* [gb](https://getgb.io/) - An easy to use project based build tool for the Go programming language.
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) - A [Yeoman](https://yeoman.io) generator to get new Go projects started.
* [go-callvis](https://github.com/TrueFurby/go-callvis):star:0 - Visualize call graph of your Go program using dot format.  *2018-04-02T19:55:18Z*
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete):star:34 - Bash completion for go and wgo.  *2017-11-17T14:00:34Z*
* [go-swagger](https://github.com/go-swagger/go-swagger):star:0 - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.  *2018-04-17T06:09:22Z*
* [OctoLinker](https://github.com/OctoLinker/browser-extension):star:0 - Navigate through go files efficiently with the OctoLinker browser extension for GitHub.  *2018-03-05T20:48:33Z*
* [richgo](https://github.com/kyoh86/richgo):star:258 - Enrich `go test` outputs with text decorations.  *2018-04-10T05:30:55Z*
* [rts](https://github.com/galeone/rts):star:170 - RTS: response to struct. Generates Go structs from server responses.  *2016-12-06T17:35:58Z*

## Software Packages

*Software written in Go.*

### DevOps Tools

* [aptly](https://github.com/smira/aptly):star:0 - aptly is a Debian repository management tool.  *2018-04-18T19:43:53Z*
* [aurora](https://github.com/xuri/aurora):star:271 - Cross-platform web-based Beanstalkd queue server console.  *2018-01-03T09:50:00Z*
* [awsenv](https://github.com/soniah/awsenv):star:12 - Small binary that loads Amazon (AWS) environment variables for a profile.  *2017-02-28T22:36:30Z*
* [Banshee](https://github.com/eleme/banshee):star:900 - Anomalies detection system for periodic metrics.  *2018-01-17T12:13:01Z*
* [Blast](https://github.com/dave/blast):star:127 - A simple tool for API load testing and batch jobs.  *2018-03-01T09:53:28Z*
* [bombardier](https://github.com/codesenberg/bombardier):star:973 - Fast cross-platform HTTP benchmarking tool.  *2018-04-10T14:25:49Z*
* [bosun](https://github.com/bosun-monitor/bosun):star:0 - Time Series Alerting Framework.  *2018-04-17T16:40:33Z*
* [dogo](https://github.com/liudng/dogo):star:169 - Monitoring changes in the source file and automatically compile and run (restart).  *2016-11-02T10:54:40Z*
* [drone-jenkins](https://github.com/appleboy/drone-jenkins):star:12 - Trigger downstream Jenkins jobs using a binary, docker or Drone CI.  *2017-10-05T03:33:53Z*
* [drone-scp](https://github.com/appleboy/drone-scp):star:28 - Copy files and artifacts via SSH using a binary, docker or Drone CI.  *2018-04-06T02:18:08Z*
* [Dropship](https://github.com/chrismckenzie/dropship):star:39 - Tool for deploying code via cdn.  *2016-06-25T05:13:43Z*
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy):star:40 - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.  *2017-11-23T13:37:01Z*
* [Gitea](https://github.com/go-gitea/gitea):star:0 - Fork of Gogs, entirely community driven.  *2018-04-21T01:23:11Z*
* [Go Metrics](https://github.com/rcrowley/go-metrics):star:0 - Go port of Coda Hale's Metrics library: https://github.com/codahale/metrics.  *2018-04-06T23:47:16Z*
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate):star:579 - Enable your Go applications to self update.  *2016-12-16T17:22:40Z*
* [gobrew](https://github.com/cryptojuice/gobrew):star:173 - gobrew lets you easily switch between multiple versions of go.  *2015-11-04T15:10:33Z*
* [godbg](https://github.com/sirnewton01/godbg):star:216 - Web-based gdb front-end application.  *2013-10-29T01:11:11Z*
* [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative):star:290 - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.  *2016-07-21T19:34:23Z*
* [govvv](https://github.com/ahmetalpbalkan/govvv):star:237 - “go build” wrapper to easily add version information into Go binaries.  *2017-11-02T04:38:29Z*
* [gox](https://github.com/mitchellh/gox):star:0 - Dead simple, no frills Go cross compile tool.  *2018-02-23T17:33:01Z*
* [goxc](https://github.com/laher/goxc):star:0 - build tool for Go, with a focus on cross-compiling and packaging.  *2018-03-02T08:56:19Z*
* [grapes](https://github.com/yaronsumel/grapes):star:101 - Lightweight tool designed to distribute commands over ssh with ease.  *2017-11-28T18:13:09Z*
* [GVM](https://github.com/moovweb/gvm):star:0 - GVM provides an interface to manage Go versions.  *2017-02-07T01:56:27Z*
* [Hey](https://github.com/rakyll/hey):star:0 - Hey is a tiny program that sends some load to a web application.  *2018-03-12T19:26:25Z*
* [kala](https://github.com/ajvb/kala):star:0 - Simplistic, modern, and performant job scheduler.  *2018-03-14T03:53:38Z*
* [kcli](https://github.com/cswank/kcli):star:17 - Command line tool for inspecting kafka topics/partitions/messages.  *2018-02-06T14:21:10Z*
* [kubernetes](https://github.com/kubernetes/kubernetes):star:0 - Container Cluster Manager from Google.  *2018-04-21T03:06:13Z*
* [manssh](https://github.com/xwjdsh/manssh):star:69 - manssh is a command line tool for managing your ssh alias config easily.  *2018-02-14T04:00:59Z*
* [Moby](https://github.com/moby/moby):star:0 - Collaborative project for the container ecosystem to assemble container-based systems.  *2018-04-21T09:03:54Z*
* [Mora](https://github.com/emicklei/mora):star:244 - REST server for accessing MongoDB documents and meta data.  *2017-01-09T17:07:37Z*
* [ostent](https://github.com/ostrost/ostent):star:150 - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.  *2018-04-03T20:54:00Z*
* [Packer](https://github.com/mitchellh/packer):star:0 - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.  *2018-04-22T01:51:17Z*
* [Pewpew](https://github.com/bengadbois/pewpew):star:149 - Flexible HTTP command line stress tester.  *2018-04-17T20:29:31Z*
* [Rodent](https://github.com/alouche/rodent):star:31 - Rodent helps you manage Go versions, projects and track dependencies.  *2017-04-22T07:48:06Z*
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r):star:911 - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.  *2017-02-10T00:40:45Z*
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli):star:418 - Manage BareMetal Servers from Command Line (as easily as with Docker).  *2018-04-20T13:49:00Z*
* [sg](https://github.com/ChristopherRabotin/sg):star:2 - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the reponse code and data between each call for specific server stress based on its previous response.  *2016-10-28T23:17:57Z*
* [skm](https://github.com/TimothyYe/skm):star:389 - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!  *2018-02-23T02:52:00Z*
* [StatusOK](https://github.com/sanathp/statusok):star:908 - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.  *2018-04-09T21:37:55Z*
* [traefik](https://github.com/containous/traefik):star:0 - Reverse proxy and load balancer with support for multiple backends.  *2018-04-19T14:30:10Z*
* [Vegeta](https://github.com/tsenart/vegeta):star:0 - HTTP load testing tool and library. It's over 9000!  *2018-03-24T20:33:31Z*
* [webhook](https://github.com/adnanh/webhook):star:0 - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.  *2018-01-11T09:34:55Z*
* [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
* [winrm-cli](https://github.com/masterzen/winrm-cli):star:33 - Cli tool to remotely execute commands on Windows machines.  *2017-09-15T06:28:57Z*

### Other Software
* [borg](https://github.com/crufter/borg):star:0 - Terminal based search engine for bash snippets.  *2018-02-07T19:40:05Z*
* [boxed](https://github.com/tejo/boxed):star:67 - Dropbox based blog engine.  *2016-04-23T21:19:47Z*
* [Cherry](https://github.com/rafael-santiago/cherry):star:157 - Tiny webchat server in Go.  *2016-12-14T15:17:46Z*
* [Circuit](https://github.com/gocircuit/circuit):star:0 - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.  *2016-11-21T20:35:24Z*
* [Comcast](https://github.com/tylertreat/Comcast):star:0 - Simulate bad network connections.  *2017-12-11T16:43:25Z*
* [confd](https://github.com/kelseyhightower/confd):star:0 - Manage local application configuration files using templates and data from etcd or consul.  *2018-03-24T15:28:39Z*
* [DDNS](https://github.com/skibish/ddns):star:43 - Personal DDNS client with Digital Ocean Networking DNS as backend.  *2018-03-11T11:33:08Z*
* [Docker](http://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
* [Documize](https://github.com/documize/community):star:400 - Modern wiki software that integrates data from SaaS tools.  *2018-04-20T13:38:35Z*
* [fleet](https://github.com/coreos/fleet):star:0 - Distributed init System.  *2018-04-12T21:54:09Z*
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store):star:844 - App that displays updates for the Go packages in your GOPATH.  *2018-02-26T22:59:04Z*
* [gocc](https://github.com/goccmack/gocc):star:166 - Gocc is a compiler kit for Go written in Go.  *2018-03-19T22:00:08Z*
* [GoDNS](https://github.com/timothyye/godns):star:199 - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.  *2018-04-19T02:57:36Z*
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip):star:11 - Chrome extension for Go Doc sites, which shows function description as tooltip at funciton list.  *2016-01-29T23:44:59Z*
* [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
* [Gor](https://github.com/buger/gor):star:0 - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.  *2018-02-27T12:25:29Z*
* [hugo](http://gohugo.io/) - Fast and Modern Static Website Engine.
* [ide](https://github.com/thestrukture/ide):star:164 - Browser accessible IDE. Designed for Go with Go.  *2018-04-15T11:05:46Z*
* [ipe](https://github.com/dimiro1/ipe):star:221 - Open source Pusher server implementation compatible with Pusher client libraries written in GO.  *2018-02-14T21:23:06Z*
* [JayDiff](https://github.com/yazgazan/jaydiff):star:10 - JSON diff utility written in Go.  *2018-01-24T12:36:13Z*
* [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [Leaps](https://github.com/jeffail/leaps):star:567 - Pair programming service using Operational Transforms.  *2018-02-19T20:38:20Z*
* [limetext](http://limetext.org/) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [LiteIDE](https://github.com/visualfc/liteide):star:0 - LiteIDE is a simple, open source, cross-platform Go IDE.  *2018-04-22T01:57:58Z*
* [mockingjay](https://github.com/quii/mockingjay-server):star:353 - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.  *2017-05-15T11:38:11Z*
* [myLG](https://github.com/mehrdadrad/mylg):star:0 - Command Line Network Diagnostic tool written in Go.  *2017-08-15T16:07:36Z*
* [naclpipe](https://github.com/unix4fun/naclpipe):star:10 - Simple NaCL EC25519 based crypto pipe tool written in Go.  *2018-04-02T11:16:55Z*
* [nes](https://github.com/fogleman/nes):star:0 - Nintendo Entertainment System (NES) emulator written in Go.  *2018-02-22T22:04:12Z*
* [orange-cat](https://github.com/noraesae/orange-cat):star:162 - Markdown previewer written in Go.  *2015-09-12T14:13:44Z*
* [Orbit](https://github.com/gulien/orbit):star:88 - A simple tool for running commands and generating files from templates.  *2018-04-13T14:47:37Z*
* [peg](https://github.com/pointlander/peg):star:480 - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.  *2017-12-12T06:07:19Z*
* [Pipe](https://github.com/b3log/pipe):star:814 - A small and beautiful blogging platform.  *2018-04-18T02:09:58Z*
* [Postman](https://github.com/zachlatta/postman):star:708 - Command-line utility for batch-sending email.  *2017-02-27T02:38:15Z*
* [restic](https://github.com/restic/restic):star:0 - De-duplicating backup program.  *2018-04-16T15:29:09Z*
* [rkt](https://github.com/coreos/rkt):star:0 - App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM.  *2018-04-16T09:50:59Z*
* [Seaweed File System](https://github.com/chrislusf/seaweedfs):star:0 - Fast, Simple and Scalable Distributed File System with O(1) disk seek.  *2018-04-17T19:53:38Z*
* [shell2http](https://github.com/msoap/shell2http):star:223 - Executing shell commands via http server (for prototyping or remote control).  *2018-04-19T12:48:07Z*
* [snap](https://github.com/intelsdi-x/snap):star:0 - Powerful telemetry framework.  *2018-02-06T09:01:08Z*
* [Snitch](https://github.com/lucasgomide/snitch):star:14 - Simple way to notify your team and many tools when someone has deployed any application via Tsuru.  *2017-11-25T02:14:12Z*
* [Stack Up](https://github.com/pressly/sup):star:0 - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.  *2018-04-18T15:51:15Z*
* [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
* [Tenyks](https://github.com/kyleterry/tenyks):star:162 - Service oriented IRC bot using Redis and JSON for messaging.  *2017-03-05T08:10:39Z*
* [term-quiz](https://github.com/crazcalm/term-quiz):star:10 - Quizzes for your terminal.  *2018-02-06T17:30:38Z*
* [toto](https://github.com/blogcin/ToTo):star:20 - Simple proxy server written in Go language, can be used together with browser.  *2016-10-31T11:20:32Z*
* [toxiproxy](https://github.com/shopify/toxiproxy):star:0 - Proxy to simulate network and system conditions for automated tests.  *2018-04-17T15:12:33Z*
* [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow):star:366 - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.  *2018-01-11T18:08:39Z*
* [websysd](https://github.com/ian-kent/websysd):star:41 - Web based process manager (like Marathon or Upstart).  *2016-05-07T21:41:27Z*
* [wellington](https://github.com/wellington/wellington):star:267 - Sass project management tool, extends the language with sprite functions (like Compass).  *2018-02-23T18:24:46Z*

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [autobench](https://github.com/davecheney/autobench):star:86 - Framework to compare the performance between different Go versions.  *2014-06-19T10:18:26Z*
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app):star:12 - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.  *2017-03-17T11:40:30Z*
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks):star:94 - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.  *2016-02-25T05:42:28Z*
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark):star:0 - Go HTTP request router benchmark and comparison.  *2017-05-31T09:54:19Z*
* [go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark):star:4 - Naive performance test of two ways to do type assertion in Go.  *2014-10-28T22:49:01Z*
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark):star:698 - Go web framework benchmark.  *2018-01-19T07:47:06Z*
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks):star:636 - Benchmarks of Go serialization methods.  *2017-09-29T02:30:39Z*
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel):star:50 - Benchmarks of common basic operations for the Go language.  *2015-01-04T15:42:27Z*
* [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks):star:14 - Tiny collection of Go micro benchmarks. The intent is to compare some language features to others.  *2015-12-18T22:27:44Z*
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark):star:43 - Collection of benchmarks for popular Go database/SQL utilities.  *2015-03-13T21:22:41Z*
* [gospeed](https://github.com/feyeleanor/GoSpeed):star:80 - Go micro-benchmarks for calculating the speed of language constructs.  *2016-01-27T16:20:40Z*
* [kvbench](https://github.com/jimrobinson/kvbench):star:12 - Key/Value database benchmark.  *2014-04-20T12:17:11Z*
* [skynet](https://github.com/atemerev/skynet):star:840 - Skynet 1M threads microbenchmark.  *2016-08-01T10:52:32Z*
* [speedtest-resize](https://github.com/fawick/speedtest-resize):star:128 - Compare various Image resize algorithms for the Go language.  *2017-03-13T22:31:46Z*

## Conferences

* [Capital Go](http://www.capitalgolang.com) - Washington, D.C., USA
* [dotGo](http://www.dotgo.eu) - Paris, France
* [GoCon](http://gocon.connpass.com/) - Tokyo, Japan
* [GoLab](http://golab.io/) - Florence, Italy
* [GolangUK](http://golanguk.com/) - London, UK
* [GopherChina](http://gopherchina.org) - Shanghai, China
* [GopherCon](http://www.gophercon.com/) - Denver, USA
* [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, BR
* [GopherCon Dubai](http://www.gophercon.ae/) - Dubai, UAE
* [GopherCon Europe](https://gophercon.is/) - Reykjavik, Iceland
* [GopherCon India](http://www.gophercon.in/) - Pune, India
* [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia
* [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore
* [GothamGo](http://gothamgo.com/) - New York City, USA

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks):star:0 - A curated list of Go books.  *2017-06-22T19:39:01Z*
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)

## Gophers

* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector):star:20 - Go gopher Vector Data [.ai, .svg]  *2018-03-04T07:19:32Z*
* [gopher-logos](https://github.com/GolangUA/gopher-logos):star:33 - adorable gopher logos  *2017-09-28T16:23:00Z*
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gopherize.me](https://github.com/matryer/gopherize.me):star:200 - Gopherize yourself  *2017-08-17T15:44:35Z*
* [gophers](https://github.com/ashleymcnamara/gophers):star:0 - Gopher artworks by Ashley McNamara  *2018-03-23T20:46:10Z*
* [gophers](https://github.com/egonelbre/gophers):star:0 - Free gophers  *2018-04-07T19:43:52Z*
* [gophers](https://github.com/rogeralsing/gophers):star:45 - random gopher graphics  *2017-03-18T21:53:19Z*
* [gophers](https://github.com/sillecelik/go-gopher):star:14 - Gopher amigurumi toy pattern  *2018-04-03T17:55:12Z*

## Meetups

* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoJakarta](https://www.meetup.com/GoJakarta/)
* [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
* [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
* [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
* [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
* [Golang Boston](https://www.meetup.com/bostongo/)
* [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
* [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
* [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
* [Golang Israel](https://www.meetup.com/Go-Israel/)
* [Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)
* [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
* [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
* [Golang Melbourne](https://www.meetup.com/golang-mel/)
* [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
* [Golang New York](https://www.meetup.com/nycgolang/)
* [Golang Paris](https://www.meetup.com/Golang-Paris/)
* [Golang Pune](https://www.meetup.com/Golang-Pune/)
* [Golang Singapore](https://www.meetup.com/golangsg/)
* [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
* [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Москва](https://www.meetup.com/Golang-Moscow/)
* [Golang Питер](https://www.meetup.com/Golang-Peter/)
* [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
* [Seattle Go Programmers](https://www.meetup.com/golang/)
* [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
* [Utah Go User Group](https://www.meetup.com/utahgophers/)
* [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

*Add the group of your city/country here (send **PR**)*

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## Websites

* [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job):star:0 - Curated list of awesome remote jobs. A lot of them are looking for Go hackers.  *2018-04-21T14:30:04Z*
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness):star:0 - List of other amazingly awesome lists.  *2018-04-12T13:19:28Z*
* [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - Collection of Go articles and tutorials.
* [Go Blog](http://blog.golang.org) - The official Go blog.
* [Go Challenge](http://golang-challenge.org/) - Learn Go by solving problems and getting feedback from Go experts.
* [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/) - 5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com) - A report card for your Go package.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp):star:26 - Collection of Go projects that needs help. Good place to start your open-source way in Go.  *2017-09-23T14:04:03Z*
* [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
* [Golang Flow](https://golangflow.io) - Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com) - Links and news about Go programming.
* [golang-graphics](https://github.com/mholt/golang-graphics):star:131 - Collection of Go images, graphics, and art.  *2015-08-24T21:29:48Z*
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [gowalker.org](https://gowalker.org) - Go Project API documentation.
* [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [r/Golang](https://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.

### Tutorials

* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go.
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang):star:0 - Golang ebook intro how to build a web app with golang.  *2018-04-20T09:01:47Z*
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Games With Go](http://gameswithgo.org/) - A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet):star:0 - Go's reference card.  *2018-03-15T14:28:31Z*
* [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
* [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
* [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests):star:0 - Learn Go with test-driven development.  *2018-04-21T06:52:11Z*
* [Working with Go](https://github.com/mkaz/working-with-go):star:953 - Intro to go for experienced programmers.  *2018-04-07T22:29:06Z*
* [Your basic Go](http://yourbasic.org/golang) - Huge collection of tutorials and how to's
