# go-csgomatchacceptor
CSGO Auto Match Acceptor (experimental)
(This is my first program in GO, may not be very good)
Program is only tested in Windows 10.

# How does it work
It fetch coordinates from "point.txt". Then it detecting pixel color at that point while in CS:GO. If pixel color is any colors of accept button, program will move mouse and click.

# How to set the point
- Take a screenshot when you see the accept button in CS:GO.
- Go to [Image-Map](https://www.image-map.net/), select your screenshot from your pc.
- Click one time to the green part of accept button in the screenshot.
- Click "Show Me The Code" button and copy coords in X,Y format here. (coords="X,Y")
- Paste the coords to points.txt

# Build
Dependencies: [robotgo](https://github.com/go-vgo/robotgo)

You can build for all platforms. But some compilers and libraries are required for build of RobotGo. RobotGo supports Mac, Windows, and Linux(X11).