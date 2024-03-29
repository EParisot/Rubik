# Rubik
Résumé: Ce projet va vous faire écrire un programme qui resoud des Rubik’s Cubes en un minimum de mouvements.
Summary: This project will make you write a program that resolves Rubik's Cubes in a minimum of movements.

# How it works
We implemented two algoritms:

The default one is a Thistlethwaite's 4 phases algorithm. 
But instead of using huge, pre-computed prunning tables, we implemented 4 heuristics matching the 4 standard algorithm groups and which could provide sufficient estimation for an IDAstar (depthFirstSearch), performing quite well on randomly scrumbled cubes.
This implementation will find short solutions (20-50 steps), but may take from a few seconds to few minutes to return.

<p align="center">
<img src="demos/Demo_Thistlethwaite.gif" width="400">
</p>

The second one is a "Human" algorithm.
Based on "CFOP" method. It uses several algorithms to solve known configurations as it evolve to solution.

See : https://ruwix.com/the-rubiks-cube/how-to-solve-the-rubiks-cube-beginners-method/

Although providing big solutions, the computational time is near from zero.

<p align="center">
<img src="demos/Demo_Human.gif" width="400">
</p>

# Metrics
All metrics mesured on 100 runs on the same machine (i5 8250U / 8Go RAM)
Timeout is 1min

Thistlethwaite IDAstar:
```
Average solution length: 37
Average compute time: 16.73s
92 Success and 8 Timeouts, rate = 0.92
```

Human CFOP:
```
Average solution length: 206
Average compute time: 0.01s
100 Success and 0 Timeouts, rate = 1.00
```

# Try it :

### Python GUI:
```
pip3 install -r requirements.txt
python3 Rubik_Cube.py [-h] ["R2 D’ B’ D F2 R F2 R2 U L’ F2 U’ B’ L2 R D B’ R’ B2 L2 F2 L2 R2 U2 D2"]
```
with -h : "human" algorithm

### or Golang CLI :
```
go build
./Rubik[.exe] [-d] [-h] "R2 D’ B’ D F2 R F2 R2 U L’ F2 U’ B’ L2 R D B’ R’ B2 L2 F2 L2 R2 U2 D2"
```
with -d : CLI debug mode
     -h : "human" algorithm


#### Credits to rabbid76 for these two awsome StackOverflow Posts about Pygames OpenGL:
https://stackoverflow.com/questions/50303616/how-to-rotate-slices-of-a-rubiks-cube-in-python-pyopengl/54953213#54953213
https://stackoverflow.com/questions/50312760/pygame-opengl-3d-cube-lag
