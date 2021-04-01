# go-csgomatchacceptor
CSGO Auto Match Acceptor (experimental)
(This is my first program in GO, may not be very good)
Program is only tested in Windows 10.

# How does it work
It is taking a screenshot of a certain area of the screen in CS:GO. Then it is getting to pixel colors from screenshot. If pixel colors is match colors of accept button, program will move mouse and click. The area of this screenshot and mouse click position varies according to the screen resolution. Therefore, it may not work correctly every time.

# Config
| key      | type   | description                                              |
|----------|--------|----------------------------------------------------------|
| tgtoken  | string | telegram bot token taken from [botfather](https://t.me/BotFather).                 |
| tguserid | int    | your telegram user id ([find](http://t.me/userinfobot)) |
| test     | bool   | if true; it just moves the mouse, don't click.           |


# Build
Dependencies: [robotgo](https://github.com/go-vgo/robotgo) [telebot](https://gopkg.in/tucnak/telebot.v2)

You can build for all platforms. But some compilers and libraries are required for build of RobotGo. RobotGo supports Mac, Windows, and Linux(X11).